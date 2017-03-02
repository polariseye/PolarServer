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
