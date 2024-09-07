[EN](./EN.md) | [简中](./zh-CN.md) | [繁中](./zh-TW.md) | [JP](./JP.md) | [RU](./RU.md) | [FR](./FR.md) | [KR](./KR.md) | [VI](./VI.md)

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