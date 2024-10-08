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
# Подготовьте среду
1. голанг >= 1.22.4
2.mysql
3. переоформить
4. bash (требуется при использовании build.sh)

## Компилировать
> Примечание. Рекомендуется выполнять компиляцию на работающем сервере самостоятельно, в противном случае могут возникнуть непредвиденные ситуации.
1. Установите зависимости
`иди в порядок`
2. Начните компилировать

#### Соберите самостоятельно
- Установите golang и версию не ниже 1.22.4.
- Установите среду gcc под Linux и затем выполните

```bash
bash ./build.sh
```

- Выполнять под окнами
```bash
bash ./build.sh
```

- После завершения работы скрипта вы сможете увидеть скомпилированный исполняемый файл в папке сборки.

### Не хочу компилировать
Перейдите на [Build-dev](https://github.com/gucooing/hkrpg-go/actions/workflows/Build.yml), чтобы загрузить.

## Бегать
### 1. Подготовьте ресурсы:
Ресурсы данных, данные могут использовать данные хранилища, но папке ресурсов необходимо предоставить права на чтение и запись.

Подготовка ресурсов:
1. Загрузите [StarRailData](https://github.com/Dimbreath/StarRailData)

2. Загрузите дополнительные файлы (файлы задач) [DanhengServer-Resources](https://github.com/EggLinks/DanhengServer-Resources)

3. Сначала распакуйте StarRailData в ресурсы, а затем один раз перезапишите его с помощью DanhengServer-Resources (охватывайте только конфигурацию обновления, не перезаписывайте несовместимость ExcelOutput)

### 2. Беги,
Параметр запуска -i appid должен передаваться при запуске, где формат appid — это формат ipv4, например: 9001.1.1.1, что означает:

```bash
9001: идентификатор окружного сервера;
1: идентификатор услуги;
1: идентификатор хоста;
1: идентификатор службы, которая будет запущена в этот раз;
```

Поняв смысл состава appid, вы можете запустить его без параметров, чтобы сгенерировать файл конфигурации каждой службы. Сгенерированный файл конфигурации находится в папке conf, а затем изменить appid в файле конфигурации по умолчанию в соответствии с вашим собственным. appid (Хотя служба использует обнаружение для добавления новых служб, все же рекомендуется, чтобы таблица конфигурации appid в каждом файле конфигурации была одинаковой), а затем измените другие параметры в файле конфигурации в соответствии с вашими собственными идеями.

### 3. Подготовка базы данных,
Установите mysql, создайте новую базу данных в mysql: hkrpg-go-account && hkrpg-go-user && hkrpg-go-player && hkrpg-go-conf (utf8mb4), затем измените учетную запись и пароль в файле конфигурации, установите Redis , и измените пароль конфигурации в файле (этот сервис можно разделить на таблицы и базы данных, но одна и та же таблица должна находиться в одной базе данных)

### 4. Старт,
Все предварительные подготовительные работы завершены и пора приступать к работе. Рекомендуемая последовательность запуска следующая:
> Метод запуска в следующем примере — это параметры запуска файла конфигурации по умолчанию.

```bash
./nodeserver -i 9001.3.1.1
./gameserver -i 9001.2.1.1
./gateserver -i 9001.1.1.1
./отправка -я 9001.4.1.1
./мультисервер -i 9001.5.1.1
./muipserver -i 9001.6.1.1
```

## Каждая сервисная функция

### nodeserver узел-сервер (с сохранением состояния, без кластеризации), обнаружение сервисов, управление сервисами

### диспетчер сервера входа в систему (без сохранения состояния, кластеризуемый)

### gateserver сервер-шлюз (с сохранением состояния, кластерный), единственный интерфейс для взаимодействия внутренней сети с внешним миром
### логический сервер игрового сервера (с сохранением состояния, кластеризуемый), обрабатывающий бизнес-логику

### multiserver Многопользовательский сервер (с сохранением состояния, без кластеризации) не имеет полезных сервисов

### muipserver в настоящее время отвечает только за API


## Расширенные операции
### Многосерверное развертывание, развертывание нескольких игровых серверов
Если взять в качестве примера сервер ворот, то по умолчанию существует только одна конфигурация 9001.1.1.1. Вы можете добавить конфигурацию 9001.1.1.2. Вы можете использовать тот же исполняемый файл при запуске. Первый сервер ворот запускается с -i 9001.1.1.1. а второй запускается с -i 9001.1.1.2 можно запустить. Если вы находитесь на одной машине, будьте осторожны, чтобы не конфликтовать с двумя настроенными портами.

ждать.........

## Примечания
Пожалуйста, хорошо обращайтесь с внутренними и внешними сетями и не позволяйте внешней сети свободно получать доступ к внутренней сети кластера.
Если пропускная способность вашей внешней сети менее 1 Гбит/с и задержка не менее 10 мс, не используйте базу данных внешней сети.

## Хотите протестировать, но не хотите настраивать сложную среду

1. Перейдите в [Build-dev](https://github.com/gucooing/hkrpg-go/actions/workflows/Build.yml), чтобы загрузить исполняемый файл hkrpg-pe.