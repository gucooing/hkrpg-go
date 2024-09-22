<div align="center">
<table>
<td valign="center"><a href="EN.md"><img src="https://github.com/twitter/twemoji/blob/master/assets/svg/1f1fa-1f1f8.svg" width="16"/> English</td>
 
<td valign="center"><a href="zh-CN.md"><img src="https://em-content.zobj.net/thumbs/120/twitter/351/flag-china_1f1e8-1f1f3.png" width="16"/> 简中</td>
 
<td valign="center"><a href="zh-TW.md"><img src="https://em-content.zobj.net/thumbs/120/twitter/351/flag-china_1f1e8-1f1f3.png" width="16"/> 繁中</td>
 
<td valign="center"><a href="JP.md"><img src="https://github.com/twitter/twemoji/blob/master/assets/svg/1f1ef-1f1f5.svg" width="16"/> 日本語</td>
 
<td valign="center"><a href="RU.md"><img src="https://github.com/twitter/twemoji/blob/master/assets/svg/1f1f7-1f1fa.svg" width="16"/> Русский</a></td>

<td valign="center"><a href="FR.md"><img src="https://em-content.zobj.net/thumbs/160/twitter/154/flag-for-france_1f1eb-1f1f7.png" width="16"/> Français</td>
 
<td valign="center"><a href="KR.md"><img src="https://em-content.zobj.net/source/twitter/53/flag-for-south-korea_1f1f0-1f1f7.png" width="16"/> 한국어</td>
 
<td valign="center"><a href="VI.md"><img src="https://em-content.zobj.net/thumbs/120/twitter/351/flag-vietnam_1f1fb-1f1f3.png" width="16"/> Tiếng Việt </a></td>
</table>
</div>

# 准备环境
1. golang >= 1.22.4
2. mysql
3. redis
~~4. linux:GCC / windows:MinGW~~

## 编译
> 注:建议自行在运行服务器上进行编译,否则可能出现意外情况

> 如果需要收集上报玩家信息请在编译时添加标签‘push‘，并在config中配置push server地址

1. 安装依赖
`go mod tidy`

2. 开始编译

#### 自行编译
1. 安装golang且版本不低于1.22.4
~~启用cgo~~
2. 运行编译脚本


linux:
```bash
bash ./build.sh
```

windows:
```bat
.\build.bat
```

- 脚本运行完毕后可在build文件夹中看到编译后的可执行文件
  
- 使用docker运行
  目前仅支持pe版本

  镜像名：gucooing/hkrpg-go-pe:latest

  默认端口：tcp 8080 udp 20041

  api暴露端口：tpc 20011

  容器目录/usr/hkrpg/conf为conf和数据库目录
  
  容器目录/usr/hkrpg/log为log目录

  容器启动时会自动下载resources，想使用自己的resources，可挂载/usr/hkrpg/resources
  
  推荐的[hkrpg-go-Resources](https://github.com/gucooing/hkrpg-go-Resources)

  如需拉取指定commit版本，可将标签改成此次commit的sha


### 不想编译
前往[Build-dev](https://github.com/gucooing/hkrpg-go/actions/workflows/Build.yml)下载

## 运行

### 1.准备资源：
data resources，data使用仓库的data即可，但资源文件夹需要给予读写权限。

resources的准备:
1. 下载[hkrpg-go-Resources](https://github.com/gucooing/hkrpg-go-Resources)

2. 将hkrpg-go-Resources解压到resources中即可

### 2.运行：
运行时需要携带启动参数 -i appid ， 其中appid格式为ipv4格式，如：9001.1.1.1 其中含义：

```bash
9001: 区服id;
1:    服务id; 
1:    主机id;
1:    本次启动服务id;
```
了解到了appid的组成含义后你可以先不携带参数启动一次，使其生成各个服务的配置文件，生成的配置文件在conf文件夹里，然后根据你自己定义的appid更改默认配置文件中的appid(虽然服务采用发现形式添加新服务，但是还是推荐每一个配置文件中的appid配置表都相同)，然后根据自己的想法更改配置文件中的其他参数。

### 3.数据库的准备：
  安装mysql，mysql中新建数据库：hkrpg-go-account && hkrpg-go-user && hkrpg-go-player && hkrpg-go-conf (utf8mb4),然后更改配置文件中的账户和密码，安装redis，更改配置文件中的密码（本服务可采用分表分库形式，但同一张表一定要是同一个数据库）

  在mysql数据库`hkrpg-go-conf`的表`region_conf`中配置相应的区服信息，默认的区服信息为`hkrpg_rel`

### 4.启动：
前期的准备工作已经全部完成了到了启动的时候了，推荐的启动顺序为：
> 下面示例的启动方法为默认配置文件的启动参数

```bash
./nodeserver -i 9001.3.1.1
./gameserver -i 9001.2.1.1
./gateserver -i 9001.1.1.1
./dispatch -i 9001.4.1.1
./muipserver -i 9001.6.1.1
```

## 各服务功能
### nodeserver 节点服务器（有状态，不可集群），服务发现，服务管理

### dispatch 登录服务器（无状态，可集群）

### gateserver 网关服务器（有状态，可集群），内部网络与外界唯一交互口

### gameserver 逻辑服务器（有状态，可集群），处理业务逻辑

### muipserver 目前仅负责api

### dispatch muipserver nodeserver不属于某一个区服，它们可以为全部区服提供服务

## 进阶操作
### 多gateserver、多gameserver部署
以gateserver为例，默认只有一个9001.1.1.1的配置，可添加一个9001.1.1.2的配置，启动时可以使用同一个可执行文件，第一个gateserver使用-i 9001.1.1.1启动，第二个使用 -i 9001.1.1.2 启动即可，如果在同一台机器注意两个配置的端口不要冲突了

等.........

## 注意事项
请处理好内外网，不要让外网可随意访问到集群内部网络
如果你的外网带宽不足 1Gpbs/s 延迟不低于10ms 请不要使用外网数据库
每次nodeserver重启时，区服 key都会强制刷新

## 想测试但不想配置复杂的环境

1.前往[Build-dev](https://github.com/gucooing/hkrpg-go/actions/workflows/Build.yml)下载hkrpg-pe可执行文件
