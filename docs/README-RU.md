![hkrpg-go](https://socialify.git.ci/gucooing/hkrpg-go/image?description=1&font=Inter&forks=1&language=1&name=1&owner=1&pattern=Circuit%20Board&stargazers=1&theme=Auto)

[EN](./README.md) | [简中](./docs/README_zh-CN.md) | [繁中](./docs/README_zh-CN.md) | [JP](./docs/README_zh-CN.md) | [RU](./docs/README_zh-CN.md) | [FR](./docs/README_zh-CN.md) | [KR](./docs/README_zh-CN.md) |  [VI](./docs/README_zh-CN.md)

# **Добро пожаловать в наш [Discord](https://discord.gg/222yVp6pUq)**

## Друг возвращается, следите за [hk4e-dmca](https://github.com/flswld/hk4e-go)

## Пожалуйста, не используйте его в производственной среде

### Документация:
* [easy-tutorial](./docs/tutorial/zh-cn.md)
* [config parsing](./docs/conf/zh-CN.md)
* [api usage](./docs/command/zh-CN.md)
* [Details](./docs/progress/zh-CN.md)

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
импорт System; импорт System.
импорт System.Windows.Forms;
импорт Fiddler; импорт System.
импорт System.Windows.Forms; импорт Fiddler; импорт System.Text.
импорт System.Text.RegularExpressions; класс Handlers
статическая функция OnBeforeRequest(oS: Session) {
статическая функция OnBeforeRequest(oS: Session) {
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