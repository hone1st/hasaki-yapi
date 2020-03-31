package yapi_interface

type YapiInterface interface {
	ProjectGet() (interface{}, error)
	InterfaceGetCateMenu() (interface{}, error)
	InterfaceListCat(catId int, params ...int) (interface{}, error)
	InterfaceListMenu() (interface{}, error)
	InterfaceGet(id int) (interface{}, error)
	InterfaceList(params ...int) (interface{}, error)
	InterfaceSave(query map[string]interface{}) (interface{}, error)
	InterfaceAdd(query map[string]interface{}) (interface{}, error)
	InterfaceUp(query map[string]interface{}) (interface{}, error)
	InterfaceAddCate(query map[string]interface{}) (int, error)
}
