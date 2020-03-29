package yapi

type Cate struct {
	Id   int    `json:"_id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}

// 状态类型
type Status string

const (
	Done   Status = "done"
	UnDone Status = "undone"
)

// 响应体类型
type ResBodyType string

const (
	ResJson ResBodyType = "json"
	ResRaw  ResBodyType = "raw"
)

// 请求方法类型
type Method string

const (
	Get     Method = "GET"
	Post    Method = "POST"
	Put     Method = "PUT"
	Delete  Method = "DELETE"
	Head    Method = "HEAD"
	Options Method = "OPTIONS"
	Patch   Method = "PATCH"
)

// 请求体类型
type ReqBodyType string

const (
	ReqForm ReqBodyType = "form"
	ReqFile ReqBodyType = "file"
	ReqRaw  ReqBodyType = "raw"
	ReqJson ReqBodyType = "json"
)

// 接口
type Api struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`  // 接口中文描述
	Path   string `json:"path"`   // 接口路径
	CatId  int    `json:"catid"`  // 所属的分类的id
	Status Status `json:"status"` // 接口状态

	Method Method `json:"method"` // 请求方法

	ReqBodyIsJsonSchema bool `json:"req_body_is_json_schema"`
	ResBodyIsJsonSchema bool `json:"res_body_is_json_schema"`

	Desc     string `json:"desc"`     // 详细信息
	Markdown string `json:"markdown"` // 备注

	ReqBodyOther string      `json:"req_body_other"` // 请求类型非form类型时候展示。最好是json字符串
	ReqBodyType  ReqBodyType `json:"req_body_type"`  // 请求类型
	ReqQuery     []*Field    `json:"req_query"`      // 请求体
	ReqBodyForm  []*Field    `json:"req_body_form"`  // 请求体
	ReqHeaders   []*header   `json:"req_headers"`    // 请求头

	ResBody     string      `json:"res_body"`      // 返回实例 最好是json字符串
	ResBodyType ResBodyType `json:"res_body_type"` // 返回值类型
}

// 参数的类型
type Type string

const (
	Text Type = "text"
	File Type = "file"
)

type Field struct {
	Required int    `json:"required"`
	Name     string `json:"name"`
	Type     Type   `json:"type"`
	Example  string `json:"example"`
	Desc     string `json:"desc"`
}

// 请求头的参数类型
type HeaderType string

const (
	Json               HeaderType = "application/json"
	XWwwFormUrlencoded HeaderType = "application/x-www-form-urlencoded"

	FormData HeaderType = "multipart/form-data"
)

// 请求头
type header struct {
	Name  string     `json:"name"`
	Value HeaderType `json:"value"`
}

func GetHeader(t HeaderType) *header {
	return &header{
		Name:  "Content-Type",
		Value: t,
	}
}
