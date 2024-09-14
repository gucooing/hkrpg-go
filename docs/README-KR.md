![hkrpg-go](https://socialify.git.ci/gucooing/hkrpg-go/image?description=1&font=Inter&forks=1&language=1&name=1&owner=1&pattern=Circuit%20Board&stargazers=1&theme=Auto)

<div align="center">
<table>
<td valign="center"><a href="README.md"><img src="https://github.com/twitter/twemoji/blob/master/assets/svg/1f1fa-1f1f8.svg" width="16"/> English</td>
 
<td valign="center"><a href="README_zh-CN.md"><img src="https://em-content.zobj.net/thumbs/120/twitter/351/flag-china_1f1e8-1f1f3.png" width="16"/> 简中</td>
 
<td valign="center"><a href="README_zh-TW.md"><img src="https://em-content.zobj.net/thumbs/120/twitter/351/flag-china_1f1e8-1f1f3.png" width="16"/> 繁中</td>
 
<td valign="center"><a href="README-JP.md"><img src="https://github.com/twitter/twemoji/blob/master/assets/svg/1f1ef-1f1f5.svg" width="16"/> 日本語</td>
 
<td valign="center"><a href="README-RU.md"><img src="https://github.com/twitter/twemoji/blob/master/assets/svg/1f1f7-1f1fa.svg" width="16"/> Русский</a></td>

<td valign="center"><a href="README-FR.md"><img src="https://em-content.zobj.net/thumbs/160/twitter/154/flag-for-france_1f1eb-1f1f7.png" width="16"/> Français</td>
 
<td valign="center"><a href="README-KR.md"><img src="https://em-content.zobj.net/source/twitter/53/flag-for-south-korea_1f1f0-1f1f7.png" width="16"/> 한국어</td>
 
<td valign="center"><a href="README-VI.md"><img src="https://em-content.zobj.net/thumbs/120/twitter/351/flag-vietnam_1f1fb-1f1f3.png" width="16"/> Tiếng Việt </a>
</td>
</table>
</div>

# **[Discord](https://discord.gg/222yVp6pUq)에 오신 것을 환영합니다**

## 친구가 돌아왔습니다. [hk4e-dmca](https://github.com/flswld/hk4e-go)를 주시하세요.

## 프로덕션 환경에서 사용하지 마세요

### 문서:
* [easy-tutorial](./docs/tutorial/zh-cn.md)
* [config parsing](./docs/conf/zh-CN.md)
* [api usage](./docs/command/zh-CN.md)
* [Details](./docs/progress/zh-CN.md)

### 참고:
* 이 프로젝트를 돕고 싶으시다면, 자유롭게 제출하세요.

 ### 완료된 콘텐츠
- **배낭**
- **전투**
- **카드 뽑기**
- **형성**
- **메일**
- **친구**
- **장면 소품/몬스터/NPC 생성** - **시나리오**
- **줄거리**
- **망각의 법정과 다른 스핀오프**
- **시뮬레이션된 우주**
- **차등 우주**
- **정규 시간 전송(부분적**

### 클라이언트(Fiddler)에 연결하기
1. [Fiddler Classic](https://www.telerik.com/fiddler)을 설치하고 실행합니다.
2. Fiddler에서 https 트래픽을 해독하도록 설정합니다(도구 -> 옵션 -> HTTPS -> HTTPS 트래픽 해독). `서버 인증서 오류 무시`가 체크되어 있는지 확인합니다.
3. 다음 코드를 Fiddler Classic의 Fiddlerscript 탭에 복사하여 붙여넣습니다.

```javascript
import System; import System.
import  System.Windows.Forms;
Fiddler 가져오기; System 가져오기.
System.Windows.Forms 가져오기; Fiddler 가져오기; System.Text 가져오기.
System.Text.RegularExpressions 가져오기; 클래스 핸들러
정적 함수 OnBeforeRequest(oS: 세션) {
정적 함수 OnBeforeRequest(oS: 세션) {
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

4. 계정 이름을 사용하여 로그인하고, 비밀번호는 어떤 값으로든 설정할 수 있습니다.