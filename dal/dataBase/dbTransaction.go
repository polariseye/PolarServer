package dataBase

import (
	"database/sql"
)

type DbTransaction struct {
	transaction *sql.Tx
}

// 事务执行函数定义
// tran:事务对象
// 返回值:
// bool:是否提交事务
// error:错误信息
type ExecuteFun func(tran *DbTransaction) (isCommit bool, errMsg error)

// 创建新的事务管理对象
func NewDbTransaction(tran *sql.Tx) *DbTransaction {
	return &DbTransaction{
		transaction: tran,
	}
}

// 执行一条sql语句
// sql:待执行的sql语句
// args:参数
// 返回值：
// int64:影响记录数
// error:错误信息
func (this *DbTransaction) Execute(sql string, args ...interface{}) (int64, error) {
	// 执行sql
	result, errMsg := this.transaction.Exec(sql, args...)
	if errMsg != nil {
		return 0, errMsg
	}

	// 返回影响记录数
	return result.RowsAffected()
}

// 以事务方式执行
// connectionName:连接名
// executeFun:执行函数
// 返回值：
// error：执行的错误信息
func ExecuteByTran(connectionName string, executeFun ExecuteFun) error {
	// 获取连接,并创建事务
	connection := GetConnection(connectionName)
	db := connection.GetDb()
	tran, errMsg := db.Begin()
	if errMsg != nil {
		return errMsg
	}

	// 事务结果处理
	isCommit := false
	defer func() {
		if isCommit {
			tran.Commit()
		} else {
			tran.Rollback()
		}
	}()

	tranInfo := NewDbTransaction(tran)

	// 执行
	isCommit, errMsg = executeFun(tranInfo)

	return errMsg
}
