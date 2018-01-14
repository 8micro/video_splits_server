package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

var FFMPEGPath string = "/data/web/ffmpeg_ubuntu/ffmpeg-git-20171206-64bit-static/ffmpeg"
var SplitTimeInterval int32 = 5

func main() {
	//todo http server

	http.HandleFunc("/split_video", SplitVideoHandler)

	log.Fatal(http.ListenAndServe("127.0.0.1:7745", nil))
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
	outputVideoListName := req.PostFormValue("video_list_name")
	//ffmpeg -i test.mp4 -c copy -map 0 -f segment -segment_list playlist.m3u8 -segment_time 5 ./video_output/output%03d.ts

	args := fmt.Sprintf("-i %s -c copy -map 0 -f segment -segment_list %s.m3u8 -segment_time %d %s'%03d'.ts", videoPath, outputVideoListName, SplitTimeInterval, outputPath)

	cmd := exec.Command(FFMPEGPath, args)
	out := new(bytes.Buffer)
	cmd.Stdout = out
	err := cmd.Run()
	if err != nil {
		log.Printf("run cmd  %s  failed", args)
		return
	}
	outStr := out.String()
	log.Printf("execute result %s", outStr)
}
