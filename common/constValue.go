package common

// 是否处于测试状态
var isTest bool = false

// 是否处于测试状态
func IsTest() bool {
	return isTest
}

// 设置是否处于测试状态
// _isTest:是否处于测试状态
func SetIsTest(_isTest bool) {
	isTest = _isTest
}
