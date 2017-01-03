package apiHandle

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Jordanzuo/goutil/logUtil"
	"github.com/Jordanzuo/goutil/webUtil"
	"github.com/polariseye/polarserver/common"
	"github.com/polariseye/polarserver/common/errorCode"
	"github.com/polariseye/polarserver/server/serverBase"
	"github.com/polariseye/polarserver/server/webServer"
)

// url形式的数据处理
type Handle4UrlStruct struct {
	// 服务对象
	server *webServer.WebServerStruct

	// Api调用对象
	caller serverBase.IApiCaller
}

// 处理请求
// response:应答对象
// request:请求对象
func (this *Handle4UrlStruct) RequestHandle(response http.ResponseWriter, request *http.Request) {
	result := common.NewResultModel(errorCode.DataError)

	// 对象序列化
	defer func() {
		if panicErr := recover(); panicErr != nil {
			logUtil.LogUnknownError(panicErr)

			// 设置数据错误
			result.SetError(errorCode.DataError, "unknown error")
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
		result.SetError(errorCode.DataError, "form parse error")
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

	requsetModel := common.NewRequestModel()
	requsetModel.Request = request
	requsetModel.Ip = webUtil.GetRequestIP(request)
	requsetModel.ModuleName = request.FormValue("ModuleName")
	requsetModel.MethodName = request.FormValue("MethodName")

	requsetModel.Data = requestInfos

	// 请求具体处理
	result = this.caller.Call(requsetModel)
}

// 设置目标服务对象
func (this *Handle4UrlStruct) SetTargetServer(server *webServer.WebServerStruct) {
	this.server = server
}

// 创建新的请求处理对象
// _caller:调用对象
func NewHandle4Url(_caller serverBase.IApiCaller) *Handle4UrlStruct {
	return &Handle4UrlStruct{}
}
