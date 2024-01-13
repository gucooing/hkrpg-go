![hkrpg-go](https://socialify.git.ci/gucooing/hkrpg-go/image?description=1&font=Inter&forks=1&language=1&name=1&owner=1&pattern=Circuit%20Board&stargazers=1&theme=Auto)

### 已实现的功能：
- 走路模拟器
- 队伍管理
- 传送
- 抽卡模拟器
- 背包
- 初级战斗
- 养成
- 初级模拟宇宙（1.6.0不支持）
- 副本/周本/挑战
- 忘却之庭
- 随机首包密钥，随机生成seed，使数据更加安全
- 基本gm功能

### 可以正常游玩的活动：
- 角色试用
- 巡星之礼
- 忘却之庭-虚构叙事

### 必须条件：
- [mysql](https://dev.mysql.com/downloads/installer/)
- [go 1.16+](https://golang.google.cn/dl/)

### 使用方法：
1. 拉取项目文件
2. 打开终端，使用`go mod tidy&&go build main.go`编译服务端核心
3. 从 [StarRailData](https://github.com/Dimbreath/StarRailData) 下载 `Config`、`TextMap` 和 `ExcelBin` 文件夹，并将它们放入`resources`目录
4. 从 [LunarCore-Configs](https://gitlab.com/Melledy/LunarCore-Configs) 下载 `Config` 文件夹，并将其放入资源文件夹。替换系统询问的任何文件。这些文件用于世界生成，对服务器非常重要。
5. 启动mysql并创建名为`hkrpg-go`的数据库，首次运行时会自动初始化所需要的表
6. 首次运行服务端后会生成一个名为config.json的文件，请修改`MysqlDsn`中password的值以及`GmKey`的值
7. 再次运行`main.exe`即可启动服务器

### 与客户端（Fiddler）连接
1. **使用客户端至少一次登录到官方服务器和Hoyoverse账户以下载游戏数据。**
2. 安装并运行 [Fiddler Classic](https://www.telerik.com/fiddler)。
3. 将Fiddler设置为解密https流量（工具 -> 选项 -> HTTPS -> 解密HTTPS流量），确保选中 `忽略服务器证书错误`。
4. 将以下代码复制并粘贴到Fiddler Classic的Fiddlerscript选项卡中：
```javascript
import System;
import System.Windows.Forms;
import Fiddler;
import System.Text.RegularExpressions;

class Handlers
{
    static function OnBeforeRequest(oS: Session) {
        if (oS.host.EndsWith(".starrails.com") || oS.host.EndsWith(".hoyoverse.com") || oS.host.EndsWith(".mihoyo.com") || oS.host.EndsWith(".bhsr.com")) {
            oS.host = "localhost:8080"; // 这也可以替换为其他IP地址。
        }
    }
};
```
5. 使用任意用户名及密码即可登录，账号默认会自动创建。

### gm功能使用方法：
- 使用浏览器访问`https://ip:port/api?cmd={cmdId}&uid={uid}&key={GmKey}&{...otherParams}`
- 支持的cmdId请阅读代码
### 注：
* 如果你想帮助此项目，欢迎提交
