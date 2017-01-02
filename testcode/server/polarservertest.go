package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/polariseye/polarserver"
	"github.com/polariseye/polarserver/common"
	"github.com/polariseye/polarserver/moduleManage"
)

func main() {
	if errMsg := polarserver.Init("app.config"); errMsg != nil {
		fmt.Printf("配置文件初始化错误,错误信息:%v", errMsg)

		return
	}

	moduleManage.InitModule()

	if errMsg := polarserver.ServerManagerObj().Start(); errMsg != nil {
		fmt.Printf("启动所有服务失败:%v", errMsg)

		return

	}

	request := common.NewRequestModel()
	request.Ip = "127.0.0.1"
	request.Data = append(request.Data, "你好哟")
	request.ModuleName = "Test"
	request.MethodName = "Hello"

	moduleManage.DefaulApiModuleManager.SetExtraObjGetFun(func(request *common.RequestModel) []interface{} {
		tmpResult := make([]interface{}, 0, 1)
		tmpResult = append(tmpResult, 2)

		return tmpResult
	})

	result := moduleManage.DefaulApiModuleManager.Call(request)
	fmt.Println(result.Value["Hello"])
	fmt.Println("Extra:", result.Value["Extra"])
	fmt.Println(result)

	polarserver.ServerManagerObj().WaitStop()
}
