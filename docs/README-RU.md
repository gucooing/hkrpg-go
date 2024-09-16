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

# **Добро пожаловать в наш [Discord](https://discord.gg/222yVp6pUq)**

## Друг возвращается, следите за [hk4e-dmca](https://github.com/flswld/hk4e-go)

## Пожалуйста, не используйте его в производственной среде

### Документация:
* [easy-tutorial](./docs/tutorial/RU.md)
* [config parsing](./docs/conf/RU.md)
* [api usage](./docs/command/RU.md)
* [Details](./docs/progress/RU.md)

### Примечания:
* Если вы хотите помочь с этим проектом, не стесняйтесь  чтобы отправить его.

 ### Завершенный контент
- **Рюкзаки**
- **Битва**
- **Вытягивание карт**
- **Формирование**
- **Почта**
- **Друзья**
- **Реквизит сцены/Генерация монстров/NPC** - **Сценарий**
- **Сюжет**
- **Суд забвения и другие спин-оффы**
- **Имитированная вселенная**
- **Дифференциальная вселенная**
- **Передача в обычном времени (частичная**

### Подключение к клиенту (Fiddler)
1. Установите и запустите [Fiddler Classic](https://www.telerik.com/fiddler).
2. Настройте Fiddler на расшифровку трафика https (Инструменты -> Параметры -> HTTPS -> Расшифровать трафик HTTPS), убедитесь, что установлен флажок «Игнорировать ошибки сертификата сервера».
3. Скопируйте и вставьте следующий код во вкладку Fiddlerscript в Fiddler  Классический:

```javascript
import System;
import System.Windows.Forms;
import Fiddler;
import System.Text.RegularExpressions;
class Handlers
{
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

4. Войдите, используя имя своей учетной записи, пароль может быть любым.