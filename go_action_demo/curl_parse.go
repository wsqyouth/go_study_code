package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"golang.org/x/net/context"
)

var finalResult string

func main() {

	curlPrefixURL := "http://featureflag.gxt.oa.com/feature/flag/list?app_key=ad_platform&env_key=production&limit=20&offset=0"
	ffStrArray := []string{
		"adq_support_pt23",
		"channles_customized_convert_component",
	}
	for _, ffStr := range ffStrArray {
		curlURL := fmt.Sprintf(curlPrefixURL+`&flag_keyword=%s&is_myflag=0&is_mycreate=0&status=1`, ffStr)
		if err := process(curlURL); err != nil {
			panic("process error")
		}
	}
	ctx := context.Background()
	if err := writeFile(ctx, "ff_res.txt", finalResult); err != nil {
		panic("writeFile error")
	}
}

func process(inputURL string) error {
	req, err := http.NewRequest("GET", inputURL, nil)
	if err != nil {
		fmt.Printf("curl url error: %+v", err)
		return err
	}
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		fmt.Printf("curl url error: %+v", err)
		return err
	}
	defer resp.Body.Close()

	respByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("curl url error: %+v", err)
		return err
	}

	//fmt.Printf("curl url succ. err: %v, body: %s\n", err, respByte)
	var response Response
	if err = json.Unmarshal(respByte, &response); err != nil {
		fmt.Printf("Unmarshal error: %+v", err)
		return err
	}
	responseData := response.ReturnMsg.Data[0]
	modMap := map[int]string{
		0: "全关",
		1: "规则匹配",
		2: "全开",
	}

	resultStr := fmt.Sprintf("   %s   %s    %s    %s", responseData.FenvKey, responseData.FflagKey, modMap[responseData.Fmode], responseData.Foperator)
	fmt.Println(resultStr)
	finalResult += "\n" + resultStr
	return nil
}

func writeFile(ctx context.Context, filePath string, content string) error {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("文件打开失败:%v", err)
		return err
	}
	// 关闭文件
	defer file.Close()

	// 字符串写入
	_, err = file.WriteString(content)
	if err != nil {
		fmt.Printf("字符串写入失败,err:%v", err)
		return err
	}
	return nil
}

type Response struct {
	ReturnCode int       `json:"return_code"`
	ReturnMsg  ReturnMsg `json:"return_msg"`
}
type Data struct {
	Fid            int    `json:"Fid"`
	FappKey        string `json:"Fapp_key"`
	FenvKey        string `json:"Fenv_key"`
	FflagKey       string `json:"Fflag_key"`
	Fname          string `json:"Fname"`
	Fstatus        int    `json:"Fstatus"`
	Fversion       string `json:"Fversion"`
	Fmode          int    `json:"Fmode"`
	Fdescription   string `json:"Fdescription"`
	FmodifyTime    string `json:"Fmodify_time"`
	FexpectEndTime string `json:"Fexpect_end_time"`
	FsyncEnvKey    string `json:"Fsync_env_key"`
	Fperson        string `json:"Fperson"`
	Foperator      string `json:"Foperator"`
	FappEditable   bool   `json:"Fapp_editable"`
	FflagEditable  bool   `json:"Fflag_editable"`
	Feditable      bool   `json:"Feditable"`
	IsSynchronized bool   `json:"is_synchronized"`
}
type ReturnMsg struct {
	Total int    `json:"total"`
	Data  []Data `json:"data"`
}
