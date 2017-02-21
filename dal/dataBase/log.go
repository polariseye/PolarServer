package dataBase

import (
	"fmt"

	"github.com/Jordanzuo/goutil/logUtil"
)

// 记录Prepare错误
// command：执行的SQL语句
// err：错误对象
func WritePrepareError(command string, err error) {
	logUtil.Log(fmt.Sprintf("Prepare失败，错误信息：%s，command:%s", err, command), logUtil.Error, true)
}

// 记录Exec错误
// command：执行的SQL语句
// err：错误对象
func WriteExecError(command string, err error) {
	logUtil.Log(fmt.Sprintf("Exec失败，错误信息：%s，command:%s", err, command), logUtil.Error, true)
}

// 记录Scan错误
// command：执行的SQL语句
// err：错误对象
func WriteScanError(command string, err error) {
	logUtil.Log(fmt.Sprintf("Scan失败，错误信息：%s，command:%s", err, command), logUtil.Error, true)
}

// 记录Query错误
// command：执行的SQL语句
// err：错误对象
func WriteQueryError(command string, err error) {
	logUtil.Log(fmt.Sprintf("Query失败，错误信息：%s，command:%s", err, command), logUtil.Error, true)
}

// 记录Query错误
// command：执行的SQL语句
// err：错误对象
func WriteTransactionError(command string, err error) {
	logUtil.Log(fmt.Sprintf("Transaction失败，错误信息：%s，command:%s", err, command), logUtil.Error, true)
}
