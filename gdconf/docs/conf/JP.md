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

# confパラメータの詳細説明

### LogLevel: ログレベル

### MaxPlayer: このゲームサーバーの最大プレイヤー数

### AutoCreate: 自動的に登録するかどうか

###アプリリスト:

 形式: appid[設定]
 port_player: ゲートサーバーに使用され、外部接続ポート、つまり kcp ポートを表します。
 port_gt: ゲームサーバーに使用され、ゲートサーバーへの接続に使用されるポートを表します。
 port_service: このサービスが他のサービスからの接続を受け入れるポートを表します。
 port_http: このポートで http サーバーを起動します

＃＃＃ アプリ：
 ポート: リスニングポート
 InnerAddr: 外部アドレス
 OuterAddr: リスニングアドレス

### ネット会議:
 形式: サービス[アドレス]
 ノード: ノードサーバーの接続アドレスを表します。

### MysqlConf:
 形式: テーブル[アドレス]
 必要な各 mysql 接続アドレス

### RedisConf:
 形式：テーブル[接続構成]
 必要な各 Redis 接続構成

### 派遣：（派遣のみ）
 形式: [] 各ノードのサーバーアドレス
 相互に通信していないサーバーが複数ある場合は、この構成を変更して接続します。

### GameDataConfigPath: (ゲームサーバーのみ) (プログラムの起動ディレクトリに基づく) 構成テーブルのパスを入力します。