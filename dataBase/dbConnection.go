package dataBase

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Jordanzuo/goutil/logUtil"
)

// 数据库连接结构体
type DbConnection struct {

	// 数据库对象
	db *sql.DB

	// 数据库连接名
	name string

	// 是否已经释放了资源
	isDisposed bool
}

// 创建一个新的数据库连接
// _name:连接名称
func NewDbConnection(_name string) *DbConnection {
	return &DbConnection{
		name: _name,
	}
}

// 连接名
func (this *DbConnection) Name() string {
	return this.name
}

// 初始化数据库连接信息
// driver：驱动名
// connectionString:链接字符串
// 返回值：
// error:初始化的错误信息
func (this *DbConnection) Init(driver string, connectionString string) error {
	connectionSlice := strings.Split(connectionString, "||")
	isHaveExtraConfig := len(connectionSlice) == 3

	// 打开连接对象
	db, errMsg := sql.Open(driver, connectionSlice[0])
	if errMsg != nil {
		return errMsg
	}

	// 设置连接池相关
	if isHaveExtraConfig {
		maxOpenConns_string := strings.Replace(connectionSlice[1], "MaxOpenConns=", "", 1)
		maxOpenCons, err := strconv.Atoi(maxOpenConns_string)
		if err != nil {
			panic(fmt.Errorf("MaxOpenConns必须为int型,连接字符串为：%s", connectionString))
		}

		maxIdleConns_string := strings.Replace(connectionSlice[2], "MaxIdleConns=", "", 1)
		maxIdleConns, err := strconv.Atoi(maxIdleConns_string)
		if err != nil {
			panic(fmt.Errorf("MaxIdleConns必须为int型,连接字符串为：%s", connectionString))
		}

		if maxOpenCons > 0 && maxIdleConns > 0 {
			db.SetMaxOpenConns(maxOpenCons)
			db.SetMaxIdleConns(maxIdleConns)

			// 设置了连接池后，一直ping
			go this.ping()
		}
	}

	if err := db.Ping(); err != nil {
		panic(fmt.Errorf("Ping数据库失败,连接字符串为：%s", connectionString))
	}

	this.db = db

	return nil
}

// 每分钟ping一次数据库
func (this *DbConnection) ping() {
	for {
		// 如果已经释放了资源，则退出循环
		if this.isDisposed {
			return
		}

		time.Sleep(time.Minute)

		if err := this.GetDb().Ping(); err != nil {
			logUtil.Log(fmt.Sprintf("游戏数据库Ping失败，错误信息为：%s", err), logUtil.Error, true)
		}
	}
}

// 获取数据库连接对象
// 返回值：
// *sql.DB:数据库对象
func (this *DbConnection) GetDb() *sql.DB {
	return this.db
}

// 释放所有资源
func (this *DbConnection) Dispose() {
	this.isDisposed = true

	if this.db != nil {
		this.db.Close()
	}
}

// 查询
// sql:待查询的sql语句
// args:参数
// 返回值：
// *sql.Rows:结果数据
// error:错误数据
func (this *DbConnection) Query(sql string, args ...interface{}) (*sql.Rows, error) {
	return this.db.Query(sql, args...)
}

// 查询并获得一条数据
// sql:待查询的sql语句
// args:参数
// 返回值：
// *sql.Row:结果数据
func (this *DbConnection) QueryRow(sql string, args ...interface{}) *sql.Row {
	return this.db.QueryRow(sql, args...)
}

// 执行一条sql语句
// sql:待执行的sql语句
// args:参数
// 返回值：
// int64:影响记录数
// error:错误信息
func (this *DbConnection) Execute(sql string, args ...interface{}) (int64, error) {

	// 执行sql
	result, errMsg := this.db.Exec(sql, args...)
	if errMsg != nil {
		return 0, errMsg
	}

	// 返回影响记录数
	return result.RowsAffected()
}

// 批量执行sql语句（内部会采用事务提交）
// sqlContent:待执行的sql语句
// paramList:sql语句对应的所有参数列表
// 返回值：
// int64：总共影响的记录集合
// error：错误信息
func (this *DbConnection) ExecuteList(sqlContent string, paramList [][]interface{}) (int64, error) {

	// 开启事务
	tran, errMsg := this.db.Begin()
	if errMsg != nil {
		return 0, errMsg
	}

	isCommit := false
	defer func() {
		if isCommit {

			// 提交事务
			tran.Commit()
		} else {

			// 事务回滚
			tran.Rollback()
		}
	}()

	// 初始化sql语句执行
	var stmt *sql.Stmt
	stmt, errMsg = tran.Prepare(sqlContent)
	if errMsg != nil {
		return 0, errMsg
	}

	var (
		result   sql.Result
		count    int64
		tmpCount int64
	)

	// 批量提交数据
	for _, paramsItem := range paramList {

		// 执行
		result, errMsg = stmt.Exec(paramsItem...)
		if errMsg != nil {
			return 0, errMsg
		}

		// 获取影响数据行数
		tmpCount, errMsg = result.RowsAffected()
		if errMsg != nil {
			return 0, errMsg
		}

		// 记录影响记录数
		count += tmpCount
	}

	isCommit = true

	return count, nil
}
