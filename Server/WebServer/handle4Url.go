package WebServer

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Jordanzuo/goutil/logUtil"
	"github.com/Jordanzuo/goutil/webUtil"
	"github.com/polariseye/PolarServer/Common"
	"github.com/polariseye/PolarServer/Common/ErrorCode"
	"github.com/polariseye/PolarServer/ModuleManage"
)

type Handle4UrlStruct struct {
}

// 处理请求
func (this *Handle4UrlStruct) RequestHandle(response http.ResponseWriter, request *http.Request) {

	result := Common.NewResultModel(ErrorCode.DataError)

	// 对象序列化
	defer func() {
		if panicErr := recover(); panicErr != nil {
			logUtil.LogUnknownError(panicErr)

			// 设置数据错误
			result.SetError(ErrorCode.DataError, "unknown error")
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
		result.SetError(ErrorCode.DataError, "form parse error")
		return
	}

	requestInfos := make([]interface{}, 0, len(request.PostForm))
	for key, val := range request.PostForm {
		if key == "ModuleName" || key == "MethodName" {
			continue
		}

		tmpVal := ""
		if val != nil && len(val) > 0 {
			tmpVal = val[0]
		}

		requestInfos = append(requestInfos, tmpVal)
	}

	requsetModel := Common.NewRequestModel()
	requsetModel.Request = request
	requsetModel.Ip = webUtil.GetRequestIP(request)
	requsetModel.ModuleName = request.FormValue("ModuleName")
	requsetModel.MethodName = request.FormValue("MethodName")

	requsetModel.Data = requestInfos

	// 请求具体处理
	result = ModuleManage.ApiModuleManager.Call(requsetModel)
}

// 创建新的请求处理对象
func NewHandle4Url() *Handle4UrlStruct {
	return &Handle4UrlStruct{}
}
