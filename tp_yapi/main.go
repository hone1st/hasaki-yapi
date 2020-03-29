package main

import (
	"encoding/json"
	"fmt"
	"hasaki-yapi/tp_yapi/tp"
	"hasaki-yapi/yapi"
	"log"
)

func main() {
	tpY := tp.InitTpYapi()
	tpApi := tpY.Scan()

	yap := yapi.Yinstace("22f43fc34baf856853e5f1c1618bf6df2e8511865b9b45ca616552cd6ded1cd7", "http://api.ouxuan.net", 229)
	yResult, err := yapi.YapiResultSetInst(yap)
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
			catId, err = yap.InterfaceAddCate(transforMap(&yapi.Cate{
				Name: k,
				Desc: k,
			}))
			if err != nil {
				fmt.Printf("添加分类：【%s】失败！\r\n", k)
				continue
			}
			fmt.Printf("添加分类：【%s】成功！\r\n", k)
		}
		for kk, vv := range v {
			fmt.Println("------------------------------------")
			vv.CatId = catId
			_, err = yap.InterfaceSave(transforMap(vv))
			if err != nil {
				fmt.Printf("添加接口：【%s】失败！原因：【%s】\r\n", kk, err.Error())
			} else {
				fmt.Printf("添加接口：【%s】成功！\r\n", kk)
			}
			fmt.Println("------------------------------------")
		}
		fmt.Println("=====================================")
	}

}

func transforMap(v interface{}) map[string]interface{} {
	b, _ := json.Marshal(v)
	mp := make(map[string]interface{})
	_ = json.Unmarshal(b, &mp)
	return mp
}
