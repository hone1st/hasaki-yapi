package yapi

import (
	"errors"
	"fmt"
	"hasaki-yapi/yapi/util"
	"reflect"
	"sync"
)

type YapiUrl string

const (
	ProjectGet YapiUrl = "/api/project/get" // 获取项目基本信息

	InterfaceAddCate     YapiUrl = "/api/interface/add_cat"    // 新增分组
	InterfaceGetCateMenu YapiUrl = "/api/interface/getCatMenu" // 获取菜单列表

	ServerImportData YapiUrl = " /api/open/import_data" // 服务端数据导入

	InterfaceListCat YapiUrl = "/api/interface/list_cat" // 获取某个分类下接口列表
	InterfaceGet     YapiUrl = "/api/interface/get"      // 获取接口数据（有详细接口数据定义文档）
	InterfaceUp      YapiUrl = "/api/interface/up"       // 如果更新
	InterfaceAdd     YapiUrl = "/api/interface/add"      // 如果新增
	InterfaceSave    YapiUrl = "/api/interface/save"     // 新增或者更新接口

	InterfaceList     YapiUrl = "/api/interface/list"      // 获取接口列表数据
	InterfaceListMenu YapiUrl = "/api/interface/list_menu" // 获取接口菜单列表

)

type YapiResult struct {
	Errcode int
	Errmsg  string
	Data    interface{}
}

var (
	yapi    *Yapi
	once    sync.Once
	yResult = &YapiResult{}
)

type Yapi struct {
	Token     string
	ProjectId int
	Host      string
}

// 获取基础的配置信息
func Yinstace(token, host string, projectId int) *Yapi {
	once.Do(func() {
		yapi = &Yapi{
			token, projectId, host,
		}
	})
	return yapi
}

// 查看项目信息
func (y *Yapi) ProjectGet() (interface{}, error) {
	req := util.Request{
		Url: fmt.Sprintf("%s%s?token=%s", y.Host, ProjectGet, y.Token),
	}
	resp, err := req.Get()
	if err != nil {
		return nil, err
	}
	err = resp.UnmarshalBody(yResult)
	if err != nil {
		return nil, err
	}
	if yResult.Errcode != 0 {
		return nil, errors.New(yResult.Errmsg)
	}
	return yResult.Data, nil

}

// 获取项目的分类目录
func (y *Yapi) InterfaceGetCateMenu() (interface{}, error) {
	req := util.Request{
		Url: fmt.Sprintf("%s%s?token=%s&prokect_id=%d", y.Host, InterfaceGetCateMenu, y.Token, y.ProjectId),
	}
	resp, err := req.Get()
	if err != nil {
		return nil, err
	}
	err = resp.UnmarshalBody(yResult)
	if err != nil {
		return nil, err
	}
	if yResult.Errcode != 0 {
		return nil, errors.New(yResult.Errmsg)
	}
	return yResult.Data, nil

}

// 获取某个分类下接口列表
func (y *Yapi) InterfaceListCat(catId int, params ...int) (interface{}, error) {
	var (
		page  = 1
		limit = 1000
	)
	if len(params) > 0 {
		page = params[0]
		if len(params) >= 2 {
			limit = params[1]
		}
	}
	req := util.Request{
		Url: fmt.Sprintf("%s%s?token=%s&prokect_id=%d&page=%d&limit=%d&catid=%d", y.Host, InterfaceListCat, y.Token, y.ProjectId, page, limit, catId),
	}
	resp, err := req.Get()
	if err != nil {
		return nil, err
	}
	err = resp.UnmarshalBody(yResult)
	if err != nil {
		return nil, err
	}
	if yResult.Errcode != 0 {
		return nil, errors.New(yResult.Errmsg)
	}
	return yResult.Data, nil

}

// 获取分类的列表
func (y *Yapi) InterfaceListMenu() (interface{}, error) {
	req := util.Request{
		Url: fmt.Sprintf("%s%s?token=%s&prokect_id=%d", y.Host, InterfaceListMenu, y.Token, y.ProjectId),
	}
	resp, err := req.Get()
	if err != nil {
		return nil, err
	}
	err = resp.UnmarshalBody(yResult)
	if err != nil {
		return nil, err
	}
	if yResult.Errcode != 0 {
		return nil, errors.New(yResult.Errmsg)
	}
	return yResult.Data, nil

}

// 获取某个接口详细数据
func (y *Yapi) InterfaceGet(id int) (interface{}, error) {
	req := util.Request{
		Url: fmt.Sprintf("%s%s?token=%s&prokect_id=%d&id=%d", y.Host, InterfaceGet, y.Token, y.ProjectId, id),
	}
	resp, err := req.Get()
	if err != nil {
		return nil, err
	}
	err = resp.UnmarshalBody(yResult)
	if err != nil {
		return nil, err
	}
	if yResult.Errcode != 0 {
		return nil, errors.New(yResult.Errmsg)
	}
	return yResult.Data, nil

}

// 获取所有的接口列表
func (y *Yapi) InterfaceList(params ...int) (interface{}, error) {
	var (
		page  = 1
		limit = 1000
	)
	if len(params) > 0 {
		page = params[0]
		if len(params) >= 2 {
			limit = params[1]
		}
	}
	req := util.Request{
		Url: fmt.Sprintf("%s%s?token=%s&prokect_id=%d&page=%d&limit=%d", y.Host, InterfaceList, y.Token, y.ProjectId, page, limit),
	}
	resp, err := req.Get()
	if err != nil {
		return nil, err
	}
	err = resp.UnmarshalBody(yResult)
	if err != nil {
		return nil, err
	}
	if yResult.Errcode != 0 {
		return nil, errors.New(yResult.Errmsg)
	}
	return yResult.Data, nil

}

// 接口路径为唯一
// 保存或者添加
func (y *Yapi) InterfaceSave(query map[string]interface{}) (interface{}, error) {
	query = deleteEmpty(query)
	query["token"] = y.Token
	req := util.Request{
		Url:   fmt.Sprintf("%s%s", y.Host, InterfaceSave),
		Query: query,
	}
	resp, err := req.Post()
	if err != nil {
		return nil, err
	}
	err = resp.UnmarshalBody(yResult)
	if err != nil {
		return nil, err
	}
	if yResult.Errcode != 0 {
		return nil, errors.New(yResult.Errmsg)
	}
	return yResult.Data, nil

}

// 添加接口
func (y *Yapi) InterfaceAdd(query map[string]interface{}) (interface{}, error) {
	query = deleteEmpty(query)
	query["token"] = y.Token
	req := util.Request{
		Url:   fmt.Sprintf("%s%s", y.Host, InterfaceAdd),
		Query: query,
	}
	resp, err := req.Post()
	if err != nil {
		return nil, err
	}
	err = resp.UnmarshalBody(yResult)
	if err != nil {
		return nil, err
	}
	if yResult.Errcode != 0 {
		return nil, errors.New(yResult.Errmsg)
	}
	return yResult.Data, nil

}

// 保存接口
func (y *Yapi) InterfaceUp(query map[string]interface{}) (interface{}, error) {
	query = deleteEmpty(query)
	query["token"] = y.Token
	req := util.Request{
		Url:   fmt.Sprintf("%s%s", y.Host, InterfaceUp),
		Query: query,
	}
	resp, err := req.Post()
	if err != nil {
		return nil, err
	}
	err = resp.UnmarshalBody(yResult)
	if err != nil {
		return nil, err
	}
	if yResult.Errcode != 0 {
		return nil, errors.New(yResult.Errmsg)
	}
	return yResult.Data, nil

}

// 添加目录 重复名字的也可以添加 在添加前需要自行过滤  返回分类的id
func (y *Yapi) InterfaceAddCate(query map[string]interface{}) (int, error) {
	query = deleteEmpty(query)
	fmt.Println(query)
	query["token"] = y.Token
	query["project_id"] = y.ProjectId
	req := util.Request{
		Url:   fmt.Sprintf("%s%s", y.Host, InterfaceAddCate),
		Query: query,
	}
	resp, err := req.Post()
	if err != nil {
		return 0, err
	}
	err = resp.UnmarshalBody(yResult)
	if err != nil {
		return 0, err
	}
	if yResult.Errcode != 0 {
		return 0, errors.New(yResult.Errmsg)
	}
	return int(yResult.Data.(map[string]interface{})["_id"].(float64)), nil

}

func deleteEmpty(query map[string]interface{}) map[string]interface{} {
	for k, v := range query {
		if reflect.TypeOf(v).Kind() == reflect.String && v.(string) == "" {
			delete(query, k)
		}
		if reflect.TypeOf(v).Kind() == reflect.Float64 && v.(float64) == 0 {
			delete(query, k)
		}
		if reflect.TypeOf(v).Kind() == reflect.Slice && len(v.([]interface{})) == 0 {
			delete(query, k)
		}
	}
	return query
}
