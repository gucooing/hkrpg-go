<div align="center">
<table>
<td valign="center"><a href="EN.md"><img src="https://github.com/twitter/twemoji/blob/master/assets/svg/1f1fa-1f1f8.svg" width="16"/> English</td>
 
<td valign="center"><a href="zh-CN.md"><img src="https://em-content.zobj.net/thumbs/120/twitter/351/flag-china_1f1e8-1f1f3.png" width="16"/> 简中</td>
 
<td valign="center"><a href="zh-TW.md"><img src="https://em-content.zobj.net/thumbs/120/twitter/351/flag-china_1f1e8-1f1f3.png" width="16"/> 繁中</td>
 
<td valign="center"><a href="JP.md"><img src="https://github.com/twitter/twemoji/blob/master/assets/svg/1f1ef-1f1f5.svg" width="16"/> 日本語</td>
 
<td valign="center"><a href="RU.md"><img src="https://github.com/twitter/twemoji/blob/master/assets/svg/1f1f7-1f1fa.svg" width="16"/> Русский</a></td>

<td valign="center"><a href="FR.md"><img src="https://em-content.zobj.net/thumbs/160/twitter/154/flag-for-france_1f1eb-1f1f7.png" width="16"/> Français</td>
 
<td valign="center"><a href="KR.md"><img src="https://em-content.zobj.net/source/twitter/53/flag-for-south-korea_1f1f0-1f1f7.png" width="16"/> 한국어</td>
 
<td valign="center"><a href="VI.md"><img src="https://em-content.zobj.net/thumbs/120/twitter/351/flag-vietnam_1f1fb-1f1f3.png" width="16"/> Tiếng Việt </a></td>
</table>
</div>

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