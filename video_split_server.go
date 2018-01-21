package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
)

var FFMPEGPath string = "/data/web/ffmpeg_ubuntu/ffmpeg-git-20171206-64bit-static/ffmpeg"
var SplitTimeInterval int32 = 5

const DbURL string = "http://127.0.0.1/save_video"

type SplitRequest struct {
	UserID        string `json:"userid"`       // user id
	UUID          string `json:"uuid"`         //file upload uuid
	FilePathName  string `json:"filepathname"` //full path name
	FileDir       string `json:"filedir"`      // file dir
	FileName      string `json:"filename"`     //file name
	VideoHeight   string `json:videoheight`
	VideoWidth    string `json:"videowidth"`
	VideoRate     string `json:"videorate"`
	VideoDuration string `json:"videoDuration"`
	VideoFileSize string `json:"videofilesize"`
}

type SaveVideoRequest struct {
	UserID                    string `json:"userid"`
	UUID                      string `json:"uuid"`
	VideoOriginFilePathName   string `json:"video_file_path_name"`
	VideoOriginFileDir        string `json:"video_file_dir"`
	VideoFileManifestPathName string `json:"video_manifest"` // full path for video splits manifest file .m3u8
	VideoManifestDir          string `json:"video_manifest_dir"`
	VideoHeight               string `json:"video_height"`
	VideoWidth                string `json:"video_width"`
	VideoRate                 string `json:"video_rate"`
	VideoDuration             string `json:"video_duration"`
	VideoFileSize             string `json:"video_file_size"`
}

func main() {
	//todo http server

	//go doSplitTest()

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

/*
	UserID        string `json:"userid"`       // user id
	UUID          string `json:"uuid"`         //file upload uuid
	FilePathName  string `json:"filepathname"` //full path name
	FileDir       string `json:"filedir"`      // file dir
	FileName      string `json:"filename"`     //file name
	VideoHeight   string `json:videoheight`
	VideoWidth    string `json:"videowidth"`
	VideoRate     string `json:"videorate"`
	VideoDuration string `json:"videoDuration"`
	VideoFileSize string `json:"videofilesize"`
*/

func doSplit(w http.ResponseWriter, req *http.Request) {
	body_bytes := make([]byte, req.ContentLength)

	_, err0 := req.Body.Read(body_bytes)

	log.Printf("recv split request : %v", body_bytes)
	var request SplitRequest
	err0 = json.Unmarshal(body_bytes, &request)
	if err0 != nil {
		log.Printf("Unmarshal failed  %v for %v ", err0.Error(), body_bytes)
		return
	}
	log.Printf("split request %v", request)
	userID := request.UserID
	uuid := request.UUID
	filePathName := request.FilePathName
	fileDir := request.FilePathName
	//fileName := request.FileName
	videoheight := request.VideoHeight
	videoWidth := request.VideoWidth
	videoRate := request.VideoRate
	videoDuration := request.VideoDuration
	videoFileSize := request.VideoFileSize

	videoManifestFilePathName := filePathName + ".m3u8"
	videoSegmentsPathNamePrefix := filePathName

	//outputVideoListName := req.PostFormValue("video_list_name")
	//ffmpeg -i test.mp4 -c copy -map 0 -f segment -segment_list playlist.m3u8 -segment_time 5 ./video_output/output%03d.ts

	args := fmt.Sprintf(" -i %s -c copy -map 0 -f segment -segment_list %s -segment_time %d %s%s03d.ts",
		filePathName, videoManifestFilePathName, SplitTimeInterval, videoSegmentsPathNamePrefix, "%")

	err1 := execute(FFMPEGPath, args, userID, uuid)
	if err1 != nil {
		log.Printf("execute split script failed for user %s  file %s \n", userID, filePathName)
	}

	//send post request to api server
	var saveRequest SaveVideoRequest
	saveRequest.UserID = userID
	saveRequest.UUID = uuid
	saveRequest.VideoDuration = videoDuration
	saveRequest.VideoFileManifestPathName = videoManifestFilePathName
	saveRequest.VideoFileSize = videoFileSize
	saveRequest.VideoHeight = videoheight
	saveRequest.VideoManifestDir = fileDir
	saveRequest.VideoOriginFileDir = fileDir
	saveRequest.VideoOriginFilePathName = filePathName
	saveRequest.VideoRate = videoRate
	saveRequest.VideoWidth = videoWidth

	b, err2 := json.Marshal(saveRequest)
	if err2 != nil {
		log.Printf("marshal json save request failed for %v", saveRequest)
		return
	}
	sendreq := bytes.NewBuffer(b)
	resp, err3 := http.Post(DbURL, "application/json;charset=utf-8", sendreq)
	if err3 != nil {
		log.Printf("send save request to db failed")
		return
	}
	if resp.Status == "200 OK" {
		log.Printf("save video to db success")
	} else {
		log.Printf("save video to db failed %v", saveRequest)
	}

}

//fmt.Sprintf(" -i %s -c copy -map 0 -f segment -segment_list %s -segment_time %d %s%s03d.ts",
//	"/data/web/video/test2.mp4", "split_test", 5, "/data/web/video/test_output/split_test_", "%")
//
func execute(cmdPath string, args string, userId string, uuid string) (err error) {

	scriptPathName := fmt.Sprintf("/tmp/%s/%s.sh", userId, uuid)
	scriptDir := fmt.Sprintf("/tmp/%s/", userId)

	if err0 := os.MkdirAll(scriptDir, 0777); err0 != nil {
		log.Printf("create dir failed %s", scriptDir)
		return errors.New("create dir failed %s")
	}

	log.Printf("create script at %s", scriptPathName)

	f, err := os.Create(scriptPathName)
	if err != nil {
		log.Printf("create script file failed %v", err.Error())
		return errors.New("create script file failed")
	}

	scriptBegin := "#!/bin/bash\n"
	scriptBeginBytes := []byte(scriptBegin)

	f.Write(scriptBeginBytes)

	cmdStr := cmdPath + args

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
