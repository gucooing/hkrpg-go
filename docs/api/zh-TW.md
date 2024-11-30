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
- 範例參數皆以 uid = 1 | sign_key = 123456 進行

### 路徑參數
- **cmd**: `int16` 必須 | **呼叫指令**
- **uid**: `uint32` 必須 | **玩家UID**
- **sign_key**: `string` 可選 | **key**

#### 請求範例:
```plaintext
GET: api?cmd=1&uid=1&sign_key=123456
```

#### 回呼解析`json`:
- **code**: 狀態 0 成功 -1 失敗
- **msg**: 回呼內容

___

### 設定世界等級 cmd 1001
**參數**:
- **world_level**: `uint32` 必須 | **要設定的世界等級**
#### 請求範例:
```plaintext
GET: api?cmd=1001&uid=1&sign_key=123456&world_level=6
```

___

### 取得帳號資料 cmd 1002
**參數**:無
#### 請求範例:
```plaintext
GET: api?cmd=1002&uid=1&sign_key=123456
```

___

### 取得伺服器狀態 cmd 1003
**參數**:無
#### 請求範例:
```plaintext
GET: api?cmd=1003&sign_key=123456
```

___

### 取得道具 cmd 1004
- **all**: `bool` 可選 | **是否取得全部物品 | 0:false|1:true**
- **id**: `uint32` 可選 | **物品id**
- **num**: `uint32` 可選 | **物品數量**
#### 請求範例:
```plaintext
GET: api?cmd=1004&uid=1&sign_key=123456&all=0&id=22&num=999
```

___

### 取得聖遺物 cmd 1005
- **all**: `bool` 可選 | **是否取得全部物品 | 0:false|1:true**
- **id**: `uint32` 必須 | **聖遺物id**
- **num**: `uint32` 必須 | **聖遺物數量**
- **main**: `uint32` 可選 | **指定聖遺物主屬性**
- **sub**: `string` 可選 | **指定聖遺物副屬性**
#### 請求範例:
```plaintext
GET: api?cmd=1005&uid=1&sign_key=123456&all=0&id=31011&num=1&main=1&sub=[2:10][3:9][4:8]
```

___