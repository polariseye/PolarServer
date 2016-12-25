package WebServer

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Jordanzuo/goutil/logUtil"
	"github.com/Jordanzuo/goutil/typeUtil"
	"github.com/Jordanzuo/goutil/webUtil"
	"github.com/polariseye/PolarServer/Common"
	"github.com/polariseye/PolarServer/ModuleManage"
)

type Handle4UrlStruct struct {
}

// 处理请求
func (this *Handle4UrlStruct) RequestHandle(response http.ResponseWriter, request *http.Request) {

	result := Common.NewResultModel(Common.DataError(""))

	// 对象序列化
	defer func() {
		if panicErr := recover(); panicErr != nil {
			logUtil.LogUnknownError(panicErr)

			// 设置数据错误
			result.SetNormalError(Common.DataError("unknown error"))
		}

		data, tmpErrMsg := json.Marshal(&result)
		if tmpErrMsg != nil {
			logUtil.NormalLog(fmt.Sprintf("应答序列化异常:%v", tmpErrMsg.Error()), logUtil.Error)
			return
		}

		response.Write(data)
	}()

	// 转换为form表单
	if tmpErr := request.ParseForm(); tmpErr != nil {
		result.SetNormalError(Common.DataError("form parse error"))
		return
	}
	requestInfos := typeUtil.NewMapData()
	for key, val := range request.PostForm {
		tmpVal := ""
		if val != nil && len(val) > 0 {
			tmpVal = val[0]
		}
		requestInfos[key] = tmpVal
	}

	requsetModel := Common.NewRequestModel()
	requsetModel.Request = request
	requsetModel.Ip = webUtil.GetRequestIP(request)
	requsetModel.ModuleName, _ = requestInfos.String("ModuleName")
	requsetModel.MethodName, _ = requestInfos.String("MethodName")
	requsetModel.Data = requestInfos

	// 请求具体处理
	result = ModuleManage.Invoke(requsetModel)
}

// 创建新的请求处理对象
func NewHandle4Url() *Handle4UrlStruct {
	return &Handle4UrlStruct{}
}
