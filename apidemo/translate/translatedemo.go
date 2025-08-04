package main

import (
	"demo/apidemo/utils"
	"demo/apidemo/utils/authv3"
	"encoding/json"
	"fmt"
	"os"
)

// 您的应用ID
var appKey = "6625a3d773ac0b7b"

// 您的应用密钥
var appSecret = "BbCFXUNQ8k8NXm1Ks2W5kDUWtERAiCRj"

type AlfredItem struct {
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Arg      string `json:"arg"`
}

type AlfredResponse struct {
	Items []AlfredItem `json:"items"`
}

func main() {
	text := os.Args[1]
	// 添加请求参数
	paramsMap := createRequestParams(text)
	header := map[string][]string{
		"Content-Type": {"application/x-www-form-urlencoded"},
	}
	// 添加鉴权相关参数
	authv3.AddAuthParams(appKey, appSecret, paramsMap)
	// 请求api服务
	result := utils.DoPost("https://openapi.youdao.com/api", header, paramsMap, "application/json")
	// 打印返回结果
	if result != nil {
		s := string(result)
		var result map[string]interface{}
		if err := json.Unmarshal([]byte(s), &result); err != nil {
			panic(err)
		}

		var items = []AlfredItem{}
		translations := result["translation"]
		queryOrigin := result["query"]
		if trVal, ok := translations.([]interface{}); ok {
			for _, val := range trVal {
				item := AlfredItem{
					Title:    val.(string),
					Subtitle: queryOrigin.(string),
					Arg:      val.(string),
				}
				items = append(items, item)
			}
		} else {
			fmt.Println("url 字段不存在或不是字符串")
		}

		response := AlfredResponse{
			Items: items,
		}
		// 生成 JSON
		data, err := json.Marshal(response)
		if err != nil {
			panic(err)
		}

		// 输出 JSON（Alfred 会自动解析）
		fmt.Println(string(data))
	}
}

func createRequestParams(text string) map[string][]string {

	/*
		note: 将下列变量替换为需要请求的参数
		取值参考文档: https://ai.youdao.com/DOCSIRMA/html/%E8%87%AA%E7%84%B6%E8%AF%AD%E8%A8%80%E7%BF%BB%E8%AF%91/API%E6%96%87%E6%A1%A3/%E6%96%87%E6%9C%AC%E7%BF%BB%E8%AF%91%E6%9C%8D%E5%8A%A1/%E6%96%87%E6%9C%AC%E7%BF%BB%E8%AF%91%E6%9C%8D%E5%8A%A1-API%E6%96%87%E6%A1%A3.html
	*/
	//q := "dog"
	from := "auto"
	to := "auto"
	//vocabId := "您的用户词表ID"

	return map[string][]string{
		"q":    {text},
		"from": {from},
		"to":   {to},
	}
}
