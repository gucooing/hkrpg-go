![hkrpg-go](https://socialify.git.ci/gucooing/hkrpg-go/image?description=1&font=Inter&forks=1&language=1&name=1&owner=1&pattern=Circuit%20Board&stargazers=1&theme=Auto)

<div align="center">
<table>
<td valign="center"><a href="README.md"><img src="https://github.com/twitter/twemoji/blob/master/assets/svg/1f1fa-1f1f8.svg" width="16"/> English</td>
 
<td valign="center"><a href="README_zh-cn.md"><img src="https://em-content.zobj.net/thumbs/120/twitter/351/flag-china_1f1e8-1f1f3.png" width="16"/> 简中</td>
</a></td>
</table>
</div>>

# **欢迎加入我们的 [Discord](https://discord.gg/222yVp6pUq)**

## 朋友要复出了，可以关注一下 [hk4e-dmca](https://github.com/flswld/hk4e-go)

# 请不要在生产环境中使用

### 文档：
* [简易教程](https://github.com/gucooing/hkrpg-go/wiki/tutorial_zh%E2%80%90cn)
* [config解析](https://github.com/gucooing/hkrpg-go/wiki/conf_zh%E2%80%90cn)
* [api使用方法](https://github.com/gucooing/hkrpg-go/wiki/command_zh%E2%80%90cn)
* [详细完善内容](https://github.com/gucooing/hkrpg-go/wiki/progress_zh%E2%80%90cn)

### 注：
* 如果你想帮助此项目，欢迎提交

### 已完成内容
- **背包**
- **战斗**
- **抽卡**
- **编队**
- **邮件**
- **好友**
- **场景道具/怪物/NPC生成**
- **剧情**
- **忘却之庭等衍生内容**
- **模拟宇宙**
- **差分宇宙**
- **常时传略(部分**

### 与客户端（Fiddler）连接
1. 安装并运行 [Fiddler Classic](https://www.telerik.com/fiddler)。
2. 将Fiddler设置为解密https流量（工具 -> 选项 -> HTTPS -> 解密HTTPS流量），确保选中 `忽略服务器证书错误`。
3. 将以下代码复制并粘贴到Fiddler Classic的Fiddlerscript选项卡中：

```javascript
import System;
import System.Windows.Forms;
import Fiddler;
import System.Text.RegularExpressions;
class Handlers
{
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

4. 使用您的帐户名称登录，密码可以设置为任何值。
