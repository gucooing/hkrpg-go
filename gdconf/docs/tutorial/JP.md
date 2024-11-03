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
# 環境を準備する
1. golang >= 1.22.4
2.mysql
3.レディス
4. bash (build.sh を使用する場合に必要)

## コンパイル
> 注: 実行中のサーバー上で自分でコンパイルすることをお勧めします。そうしないと、予期しない状況が発生する可能性があります。
1. 依存関係をインストールする
「整頓して行きましょう」
2.コンパイルを開始します

#### 自分でコンパイルする
- golang をインストールし、バージョンが 1.22.4 以上であること
- Linux に gcc 環境をインストールして実行します。

```bash
bash ./build.sh
```

- Windows下で実行
```bash
bash ./build.sh
```

- スクリプトの実行が完了すると、ビルド フォルダーにコンパイルされた実行可能ファイルが表示されます。

### コンパイルしたくない
[Build-dev](https://github.com/gucooing/hkrpg-go/actions/workflows/Build.yml) にアクセスしてダウンロードします。

＃＃ 走る
### 1. リソースを準備します。
データ リソース、データはウェアハウスのデータを使用できますが、リソース フォルダーには読み取りおよび書き込み権限が付与されている必要があります。

リソースの準備:
1. [StarRailData](https://github.com/Dimbreath/StarRailData)をダウンロード

2. 補足ファイル（タスクファイル）をダウンロードする [DanhengServer-Resources](https://github.com/EggLinks/DanhengServer-Resources)

3. まず StarRailData をリソースに解凍し、DanhengServer-Resources で一度上書きします (更新構成のみをカバーし、ExcelOutput の非互換性は上書きしないでください)。

### 2. 走ってください。
実行時に起動パラメータ -i appid を指定する必要があります。appid 形式は 9001.1.1.1 などの ipv4 形式です。これは次のことを意味します。

```bash
9001: 地区サーバー ID。
1: サービス ID;
1: ホスト ID。
1: 今回開始するサービス ID。
```

appid の構成の意味を理解した後、パラメータなしで起動して各サービスの設定ファイルを生成し、生成された設定ファイルは conf フォルダーにあり、独自の定義に従ってデフォルト設定ファイルの appid を変更できます。 appid (サービスはディスカバリーを使用して新しいサービスを追加しますが、各構成ファイルの appid 構成テーブルを同じにすることをお勧めします)、独自のアイデアに従って構成ファイル内の他のパラメーターを変更します。

### 3. データベースの準備、
mysql をインストールし、mysql で新しいデータベースを作成します: hkrpg-go-account && hkrpg-go-user && hkrpg-go-player && hkrpg-go-conf (utf8mb4)。次に、設定ファイルのアカウントとパスワードを変更し、redis をインストールします。 、ファイル内の構成パスワードを変更します (このサービスはテーブルとデータベースに分割できますが、同じテーブルは同じデータベース内に存在する必要があります)

### 4. 始めます。
すべての事前準備作業が完了したので、開始するときが来ました。推奨される起動シーケンスは次のとおりです。
> 以下の例の起動方法は、デフォルト設定ファイルの起動パラメータです。

```bash
./nodeserver -i 9001.3.1.1
./gameserver -i 9001.2.1.1
./gateserver -i 9001.1.1.1
./dispatch -i 9001.4.1.1
./multiserver -i 9001.5.1.1
./muipserver -i 9001.6.1.1
```

## 各サービスの機能
### ノードサーバー ノードサーバー (ステートフル、クラスタ化不可)、サービス検出、サービス管理

### ディスパッチ ログイン サーバー (ステートレス、クラスタ化可能)

### ゲートサーバー ゲートウェイ サーバー (ステートフル、クラスタ化可能)、内部ネットワークと外部との間の対話のための唯一のインターフェイス

### ゲームサーバー論理サーバー (ステートフル、クラスター化可能)、ビジネス ロジックの処理

### マルチサーバー マルチプレイヤー サーバー (ステートフル、クラスタ化不可) には有用なサービスがありません

### muipserver は現在 API のみを担当します

## 高度な操作

### マルチゲートサーバー、マルチゲームサーバーの展開
ゲートサーバーを例にとると、デフォルトでは 9001.1.1.1 の構成が 1 つだけあり、起動時に同じ実行可能ファイルを使用できます。 2 番目のポートは - i 9001.1.1.2 で開始できます。同じマシン上で使用している場合は、構成された 2 つのポートが競合しないように注意してください。

待って.........

## 注意事項
内部ネットワークと外部ネットワークを適切に管理し、外部ネットワークがクラ​​スターの内部ネットワークに自由にアクセスできないようにしてください。
外部ネットワークの帯域幅が 1Gpbs/s 未満で、遅延が 10ms 以上の場合は、外部ネットワーク データベースを使用しないでください。

## テストはしたいが、複雑な環境を構築したくない

1. [Build-dev](https://github.com/gucooing/hkrpg-go/actions/workflows/Build.yml) に移動して、hkrpg-pe 実行可能ファイルをダウンロードします