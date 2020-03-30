

#### yapi-go 对开放的api进行封装

##### yapi.go
```golang
// 封装yapi的内网部署的开放接口
```

##### api.gp
```golang
// 用于解析内网部署的开放接口返回的数据
```

##### yapi_reulst.go
```golang
// 用于服务器上的整个接口项目的所有数据

type YapiResultSet struct {
	CateSet        map[string]*Cate
	ApiSet         map[string]*Api               // path
	ReqQuerySet    map[string]map[string]*Field  // path:rq:field
	ReqBodyFormSet map[string]map[string]*Field  // path:rbf:field
	ReqHeadersSet  map[string]map[string]*header // path:field
}

// CateSet 分组的名字作为唯一存储分组的信息
// ApiSet  接口路径作为唯一存储接口的信息
// ReqQuerySet  接口路径+(:rq:)+字段名字作为唯一存储字段的请求的信息
// ReqBodyFormSet  接口路径+(:rbf:)+字段名字作为唯一存储字段的请求的信息
// ReqHeadersSet  接口路径+(:)+请求头参数名字作为唯一存储请求头的的请求的信息

```


#### tp-yapi

##### tp-yapi.go
```yaml
tp:
  scandir: F:\OneDrive\huishi-shop\application
  pathsuffix: .json
  filesuffix: .class.php
  controller: controller
yapi:
  token: 5af10be82269b7f3dfae11ad76f086fd0909223962ebbda6101
  projectid: 23
  host: http://api.example.net
```
```golang
// 根据配置文件来扫描scandir目录
// 规则是tp框架使用自动路由
// pathsuffix 输出的接口路径后缀带上.json
// filesuffix 扫描的控制器文件的后缀
// controller 存放控制器的目录用于定点扫描

// token/projectid/host必须的配置
// 被扫描的文件 类注释 必须带@api  不存在则忽略掉该文件 @api作为接口的分组名字
/**
 * Class Activity
 * @package app\admino\controller
 * @api 后台活动模块
 */

// 被扫描的接口 注释 与平常注释一样 需要在首行注释中文 @param用于读取字段的描述
/**
 * 活动列表带筛选
 * @param int $page
 * @param int $limit
 * @param string $state 未开始 已结束 进行中
 * @throws DataNotFoundException
 * @throws ModelNotFoundException
 * @throws DbException
 */
// 接口参数 用于读取参数和参数的默认值
//public function activityList($page = 1, $limit = 10, $state = "")
```