![hkrpg-go](https://socialify.git.ci/gucooing/hkrpg-go/image?description=1&font=Inter&forks=1&language=1&name=1&owner=1&pattern=Circuit%20Board&starzers=1&theme=Auto)

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
 
 # **歡迎來到我們的 [Discord](https://discord.gg/222yVp6pUq)**

 ## 有朋友捲土重來，請關注 [hk4e-dmca](https://github.com/flswld/hk4e-go)

 ## 請不要在生產環境中使用

 ### 文件:
 * [簡易教學](./docs/tutorial/zh-cn.md)
 * [配置解析](./docs/conf/zh-CN.md)
 * [api使用](./docs/command/zh-CN.md)
 * [詳情](./docs/progress/zh-CN.md)

 ### 注意：
 * 如果您想幫助該項目，請隨時提交。

 ### 已完成的內容
 - **背包**
 - **戰鬥**
 - **抽卡**
 - **陣型**
 - **郵件**
 - **朋友們**
 - **場景道具/怪物/NPC生成** - **場景**
 - **陰謀**
 - **遺忘法庭和其他衍生性商品**
 - **模擬宇宙**
 - **微分宇宙**
 - **定期時間傳輸（部分**

 ### 與客戶端連線（Fiddler）
 1.安裝並執行[Fiddler Classic](https://www.telerik.com/fiddler)。
 2. 設定 Fiddler 解密 https 流量（工具 -> 選項 -> HTTPS -> 解密 HTTPS 流量），確保勾選「忽略伺服器憑證錯誤」。
 3. 複製以下程式碼並貼上到 Fiddler Classic 的 Fiddlerscript 標籤中：

 ```javascript
 導入系統； 導入系統。
 導入 System.Windows.Forms；
 導入提琴手； 導入系統。
 導入 System.Windows.Forms； 導入提琴手； 導入系統.文字。
 導入 System.Text.RegularExpressions； 類別處理程序
 靜態函數 OnBeforeRequest(oS: Session) {
     靜態函數 OnBeforeRequest(oS: Session) {
     如果（
         oS.host.EndsWith(".yuanshen.com") ||
         oS.host.EndsWith(".hoyoverse.com") ||
         oS.host.EndsWith(".mihoyo.com") ||
         oS.host.EndsWith(".zenlesszonezero.com") ||
         oS.host.EndsWith(".honkaiimpact3.com") ||
         oS.host.EndsWith(".bhsr.com") ||
         oS.host.EndsWith(".starrails.com") ||
         oS.uriContains("http://overseauspider.yuanshen.com:8888/log")
     ）{
         var newUrl = "http://" + oS.host + oS.PathAndQuery;
         oS.fullUrl = newUrl;
         oS.host = "127.0.0.1:8080";
     }
 }
 };
 ````

 4. 使用您的帳戶名稱登錄，密碼可以設定為任意值。