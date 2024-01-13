![hkrpg-go](https://socialify.git.ci/gucooing/hkrpg-go/image?description=1&font=Inter&forks=1&language=1&name=1&owner=1&pattern=Circuit%20Board&stargazers=1&theme=Auto)

**注意:** 如果需要支持，请加入我们的 [Discord](https://discord.gg/ZJGTU8ZFGW).

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
3. 下载resources
4. 启动mysql并创建名为`hkrpg-go`的数据库，首次运行时会自动初始化所需要的表
5. 首次运行服务端后会生成一个名为config.json的文件，请修改`MysqlDsn`中password的值以及`GmKey`的值
6. 再次运行`main.exe`即可启动服务器

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
        if (oS.host.EndsWith("host自己找，一共四个")) {
            oS.host = "localhost:8080"; // 这也可以替换为其他IP地址。
        }
    }
};
```
4. 使用任意用户名及密码即可登录，账号默认会自动创建。

### gm功能使用方法：
- 自己看代码去❤️看不懂不怪我
### 注：
* 如果你想帮助此项目，欢迎提交
