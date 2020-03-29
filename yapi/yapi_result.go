package yapi

import (
	"encoding/json"
	"sync"
)
type YapiResultSet struct {
	CateSet        map[string]*Cate
	ApiSet         map[string]*Api               // path
	ReqQuerySet    map[string]map[string]*Field  // path:rq:field
	ReqBodyFormSet map[string]map[string]*Field  // path:rbf:field
	ReqHeadersSet  map[string]map[string]*header // path:field
}

var (
	yapiResultSet *YapiResultSet
	yOnce         sync.Once
)
// 获取服务器上的所有分类和接口
func YapiResultSetInst(yapi *Yapi) (*YapiResultSet, error) {
	var err error
	yOnce.Do(func() {
		yapiResultSet = &YapiResultSet{
			CateSet:        make(map[string]*Cate),
			ApiSet:         make(map[string]*Api),
			ReqQuerySet:    make(map[string]map[string]*Field),
			ReqBodyFormSet: make(map[string]map[string]*Field),
			ReqHeadersSet:  make(map[string]map[string]*header),
		}
		resp, err := yapi.InterfaceGetCateMenu()
		if err != nil {
			return
		}
		bytes, _ := json.Marshal(resp)
		cateMp := make([]*Cate, 0)
		_ = json.Unmarshal(bytes, &cateMp)
		for _, v := range cateMp {
			if _, ex := yapiResultSet.CateSet[v.Name]; !ex {
				yapiResultSet.CateSet[v.Name] = v
			}
		}
		resp, err = yapi.InterfaceList(1, 1000)
		if err != nil {
			return
		}
		apiMp := resp.(map[string]interface{})["list"].([]interface{})
		for _, v := range apiMp {
			apiId := int(v.(map[string]interface{})["_id"].(float64))
			resp, err = yapi.InterfaceGet(apiId)
			if err != nil {
				return
			}
			api := &Api{}
			bytes, _ := json.Marshal(resp)
			_ = json.Unmarshal(bytes, api)
			if _, ex := yapiResultSet.ApiSet[api.Path]; !ex {
				yapiResultSet.ApiSet[api.Path] = api
			}
			api = yapiResultSet.ApiSet[api.Path]
			// 处理ReqformBody
			for _, vv := range api.ReqBodyForm {
				if _, ex := yapiResultSet.ReqBodyFormSet[api.Path]; !ex {
					yapiResultSet.ReqBodyFormSet[api.Path] = map[string]*Field{vv.Name: vv}
				}
				if _, ex := yapiResultSet.ReqBodyFormSet[api.Path][vv.Name]; !ex {
					yapiResultSet.ReqBodyFormSet[api.Path][vv.Name] = vv
				}
			}

			// 处理ReqQuery
			for _, vv := range api.ReqQuery {
				if _, ex := yapiResultSet.ReqQuerySet[api.Path]; !ex {
					yapiResultSet.ReqQuerySet[api.Path] = map[string]*Field{vv.Name: vv}
				}
				if _, ex := yapiResultSet.ReqQuerySet[api.Path][vv.Name]; !ex {
					yapiResultSet.ReqQuerySet[api.Path][vv.Name] = vv
				}
			}

			// 处理ReqHeaders
			for _, vv := range api.ReqHeaders {
				if _, ex := yapiResultSet.ReqHeadersSet[api.Path]; !ex {
					yapiResultSet.ReqHeadersSet[api.Path] = map[string]*header{vv.Name: vv}
				}
				if _, ex := yapiResultSet.ReqHeadersSet[api.Path][vv.Name]; !ex {
					yapiResultSet.ReqHeadersSet[api.Path][vv.Name] = vv
				}
			}
		}
	})

	return yapiResultSet, err
}

// 过滤只替换一些一些字段
func (y *YapiResultSet) FilterOnly(api *Api) *Api{
	// 如果结果集不存在直接返回
	if _, ex := y.ApiSet[api.Path]; !ex {
		return api
	}
	// 处理ReqformBody
	if v, ex := yapiResultSet.ReqBodyFormSet[api.Path]; ex {
		keys := make(map[string]struct{})
		// 遍历最新的
		for _, vv := range api.ReqBodyForm {
			keys[vv.Name] = struct{}{}
		}
		for key, vv := range v {
			// 如果结果集的不在需要更新的map中
			if _, exx := keys[key]; !exx {
				api.ReqBodyForm = append(api.ReqBodyForm, vv)
			}
		}

	}

	// 处理ReqQuery
	if v, ex := yapiResultSet.ReqQuerySet[api.Path]; ex {
		keys := make(map[string]struct{})
		// 遍历最新的
		for _, vv := range api.ReqQuery {
			keys[vv.Name] = struct{}{}
		}
		for key, vv := range v {
			// 如果结果集的不在需要更新的map中
			if _, exx := keys[key]; !exx {
				api.ReqQuery = append(api.ReqQuery, vv)
			}
		}

	}

	// 处理ReqHeaders
	if v, ex := yapiResultSet.ReqHeadersSet[api.Path]; ex {
		keys := make(map[string]struct{})
		// 遍历最新的
		for _, vv := range api.ReqHeaders {
			keys[vv.Name] = struct{}{}
		}
		for key, vv := range v {
			// 如果结果集的不在需要更新的map中
			if _, exx := keys[key]; !exx {
				api.ReqHeaders = append(api.ReqHeaders, vv)
			}
		}

	}
	return api
}
