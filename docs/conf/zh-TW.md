[EN](./EN.md) | [简中](./zh-CN.md) | [繁中](./zh-TW.md) | [JP](./JP.md) | [RU](./RU.md) | [FR](./FR.md) | [KR](./KR.md) | [VI](./VI.md)

# conf參數詳解

### LogLevel：日誌等級

### MaxPlayer：本gameserver最大玩家數

### AutoCreate：是否自動註冊

### AppList：

 格式：appid[配置]
 port_player：用於gateserver，代表對外連接端口，即kcp端口
 port_gt：用於gameserver，代表與gateserver構成連接所使用的連接埠
 port_service：代表本服務接受其他服務連線的連接埠
 port_http：在這個連接埠上啟動一個http伺服器

### App:
 Port:監聽埠
 InnerAddr:外部位址
 OuterAddr:監聽地址

### NetConf：
 格式：服務[地址]
 Node：表示nodeserver的連線位址

### MysqlConf:
 格式：表[地址]
 各個需要的mysql連接位址

### RedisConf:
 格式：表[連線配置]
 各個需要的redis連接配置

### Dispatch:(dispatch專用)
 格式：[]各節點伺服器位址
 當有多個不互通的伺服器時，修改此配置可連接

### GameDataConfigPath:(gameserver專用)填入設定表路徑（以程式啟動目錄為準）