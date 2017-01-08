package polarserver

import (
	"fmt"

	"github.com/Jordanzuo/goutil/configUtil"
	"github.com/Jordanzuo/goutil/debugUtil"
	"github.com/Jordanzuo/goutil/logUtil"
	"github.com/Jordanzuo/goutil/xmlUtil"
	"github.com/polariseye/polarserver/common"
	"github.com/polariseye/polarserver/dataBase"
	"github.com/polariseye/polarserver/moduleManage"
	"github.com/polariseye/polarserver/server"
	"github.com/polariseye/polarserver/server/webServer"
	"github.com/polariseye/polarserver/server/webServer/apiHandle"
)

var (
	// 服务管理对象
	serverManagerObj *server.ServerManagerStruct = server.NewServerManager()

	// web服务对象
	webServerObj *webServer.WebServerStruct

	// 配置对象
	configObj *configUtil.XmlConfig

	// 处理对象
	clientHandler *apiHandle.Handle4JsonStruct

	// 客户端Api接口
	clientApiManager *moduleManage.ApiModuleManagerStruct = moduleManage.NewApiModuleManager()

	// 处理对象
	manageHandler *apiHandle.Handle4JsonStruct

	// 后端Api接口
	mangageApiManager *moduleManage.ApiModuleManagerStruct = moduleManage.NewApiModuleManager()
)

func init() {
	mangageApiManager.SetMethodPrefix("M_")
	mangageApiManager.SetTestMethodPrefix("MTest_")

	moduleManage.AddApiManager(clientApiManager)
	moduleManage.AddApiManager(mangageApiManager)
}

// 初始化
// configFileName:配置文件名
// errMsg:错误信息
func Init(configFileName string) (errMsg error) {
	root, errMsg := xmlUtil.LoadFromFile(configFileName)
	if errMsg != nil {
		return
	}

	configObj = configUtil.NewXmlConfig(root)

	// 初始化日志记录
	logUtil.SetLogPath(configObj.DefaultString("Config/LogPath", "", "DefaultLogPath/"))

	// 初始化是否是测试模式
	common.SetIsTest(configObj.DefaultBool("Config/IsTest", "", false))

	// 初始化是否是调试模式
	debugUtil.SetDebug(configObj.DefaultBool("Config/IsDebug", "", false))

	// 数据库配置初始化
	errMsg = initDataBaseFromConfig(configObj)
	if errMsg != nil {
		return errMsg
	}

	// web服务配置初始化
	initWebServerFromConfig(configObj)

	return nil
}

// web服务对象
func WebServerObj() *webServer.WebServerStruct {
	return webServerObj
}

// 服务管理对象
func ServerManagerObj() *server.ServerManagerStruct {
	return serverManagerObj
}

// 配置对象
func ConfigObj() *configUtil.XmlConfig {
	return configObj
}

// 初始化web服务
func initWebServerFromConfig(config *configUtil.XmlConfig) {
	// web服务初始化
	port := config.DefaultInt("Config/WebPort", "", 2017)
	webServerObj = webServer.NewWebServer(int32(port), "web 服务")

	// 初始化客户端Api处理
	clientHandler := apiHandle.NewHandle4Json(clientApiManager)
	webServerObj.AddRouter("/Api/Client", clientHandler.RequestHandle)

	// 初始化后台api处理
	manageHandler := apiHandle.NewHandle4Json(mangageApiManager)
	webServerObj.AddRouter("/Api/Manage", manageHandler.RequestHandle)

	// 注册模块
	ServerManagerObj().Register(webServerObj)
}

// 从配置文件初始化数据库
func initDataBaseFromConfig(config *configUtil.XmlConfig) error {
	dbConnectionNode := config.Node("Config/DbConnection")
	if dbConnectionNode == nil {
		return fmt.Errorf("未找到配置节点:Config/DbConnection")
	}

	connectionData := dbConnectionNode.Children()
	for _, connectionInfo := range connectionData {
		var driver, connectionString string

		driver = connectionInfo.SelectAttr("Driver")
		if driver == "" {
			return fmt.Errorf("数据库配置不合法,不存在节点或其空：Config/DbConnection/%v的属性:Driver", connectionInfo.NodeName)
		}

		connectionString = connectionInfo.InnerText()
		if connectionString == "" {
			return fmt.Errorf("数据库配置不合法,不存在节点：Config/DbConnection/%v不存在连接配置文本", connectionInfo.NodeName)
		}

		errMsg := dataBase.AddConnection(connectionInfo.NodeName, driver, connectionString)
		if errMsg != nil {
			return errMsg
		}
	}

	return nil
}

// 客户端管理对象
func ClientApiManager() *moduleManage.ApiModuleManagerStruct {
	return clientApiManager
}

// 后端管理对象
func MangageApiManager() *moduleManage.ApiModuleManagerStruct {
	return clientApiManager
}
