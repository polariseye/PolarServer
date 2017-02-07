/*
使用单独的init函数初始化
*/
package localtest

import (
	"github.com/polariseye/polarserver/testcode/server/common"
)

type TestGlobalStruct struct {
}

var TestGlobalBLL *TestGlobalStruct

func init() {
	TestGlobalBLL = newTestGlobal()
	common.RegisterApiModule(TestGlobalBLL)
	common.RegisterInit(TestGlobalBLL)
}

func (this *TestGlobalStruct) Init() {

}

func newTestGlobal() *TestGlobalStruct {
	return &TestGlobalStruct{}
}
