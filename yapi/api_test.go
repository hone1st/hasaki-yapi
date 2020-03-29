package yapi

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestApi(t *testing.T) {
	Init()
	desc := `/**
 * 获取账户余额
 * @param int $type 账户标识 1网站2红包3打赏4霸屏
 */`

	api := &Api{
		//Id:                  5839,
		Title:               "测试api",
		Path:                "/Api/WalletAccount/test.json",
		CatId:               750,
		Status:              Done,
		Method:              Post,
		ReqBodyIsJsonSchema: false,
		ResBodyIsJsonSchema: false,
		Desc:                desc,
		Markdown:            desc,
		ReqBodyOther:        "{a:1}",
		ReqBodyType:         ReqForm,
		ReqQuery:            make([]*Field, 0),
		ReqBodyForm: []*Field{
			&Field{
				Required: 1,
				Name:     "name",
				Type:     Text,
				Example:  "李四",
				Desc:     "名字",
			},
			&Field{
				Required: 1,
				Name:     "file",
				Type:     File,
				Example:  "",
				Desc:     "上传文件",
			},
		},
		ReqHeaders: []*header{
			GetHeader(XWwwFormUrlencoded),
		},
		ResBody:     `{
  "catid": 2334,
  "desc": "/** * 团队设置添加大区或者修改 * /admino/SaleGoods/aOsBigSaleA.json * @param string $name 大区名字 * @param string $sale_dest 销售指标 * @param string $sale_actual 实际销售 * @param int $id 默认为0则添加 */",
  "id": 4933,
  "markdown": "/** * 团队设置添加大区或者修改 * /admino/SaleGoods/aOsBigSaleA.json * @param string $name 大区名字 * @param string $sale_dest 销售指标 * @param string $sale_actual 实际销售 * @param int $id 默认为0则添加 */",
  "method": "",
  "path": "/admino/SaleGoods/aOsBigSalexxx.json",
  "req_body_form": [
   {
    "desc": "名字",
    "example": "",
    "name": "name",
    "required": 1,
    "type": "text"
   },
   {
    "desc": "年龄",
    "example": "",
    "name": "age",
    "required": 1,
    "type": "text"
   },
   {
    "desc": "性别",
    "example": "",
    "name": "sex",
    "required": 1,
    "type": "text"
   }
  ],
  "req_query": null,
  "res_body": "",
  "res_body_type": "",
  "status": "done",
  "title": "测试接口",
  "token": "01ea39a61a7b4d833bbad2bbe1ae41c54880b614b02dfcd80102a1f80063ebdd"
 }`,
		ResBodyType: ResJson,
	}
	b, _ := json.Marshal(api)
	mp := make(map[string]interface{})
	_ = json.Unmarshal(b, &mp)


	//d, err := yapi.InterfaceAdd(mp)
	d, err := yapi.InterfaceSave(mp)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(d)
}

func TestInterfaceAddCate(t *testing.T) {
	Init()

	cate := &Cate{
		Id:   0,
		Name: "测试目录",
		Desc: "测试目录",
	}
	b, _ := json.Marshal(cate)
	mp := make(map[string]interface{})
	_ = json.Unmarshal(b, &mp)
	d, err := yapi.InterfaceAddCate(mp)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(d)
}
