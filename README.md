[PolarServer介绍](https://github.com/polariseye/PolarServer)
==============================

当然，这也算不上一个server，叫做util工具集更为恰当，就像[beego](http://github.com\astaxie\beego)一样。由于各种原因，把框架类代码写在项目内部导致各种维护问题。所以建了此项目

这个项目主要是构建一个方便进行后续业务功能开发的框架。会集成业务型功能开发的各个方面，包含：数据库、web服务、rpc服务、rpc调用的各个方面。

##功能预期##
我最终对这个项目的预期是：

1. 作为项目开发中的核心服务框架，能够方便进行web站点开发、apiserver开发、兼容socket级别的实时通信类开发
2. 提供的各个具体server只需要像使用mysql驱动一样，import即可使用，而不需要做过多配置
3. 最终将配备DAL层的代码生成工具
4. 对ApiServer而言，只需要按照指定规则命名函数，即可被正常调用

##PolarServer目录结构一览##

* [Common](/Common/): 公共包，提供一些常用函数
* [Config](/Config/): 配置包，提供配置文件操作的包
* [DataBase](/DataBase/): 数据库工具包，提供数据库操作的公共函数
* [ModuleManage](/ModuleManage/)：模块管理包，负责各个模块的初始化、数据检查等工作
* [Server](/Server/)：服务包，提供各种协议的服务，比如：站点服务、定时调度服务等

##现在存在的问题##
* 方法调用的权限验证
 +  用户验证
 +  用户调用接口的验证

* 模块重新加载
 + 暂未考虑到模块重新加载的问题（已调整）
  

* 数据重新加载
  + 数据重新加载一般是带有一定条件的，当满足条件后才需要重新加载
  + 重新加载时，要读取重新加载的其他配置，不能从老的数据中读取
* xml配置文件的支持(已调整)
* gameserver的reloadplayer
