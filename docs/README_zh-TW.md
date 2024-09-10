![hkrpg-go](https://socialify.git.ci/gucooing/hkrpg-go/image?description=1&font=Inter&forks=1&language=1&name=1&owner=1&pattern=Circuit%20Board&starzers=1&theme=Auto)

 [EN](./README.md) |  [簡中](./docs/README_zh-CN.md) |  [繁中](./docs/README_zh-CN.md) |  [JP](./docs/README_zh-CN.md) |  [RU](./docs/README_zh-CN.md) |  [FR](./docs/README_zh-CN.md) |  [KR](./docs/README_zh-CN.md) |  [KI](./docs/README_zh-CN.md)
 
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