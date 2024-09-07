[EN](./EN.md) | [简中](./zh-CN.md) | [繁中](./zh-TW.md) | [JP](./JP.md) | [RU](./RU.md) | [FR](./FR.md) | [KR](./KR.md) | [VI](./VI.md)

# conf parameter details

### LogLevel: log level

### MaxPlayer: maximum number of players for this gameserver

### AutoCreate: whether to automatically register

### AppList:

Format: appid[configuration]
port_player: used for gateserver, represents the external connection port, i.e. kcp port
port_gt: used for gameserver, represents the port used to connect to gateserver
port_service: represents the port that this service accepts connections from other services
port_http: start an http server on this port

### App:
Port: listening port
InnerAddr: external address
OuterAddr: listening address

### NetConf:
Format: service [address]
Node: indicates the connection address of nodeserver

### MysqlConf:
Format: table [address]
Each required mysql connection address

### RedisConf:
Format: table [connection configuration]
Each required redis connection configuration

### Dispatch: (dispatch only)
Format: [] each node server address
When there are multiple servers that are not interoperable, modify this configuration to connect

### GameDataConfigPath: (gameserver only) Fill in the configuration table path (based on the program startup directory)