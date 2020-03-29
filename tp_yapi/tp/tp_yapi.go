package tp

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"hasaki-yapi/yapi"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"sync"
)

/**
 * 获取门店下的活动列表
 * @path /home/Activity/listActivity.json
 * @throws \think\db\exception\DataNotFoundException
 * @throws \think\db\exception\ModelNotFoundException
 * @throws \think\exception\DbException
 */

/**
 * @api 前台的活动的模块
 * Class Activity
 * @package app\home\controller
 */

// 仅对自动路由匹配
// 扫描规则
// 控制器类的方法声明上@api 表示属于某个分类
// 方法上的注释表示接口路径 @path
// 方法上的@param表示参数 @param
// 目录结构
/**

-- app
	-- home
		-- controller
			-- Index		/home/Index/xxx方法
*/

type TpYapi struct {
	ScanDir    string `comment:"扫描的目录"`
	PathSuffix string `comment:"路径后缀比如.json .html"`
	FileSuffix string `comment:"扫描的文件后缀 .class.php  .php"`
	Controller string `comment:"扫描的文件目录"`
}

var (
	tp   *TpYapi
	once sync.Once
)

func InitTpYapi() *TpYapi {
	once.Do(func() {
		tp = readConfig()
	})
	return tp
}

func readConfig() *TpYapi {
	args := os.Args[1:]
	var configPath string
	if len(args) >= 2 && (args[0] == "-c" || args[0] == "-C") {
		configPath = args[1]
		if filepath.Ext(configPath) != "yaml" {
			log.Fatal("配置文件必须是yaml格式！")
		}
	}
	if configPath == "" {
		executePath := os.Args[0]
		configPath = Append(false, filepath.Dir(executePath), "yapi.yaml")
	}
	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatalf("配置文件：%s不存在！", configPath)
	}
	tp = &TpYapi{}
	err = yaml.Unmarshal(yamlFile, tp)
	if err != nil {
		log.Fatalf("配置文件解析失败：%s", err)
	}
	return tp
}

type TpApiCollect struct {
	TpApi        map[string]map[string]*yapi.Api `json:"tp_api" comment:"搜集的某个分类下的api接口"`
	FailureTpApi map[string]map[string]*yapi.Api `json:"failure_tp_api" comment:"成功的"`
	SuccessTpApi map[string]map[string]*yapi.Api `json:"success_tp_api" comment:"失败的"`
}

func (t *TpYapi) Scan() {
	fmt.Println(t.ScanDir)
	_ = filepath.Walk(t.ScanDir, t.walk)
}

func (t *TpYapi) walk(filename string, fi os.FileInfo, err error) error {
	if strings.Contains(filepath.Dir(filename), t.Controller) {
		// 获取目标文件的方法
		b, _ := ioutil.ReadFile(filename)
		// 获取@api
		re, _ := regexp.Compile("@api[^\r\n]+")
		cateName := string(re.Find(b))
		if cateName == "" {
			return nil
		}
		cateName = Replace(cateName, "", "@api")

		modules := Split(strings.Replace(filename, Append(true, t.Controller, fi.Name()), "", -1))
		module := modules[len(modules)-1]
		control := Replace(fi.Name(), "", t.FileSuffix, ".php", ".class.php", "controller", "Controller")

		dealFunction(b)
		fmt.Println(cateName)
		fmt.Println(control)
		fmt.Println(module)
		os.Exit(9)
	}
	return nil
}

func dealFunction(b []byte) []*yapi.Api {
	apis := make([]*yapi.Api, 0)

	re, _ := regexp.Compile("\\s+\\/\\*\\*[^`]*?\\)")
	bs := re.FindAll(b, -1)
	for _, v := range bs {
		if Contains(string(v), true, "namespace", "class") || (!Contains(string(v), true, "public", "function")) {
			continue
		}
		reName, _ := regexp.Compile(`function\s*?([^-~\s]*?)\(`)
		name := Replace(string(reName.Find(v)), "", "function", " ", "(")
		reDesc, _ := regexp.Compile("[^/\\*\n(\\s+)]+[\\s]*?[^/*\n(\\s+)]+")
		desc := string(reDesc.Find(v))
		dealFiled(v)
		fmt.Println(name)
		fmt.Println(desc)
	}

}

// 获取字段信息
func dealFiled(b []byte) []*yapi.Field {
	fs := make([]*yapi.Field, 0)
	// 获取字段
	reFiled, _ := regexp.Compile("\\(([^`]*?)\\)")
	fieldsStr := reFiled.Find(b)
	if !Contains(string(fieldsStr), true, "$") {
		return fs
	}
	fields := strings.Split(Replace(string(fieldsStr), "", " int", " int ", " string", " string ", " array", " array ", "(", ")", "\r\n", " "), ",")
	for _, v := range fields {
		f := &yapi.Field{}
		tempS := strings.Split(v, "=")
		if len(tempS) == 1 {
			f.Name = Replace(tempS[0], "", "$")
			f.Example = ""
		} else if len(tempS) == 2 {
			f.Name = Replace(tempS[0], "", "$")
			f.Example = tempS[1]
		}
		if f.Name == "" {
			continue
		}
		f.Required = 1
		f.Type = yapi.Text
		// 获取中文注释
		reDesc, _ := regexp.Compile(fmt.Sprintf("@param[^(\r\n)]+\\$%s[^(\r\n)]+", f.Name))
		f.Desc = ""
		if string(reDesc.Find(b)) != "" {
			descs := strings.Split(string(reDesc.Find(b)), f.Name)
			if len(descs) == 2 {
				f.Desc = strings.Split(string(reDesc.Find(b)), f.Name)[1]
			}
		}
		fs = append(fs, f)
	}
	return fs
}

// b是否在前面加分隔符
func Append(b bool, params ...string) string {
	var s = "/"
	if runtime.GOOS == "windows" {
		s = "\\"
	}
	if b {
		return s + strings.Join(params, s)
	}
	return strings.Join(params, s)
}

func Split(str string) []string {
	var s = "/"
	if runtime.GOOS == "windows" {
		s = "\\"
	}
	return strings.Split(str, s)
}

func Replace(need string, new string, params ...string) string {
	if len(params) == 0 {
		return need
	}

	for _, v := range params {
		need = strings.Replace(need, v, new, -1)
	}
	return need
}

// b true的时候都要包含
// b false 包含其中一个
func Contains(need string, b bool, params ...string) bool {
	for _, v := range params {
		if (strings.Contains(need, v) && !b) || (!strings.Contains(need, v) && b) {
			return !b
		}
	}
	return b
}
