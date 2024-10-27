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

# API
- サンプルパラメータはすべて uid = 1 |sign_key = 123456 に基づいています。

### パスパラメータ
- **cmd**: `int16` が必要です | **コマンドを呼び出します**
- **uid**: `uint32` が必要です | **プレイヤー UID**
- **sign_key**: `文字列` オプション | **キー**

#### リクエストの例:
```平文
GET: api?cmd=1&uid=1&sign_key=123456
「」

#### `json` を解析するコールバック:
- **コード**: ステータス 0 成功 -1 失敗
- **msg**: コールバックの内容

___

### ワールドレベルの設定 cmd 1001
**パラメータ**:
- **world_level**: `uint32` が必要です **設定するワールド レベル**
#### リクエストの例:
```平文
GET: api?cmd=1001&uid=1&sign_key=123456&world_level=6
「」

___

### アカウント データの取得 cmd 1002
**パラメータ**: なし
#### リクエストの例:
```平文
GET: api?cmd=1002&uid=1&sign_key=123456
「」

___

### サーバーステータスの取得 cmd 1003
**パラメータ**: なし
#### リクエストの例:
```平文
GET: api?cmd=1003&sign_key=123456
「」

___

### 小道具の取得 cmd 1004
- **all**: `bool` オプション | **すべての項目を取得するかどうか | 0:false|1:true
- **id**: `uint32` オプション | **アイテム ID**
- **num**: `uint32` オプション | **項目数**
#### リクエストの例:
```平文
GET: api?cmd=1004&uid=1&sign_key=123456&all=0&id=22&num=999
「」

___

### 聖遺物入手cmd 1005
- **all**: `bool` オプション | **すべての項目を取得するかどうか | 0:false|1:true
- **id**: `uint32` が必要です | **聖遺物 ID**
- **num**: `uint32` が必要 | **聖遺物の数**
- **main**: `uint32` オプション | **聖遺物の主な属性を指定します**
- **sub**: `string` オプション | **聖遺物の二次属性を指定します**
#### リクエストの例:
```平文
GET: api?cmd=1005&uid=1&sign_key=123456&all=0&id=31011&num=1&main=1&sub=[2:10][3:9][4:8]
「」

___