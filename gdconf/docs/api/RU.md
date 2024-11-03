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
- Все параметры образца основаны на uid = 1 |sign_key = 123456.

### Параметры пути
- **cmd**: требуется `int16` | **Вызов команды**
- **uid**: требуется `uint32` | **UID игрока**
- **sign_key**: `строка` необязательно | **ключ**

#### Пример запроса:
```открытый текст
ПОЛУЧИТЬ: api?cmd=1&uid=1&sign_key=123456
```

#### Анализ обратного вызова `json`:
- **код**: статус 0 успех -1 неудача
- **msg**: содержимое обратного вызова.

___

### Установить мировой уровень cmd 1001
**параметр**:
- **world_level**: требуется `uint32` **Мировой уровень должен быть установлен**
#### Пример запроса:
```открытый текст
ПОЛУЧИТЬ: api?cmd=1001&uid=1&sign_key=123456&world_level=6
```

___

### Получить данные аккаунта cmd 1002
**Параметры**: Нет
#### Пример запроса:
```открытый текст
ПОЛУЧИТЬ: api?cmd=1002&uid=1&sign_key=123456
```

___

### Получить статус сервера cmd 1003
**Параметры**: Нет
#### Пример запроса:
```открытый текст
ПОЛУЧИТЬ: api?cmd=1003&sign_key=123456
```

___

### Получить реквизиты cmd 1004
- **all**: `bool` необязательно | **Получать ли все элементы | 0:false|1:true**
- **id**: `uint32` необязательно | **идентификатор элемента**
- **num**: `uint32` необязательно | **количество элементов**
#### Пример запроса:
```открытый текст
ПОЛУЧИТЬ: api?cmd=1004&uid=1&sign_key=123456&all=0&id=22&num=999
```

___

### Получить святую реликвию cmd 1005
- **all**: `bool` необязательно | **Получать ли все элементы | 0:false|1:true**
- **id**: требуется `uint32` | **Идентификатор священной реликвии**
- **num**: требуется `uint32` | **Количество святых мощей**
- **main**: `uint32` необязательно | **Укажите основной атрибут святой реликвии**
- **sub**: `string` необязательно | **Укажите второстепенные атрибуты святой реликвии**
#### Пример запроса:
```открытый текст
GET: api?cmd=1005&uid=1&sign_key=123456&all=0&id=31011&num=1&main=1&sub=[2:10][3:9][4:8]
```

___