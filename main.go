package main

import (
	"encoding/json"
	"errors"
	"github.com/hsw409328/gofunc/go_http"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
	"time"
)

// GO版本上传zimg方法
func main() {
	f, _ := os.Open("/Users/xxxxx/Downloads/188843622.jpg")
	body, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	accessZIMGUrl := "http://127.0.0.1:4869/"
	resp, err := go_http.HttpPost(&go_http.RequestOptions{
		RequestTimeout: 10 * time.Second,
		DialTimeout:    10 * time.Second,
		ContentType:    strings.Replace(strings.ToLower(path.Ext(f.Name())), ".", "", -1),
		ReqData:        string(body),
		UrlStr:         accessZIMGUrl + "upload",
	}, nil)
	if err != nil {
		panic(err)
	}

	var tmp = struct {
		Ret  bool
		Info struct {
			Md5  string
			Size int
		}
		Error struct {
			Code    int
			Message string
		}
	}{}
	err = json.Unmarshal([]byte(resp.GetBodyString()), &tmp)
	if err != nil {
		panic(err)
	}
	if tmp.Ret == false {
		panic(errors.New(tmp.Error.Message))
	}
	log.Println(accessZIMGUrl + tmp.Info.Md5)
}
