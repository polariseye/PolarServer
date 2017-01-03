package apiHandle

import (
	"bytes"
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

// 使用Json形式进行数据格式解析
type Handle4JsonStruct struct {
	// 服务对象
	server *webServer.WebServerStruct

	// Api调用对象
	caller serverBase.IApiCaller
}

// 处理请求
// response:应答对象
// request:请求对象
func (this *Handle4JsonStruct) RequestHandle(response http.ResponseWriter, request *http.Request) {
	result := common.NewResultModel(errorCode.ClientDataError)

	// 对象序列化
	defer func() {
		if panicErr := recover(); panicErr != nil {
			logUtil.LogUnknownError(panicErr)

			// 设置数据错误
			result.SetNormalError(errorCode.DataError)
		}

		data, tmpErrMsg := json.Marshal(&result)
		if tmpErrMsg != nil {
			logUtil.NormalLog(fmt.Sprintf("应答序列化异常:%v", tmpErrMsg.Error()), logUtil.Error)
			return
		}

		response.Write(data)
	}()

	// 读取数据
	buf := bytes.NewBuffer(nil)
	dataLen, err := buf.ReadFrom(request.Body)
	if err != nil {
		result.SetError(errorCode.DataError, "read request data error")
		return
	} else if dataLen <= 0 {
		result.SetError(errorCode.DataError, "have no request data")
		return
	}

	// 反序列化
	requestModel := common.NewRequestModel()
	if err = json.Unmarshal(buf.Bytes(), &requestModel); err != nil {
		result.SetError(errorCode.DataError, "json format error")
		return
	}

	// 设置请求参数
	requestModel.Request = request
	requestModel.Ip = webUtil.GetRequestIP(request)

	// 请求具体处理
	result = this.caller.Call(requestModel)
}

// 设置目标服务对象
func (this *Handle4JsonStruct) SetTargetServer(server *webServer.WebServerStruct) {
	this.server = server
}

// 创建新的请求处理对象
// _caller:调用对象
func NewHandle4Json(_caller serverBase.IApiCaller) *Handle4JsonStruct {
	return &Handle4JsonStruct{
		caller: _caller,
	}
}
