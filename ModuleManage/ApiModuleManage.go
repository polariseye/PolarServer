package ModuleManage

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/polariseye/PolarServer/Common"
	"github.com/polariseye/PolarServer/Common/ErrorCode"
)

// api模块结构
type apiModuleManagerStruct struct {
	// 客户端函数列表
	clientMethodData map[string]*MethodAndInOutTypes

	// 额外对象获取函数
	extraObjGetFun func(*Common.RequestModel) []interface{}
}

const (
	// 模块名后缀
	moduleSuffix = "BLL"

	// 客户端函数前缀
	methodPrefix = "C_"

	// 分隔符
	seperator = "_"
)

var (
	// 结果模型类型
	resultModelType reflect.Type
	// Api模块管理对象
	ApiModuleManager *apiModuleManagerStruct
)

func init() {
	// 初始化结果类型
	result := Common.NewResultModel(ErrorCode.Success)
	resultModelType = reflect.TypeOf(result)
	ApiModuleManager = NewApiModuleManager()
}

// 模块初始化
// module:待初始化的模块
func (this *apiModuleManagerStruct) AddApiModule(module IModule) {
	// 获取structObject对应的反射 Type 和 Value
	reflectValue := reflect.ValueOf(module)
	reflectType := reflect.TypeOf(module)

	// 检查模块名的后缀
	if strings.HasSuffix(module.Name(), moduleSuffix) == false {
		return
	}

	// 获取structObject中返回值为ResponseObject的方法
	for i := 0; i < reflectType.NumMethod(); i++ {
		// 获得方法名称
		methodName := reflectType.Method(i).Name

		// 判断是否为导出的方法
		if strings.HasPrefix(methodName, methodPrefix) == false {
			continue
		}

		// 获得方法及其输入参数的类型列表
		method := reflectValue.MethodByName(methodName)
		inTypes, outTypes := this.resolveMethodInOutParams(method)

		// 判断输出参数数量是否正确
		if len(outTypes) != 1 {
			continue
		}

		// 判断返回值是否为resultmodel
		outType := outTypes[0]
		if outType != resultModelType {
			continue
		}

		// 添加到列表中
		this.clientMethodData[this.getFullMethodName(module.Name(), methodName)] = NewMethodAndInOutTypes(method, inTypes, outTypes)
	}
}

// 方法调用
// request:请求参数
func (this *apiModuleManagerStruct) Call(request *Common.RequestModel) (result *Common.ResultModel) {
	result = Common.NewResultModel(ErrorCode.ClientDataError)

	// 获取方法
	methodFullName := this.getFullMethodName(request.ModuleName, request.MethodName)
	targetMethod, isExist := this.clientMethodData[methodFullName]
	if isExist == false {
		result.SetError(ErrorCode.MethodNoExist, fmt.Sprintf("未找到module"))
		return
	}

	extraObj := make([]interface{}, 0, 5)
	if this.extraObjGetFun != nil {
		extraObj = this.extraObjGetFun(request)
	}

	// 组装请求参数
	requestparam := make([]interface{}, 1, len(request.Data)+1+len(extraObj))
	requestparam[0] = request
	requestparam = append(requestparam, extraObj...)
	requestparam = append(requestparam, request.Data...)

	// 请求参数转换成调用参数
	invokeParam, errMsg := targetMethod.GetCallParams(requestparam)
	if errMsg != nil {
		result.SetError(ErrorCode.DataError, errMsg.Error())

		return
	}

	// 调用函数
	var invokeResult []reflect.Value
	invokeResult = targetMethod.Call(invokeParam)

	// 获取返回结果
	methodResult := invokeResult[0].Interface()

	return methodResult.(*Common.ResultModel)
}

// 设置获取额外对象的函数
func (this *apiModuleManagerStruct) SetExtraObjGetFun(_extraObjGetFun func(*Common.RequestModel) []interface{}) {
	this.extraObjGetFun = _extraObjGetFun
}

// 解析方法的输入输出参数
// method：方法对应的反射值
// 返回值：
// 输入参数类型集合
// 输出参数类型集合
func (this *apiModuleManagerStruct) resolveMethodInOutParams(method reflect.Value) (inTypes []reflect.Type, outTypes []reflect.Type) {
	methodType := method.Type()
	for i := 0; i < methodType.NumIn(); i++ {
		inTypes = append(inTypes, methodType.In(i))
	}

	for i := 0; i < methodType.NumOut(); i++ {
		outTypes = append(outTypes, methodType.Out(i))
	}

	return
}

// 获取结构体类型的名称
// structType：结构体类型
// 返回值：
// 结构体类型的名称
func (this *apiModuleManagerStruct) getStructName(structType reflect.Type) string {
	reflectTypeStr := structType.String()
	reflectTypeArr := strings.Split(reflectTypeStr, ".")

	return reflectTypeArr[len(reflectTypeArr)-1]
}

// 获取完整的模块名称
// moduleName：模块名称
// 返回值：
// 完整的模块名称
func (this *apiModuleManagerStruct) getFullModuleName(moduleName string) string {
	return moduleName + moduleSuffix
}

// 获取完整的方法名称
// structName：结构体名称
// methodName：方法名称
// 返回值：
// 完整的方法名称
func (this *apiModuleManagerStruct) getFullMethodName(structName, methodName string) string {
	return structName + seperator + methodName
}

// 创建新的api模块管理对象
func NewApiModuleManager() *apiModuleManagerStruct {
	return &apiModuleManagerStruct{
		clientMethodData: make(map[string]*MethodAndInOutTypes, 0),
		extraObjGetFun:   nil,
	}
}
