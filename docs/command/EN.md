[EN](./EN.md) | [简中](./zh-CN.md) | [繁中](./zh-TW.md) | [JP](./JP.md) | [RU](./RU.md) | [FR](./FR.md) | [KR](./KR.md) | [VI](./VI.md)

# API
- All example parameters are based on uid = 1 | sign_key = 123456

### Path parameters
- **cmd**: `int16` required | **Call command**
- **uid**: `uint32` required | **Player UID**
- **sign_key**: `string` optional | **key**

#### Request example:
```plaintext
GET: api?cmd=1&uid=1&sign_key=123456
```

#### callback parsing `json`:
- **code**: status 0 success -1 failure
- **msg**: callback content

___

### set world level cmd 1001
**parameter**:
- **world_level**: `uint32` required | **world level to be set**
#### request example:
```plaintext
GET: api?cmd=1001&uid=1&sign_key=123456&world_level=6
```

___

### get account data cmd 1002
**parameter**: none
#### request example:
```plaintext
GET: api?cmd=1002&uid=1&sign_key=123456
```

___

### Get server status cmd 1003
**Parameter**: None
#### Request example:
```plaintext
GET: api?cmd=1003&sign_key=123456
```

___

### Get props cmd 1004
- **all**: `bool` Optional | **Whether to get all items | 0:false|1:true**
- **id**: `uint32` Optional | **Item id**
- **num**: `uint32` Optional | **Item quantity**
#### Request example:
```plaintext
GET: api?cmd=1004&uid=1&sign_key=123456&all=0&id=22&num=999
```

___

### Get holy relics cmd 1005
- **all**: `bool` optional | **Whether to get all items | 0:false|1:true**
- **id**: `uint32` required | **Holy relic id**
- **num**: `uint32` required | **Number of holy relics**
- **main**: `uint32` optional | **Specify the main attribute of the holy relic**
- **sub**: `string` optional | **Specify the secondary attribute of the holy relic**
#### Request example:
```plaintext
GET: api?cmd=1005&uid=1&sign_key=123456&all=0&id=31011&num=1&main=1&sub=[2:10][3:9][4:8] ``` ___