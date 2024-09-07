[EN](./EN.md) | [简中](./zh-CN.md) | [繁中](./zh-TW.md) | [JP](./JP.md) | [RU](./RU.md) | [FR](./FR.md) | [KR](./KR.md) | [VI](./VI.md)

#Chuẩn bị môi trường
1. golang >= 1.22.4
2.mysql
3. làm lại
4. bash (bắt buộc khi sử dụng build.sh)

## Biên dịch
> Lưu ý: Nên tự biên dịch trên máy chủ đang chạy, nếu không có thể xảy ra những tình huống không mong muốn.
1. Cài đặt phụ thuộc
`đi mod gọn gàng`
2. Bắt đầu biên dịch
####Tự biên dịch
- Cài đặt golang và phiên bản không dưới 1.22.4
- Cài đặt môi trường gcc trên linux rồi thực thi

``` bash
bash ./build.sh
```

- Thực thi dưới cửa sổ
``` bash
.\build.bat
```

- Sau khi script chạy xong, bạn có thể thấy file thực thi đã biên dịch trong thư mục build

### Không muốn biên dịch
Truy cập [Build-dev](https://github.com/gucooing/hkrpg-go/actions/workflows/Build.yml) để tải xuống

## Chạy
### 1. Chuẩn bị nguồn lực:
Tài nguyên dữ liệu, dữ liệu có thể sử dụng dữ liệu của kho nhưng thư mục tài nguyên cần được cấp quyền đọc và ghi.

Chuẩn bị nguồn lực:
1. Tải xuống [StarRailData](https://github.com/Dimbreath/StarRailData)
2. Tải file bổ sung (tệp tác vụ) [DanhengServer-Resources](https://github.com/EggLinks/DanhengServer-Resources)
3. Đầu tiên giải nén StarRailData thành tài nguyên, sau đó ghi đè một lần bằng DanhengServer-Resources (chỉ che phần cập nhật Config, không ghi đè lên phần không tương thích ExcelOutput)
### 2. Chạy đi,
Tham số khởi động -i appid cần được mang theo khi chạy, trong đó định dạng appid là định dạng ipv4, chẳng hạn như: 9001.1.1.1, có nghĩa là:

``` bash
9001: Id máy chủ huyện;
1: id dịch vụ;
1: id máy chủ;
1: Id dịch vụ sẽ bắt đầu lần này;
```

Sau khi hiểu ý nghĩa thành phần của appid, bạn có thể khởi động nó mà không cần tham số để tạo tệp cấu hình của từng dịch vụ. Tệp cấu hình được tạo nằm trong thư mục conf, sau đó thay đổi appid trong tệp cấu hình mặc định theo định nghĩa của riêng bạn. appid (Mặc dù dịch vụ sử dụng tính năng khám phá để thêm dịch vụ mới nhưng vẫn khuyến nghị bảng cấu hình appid trong mỗi tệp cấu hình phải giống nhau), sau đó thay đổi các tham số khác trong tệp cấu hình theo ý tưởng của riêng bạn.

### 3. Chuẩn bị cơ sở dữ liệu,
Cài đặt mysql, tạo cơ sở dữ liệu mới trong mysql: hkrpg-go-account && hkrpg-go-user && hkrpg-go-player && hkrpg-go-conf (utf8mb4), sau đó thay đổi tài khoản và mật khẩu trong file cấu hình, cài đặt redis , và thay đổi mật khẩu cấu hình trong file (dịch vụ này có thể chia thành các bảng và cơ sở dữ liệu, nhưng cùng một bảng phải nằm trong cùng một cơ sở dữ liệu)

### 4. Bắt đầu,
Mọi công việc chuẩn bị sơ bộ đã hoàn tất và đã đến lúc bắt đầu. Trình tự khởi động được đề xuất là:
> Phương pháp khởi động trong ví dụ sau là các tham số khởi động của tệp cấu hình mặc định.

``` bash
./nodeserver -i 9001.3.1.1
./gameserver -i 9001.2.1.1
./gateserver -i 9001.1.1.1
./dispatch -i 9001.4.1.1
./multiserver -i 9001.5.1.1
./muipserver -i 9001.6.1.1
```

##Mỗi chức năng dịch vụ
### máy chủ nút máy chủ nút (có trạng thái, không thể phân cụm), khám phá dịch vụ, quản lý dịch vụ

### máy chủ đăng nhập gửi đi (không trạng thái, có thể phân cụm)

### Máy chủ cổng máy chủ cổng (trạng thái, có thể phân cụm), giao diện duy nhất để tương tác giữa mạng nội bộ và thế giới bên ngoài

### máy chủ logic máy chủ trò chơi (có trạng thái, có thể phân cụm), xử lý logic nghiệp vụ

### multiserver Multiplayer server (có trạng thái, không thể phân cụm) không có dịch vụ hữu ích

### muipserver hiện chỉ chịu trách nhiệm về api

## Các thao tác nâng cao
### Triển khai nhiều máy chủ, nhiều máy chủ trò chơi
Lấy máy chủ cổng làm ví dụ, theo mặc định chỉ có một cấu hình 9001.1.1.1. Bạn có thể thêm cấu hình 9001.1.1.2. Bạn có thể sử dụng cùng một tệp thực thi khi khởi động máy chủ cổng đầu tiên. và cái thứ hai được bắt đầu bằng - i 9001.1.1.2 có thể được khởi động. Nếu bạn ở trên cùng một máy, hãy cẩn thận để không xung đột với hai cổng được định cấu hình.

Chờ đợi.........

## Ghi chú
Hãy xử lý tốt mạng nội bộ và bên ngoài và không cho phép mạng bên ngoài tự do truy cập vào mạng nội bộ của cụm.
Nếu băng thông mạng bên ngoài của bạn nhỏ hơn 1Gpbs/s và độ trễ không nhỏ hơn 10ms, vui lòng không sử dụng cơ sở dữ liệu mạng bên ngoài.

## Muốn thử nghiệm nhưng không muốn định cấu hình môi trường phức tạp

1. Truy cập [Build-dev](https://github.com/gucooing/hkrpg-go/actions/workflows/Build.yml) để tải xuống tệp thực thi hkrpg-pe