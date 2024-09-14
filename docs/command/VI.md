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