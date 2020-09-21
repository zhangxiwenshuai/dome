package main

import "net/http"

func main() {
	http.Handle("/",http.FileServer(http.Dir("C:")))
	http.ListenAndServe("127.1.1.1:9000",nil)

}
