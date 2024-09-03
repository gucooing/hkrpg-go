![hkrpg-go](https://socialify.git.ci/gucooing/hkrpg-go/image?description=1&font=Inter&forks=1&language=1&name=1&owner=1&pattern=Circuit%20Board&stargazers=1&theme=Auto)

[EN](./README.md) | [简中](./docs/README_zh-CN.md) | [繁中](./docs/README_zh-CN.md) | [JP](./docs/README_zh-CN.md) | [RU](./docs/README_zh-CN.md) | [FR](./docs/README_zh-CN.md) | [KR](./docs/README_zh-CN.md) | [VI](./docs/README_zh-CN.md)
 
# **Welcome to our [Discord](https://discord.gg/222yVp6pUq)**

## A friend is making a comeback, keep an eye out [hk4e-dmca](https://github.com/flswld/hk4e-go)

## Please don't use it in a production environment

### Documentation:
* [easy-tutorial](./docs/tutorial/zh-cn.md)
* [config parsing](./docs/conf/zh-CN.md)
* [api usage](./docs/command/zh-CN.md)
* [Details](./docs/progress/zh-CN.md)

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