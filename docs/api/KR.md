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

# API
- 샘플 매개변수는 모두 uid = 1 | sign_key = 123456을 기반으로 합니다.

### 경로 매개변수
- **cmd**: `int16` 필요 **호출 명령**
- **uid**: `uint32` 필요 **플레이어 UID**
- **sign_key**: `문자열` 선택 사항 **키**

#### 요청 예시:
``일반 텍스트
가져오기: api?cmd=1&uid=1&sign_key=123456
````

#### `json`을 구문 분석하는 콜백:
- **코드**: 상태 0 성공 -1 실패
- **msg**: 콜백 내용

___

### 월드 레벨 cmd 1001 설정
**매개변수**:
- **world_level**: `uint32` 필요 **월드 레벨 설정**
#### 요청 예시:
``일반 텍스트
가져오기: api?cmd=1001&uid=1&sign_key=123456&world_level=6
````

___

### 계정 데이터 가져오기 cmd 1002
**매개변수**: 없음
#### 요청 예시:
``일반 텍스트
가져오기: api?cmd=1002&uid=1&sign_key=123456
````

___

### 서버 상태 가져오기 cmd 1003
**매개변수**: 없음
#### 요청 예시:
``일반 텍스트
가져오기: api?cmd=1003&sign_key=123456
````

___

### 소품 cmd 1004 가져오기
- **all**: `bool` 선택 사항 | **모든 항목을 가져올지 여부 | 0:false|1:true**
- **id**: `uint32` 선택 사항 **항목 ID**
- **num**: `uint32` 선택 사항 | **항목 수**
#### 요청 예시:
``일반 텍스트
가져오기: api?cmd=1004&uid=1&sign_key=123456&all=0&id=22&num=999
````

___

### 성물 얻기 cmd 1005
- **all**: `bool` 선택 사항 | **모든 항목을 가져올지 여부 | 0:false|1:true**
- **id**: `uint32` 필요 **성유물 ID**
- **num**: `uint32` 필요 **성물 수**
- **main**: `uint32` 선택사항 | **성물의 주요 속성을 지정**
- **sub**: `string` 선택사항 | **성유물의 보조 속성을 지정**
#### 요청 예시:
``일반 텍스트
가져오기: api?cmd=1005&uid=1&sign_key=123456&all=0&id=31011&num=1&main=1&sub=[2:10][3:9][4:8]
````

___