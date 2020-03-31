package main

import (
	"encoding/json"
	"fmt"
	"github.com/liangsssttt/hasaki-yapi/tp_yapi/tp"
	"github.com/liangsssttt/hasaki-yapi/yapi"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func main() {
	c := tp.InitTpYapi()
	tpApi := c.Tp.Scan()
	yResult, err := yapi.YapiResultSetInst(c.Yapi)
	if err != nil {
		log.Fatal(err)
	}
	for k, v := range tpApi.TpApi {
		fmt.Println("=====================================")
		var catId = 0
		if cate, ex := yResult.CateSet[k]; ex {
			catId = cate.Id
			fmt.Printf("分类已存在：【%s】！\r\n", k)
		} else {
			catId, err = c.Yapi.InterfaceAddCate(transforMap(&yapi.Cate{
				Name: k,
				Desc: k,
			}))
			if err != nil {
				fmt.Printf("添加分类：【%s】失败！\r\n", k)
				continue
			}
			fmt.Printf("添加分类：【%s】成功！\r\n", k)
		}

		if _, ex := tpApi.SuccessTpApi[k]; !ex {
			tpApi.SuccessTpApi[k] = make(map[string]*yapi.Api)
		}
		if _, ex := tpApi.FailureTpApi[k]; !ex {
			tpApi.FailureTpApi[k] = make(map[string]*yapi.Api)
		}

		for kk, vv := range v {
			fmt.Println("------------------------------------")
			vv.CatId = catId
			_, err = c.Yapi.InterfaceSave(transforMap(vv))
			if err != nil {
				tpApi.FailureTpApi[k][kk] = vv
				fmt.Printf("添加接口失败！\r\n路径：【%s】\r\n名字：【%s】\r\n原因：【%s】\r\n", kk, vv.Title, err.Error())
			} else {
				tpApi.SuccessTpApi[k][kk] = vv
				fmt.Printf("添加接口成功！\r\n路径：【%s】\r\n名字：【%s】\r\n", kk, vv.Title)
			}
			fmt.Println("------------------------------------")
		}
		fmt.Println("=====================================")
	}
	t := time.Now().Format("2006_01_02_15_04_05")
	os.MkdirAll("./log/"+t, 0777)
	success, _ := json.MarshalIndent(tpApi.SuccessTpApi, " ", " ")
	_ = ioutil.WriteFile("./log/"+t+"/success_.json", success, 0777)
	failure, _ := json.MarshalIndent(tpApi.FailureTpApi, " ", " ")
	_ = ioutil.WriteFile("./log/"+t+"/failure_.json", failure, 0777)

	fmt.Println("按下任意键结束：")
	var i int
	fmt.Scanf("按下任意键后enter结束：%i", i)
}

func transforMap(v interface{}) map[string]interface{} {
	b, _ := json.Marshal(v)
	mp := make(map[string]interface{})
	_ = json.Unmarshal(b, &mp)
	return mp
}
