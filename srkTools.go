package srkTools

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"
)

type SrkTools struct {

}

func DebugLog(debug, msg string)  {
	if debug != "" {
		fmt.Println(msg)
	}
}

//func DecodeJson(resp *http.Response, sct interface{}) map[string]interface{} {
func DecodeJson(logPrefix, debugLevel string, resp *http.Response, sct interface{}) {

	body, _ := ioutil.ReadAll(resp.Body)

	_ = json.Unmarshal(body, &sct)
	defer resp.Body.Close()

	DebugLog(debugLevel, fmt.Sprintf("%s %s",logPrefix, body))

	/*
	result := sct.(map[string]interface{})
	return result
	*/
}

func GetCstTime() time.Time {
	TZ := time.FixedZone("CST", 8*3600)
	return time.Now().In(TZ)
}

func (srk SrkTools) CustomCmd(command string) {
	/*
	执行linux命令
	 */
	cmd := exec.Command("sh", "-c", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%s\n", output)
	}
}

func (srk SrkTools) Download(filePath string, url string) {
	/*
	下载文件
	 */
	resp := srk.HttpReq(url)

	f, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(f, resp.Body)
	if err != nil {
		panic(err)
	}

	//if strings.Contains(filePath, "sbin") {
	//	cmd := fmt.Sprintf("chmod +x %s", filePath)
	//	customCmd(cmd)
	//}

	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
}

func (srk SrkTools) HttpReq(url string) *http.Response {
	/*
	http 请求
	return response
	 */

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return resp
}
