package yapi

import (
	"fmt"
	"testing"
)

func Init() {
	Yinstace("", "", 149)

}

func TestProjectGet(t *testing.T) {
	Init()
	resp, err := yapi.ProjectGet()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)

}
func TestInterfaceGetCateMenu(t *testing.T) {
	Init()
	resp, err := yapi.InterfaceGetCateMenu()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)

}
func TestInterfaceListCat(t *testing.T) {
	Init()
	resp, err := yapi.InterfaceListCat(750)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)

}
func TestInterfaceGet(t *testing.T) {
	Init()
	resp, err := yapi.InterfaceGet(1172)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)

}
func TestInterfaceList(t *testing.T) {
	Init()
	resp, err := yapi.InterfaceList()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)

}
func TestInterfaceListMenu(t *testing.T) {
	Init()
	resp, err := yapi.InterfaceList()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)

}
