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

# **[Discord](https://discord.gg/222yVp6pUq) へようこそ**

## 友人が復帰しています。[hk4e-dmca](https://github.com/flswld/hk4e-go) に注目してください

## 本番環境では使用しないでください

### ドキュメント:
* [easy-tutorial](./docs/tutorial/zh-cn.md)
* [config 解析](./docs/conf/zh-CN.md)
* [api の使用法](./docs/command/zh-CN.md)
* [詳細](./docs/progress/zh-CN.md)

### 注:
* このプロジェクトに協力したい場合は、お気軽に提出してください。

 ### 完了したコンテンツ
- **バックパック**
- **戦闘**
- **カード ドロー**
- **フォーメーション**
- **メール**
- **友達**
- **シーン プロップ/モンスター/NPC 生成** - **シナリオ**
- **プロット**
- **忘却の宮廷とその他のスピンオフ**
- **シミュレートされた宇宙**
- **差分宇宙**
- **定期的な時間送信 (部分的**

### クライアント (Fiddler) との接続
1. [Fiddler Classic](https://www.telerik.com/fiddler) をインストールして実行します。
2. Fiddler を https トラフィックの暗号化解除に設定し ([ツール] -> [オプション] -> [HTTPS] -> [HTTPS トラフィックの暗号化解除])、[サーバー証明書エラーを無視] がオンになっていることを確認します。
3. 次のコードをコピーして、Fiddler の Fiddlerscript タブに貼り付けます。 クラシック:

```javascript
import System; import System.
import System.Windows.Forms;
import Fiddler; import System.
import System.Windows.Forms; import Fiddler; import System.Text.
import System.Text.RegularExpressions; class Handlers
static function OnBeforeRequest(oS: Session) {
static function OnBeforeRequest(oS: Session) {
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

4. アカウント名を使用してログインします。パスワードは任意の値に設定できます。