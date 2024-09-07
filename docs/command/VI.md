[EN](./EN.md) | [简中](./zh-CN.md) | [繁中](./zh-TW.md) | [JP](./JP.md) | [RU](./RU.md) | [FR](./FR.md) | [KR](./KR.md) | [VI](./VI.md)

#API
- Các tham số mẫu đều dựa trên uid = 1 |

### Tham số đường dẫn
- **cmd**: `int16` bắt buộc | **Gọi lệnh**
- **uid**: bắt buộc phải có `uint32` | **UID của người chơi**
- **sign_key**: `string` tùy chọn **key**

#### Ví dụ về yêu cầu:
``` bản rõ
NHẬN: api?cmd=1&uid=1&sign_key=123456
```

#### Phân tích cú pháp gọi lại `json`:
- **code**: trạng thái 0 thành công -1 thất bại
- **tin nhắn**: nội dung gọi lại

___

### Đặt cấp độ thế giới cmd 1001
**tham số**:
- **world_level**: `uint32` bắt buộc | **Cấp độ thế giới được thiết lập**
#### Ví dụ về yêu cầu:
``` bản rõ
NHẬN: api?cmd=1001&uid=1&sign_key=123456&world_level=6
```

___

### Lấy dữ liệu tài khoản cmd 1002
**Thông số**: Không có
#### Ví dụ về yêu cầu:
``` bản rõ
NHẬN: api?cmd=1002&uid=1&sign_key=123456
```

___

### Lấy trạng thái máy chủ cmd 1003
**Thông số**: Không có
#### Ví dụ về yêu cầu:
``` bản rõ
NHẬN: api?cmd=1003&sign_key=123456
```

___

### Lấy đạo cụ cmd 1004
- **all**: `bool` tùy chọn | **Có lấy tất cả các mục hay không | 0:false|1:true**
- **id**: `uint32` tùy chọn | id mục**
- **num**: `uint32` tùy chọn | **số lượng mục**
#### Ví dụ về yêu cầu:
``` bản rõ
NHẬN: api?cmd=1004&uid=1&sign_key=123456&all=0&id=22&num=999
```

___

### Nhận thánh tích cmd 1005
- **all**: `bool` tùy chọn | **Có lấy tất cả các mục hay không | 0:false|1:true**
- **id**: `uint32` bắt buộc | **ID thánh tích**
- **num**: `uint32` bắt buộc | **Số lượng thánh tích**
- **chính**: `uint32` tùy chọn | **Chỉ định thuộc tính chính của thánh tích**
- **sub**: `string` tùy chọn | **Chỉ định các thuộc tính phụ của thánh tích**
#### Ví dụ về yêu cầu:
``` bản rõ
NHẬN: api?cmd=1005&uid=1&sign_key=123456&all=0&id=31011&num=1&main=1&sub=[2:10][3:9][4:8]
```

___