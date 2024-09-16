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

# Giải thích chi tiết các tham số conf

### LogLevel: Cấp độ nhật ký

### MaxPlayer: Số lượng người chơi tối đa trong gameserver này

### AutoCreate: Có đăng ký tự động không

###Danh sách ứng dụng:
 Định dạng: appid[cấu hình]
 port_player: dùng cho cổng server, đại diện cho cổng kết nối bên ngoài, tức là cổng kcp
 port_gt: dùng cho gameserver, đại diện cho port dùng để kết nối với Gateserver
 port_service: đại diện cho cổng mà dịch vụ này chấp nhận kết nối từ các dịch vụ khác.
 port_http: Khởi động máy chủ http trên cổng này

### Ứng dụng:
 Cổng: cổng nghe
 InnerAddr: địa chỉ bên ngoài
 OuterAddr: địa chỉ nghe

### NetConf:
 Định dạng: dịch vụ[địa chỉ]
 Nút: Đại diện cho địa chỉ kết nối của máy chủ nút

### MysqlConf:
 Định dạng: bảng[địa chỉ]
 Mỗi địa chỉ kết nối mysql được yêu cầu

### RedisConf:
 Định dạng: bảng [cấu hình kết nối]
 Mỗi cấu hình kết nối redis cần thiết

### Công văn: (chỉ công văn)
 Định dạng: [] Địa chỉ máy chủ của mỗi nút
 Khi có nhiều máy chủ không liên lạc với nhau, hãy sửa đổi cấu hình này để kết nối

### GameDataConfigPath: (chỉ dành cho gameserver) điền đường dẫn bảng cấu hình (dựa trên thư mục khởi động chương trình)