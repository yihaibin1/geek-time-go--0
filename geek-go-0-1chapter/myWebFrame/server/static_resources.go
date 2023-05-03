package server

import (
	"fmt"
	lru "github.com/hashicorp/golang-lru"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type StaticResourceHandlerOption func(h *StaticResourceHandler)

type StaticResourceHandler struct {
	//去哪找静态文件
	dir string
	//路由前缀
	pathPrefix string
	//支持文件类型map
	extensionContentTypeMap map[string]string


	//缓存静态资源的限制
	cache *lru.Cache
	maxFileSize int
}

type fileCacheItem struct {
	filename string
	filesize int
	contentType string
	data []byte
}

func NewStaticResourceHandler(dir string,
	pathPrefix string,
	opts...StaticResourceHandlerOption)*StaticResourceHandler{
	s:= &StaticResourceHandler{
		dir:                     dir,
		pathPrefix:              pathPrefix,
		extensionContentTypeMap: map[string]string{
			// 这里根据自己的需要不断添加
			"jpeg": "image/jpeg",
			"jpe": "image/jpeg",
			"jpg": "image/jpeg",
			"png": "image/png",
			"pdf": "image/pdf",
		},
	}
	for _,o:=range opts{
		o(s)
	}
	return s
}

// WithFileCache 静态文件将会被缓存
// maxFileSizeThreshold 超过这个大小的文件，就被认为是大文件，我们将不会缓存
// maxCacheFileCnt 最多缓存多少个文件
// 所以我们最多缓存 maxFileSizeThreshold * maxCacheFileCnt
func WithFileCache(maxFileSizeThreshold int,maxCacheFileCnt int)StaticResourceHandlerOption{
	return func(h *StaticResourceHandler) {
		c,err:=lru.New(maxCacheFileCnt)
		if err!=nil{
			fmt.Println("something wrong happened with the cache")
		}
		h.maxFileSize=maxFileSizeThreshold
		h.cache=c
	}
}

func WithMoreExtension(extensionMap map[string]string)StaticResourceHandlerOption{
	return func(h *StaticResourceHandler){
		for key,val:=range extensionMap{
			h.extensionContentTypeMap[key]=val
		}
	}
}

func(h *StaticResourceHandler)ServerStaticResource(c *Context){
	//去掉路径的前缀
	req:=strings.TrimPrefix(c.R.URL.Path,h.pathPrefix)
	if item,ok:=h.readFileFromDataI(req);ok{
		//读取成功
		fmt.Println("read data from cache")
		h.writeItemAsResponse(item,c.W)
		return
	}
	//没读到就去静态文件路径寻找并加入
	path:=filepath.Join(h.dir,req)
	f,err:=os.Open(path)
	if err!=nil{
		//没读到想要的文件
		c.W.WriteHeader(http.StatusInternalServerError)
		return
	}
	ext:=getFileExt(path)
	t,ok:=h.extensionContentTypeMap[ext]
	if !ok{
		c.W.WriteHeader(http.StatusBadRequest)
		return
	}


	data,err:=ioutil.ReadAll(f)
	if err!=nil{
		c.W.WriteHeader(http.StatusInternalServerError)
		return
	}

	item:=&fileCacheItem{
		filename:    req,
		filesize:    len(data),
		contentType: t,
		data:        data,
	}

	h.cache.Add(item.filename,item)
	h.writeItemAsResponse(item,c.W)
}

func(h *StaticResourceHandler)readFileFromDataI(fileName string)(*fileCacheItem,bool){
	if item,ok:=h.cache.Get(fileName);ok{
		return item.(*fileCacheItem),true
	}
	return nil,false
}

func(h *StaticResourceHandler)writeItemAsResponse(item *fileCacheItem,writer http.ResponseWriter){
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type",item.contentType)
	writer.Header().Set("Content-Length",fmt.Sprintf("%d",item.filesize))
	writer.Write(item.data)
}


func getFileExt(name string)string{
	index:=strings.LastIndex(name,".")
	if index==len(name)-1{
		return ""
	}
	return name[index+1:]
}