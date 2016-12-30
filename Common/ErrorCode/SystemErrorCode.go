package ErrorCode

var (
	// 调用成功
	Success ErrorCode = NewErrorCode(0, "调用成功")
	// 客户端数据错误
	ClientDataError ErrorCode = NewErrorCode(-1, "客户端数据错误")
	// 数据错误
	DataError ErrorCode = NewErrorCode(-2, "数据错误")
	// 客户端调用参数不足
	paramNotEnough ErrorCode = NewErrorCode(-1, "客户端调用参数不足")
)
