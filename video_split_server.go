package main

import (
	"log"
	"net/http"
)

func main() {
	//todo http server
	http.HandleFunc("/split_video", SplitVideoHandler)

	log.Fatal(http.ListenAndServe("http://127.0.0.1:7745", nil))
}

func SplitVideoHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		doSplit(w, req)
		return
	default:
		log.Printf("unknow method")
		return
	}
}

func doSplit(w http.ResponseWriter, req *http.Request) {
	videoPath := req.PostFormValue("video_path")
	outputPath := req.PostFormValue("output_path")

}
