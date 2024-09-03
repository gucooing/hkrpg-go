![hkrpg-go](https://socialify.git.ci/gucooing/hkrpg-go/image?description=1&font=Inter&forks=1&language=1&name=1&owner=1&pattern=Circuit%20Board&stargazers=1&theme=Auto)

[EN](./README.md) | [简中](./docs/README_zh-CN.md) | [繁中](./docs/README_zh-CN.md) | [JP](./docs/README_zh-CN.md) | [RU](./docs/README_zh-CN.md) | [FR](./docs/README_zh-CN.md) | [KR](./docs/README_zh-CN.md) |  [VI](./docs/README_zh-CN.md)

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