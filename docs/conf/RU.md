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

# Подробное объяснение параметров конфигурации

### LogLevel: уровень журнала.

### MaxPlayer: Максимальное количество игроков на этом игровом сервере.

### AutoCreate: следует ли регистрироваться автоматически

###Список приложений:

 Формат: appid[конфигурация]
 port_player: используется для сервера шлюзов, представляющего порт внешнего подключения, т. е. порт kcp.
 port_gt: используется для игрового сервера, представляет порт, используемый для подключения к серверу ворот.
 port_service: представляет порт, на котором эта служба принимает соединения от других служб.
 port_http: запустить http-сервер на этом порту.

###Приложение:
 Порт: порт прослушивания
 InnerAddr: внешний адрес
 OuterAddr: адрес прослушивания

### НетКонф:
 Формат: услуга[адрес]
 Узел: представляет адрес подключения узла-сервера.

### MysqlConf:
 Формат: таблица[адрес]
 Каждый требуемый адрес подключения MySQL

### RedisConf:
 Формат: таблица[конфигурация соединения]
 Каждая необходимая конфигурация подключения Redis

### Отправка:(только отправка)
 Формат: [] Адрес сервера каждого узла.
 Если есть несколько серверов, которые не обмениваются данными друг с другом, измените эту конфигурацию для подключения.

### GameDataConfigPath: (только для игрового сервера) заполните путь к таблице конфигурации (на основе каталога запуска программы)