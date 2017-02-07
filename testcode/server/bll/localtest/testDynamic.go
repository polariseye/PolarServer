package localtest

import (
	"github.com/polariseye/polarserver/testcode/server/common"
)

type TestDynamicStruct struct {
	dynamic     interface{}
	dynamicData map[interface{}]TestDynamicStruct
	Data        interface{}
}

var TestDynamicBLL *TestDynamicStruct

func init() {
	TestDynamicBLL = newTestDynamicBLL()
	common.RegisterApiModule(TestDynamicBLL)
	common.RegisterInit(TestDynamicBLL)
}

func (this *TestDynamicStruct) Init(tmpDynamic interface{}) {
	instance := this
	if this.initInstance != nil {
		//throw new exception:初始化中,不能反复初始化
		return
	}

	instance = this.initInstance

	// 初始化instance
	instance.Init()
	this.GetItem()
}

func (this *TestDynamicStruct) Check(tmpDynamic interface{}) {
	instance := this
	if this.initInstance != nil {
		//throw new exception:初始化中,不能反复初始化
		return
	}

	instance = this.initInstance

	// 初始化instance
	instance.Check()
}

func (this *TestDynamicStruct) Convert(tmpDynamic interface{}) {

}

func (this *TestDynamicStruct) Confirm(tmpDynamic interface{}) {
	if this.initInstance == tmpDynamic {

	}
}

func (this *TestDynamicStruct) GetData(tmpDynamic interface{}) {
	if tmpDynamic == this.dynamic {
		return this.Data
	}

	return this.initInstance.Data
}

func (this *TestDynamicStruct) GetItem(tmpDynamic interface{}) interface{} {

	return this
}

func newTestDynamicBLL() *TestDynamicStruct {
	return &TestDynamicStruct{}
}
