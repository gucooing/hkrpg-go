[EN](./EN.md) | [简中](./zh-CN.md) | [繁中](./zh-TW.md) | [JP](./JP.md) | [RU](./RU.md) | [FR](./FR.md) | [KR](./KR.md) | [VI](./VI.md)

# conf参数详解

### LogLevel：日志等级

### MaxPlayer：单进程最大支持玩家数量（仅gateserver和gameserver有效

### AutoCreate：仅sdk有效

### AppList：
    格式：appid[配置]
    port_player：用于gateserver，代表对外连接端口，即kcp端口
    port_gt：用于gameserver，代表与gateserver构成连接所使用的端口
    port_service：代表本服务接受其他服务连接的端口
    port_http：在这个端口上启动一个http服务器
    MqAddr: gate的消息队列服务地址
    RegionName: 区服名称

### AppNet:
    InnerAddr: 启动地址
    InnerPort: 启动端口
    OuterAddr: 暴露地址
    OuterPort: 暴露端口

### NetConf：
    格式：服务[地址]
    Node：表示nodeserver的连接地址
    仅Node一个地址

### MysqlConf:
    格式：表[地址]
    各个需要的mysql连接地址

### RedisConf:
    格式：表[连接配置]
    各个需要的redis连接配置


### RedisConf:
    Name: 区服名称
    AutoCreate: 是否自动注册
    Title: 地区
    Type: sdk类型
    ClientSecretKey: 不需要填写（由nodeserver启动时自动生成
### GameDataConfigPath:(gameserver专用)填写配置表路径（以程序启动目录为准）