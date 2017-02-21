package dataBase

// 游戏库获取
// 返回值:
// *DbConnection:游戏库对象
func GameDb() *DbConnection {
	return GetConnection("GameDb")
}

// model库连接对象获取
// 返回值:
// *DbConnection:model库连接对象
func ModelDb() *DbConnection {
	return GetConnection("ModelDb")
}

// 日志库获取
// 返回值:
// *DbConnection:日志库连接对象
func LogDb() *DbConnection {
	return GetConnection("LogDb")
}
