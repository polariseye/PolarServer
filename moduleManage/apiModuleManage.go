package moduleManage

import (
	"fmt"
	"reflect"
	"strings"
	"sync"

	"github.com/Jordanzuo/goutil/stringUtil"
	"github.com/polariseye/polarserver/common"
	"github.com/polariseye/polarserver/common/errorCode"
)

// api模块结构
type ApiModuleManagerStruct struct {
	// 客户端函数列表
	clientMethodData map[string]*MethodAndInOutTypes

	// 客户端函数同步锁
	methodDataLocker *sync.RWMutex

	// 额外对象获取函数
	extraObjGetFun func(*common.RequestModel) []interface{}

	// 模块名后缀
	moduleSuffix string

	// 客户端函数前缀
	methodPrefix string

	// 分隔符
	seperator string

	// 测试函数的后缀
	testMethodPrefix string
}

var (
	// 结果模型类型
	resultModelType reflect.Type
)

// 初始化
func init() {
	// 初始化结果类型
	result := common.NewResultModel(errorCode.Success)
	resultModelType = reflect.TypeOf(result)
}

// 添加模块(可重复调用)
func (this *ApiModuleManagerStruct) AddApiModule(module IModule) {
	// 获取structObject对应的反射 Type 和 Value
	reflectValue := reflect.ValueOf(module)
	reflectType := reflect.TypeOf(module)

	// 检查模块名的后缀
	if strings.HasSuffix(module.Name(), this.moduleSuffix) == false {
		return
	}

	// 获取structObject中返回值为ResponseObject的方法
	for i := 0; i < reflectType.NumMethod(); i++ {
		// 获得方法名称
		methodName := reflectType.Method(i).Name

		// 判断是否为导出的方法
		if strings.HasPrefix(methodName, this.methodPrefix) == false &&
			common.IsTest() &&
			strings.HasPrefix(methodName, this.testMethodPrefix) == false {

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
		methodName = strings.TrimLeft(methodName, this.methodPrefix)
		if common.IsTest() {
			methodName = strings.TrimLeft(methodName, this.testMethodPrefix)
		}

		tmpModuleName := strings.TrimRight(module.Name(), this.moduleSuffix)

		// 更新缓存
		func() {
			this.methodDataLocker.Lock()
			defer this.methodDataLocker.Unlock()

			this.clientMethodData[this.getFullMethodName(tmpModuleName, methodName)] = NewMethodAndInOutTypes(method, inTypes, outTypes)
		}()
	}
}

// 方法调用
// request:请求参数
// 返回值:
// result:结果对象
func (this *ApiModuleManagerStruct) Call(request *common.RequestModel) (result *common.ResultModel) {
	result = common.NewResultModel(errorCode.ClientDataError)

	// 获取方法
	targetMethod, isExist := this.getCallMethod(request.ModuleName, request.MethodName)
	if isExist == false {
		result.SetError(errorCode.MethodNoExist, fmt.Sprintf("未找到调用方法"))
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
		result.SetError(errorCode.DataError, errMsg.Error())

		return
	}

	// 调用函数
	var invokeResult []reflect.Value
	invokeResult = targetMethod.Call(invokeParam)

	// 获取返回结果
	methodResult := invokeResult[0].Interface()

	return methodResult.(*common.ResultModel)
}

// 设置获取额外对象的函数
// _extraObjGetFun:额外实体获取函数
func (this *ApiModuleManagerStruct) SetExtraObjGetFun(_extraObjGetFun func(*common.RequestModel) []interface{}) {
	this.extraObjGetFun = _extraObjGetFun
}

// 获取调用函数
// moduleName:模块名
// methodName:方法名
// 返回值:
// *MethodAndInOutTypes:调用方法
// bool:是否存在
func (this *ApiModuleManagerStruct) getCallMethod(moduleName, methodName string) (*MethodAndInOutTypes, bool) {
	// 方法完整名
	methodFullName := this.getFullMethodName(moduleName, methodName)

	// 进行读锁
	this.methodDataLocker.RLock()
	defer this.methodDataLocker.RUnlock()

	result, isExist := this.clientMethodData[methodFullName]
	return result, isExist
}

// 解析方法的输入输出参数
// method：方法对应的反射值
// 返回值：
// 输入参数类型集合
// 输出参数类型集合
func (this *ApiModuleManagerStruct) resolveMethodInOutParams(method reflect.Value) (inTypes []reflect.Type, outTypes []reflect.Type) {
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
func (this *ApiModuleManagerStruct) getStructName(structType reflect.Type) string {
	reflectTypeStr := structType.String()
	reflectTypeArr := strings.Split(reflectTypeStr, ".")

	return reflectTypeArr[len(reflectTypeArr)-1]
}

// 获取完整的模块名称
// moduleName：模块名称
// 返回值：
// 完整的模块名称
func (this *ApiModuleManagerStruct) getFullModuleName(moduleName string) string {
	return moduleName + this.moduleSuffix
}

// 获取完整的方法名称
// moduleName：结构体名称
// methodName：方法名称
// 返回值：
// 完整的方法名称
func (this *ApiModuleManagerStruct) getFullMethodName(moduleName, methodName string) string {
	fmt.Println(moduleName + this.moduleSuffix + this.seperator + methodName)
	return moduleName + this.moduleSuffix + this.seperator + methodName
}

// 设置函数前缀名
// prefix:函数前缀
func (this *ApiModuleManagerStruct) SetMethodPrefix(prefix string) {
	if stringUtil.IsEmpty(prefix) {
		return
	}

	this.methodPrefix = prefix
}

// 设置函数前缀名
// prefix:函数前缀
func (this *ApiModuleManagerStruct) SetTestMethodPrefix(prefix string) {
	if stringUtil.IsEmpty(prefix) {
		return
	}

	this.testMethodPrefix = prefix
}

// 创建新的api模块管理对象
func NewApiModuleManager() *ApiModuleManagerStruct {
	return &ApiModuleManagerStruct{
		clientMethodData: make(map[string]*MethodAndInOutTypes, 0),
		extraObjGetFun:   nil,
		moduleSuffix:     "BLL",
		methodPrefix:     "C_",
		seperator:        "_",
		testMethodPrefix: "Test_",
		methodDataLocker: &sync.RWMutex{},
	}
}
