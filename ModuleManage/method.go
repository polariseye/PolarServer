package moduleManage

import (
	"fmt"
	"reflect"

	"github.com/Jordanzuo/goutil/typeUtil"
)

// 反射的方法和输入、输出参数类型组合类型
type MethodAndInOutTypes struct {
	// 方法名
	methodName string

	// 反射出来的对应方法对象
	method reflect.Value

	// 反射出来的方法的输入参数的类型集合
	inTypes []reflect.Type

	// 反射出来的方法的输出参数的类型集合
	outTypes []reflect.Type
}

// 获取输入参数列表（返回的是副本）
func (this *MethodAndInOutTypes) InTypes() []reflect.Type {
	return this.inTypes[:]
}

// 获取函数参数个数
func (this *MethodAndInOutTypes) inLen() int {
	return len(this.inTypes)
}

// 把数据列表转换成调用参数列表
// paramData:调用的参数列表
// 返回值:
// []reflect.Value:调用的参数列表
// error:返回的错误信息
func (this *MethodAndInOutTypes) GetCallParams(paramData []interface{}) ([]reflect.Value, error) {
	// 是否存在参数
	inTypesLength := len(this.inTypes)
	if inTypesLength <= 0 {
		return make([]reflect.Value, 0), nil
	}

	// 判断参数数量是否相同
	paramLength := len(paramData)
	if paramLength < inTypesLength {
		return nil, fmt.Errorf("传入的参数数量不符，本地方法%s的参数数量：%d，传入的参数数量为：%d", this.methodName, inTypesLength, paramLength)
	}

	// 构造参数
	in := make([]reflect.Value, inTypesLength)
	for i := 0; i < inTypesLength; i++ {
		inTypeItem := this.inTypes[i]
		paramItem := paramData[i]

		// 已支持类型：Common.RequestModel,Player(非基本类型)
		// 已支持类型：Bool,Int,Int8,Int16,Int32,Int64,Uint,Uint8,Uint16,Uint32,Uint64,Float32,Float64,String
		// 已支持类型：以及上面所列出类型的Slice类型
		// 未支持类型：Uintptr,Complex64,Complex128,Array,Chan,Func,Interface,Map,Ptr,Struct,UnsafePointer
		// 由于byte与int8同义，rune与int32同义，所以并不需要单独处理

		// 枚举参数的类型，并进行类型转换
		switch inTypeItem.Kind() {
		case reflect.Ptr:
			in[i] = reflect.ValueOf(paramItem)
			/*
				if param_request, ok := paramItem.(*Common.RequestModel); ok {
					in[i] = reflect.ValueOf(param_request)
				}
				if param_client, ok := paramItem.(*Client); ok {
					in[i] = reflect.ValueOf(param_client)
				}
				if param_player, ok := paramItem.(*player.Player); ok {
					in[i] = reflect.ValueOf(param_player)
				}
			*/
		case reflect.Bool:
			if param_bool, ok := typeUtil.Bool(paramItem); ok {
				in[i] = reflect.ValueOf(param_bool)
			}
		case reflect.Int:
			if param_float64, ok := typeUtil.Float64(paramItem); ok {
				in[i] = reflect.ValueOf(int(param_float64))
			}
		case reflect.Int8:
			if param_float64, ok := typeUtil.Float64(paramItem); ok {
				in[i] = reflect.ValueOf(int8(param_float64))
			}
		case reflect.Int16:
			if param_float64, ok := typeUtil.Float64(paramItem); ok {
				in[i] = reflect.ValueOf(int16(param_float64))
			}
		case reflect.Int32:
			if param_float64, ok := typeUtil.Float64(paramItem); ok {
				in[i] = reflect.ValueOf(int32(param_float64))
			}
		case reflect.Int64:
			if param_float64, ok := typeUtil.Float64(paramItem); ok {
				in[i] = reflect.ValueOf(int64(param_float64))
			}
		case reflect.Uint:
			if param_float64, ok := typeUtil.Float64(paramItem); ok {
				in[i] = reflect.ValueOf(uint(param_float64))
			}
		case reflect.Uint8:
			if param_float64, ok := typeUtil.Float64(paramItem); ok {
				in[i] = reflect.ValueOf(uint8(param_float64))
			}
		case reflect.Uint16:
			if param_float64, ok := typeUtil.Float64(paramItem); ok {
				in[i] = reflect.ValueOf(uint16(param_float64))
			}
		case reflect.Uint32:
			if param_float64, ok := typeUtil.Float64(paramItem); ok {
				in[i] = reflect.ValueOf(uint32(param_float64))
			}
		case reflect.Uint64:
			if param_float64, ok := typeUtil.Float64(paramItem); ok {
				in[i] = reflect.ValueOf(uint64(param_float64))
			}
		case reflect.Float32:
			if param_float64, ok := typeUtil.Float64(paramItem); ok {
				in[i] = reflect.ValueOf(float32(param_float64))
			}
		case reflect.Float64:
			if param_float64, ok := typeUtil.Float64(paramItem); ok {
				in[i] = reflect.ValueOf(param_float64)
			}
		case reflect.String:
			if param_string, ok := typeUtil.String(paramItem); ok {
				in[i] = reflect.ValueOf(param_string)
			}
		case reflect.Slice:
			// 如果是Slice类型，则需要对其中的项再次进行类型判断及类型转换
			if param_interface, ok := paramItem.([]interface{}); ok {
				switch inTypeItem.String() {
				case "[]bool":
					params_inner := make([]bool, len(param_interface))
					for i := 0; i < len(param_interface); i++ {
						if param_bool, ok := typeUtil.Bool(param_interface[i]); ok {
							params_inner[i] = param_bool
						}
					}
					in[i] = reflect.ValueOf(params_inner)

				case "[]int":
					params_inner := make([]int, len(param_interface))
					for i := 0; i < len(param_interface); i++ {
						if param_float64, ok := typeUtil.Float64(param_interface[i]); ok {
							params_inner[i] = int(param_float64)
						}
					}
					in[i] = reflect.ValueOf(params_inner)
				case "[]int8":
					params_inner := make([]int8, len(param_interface))
					for i := 0; i < len(param_interface); i++ {
						if param_float64, ok := typeUtil.Float64(param_interface[i]); ok {
							params_inner[i] = int8(param_float64)
						}
					}
					in[i] = reflect.ValueOf(params_inner)
				case "[]int16":
					params_inner := make([]int16, len(param_interface))
					for i := 0; i < len(param_interface); i++ {
						if param_float64, ok := typeUtil.Float64(param_interface[i]); ok {
							params_inner[i] = int16(param_float64)
						}
					}
					in[i] = reflect.ValueOf(params_inner)
				case "[]int32":
					params_inner := make([]int32, len(param_interface))
					for i := 0; i < len(param_interface); i++ {
						if param_float64, ok := typeUtil.Float64(param_interface[i]); ok {
							params_inner[i] = int32(param_float64)
						}
					}
					in[i] = reflect.ValueOf(params_inner)
				case "[]int64":
					params_inner := make([]int64, len(param_interface))
					for i := 0; i < len(param_interface); i++ {
						if param_float64, ok := typeUtil.Float64(param_interface[i]); ok {
							params_inner[i] = int64(param_float64)
						}
					}
					in[i] = reflect.ValueOf(params_inner)
				case "[]uint":
					params_inner := make([]uint, len(param_interface))
					for i := 0; i < len(param_interface); i++ {
						if param_float64, ok := typeUtil.Float64(param_interface[i]); ok {
							params_inner[i] = uint(param_float64)
						}
					}
					in[i] = reflect.ValueOf(params_inner)
				// case "[]uint8": 特殊处理
				case "[]uint16":
					params_inner := make([]uint16, len(param_interface))
					for i := 0; i < len(param_interface); i++ {
						if param_float64, ok := typeUtil.Float64(param_interface[i]); ok {
							params_inner[i] = uint16(param_float64)
						}
					}
					in[i] = reflect.ValueOf(params_inner)
				case "[]uint32":
					params_inner := make([]uint32, len(param_interface))
					for i := 0; i < len(param_interface); i++ {
						if param_float64, ok := typeUtil.Float64(param_interface[i]); ok {
							params_inner[i] = uint32(param_float64)
						}
					}
					in[i] = reflect.ValueOf(params_inner)
				case "[]uint64":
					params_inner := make([]uint64, len(param_interface))
					for i := 0; i < len(param_interface); i++ {
						if param_float64, ok := typeUtil.Float64(param_interface[i]); ok {
							params_inner[i] = uint64(param_float64)
						}
					}
					in[i] = reflect.ValueOf(params_inner)
				case "[]float32":
					params_inner := make([]float32, len(param_interface))
					for i := 0; i < len(param_interface); i++ {
						if param_float64, ok := typeUtil.Float64(param_interface[i]); ok {
							params_inner[i] = float32(param_float64)
						}
					}
					in[i] = reflect.ValueOf(params_inner)
				case "[]float64":
					params_inner := make([]float64, len(param_interface))
					for i := 0; i < len(param_interface); i++ {
						if param_float64, ok := typeUtil.Float64(param_interface[i]); ok {
							params_inner[i] = param_float64
						}
					}
					in[i] = reflect.ValueOf(params_inner)
				case "[]string":
					params_inner := make([]string, len(param_interface))
					for i := 0; i < len(param_interface); i++ {
						if param_string, ok := typeUtil.String(param_interface[i]); ok {
							params_inner[i] = param_string
						}
					}
					in[i] = reflect.ValueOf(params_inner)
				}
			} else if inTypeItem.String() == "[]uint8" { // 由于[]uint8在传输过程中会被转化成字符串，所以单独处理;
				if param_string, ok := typeUtil.String(paramItem); ok {
					param_uint8 := ([]uint8)(param_string)
					in[i] = reflect.ValueOf(param_uint8)
				}
			}
		}
	}

	// 判断是否有无效的参数（传入的参数类型和方法定义的类型不匹配导致没有赋值）
	for _, item := range in {
		if reflect.Value.IsValid(item) == false {
			return nil, fmt.Errorf("方法%s传入的参数%v无效", this.methodName, paramData)
		}
	}

	return in, nil
}

// 调用函数
// in:调用的参数信息
// 返回值:
// []reflect.Value:返回值
func (this *MethodAndInOutTypes) Call(in []reflect.Value) []reflect.Value {
	// 传入参数，调用方法
	out := this.method.Call(in)

	return out
}

// 创建方法实例对象
// method:方法对象
// inTypes:参数列表
// outTypes:返回值
func NewMethodAndInOutTypes(method reflect.Value, inTypes []reflect.Type, outTypes []reflect.Type) *MethodAndInOutTypes {
	return &MethodAndInOutTypes{
		method:   method,
		inTypes:  inTypes,
		outTypes: outTypes,
	}
}
