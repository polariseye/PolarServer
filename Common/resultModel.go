package Common

// 请求结果实体对象
type ResultModel struct {

	// 状态
	ErrorInfo

	// 数据
	Value map[string]interface{}

	// 附加数据
	Avatar map[string]interface{}
}

// 创建新的实体对象
// defaultErrorCode:默认的错误码信息
func NewResultModel(defaultErrorCode ErrorCode) (result *ResultModel) {
	result = &ResultModel{
		Value:  make(map[string]interface{}),
		Avatar: make(map[string]interface{}),
	}

	result.SetNormalError(defaultErrorCode)

	return
}
