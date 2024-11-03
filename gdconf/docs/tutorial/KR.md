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
# 환경 준비
1. 골랭 >= 1.22.4
2.mysql
3. 레디스
4. bash(build.sh 사용 시 필수)

## 컴파일
> 참고: 실행 중인 서버에서 직접 컴파일하는 것을 권장합니다. 그렇지 않으면 예상치 못한 상황이 발생할 수 있습니다.
1. 종속성 설치
`모드 정리하러 가세요`
2. 컴파일 시작

#### 직접 컴파일
- golang을 설치하고 버전이 1.22.4 이상입니다.
- Linux에 gcc 환경을 설치한 후 실행

```bash
bash ./build.sh
```

- 윈도우에서 실행
```bash
bash ./build.sh
```

- 스크립트 실행이 완료되면 빌드 폴더에 컴파일된 실행 파일을 확인할 수 있습니다.

### 컴파일하고 싶지 않음
[Build-dev](https://github.com/gucooing/hkrpg-go/actions/workflows/Build.yml)로 이동하여 다운로드하세요.

## 달리다

### 1. 리소스 준비:
데이터 리소스, 데이터는 웨어하우스의 데이터를 사용할 수 있지만 리소스 폴더에는 읽기 및 쓰기 권한이 부여되어야 합니다.

자원 준비:
1. [StarRailData](https://github.com/Dimbreath/StarRailData) 다운로드

2. 보조파일(작업파일) 다운로드 [DanhengServer-Resources](https://github.com/EggLinks/DanhengServer-Resources)

3. 먼저 StarRailData를 리소스로 압축 해제한 다음 DanhengServer-Resources로 한 번 덮어씁니다(업데이트 구성만 다루고 ExcelOutput 비호환성을 덮어쓰지 않음)

### 2. 실행,
실행 시 시작 매개변수 -i appid를 전달해야 합니다. 여기서 appid 형식은 ipv4 형식입니다(예: 9001.1.1.1). 이는 다음을 의미합니다.

```bash
9001: 구역 서버 ID;
1: 서비스 ID;
1: 호스트 ID;
1: 이번에 시작할 서비스 ID.
```

appid 구성의 의미를 이해한 후 매개변수 없이 시작하여 각 서비스의 구성 파일을 생성할 수 있습니다. 생성된 구성 파일은 conf 폴더에 있으며, 기본 구성 파일에서 appid를 사용자가 정의한 대로 변경할 수 있습니다. appid(서비스가 검색을 사용하여 새 서비스를 추가하더라도 각 구성 파일의 appid 구성 테이블이 동일한 것이 좋습니다.) 그런 다음 자신의 아이디어에 따라 구성 파일의 다른 매개변수를 변경합니다.

### 3. 데이터베이스 준비,
mysql을 설치하고 mysql에 새 데이터베이스를 생성합니다: hkrpg-go-account && hkrpg-go-user && hkrpg-go-player && hkrpg-go-conf (utf8mb4) 그런 다음 구성 파일에서 계정과 비밀번호를 변경하고 redis를 설치합니다. , 파일의 구성 비밀번호를 변경합니다. (이 서비스는 테이블과 데이터베이스로 나눌 수 있지만 동일한 테이블은 동일한 데이터베이스에 있어야 합니다.)

### 4. 시작,
모든 예비 준비 작업이 완료되었으며 이제 시작할 시간입니다. 권장되는 시작 순서는 다음과 같습니다.
> 다음 예의 시작 방법은 기본 구성 파일의 시작 매개변수입니다.

```bash
./nodeserver -i 9001.3.1.1
./gameserver -i 9001.2.1.1
./gateserver -i 9001.1.1.1
./dispatch -i 9001.4.1.1
./multiserver -i 9001.5.1.1
./muipserver -i 9001.6.1.1
```

## 각 서비스 기능

### nodeserver 노드 서버(상태 저장, 클러스터 불가능), 서비스 검색, 서비스 관리

### 로그인 서버 디스패치(상태 비저장, 클러스터 가능)

### Gateserver 게이트웨이 서버(상태 저장, 클러스터 가능), 내부 네트워크와 외부 세계 간의 상호 작용을 위한 유일한 인터페이스

### gameserver 논리 서버(상태 저장, 클러스터 가능), 비즈니스 논리 처리

### 멀티서버 멀티플레이어 서버(상태 저장형, 클러스터형 아님)에는 유용한 서비스가 없습니다.

### muipserver는 현재 API만 담당합니다.

## 고급 작업

### 다중 게이트서버, 다중 게임서버 배포
게이트 서버를 예로 들면 기본적으로 9001.1.1.1의 구성만 있습니다. 9001.1.1.2의 구성을 추가할 수 있습니다. 첫 번째 게이트 서버는 -i 9001.1.1.1로 시작됩니다. 두 번째 포트는 - i 9001.1.1.2로 시작됩니다. 동일한 시스템에 있는 경우 구성된 두 포트와 충돌하지 않도록 주의하십시오.

기다리다.........

## 메모
내부 및 외부 네트워크를 잘 처리하고 외부 네트워크가 클러스터 내부 네트워크에 자유롭게 접근하는 것을 허용하지 마십시오.
외부 네트워크 대역폭이 1Gpbs/s 미만이고 지연 시간이 10ms 이상인 경우 외부 네트워크 데이터베이스를 사용하지 마십시오.

## 테스트하고 싶지만 복잡한 환경을 구성하고 싶지 않음

1. [Build-dev](https://github.com/gucooing/hkrpg-go/actions/workflows/Build.yml)로 이동하여 hkrpg-pe 실행 파일을 다운로드합니다.