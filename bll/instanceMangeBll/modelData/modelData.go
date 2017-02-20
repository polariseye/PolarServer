package modelData

import (
	"fmt"

	"github.com/Jordanzuo/goutil/stringUtil"
)

var moduleData map[string]IModel

func Register(modelItem IModel) {
	moduleName := modelItem.ModuleName()
	if stringUtil.IsEmpty(moduleName) {
		panic(fmt.Errorf("model模块的模块名不能为空"))
	}

	_, isExist := moduleData[moduleName]
	if isExist {
		panic(fmt.Errorf("model模块的模块名重复,ModuleName:%v", moduleName))
	}

	moduleData[moduleName] = modelItem
}

func Load() []error {
	errList := make([]error, 0, 100)
	for _, item := range moduleData {
		tmpErrList := item.Init()
		if tmpErrList != nil && len(tmpErrList) > 0 {
			errList = append(errList, tmpErrList...)
		}
	}

	if len(errList) > 0 {
		return errList
	}

	for _, item := range moduleData {
		tmpErrList := item.Check()
		if tmpErrList != nil && len(tmpErrList) > 0 {
			errList = append(errList, tmpErrList...)
		}

		tmpErrList = item.Convert()
		if tmpErrList != nil && len(tmpErrList) > 0 {
			errList = append(errList, tmpErrList...)
		}
	}

	if len(errList) > 0 {
		return errList
	}

	return nil
}

// 模块集合
func Modules() map[string]IModel {
	return moduleData
}
