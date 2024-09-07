[EN](./EN.md) | [简中](./zh-CN.md) | [繁中](./zh-TW.md) | [JP](./JP.md) | [RU](./RU.md) | [FR](./FR.md) | [KR](./KR.md) | [VI](./VI.md)

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