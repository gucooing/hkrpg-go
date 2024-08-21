[EN](./zh-CN.md) | [简中](./zh-CN.md) | [繁中](./zh-CN.md) | [JP](./zh-CN.md) | [RU](./zh-CN.md) | [FR](./zh-CN.md) | [KR](./zh-CN.md) | [VI](./zh-CN.md)

# API

### 路径参数
- **cmd**: `int16` 必须
- **uid**: `uint32` 必须
- **sign_key**: `string` 可选

#### 请求示例:
```plaintext
GET: api?cmd=1&uid=1
```

#### 回调解析`json`:
- **code**:  状态 0 成功 -1 失败
- **msg**:  回调内容

___

### 设置世界等级 cmd 1001
**参数**:
- **world_level**: `uint32` 必须

___

### 获取账号数据 cmd 1002
**参数**:无

___

### 获取服务器状态 cmd 1003
**参数**:无