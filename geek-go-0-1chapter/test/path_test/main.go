package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"%s\n",r.URL.Path)
	fmt.Fprintf(w,"%s\n",r.URL.Query())

}

func main(){
	http.HandleFunc("/",handler)
	http.ListenAndServe(":8080",nil)
}
