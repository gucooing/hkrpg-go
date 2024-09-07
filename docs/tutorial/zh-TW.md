[EN](./EN.md) | [簡中](./zh-CN.md) | [繁中](./zh-TW.md) | [JP](./JP.md) | [RU](./RU.md) | [FR](./FR.md) | [KR](./KR.md) | [VI](./VI.md)

# 準備環境
1. golang >= 1.22.4
2. mysql
3. redis
4. bash(使用build.sh時才需要)

## 編譯
> 註:建議自行在執行伺服器上進行編譯,否則可能出現意外狀況
1. 安裝依賴
`go mod tidy`

2. 開始編譯

#### 自行編譯
- 安裝golang且版本不低於1.22.4
- linux下安裝gcc環境，然後執行

```bash
bash ./build.sh
```

- windows下執行
```bash
.\build.bat
```

- 腳本運行完畢後可在build資料夾中看到編譯後的可執行文件

### 不想編譯
前往[Build-dev](https://github.com/gucooing/hkrpg-go/actions/workflows/Build.yml)下載

## 運行

### 1.準備資源：
data resources，data使用倉庫的data即可，但資源資料夾需要給予讀寫權限。

resources的準備:
1. 下載[StarRailData](https://github.com/Dimbreath/StarRailData)

2. 下載補充文件（任務文件）[DanhengServer-Resources](https://github.com/EggLinks/DanhengServer-Resources)

3. 先將StarRailData解壓到resources中，然後用DanhengServer-Resources覆寫一次(只覆寫更新Config即可，不要覆寫ExcelOutput不相容)

### 2.運行，
運行時需攜帶啟動參數 -i appid ， 其中appid格式為ipv4格式，如：9001.1.1.1 其中意義：

```bash
9001: 區服id;
1: 服務id;
1: 主機id;
1: 本次啟動服務id;
```
了解到了appid的組成含義後你可以先不攜帶參數啟動一次，使其生成各個服務的配置文件，生成的配置文件在conf文件夾裡，然後根據你自己定義的appid更改默認配置文件中的appid(雖然服務採用發現形式新增服務，但還是推薦每一個設定檔中的appid配置表都相同)，然後根據自己的想法更改設定檔中的其他參數。

### 3.資料庫的準備，
安裝mysql，mysql中新建資料庫：hkrpg-go-account && hkrpg-go-user && hkrpg-go-player && hkrpg-go-conf (utf8mb4),然後更改設定檔中的帳號和密碼，安裝redis，變更設定文件中的密碼（本服務可採用分錶分庫形式，但同一張表一定要是同一個資料庫）

### 4.啟動，
前期的準備工作已經全部完成了到了啟動的時候了，建議的啟動順序為：
> 以下範例的啟動方法為預設設定檔的啟動參數

```bash
./nodeserver -i 9001.3.1.1
./gameserver -i 9001.2.1.1
./gateserver -i 9001.1.1.1
./dispatch -i 9001.4.1.1
./multiserver -i 9001.5.1.1
./muipserver -i 9001.6.1.1
```

## 各服務功能
### nodeserver 節點伺服器（有狀態，不可叢集），服務發現，服務管理

### dispatch 登入伺服器（無狀態，可叢集）

### gateserver 閘道伺服器（有狀態，可叢集），內部網路與外界唯一交互口

### gameserver 邏輯伺服器（有狀態，可叢集），處理業務邏輯

### multiserver 多人伺服器（有狀態，不可叢集）沒有用的服務

### muipserver 目前只負責api

## 進階操作

### 多gateserver、多gameserver部署
以gateserver為例，預設只有一個9001.1.1.1的配置，可新增一個9001.1.1.2的配置，啟動時可以使用同一個可執行文件，第一個gateserver使用-i 9001.1.1.1啟動，第二個使用- i 9001.1.1.2 啟動即可，如果在同一台機器注意兩個配置的連接埠不要衝突了

等等.........

## 注意事項
請處理好內外網，不要讓外網可隨意存取到群集內部網絡
如果你的外網頻寬不足 1Gpbs/s 延遲不低於10ms 請不要使用外網資料庫

## 想測試但不想配置複雜的環境

1.前往[Build-dev](https://github.com/gucooing/hkrpg-go/actions/workflows/Build.yml)下載hkrpg-pe執行檔