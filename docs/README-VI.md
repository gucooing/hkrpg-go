![hkrpg-go](https://socialify.git.ci/gucooing/hkrpg-go/image?description=1&font=Inter&forks=1&language=1&name=1&owner=1&pattern=Circuit%20Board&stargazers=1&theme=Auto)

<div align="center">
<table>
<td valign="center"><a href="README.md"><img src="https://github.com/twitter/twemoji/blob/master/assets/svg/1f1fa-1f1f8.svg" width="16"/> English</td>
 
<td valign="center"><a href="README_zh-CN.md"><img src="https://em-content.zobj.net/thumbs/120/twitter/351/flag-china_1f1e8-1f1f3.png" width="16"/> 简中</td>
 
<td valign="center"><a href="README_zh-TW.md"><img src="https://em-content.zobj.net/thumbs/120/twitter/351/flag-china_1f1e8-1f1f3.png" width="16"/> 繁中</td>
 
<td valign="center"><a href="README-JP.md"><img src="https://github.com/twitter/twemoji/blob/master/assets/svg/1f1ef-1f1f5.svg" width="16"/> 日本語</td>
 
<td valign="center"><a href="README-RU.md"><img src="https://github.com/twitter/twemoji/blob/master/assets/svg/1f1f7-1f1fa.svg" width="16"/> Русский</a></td>

<td valign="center"><a href="README-FR.md"><img src="https://em-content.zobj.net/thumbs/160/twitter/154/flag-for-france_1f1eb-1f1f7.png" width="16"/> Français</td>
 
<td valign="center"><a href="README-KR.md"><img src="https://em-content.zobj.net/source/twitter/53/flag-for-south-korea_1f1f0-1f1f7.png" width="16"/> 한국어</td>
 
<td valign="center"><a href="README-VI.md"><img src="https://em-content.zobj.net/thumbs/120/twitter/351/flag-vietnam_1f1fb-1f1f3.png" width="16"/> Tiếng Việt </a>
</td>
</table>
</div>

# **Chào mừng đến với [Discord](https://discord.gg/222yVp6pUq)** của chúng tôi

## Một người bạn sắp trở lại, hãy chú ý theo dõi [hk4e-dmca](https://github.com/flswld/hk4e-go)

## Vui lòng không sử dụng nó trong môi trường sản xuất

### Tài liệu:
* [easy-tutorial](./docs/tutorial/zh-cn.md)
* [config parsing](./docs/conf/zh-CN.md)
* [api usage](./docs/command/zh-CN.md)
* [Details](./docs/progress/zh-CN.md)

### Lưu ý:
* Nếu bạn muốn giúp đỡ dự án này, hãy gửi dự án.

 ### Nội dung đã hoàn thành
- **Ba lô**
- **Trận chiến**
- **Rút bài**
- **Đội hình**
- **Thư**
- **Bạn bè**
- **Đạo cụ cảnh/Quái vật/Thế hệ NPC** - **Kịch bản**
- **Cốt truyện**
- **Tòa án lãng quên và các phần phụ khác**
- **Vũ trụ mô phỏng**
- **Vũ trụ vi phân**
- **Truyền thời gian thông thường (Một phần**

### Kết nối với máy khách (Fiddler)
1. Cài đặt và chạy [Fiddler Classic](https://www.telerik.com/fiddler).
2. Đặt Fiddler để giải mã lưu lượng https (Công cụ -> Tùy chọn -> HTTPS -> Giải mã lưu lượng HTTPS), đảm bảo rằng `Bỏ qua lỗi chứng chỉ máy chủ` được chọn.
3. Sao chép và dán mã sau vào tab Fiddlerscript của Fiddler Classic:

```javascript
import System; import System.
import System.Windows.Forms;
import  Fiddler; nhập System.
nhập System.Windows.Forms; nhập Fiddler; nhập System.Text.
nhập System.Text.RegularExpressions; lớp Handlers
hàm tĩnh OnBeforeRequest(oS: Phiên) {
hàm tĩnh OnBeforeRequest(oS: Phiên) {
if(
oS.host.EndsWith(".yuanshen.com") ||
oS.host.EndsWith(".hoyoverse.com") ||
oS.host.EndsWith(".mihoyo.com") ||
oS.host.EndsWith(".zenlesszonezero.com") ||
oS.host.EndsWith(".honkaiimpact3.com") ||
oS.host.EndsWith(".bhsr.com") ||
oS.host.EndsWith(".starrails.com") ||
 oS.uriContains("http://overseauspider.yuanshen.com:8888/log")
) {
var newUrl = "http://" + oS.host + oS.PathAndQuery;
oS.fullUrl = newUrl;
oS.host = "127.0.0.1:8080";
}
}
};
```

4. Đăng nhập bằng tên tài khoản của bạn, mật khẩu có thể được đặt thành bất kỳ giá trị nào..1