![hkrpg-go](https://socialify.git.ci/gucooing/hkrpg-go/image?description=1&font=Inter&forks=1&language=1&name=1&owner=1&pattern=Circuit%20Board&stargazers=1&theme=Auto)

<div align="center">
<table>
<td valign="center"><a href="README.md"><img src="https://github.com/twitter/twemoji/blob/master/assets/svg/1f1fa-1f1f8.svg" width="16"/> English</td>
 
<td valign="center"><a href="README_zh-cn.md"><img src="https://em-content.zobj.net/thumbs/120/twitter/351/flag-china_1f1e8-1f1f3.png" width="16"/> 简中</td>
</a></td>
</table>
</div>
 
# **Welcome to our [Discord](https://discord.gg/222yVp6pUq)**

## A friend is making a comeback, keep an eye out [hk4e-dmca](https://github.com/flswld/hk4e-go)

## Please don't use it in a production environment

### Documentation:
* [easy-tutorial](https://github.com/gucooing/hkrpg-go/wiki/tutorial)
* [config parsing](https://github.com/gucooing/hkrpg-go/wiki/conf)
* [api usage](https://github.com/gucooing/hkrpg-go/wiki/command)
* [Details](https://github.com/gucooing/hkrpg-go/wiki/progress)

### Notes:
* If you want to help with this project, feel free to submit it.

### Completed content
- **Backpacks**
- **Battle**
- **Card draw**
- **Formation**
- **Mail**
- **Friends**
- **Scene Props/Monster/NPC Generation** - **Scenario**
- **Plot**
- **The Court of Forgetfulness and other spin-offs**
- **Simulated Universe**
- **Differential Universe**
- **Regular Time Transmission (Partial**

### Connecting with the client (Fiddler)
1. Install and run [Fiddler Classic](https://www.telerik.com/fiddler).
2. Set Fiddler to decrypt https traffic (Tools -> Options -> HTTPS -> Decrypt HTTPS Traffic), make sure `Ignore server certificate errors` is checked.
3. Copy and paste the following code into the Fiddlerscript tab of Fiddler Classic:

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

4. Log in using your account name, the password can be set to any value.
