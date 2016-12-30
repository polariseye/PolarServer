package ModuleManage

import (
	"reflect"
	"strings"

	_ "github.com/Jordanzuo/goutil/debugUtil"
)

// api模块结构
type apiModuleStruct struct {
	// 目标模块
	targetModule IModule

	// 客户端函数列表
	clientMethodData map[string]reflect.Method
}

const (
	// 模块名后缀
	moduleSuffix = "BLL"

	// 客户端函数前缀
	methodPrefix = "C_"
)

func initApiModule(module IModule) {
	tp := reflect.TypeOf(module)

	methodNum := tp.NumMethod()
	if methodNum <= 0 {
		return
	}

	// 提取所有C函数
	for i := 0; i < methodNum; i++ {
		methodItem := tp.Method(i)

		// 必须以指定函数命名规则结尾
		if strings.HasPrefix(methodItem.Name, methodPrefix) == false {
			continue
		}

		// 检查返回值类型

	}
}
