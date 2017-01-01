package dataBase

import (
	"errors"
	"fmt"
	"sync"
)

// 连接信息获取函数
// connectionName:连接名称
// 返回值：
// driverName:驱动名
// connectionString:数据库字符串
type connectionFun func(connectionName string) (driverName string, connectionString string)

var (
	// 连接对象集合
	connectionDic map[string]*DbConnection

	// 数据库集合锁对象
	connectionLock sync.RWMutex
)

func init() {
	// 初始化集合对象
	connectionDic = make(map[string]*DbConnection, 0)
}

// 初始化数据库连接
func AddConnectionByFun(dbNames []string, getConnection connectionFun) {
	for _, dbName := range dbNames {
		// 初始化game库
		driverName, connectionString := getConnection(dbName)
		if errMsg := AddConnection(dbName, driverName, connectionString); errMsg != nil {
			panic(errors.New("数据库初始化失败" + dbName + ":" + errMsg.Error()))
		}
	}
}

// 添加数据库连接
// dbName:数据库名
// driverName:驱动名
// connectionString:连接字符串
// 返回值：
// error:错误数据
func AddConnection(dbName string, driverName string, connectionString string) error {
	// 创建数据库连接
	connection := NewDbConnection(dbName)
	errMsg := connection.Init(driverName, connectionString)
	if errMsg != nil {
		return errMsg
	}

	connectionLock.Lock()
	defer connectionLock.Unlock()

	// 添加到集合
	connectionDic[dbName] = connection

	return nil
}

// 删除数据库连接
// dbName:数据库名
func removeConnection(dbName string) {
	connectionLock.Lock()
	defer connectionLock.Unlock()

	// 获取连接对象
	item, isFind := connectionDic[dbName]
	if isFind == false {
		return
	}

	// 删除连接
	delete(connectionDic, dbName)

	// 释放连接
	item.Dispose()
}

// 获取连接对象
// connectionName:连接名
// *DbConnection:数据库连接工具类
func GetConnection(connectionName string) *DbConnection {
	connectionLock.RLock()
	defer connectionLock.RUnlock()

	// 获取连接对象
	item, isFind := connectionDic[connectionName]
	if isFind == false {
		panic(fmt.Errorf("不存连接对象，connectionName:%v", connectionName))
	}

	return item
}
