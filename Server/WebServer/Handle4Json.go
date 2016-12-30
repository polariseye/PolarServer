package WebServer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Jordanzuo/goutil/logUtil"
	"github.com/Jordanzuo/goutil/webUtil"
	"github.com/polariseye/PolarServer/Common"
	"github.com/polariseye/PolarServer/ModuleManage"
)

// 使用Json形式进行数据格式解析
type Handle4JsonStruct struct {
	server *WebServerStruct
}

// 处理请求
func (this *Handle4JsonStruct) RequestHandle(response http.ResponseWriter, request *http.Request) {

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

	// 读取数据
	buf := bytes.NewBuffer(nil)
	dataLen, err := buf.ReadFrom(request.Body)
	if err != nil {
		result.SetNormalError(Common.DataError("read request data error"))
		return
	} else if dataLen <= 0 {
		result.SetNormalError(Common.DataError("have no request data"))
		return
	}

	// 反序列化
	requestModel := Common.NewRequestModel()
	if err = json.Unmarshal(buf.Bytes(), &requestModel); err != nil {
		result.SetNormalError(Common.DataError("json format error"))
		return
	}

	// 设置请求参数
	requestModel.Request = request
	requestModel.Ip = webUtil.GetRequestIP(request)

	// 请求具体处理
	result = ModuleManage.Invoke(requestModel)
}

// 创建新的请求处理对象
// webServer:服务对象
func NewHandle4Json(webServer *WebServerStruct) *Handle4JsonStruct {
	return &Handle4JsonStruct{
		server: webServer,
	}
}
