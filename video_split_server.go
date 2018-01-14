package main

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"

	uuid "github.com/satori/go.uuid"
)

var FFMPEGPath string = "/data/web/ffmpeg_ubuntu/ffmpeg-git-20171206-64bit-static/ffmpeg"
var SplitTimeInterval int32 = 5

func main() {
	//todo http server

	go doSplitTest()

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

func doSplitTest() {
	args := fmt.Sprintf("-i %s -c copy -map 0 -f segment -segment_list /data/web/video/test_output/%s.m3u8 -segment_time %d %s%s03d.ts",
		"/data/web/video/test2.mp4", "split_test", 5, "/data/web/video/test_output/split_test_", "%")
	//args := EncodingArgs("/data/web/video/test2.mp4", "/data/web/video/test_output/", "test_video_manifest", 5, "/data/web/video/test_output/", "split_test_")

	err := execute(FFMPEGPath, []string{args})
	if err != nil {
		log.Fatal(err)
	}

}

func execute(cmdPath string, args []string) (err error) {

	u1 := uuid.Must(uuid.NewV4())
	uuidStr := u1.String()
	scriptPathName := fmt.Sprintf("/tmp/%s.sh", uuidStr)

	log.Printf("create script at %s", scriptPathName)

	f, err := os.Create(scriptPathName)
	if err != nil {
		log.Printf("create script file faild")
		return errors.New("create script file failed")
	}

	scriptBegin := "#!/bin/bash\n"
	scriptBeginBytes := []byte(scriptBegin)

	f.Write(scriptBeginBytes)

	cmdArgs := fmt.Sprintf(" -i %s -c copy -map 0 -f segment -segment_list /data/web/video/test_output/%s.m3u8 -segment_time %d %s%s03d.ts",
		"/data/web/video/test2.mp4", "split_test", 5, "/data/web/video/test_output/split_test_", "%")
	cmdStr := cmdPath + cmdArgs

	cmdBytes := []byte(cmdStr)
	var writeBytes int
	writeBytes, err = f.Write(cmdBytes)
	if err != nil {
		log.Printf("write scripts failed")
		return errors.New("write scripts failed")
	}
	log.Printf("write script file bytes %v", writeBytes)
	f.Close()

	err = os.Chmod(scriptPathName, 777)
	if err != nil {
		log.Printf("change file to have write failed")
		return errors.New("change scripts failed")
	}

	cmd := exec.Command(scriptPathName)
	err = cmd.Run()
	if err != nil {
		log.Printf("run command failed %v", err.Error())
	}
	log.Printf("execute finished")

	return nil
}

func EncodingArgs(videoInput string, videoManifestOutputPath string, videoManifestPrefix string, segmentInterval int, videoSegOutputPath string, videoSegOutputPrefix string) []string {
	//args := fmt.Sprintf("-i %s -c copy -map 0 -f segment -segment_list /data/web/video/test_output/%s.m3u8 -segment_time %d %s%s03d.ts",
	//	"/data/web/video/test2.mp4", "split_test", 5, "/data/web/video/test_output/split_test_", "%")
	// see http://superuser.com/questions/908280/what-is-the-correct-way-to-fix-keyframes-in-ffmpeg-for-dash
	return []string{
		// Prevent encoding to run longer than 30 seonds
		"-i", videoInput,
		"-c copy -map 0 -f segment  -segment_list", videoManifestOutputPath, videoManifestPrefix, ".m3u8",
		"-segment_time",
		strconv.Itoa(segmentInterval),
		videoSegOutputPath,
		videoSegOutputPrefix,
		"%d03d.ts",
		//"pipe:out%03d.ts",
	}
}
