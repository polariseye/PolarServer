package dataBase

// 游戏库获取
// 返回值:
// *DbSession:游戏库会话对象
func GameDb() *DbSession {
	return NewDbSession(GetConnection("GameDb"))
}

// model库连接对象获取
// 返回值:
// *DbSession:model库会话对象
func ModelDb() *DbSession {
	return NewDbSession(GetConnection("ModelDb"))
}

// 日志库获取
// 返回值:
// *DbSession:日志库会话对象
func LogDb() *DbSession {
	return NewDbSession(GetConnection("LogDb"))
}
