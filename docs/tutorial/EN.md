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

# Prepare the environment
1. golang >= 1.22.4
2. mysql
3. redis
4. bash (required only when using build.sh)

## Compile
> Note: It is recommended to compile on the running server by yourself, otherwise unexpected situations may occur
1. Install dependencies
`go mod tidy`

2. Start compiling

#### Compile by yourself
- Install golang and the version should not be lower than 1.22.4
- Install gcc environment under linux, and then execute
```bash
bash ./build.sh
```

- Execute under windows
```bash
.\build.bat
```

- After the script is run, you can see the compiled executable file in the build folder

### Don't want to compile
Go to [Build-dev](https://github.com/gucooing/hkrpg-go/actions/workflows/Build.yml) to download

## Run
### 1. Prepare resources:
data resources, data can use the data in the warehouse, but the resource folder needs to be given read and write permissions

Preparation of resources:
1. Download [StarRailData](https://github.com/Dimbreath/StarRailData)

2. Download the supplementary file (task file) [DanhengServer-Resources](https://github.com/EggLinks/DanhengServer-Resources)

3. Unzip StarRailData to resources first, and then overwrite it with DanhengServer-Resources once (only overwrite and update Config, do not overwrite ExcelOutput is incompatible)

### 2. Run,
When running, you need to carry the startup parameter -i appid, where the appid format is ipv4 format, such as: 9001.1.1.1 Meaning:
```bash
9001: zone id;
1: service id;
1: host id;
1: This time start the service id;
```
After understanding the composition and meaning of appid, you can start it once without parameters to generate configuration files for each service. The generated configuration files are in the conf folder, and then change the appid in the default configuration file according to your own defined appid (although the service uses discovery to add new services, it is still recommended that the appid configuration table in each configuration file is the same), and then change other parameters in the configuration file according to your own ideas.

### 3. Database preparation,
Install mysql, create a new database in mysql: hkrpg-go-account && hkrpg-go-user && hkrpg-go-player && hkrpg-go-conf (utf8mb4), then change the account and password in the configuration file, install redis, and change the password in the configuration file (this service can be divided into tables and databases, but the same table must be in the same database)

### 4. Start,
The preliminary preparations have been completed and it is time to start. The recommended startup sequence is:
> The startup method in the following example is the startup parameter of the default configuration file

```bash
./nodeserver -i 9001.3.1.1
./gameserver -i 9001.2.1.1
./gateserver -i 9001.1.1.1
./dispatch -i 9001.4.1.1
./multiserver -i 9001.5.1.1
./muipserver -i 9001.6.1.1
```

## Each service function
### nodeserver node server (stateful, non-clusterable), service discovery, service management

### dispatch login server (stateless, clusterable)

### gateserver gateway server (stateful, clusterable), the only interaction interface between the internal network and the outside world

### gameserver logic server (stateful, clusterable), processing business logic

### multiserver multi-player server (stateful, non-clusterable) useless services

### muipserver is currently only responsible for api

## Advanced operation

### Multi-gateserver, multi-gameserver deployment
Take gateserver as an example. By default, there is only one 9001.1.1.1 configuration. You can add a 9001.1.1.2 configuration. You can use the same executable file when starting. The first gateserver is started with -i 9001.1.1.1, and the second one is started with -i 9001.1.1.2, and start it. If you are on the same machine, make sure that the two configured ports do not conflict.

And so on.........

## Notes
Please handle the internal and external networks properly, and do not allow the external network to access the internal network of the cluster at will.
If your external network bandwidth is less than 1Gpbs/s and the delay is not less than 10ms, please do not use the external network database.

## Want to test but do not want to configure a complex environment

1. Go to [Build-dev](https://github.com/gucooing/hkrpg-go/actions/workflows/Build.yml) to download the hkrpg-pe executable file