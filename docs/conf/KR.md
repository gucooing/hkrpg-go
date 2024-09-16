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

# conf 매개변수에 대한 자세한 설명

### LogLevel: 로그 수준

### MaxPlayer: 이 게임 서버의 최대 플레이어 수

### AutoCreate: 자동으로 등록할지 여부

###앱 목록:

 형식: appid[구성]
 port_player: 게이트서버에 사용되며 외부 연결 포트(예: kcp 포트)를 나타냅니다.
 port_gt: 게임 서버에 사용되며, 게이트 서버에 연결하는 데 사용되는 포트를 나타냅니다.
 port_service: 이 서비스가 다른 서비스의 연결을 수락하는 포트를 나타냅니다.
 port_http: 이 포트에서 http 서버를 시작합니다.

### 앱:
 포트: 청취 포트
 InnerAddr: 외부 주소
 OuterAddr: 청취 주소

### NetConf:
 형식: 서비스[주소]
 Node: 노드 서버의 연결 주소를 나타냅니다.

### MysqlConf:
 형식: 테이블[주소]
 각각의 필수 mysql 연결 주소

### RedisConf:
 형식: 테이블[연결 구성]
 각 필수 Redis 연결 구성

### 파견: (배달만)
 형식: [] 각 노드의 서버 주소
 서로 통신하지 않는 서버가 여러 대 있는 경우 이 구성을 수정하여 연결하세요.

### GameDataConfigPath: (게임 서버에만 해당) 구성 테이블 경로를 입력합니다(프로그램 시작 디렉터리 기준).