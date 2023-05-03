package main

import "net/http"

func handler1(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("这里是handler1\n"))
}


func handler2(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("这里是handler2\n"))
}

func main(){
	myHandler:=&HandleBasesHttp{}
	http.Handle("/test/",myHandler)
	http.ListenAndServe(":8080",nil)
}

type HandleBasesHttp struct {

}

func (m *HandleBasesHttp) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path=="/"{
		writer.Write([]byte("这里是handler1\n"))
	} else{
		writer.Write([]byte(request.URL.Path))
	}
}
