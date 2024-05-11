# conf参数详解

### LogLevel：日志等级
### MaxPlayer：本gameserver最大玩家数
### AutoCreate：是否自动注册
### OuterIp：本服务启动地址
### AppList：
    格式：appid[配置]
    port_player：用于gateserver，代表对外连接端口，即kcp端口
    port_gt：用于gameserver，代表与gateserver构成连接所使用的端口
    port_service：代表本服务接受其他服务连接的端口
    port_http：在这个端口上启动一个http服务器
### NetConf：
    格式：服务[地址]
    Node：表示nodeserver的连接地址
### MysqlConf:
    格式：表[地址]
    各个需要的mysql连接地址
### RedisConf:
    格式：表[连接配置]
    各个需要的redis连接配置
### Dispatch:(dispatch专用)
    格式：[]各个节点服务器地址
    当有多个不互通的服务器时，修改此配置可连接
### GameDataConfigPath:(gameserver专用)填写配置表路径（以程序启动目录为准）