package ModuleManage

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/Jordanzuo/goutil/logUtil"
)

var (
	// 供客户端访问的模块的后缀
	mModuleSuffix = "Module"

	// 定义用于分隔模块名称和方法名称的分隔符
	mDelimeterOfObjAndMethod = "_"

	// 定义存放所有方法映射的变量
	mMethodMap = make(map[string]*MethodAndInOutTypes)
)

// 获取结构体类型的名称
// structType：结构体类型
// 返回值：
// 结构体类型的名称
func getStructName(structType reflect.Type) string {
	reflectTypeStr := structType.String()
	reflectTypeArr := strings.Split(reflectTypeStr, ".")

	return reflectTypeArr[len(reflectTypeArr)-1]
}

// 获取完整的模块名称
// moduleName：模块名称
// 返回值：
// 完整的模块名称
func getFullModuleName(moduleName string) string {
	return moduleName + mModuleSuffix
}

// 获取完整的方法名称
// structName：结构体名称
// methodName：方法名称
// 返回值：
// 完整的方法名称
func getFullMethodName(structName, methodName string) string {
	return structName + mDelimeterOfObjAndMethod + methodName
}

// 解析方法的输入输出参数
// method：方法对应的反射值
// 返回值：
// 输入参数类型集合
// 输出参数类型集合
func resolveMethodInOutParams(method reflect.Value) (inTypes []reflect.Type, outTypes []reflect.Type) {
	methodType := method.Type()
	for i := 0; i < methodType.NumIn(); i++ {
		inTypes = append(inTypes, methodType.In(i))
	}

	for i := 0; i < methodType.NumOut(); i++ {
		outTypes = append(outTypes, methodType.Out(i))
	}

	return
}

// 将需要对客户端提供方法的对象进行注册
// structObject：对象
func RegisterFunction(structObject interface{}) {
	// 获取structObject对应的反射 Type 和 Value
	reflectValue := reflect.ValueOf(structObject)
	reflectType := reflect.TypeOf(structObject)

	// 提取对象类型名称
	structName := getStructName(reflectType)

	// 获取structObject中返回值为ResponseObject的方法
	for i := 0; i < reflectType.NumMethod(); i++ {
		// 获得方法名称
		methodName := reflectType.Method(i).Name

		// 判断是否为导出的方法

		// 获得方法及其输入参数的类型列表
		method := reflectValue.MethodByName(methodName)
		inTypes, outTypes := resolveMethodInOutParams(method)

		// 判断输出参数数量是否正确
		if len(outTypes) != 1 {
			continue
		}

		// 判断返回值是否为ResponseObject
		outType := outTypes[0]
		if _, ok := outType.FieldByName("Code"); !ok {
			continue
		}
		if _, ok := outType.FieldByName("Message"); !ok {
			continue
		}
		if _, ok := outType.FieldByName("Data"); !ok {
			continue
		}

		// 添加到列表中
		mMethodMap[getFullMethodName(structName, methodName)] = NewMethodAndInOutTypes(method, inTypes, outTypes)
	}
}

// 调用方法
// clientObj：客户端对象
// requestObj：请求对象
func callFunction(clientObj *Client, requestObj *RequestObject) {
	responseObj := GetInitResponseObj()

	var methodAndInOutTypes *MethodAndInOutTypes
	var ok bool

	// 根据传入的ModuleName和MethodName找到对应的方法对象
	key := getFullMethodName(getFullModuleName(requestObj.ModuleName), requestObj.MethodName)
	if methodAndInOutTypes, ok = mMethodMap[key]; !ok {
		logUtil.Log(fmt.Sprintf("找不到指定的方法：%s", key), logUtil.Error, true)
		ResponseResult(clientObj, requestObj, responseObj.SetResultStatus(NoTargetMethod))
		return
	}

	// 判断参数数量是否相同
	inTypesLength := len(methodAndInOutTypes.inTypes)
	paramLength := len(requestObj.Parameters)
	if paramLength != inTypesLength {
		logUtil.Log(fmt.Sprintf("传入的参数数量不符，本地方法%s的参数数量：%d，传入的参数数量为：%d", key, inTypesLength, paramLength), logUtil.Error, true)
		ResponseResult(clientObj, requestObj, responseObj.SetResultStatus(ParamNotMatch))
		return
	}

	// 构造参数
	in := methodAndInOutTypes.GetCallParams(requestObj.Parameters)

	// 判断是否有无效的参数（传入的参数类型和方法定义的类型不匹配导致没有赋值）
	for _, item := range in {
		if reflect.Value.IsValid(item) == false {
			logUtil.Log(fmt.Sprintf("方法%s传入的参数%v无效", key, requestObj.Parameters), logUtil.Error, true)
			ResponseResult(clientObj, requestObj, responseObj.SetResultStatus(ParamInValid))
			return
		}
	}

	// 传入参数，调用方法
	out := methodAndInOutTypes.method.Call(in)

	// 并输出结果到客户端（由于只有一个返回值，所以取out[0]）
	ResponseResult(clientObj, requestObj, (&out[0]).Interface())
}
