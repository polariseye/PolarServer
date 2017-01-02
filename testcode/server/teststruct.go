package main

import (
	"fmt"

	"github.com/polariseye/polarserver/common"
	"github.com/polariseye/polarserver/common/errorCode"
	"github.com/polariseye/polarserver/moduleManage"
)

type testStruct struct {
	className string
}

var TestBLL *testStruct

func init() {
	TestBLL = NewTestStruct()
	moduleManage.RegisterModule(TestBLL, moduleManage.Normal, moduleManage.ApiModule)
}

// 类名
func (this *testStruct) Name() string {
	return this.className
}

func (this *testStruct) InitModule() []error {
	fmt.Println("初始化")

	return nil
}

func (this *testStruct) CheckModule() []error {
	fmt.Println("check")

	return nil
}

func (this *testStruct) ConvertModule() []error {
	fmt.Println("数据转换")

	return nil
}

// 接口调用
func (this *testStruct) C_Hello(request *common.RequestModel, d int, name string) *common.ResultModel {
	result := common.NewResultModel(errorCode.ClientDataError)

	result.Value["Hello"] = name + "_" + this.Name()
	result.Value["Extra"] = d

	result.SetNormalError(errorCode.Success)
	return result
}

func NewTestStruct() *testStruct {
	return &testStruct{
		className: "TestBLL",
	}
}
