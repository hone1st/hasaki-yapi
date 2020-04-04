package yapi

import (
	"fmt"
	"testing"
)

func init() {
	Yinstace("22f43fc34baf856853e5f1c1618bf6df2e8511865b9d6ded1cd7", "http://api.exmaple.net", 19)

}

func TestProjectGet(t *testing.T) {
	resp, err := yapi.ProjectGet()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)

}
func TestInterfaceGetCateMenu(t *testing.T) {
	resp, err := yapi.InterfaceGetCateMenu()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)

}
func TestInterfaceListCat(t *testing.T) {
	resp, err := yapi.InterfaceListCat(750)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)

}
func TestInterfaceGet(t *testing.T) {
	resp, err := yapi.InterfaceGet(1172)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)

}
func TestInterfaceList(t *testing.T) {
	resp, err := yapi.InterfaceList()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)

}
func TestInterfaceListMenu(t *testing.T) {
	resp, err := yapi.InterfaceList()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)

}
