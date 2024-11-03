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
- 示例参数均以 uid = 1 | sign_key = 123456 | region = hkrpg_rel进行

### 路径参数
- **cmd**: `int16` 必须 | **调用指令**
- **uid**: `uint32` 必须 | **玩家UID**
- **region**: `string` 必须 | **区服**
- **sign_key**: `string` 可选 | **key**

#### 请求示例:
```plaintext
GET: api?region=hkrpg_rel&cmd=1&uid=1&sign_key=123456
```

#### 回调解析`json`:
- **code**:  状态 0 成功 -1 失败
- **msg**:  回调内容

___

### 设置世界等级 cmd 1001
**参数**:
- **world_level**: `uint32` 必须 | **要设置的世界等级**
#### 请求示例:
```plaintext
GET: api?region=hkrpg_rel&cmd=1001&uid=1&sign_key=123456&world_level=6
```

___

### 获取账号数据 cmd 1002
**参数**:
- **bin**: `uint32` 可选 | **1:二进制**
#### 请求示例:
```plaintext
GET: api?region=hkrpg_rel&cmd=1002&uid=1&sign_key=123456
```

___

### 获取服务器状态 cmd 1003
**参数**:无
#### 请求示例:
```plaintext
GET: api?region=hkrpg_rel&cmd=1003&sign_key=123456&bin=1
```

___

### 获取道具 cmd 1004
- **all**: `bool` 可选 | **是否获取全部物品 | 0:false|1:true**
- **id**: `uint32` 可选 | **物品id**
- **num**: `uint32` 可选 | **物品数量**
#### 请求示例:
```plaintext
GET: api?region=hkrpg_rel&cmd=1004&uid=1&sign_key=123456&all=0&id=22&num=999
```

___

### 获取圣遗物 cmd 1005
- **all**: `bool` 可选 | **是否获取全部物品 | 0:false|1:true**
- **id**: `uint32` 必须 | **圣遗物id**
- **num**: `uint32` 必须 | **圣遗物数量**
- **main**: `uint32` 可选 | **指定圣遗物主属性**
- **sub**: `string` 可选 | **指定圣遗物副属性**
#### 请求示例:
```plaintext
GET: api?region=hkrpg_rel&cmd=1005&uid=1&sign_key=123456&all=0&id=31011&num=1&main=1&sub=[2:10][3:9][4:8]
```

___

###  修改是否跳过剧情 cmd 1007
- **is**: `bool` 可选 | **修改是否跳过剧情 | 0:false|1:true**
#### 请求示例:
```plaintext
GET: api?region=hkrpg_rel&cmd=1007&uid=1&sign_key=123456&is=1
```

> 需要注意的是，如果你开启了一次跳过剧情，你的所有教程会直接被完成，并且后续无法触发，这会造成未知后果，包括但不限于某些任务无法正常进行，某些事件无法正常触发

___