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

# Préparer l'environnement
1. golang >= 1.22.4
2.mysql
3. redis
4. bash (obligatoire lors de l'utilisation de build.sh)

## Compiler
> Remarque : Il est recommandé de compiler vous-même sur le serveur en cours d'exécution, sinon des situations inattendues pourraient survenir.
1. Installer les dépendances
`allez ranger le mod`
2. Commencez à compiler

#### Compilez vous-même
- Installez Golang et la version n'est pas inférieure à 1.22.4
- Installer l'environnement gcc sous linux puis exécuter

```bash
bash ./build.sh
```

- Exécuter sous Windows
```bash
.\build.bat
```

- Une fois l'exécution du script terminée, vous pouvez voir le fichier exécutable compilé dans le dossier build

### Je ne veux pas compiler
Accédez à [Build-dev](https://github.com/gucooing/hkrpg-go/actions/workflows/Build.yml) pour télécharger

## Courir
### 1. Préparez les ressources :
Ressources de données, les données peuvent utiliser les données de l'entrepôt, mais le dossier de ressources doit disposer d'autorisations de lecture et d'écriture.

Préparation des ressources :
1. Téléchargez [StarRailData](https://github.com/Dimbreath/StarRailData)

2. Téléchargez des fichiers supplémentaires (fichiers de tâches) [DanhengServer-Resources](https://github.com/EggLinks/DanhengServer-Resources)

3. Décompressez d'abord StarRailData en ressources, puis écrasez-le une fois avec DanhengServer-Resources (couvrez uniquement la configuration de mise à jour, n'écrasez pas l'incompatibilité ExcelOutput)
### 2. Courez,

Le paramètre de démarrage -i appid doit être transporté lors de l'exécution, où le format appid est le format ipv4, tel que : 9001.1.1.1, ce qui signifie :

```bash
9001 : identifiant du serveur de district ;
1 : identifiant du service ;
1 : identifiant de l’hôte ;
1 : l'identifiant du service à démarrer cette fois ;
```
Après avoir compris la signification de la composition de l'appid, vous pouvez le démarrer sans paramètres pour générer le fichier de configuration de chaque service. Le fichier de configuration généré se trouve dans le dossier conf, puis modifier l'appid dans le fichier de configuration par défaut selon votre propre définition. appid (Bien que le service utilise la découverte pour ajouter de nouveaux services, il est toujours recommandé que la table de configuration appid dans chaque fichier de configuration soit la même), puis modifiez les autres paramètres du fichier de configuration selon vos propres idées.

### 3. Préparation de la base de données,
Installez mysql, créez une nouvelle base de données dans mysql : hkrpg-go-account && hkrpg-go-user && hkrpg-go-player && hkrpg-go-conf (utf8mb4), puis changez le compte et le mot de passe dans le fichier de configuration, installez redis , et modifiez le mot de passe de configuration dans le fichier (ce service peut être divisé en tables et bases de données, mais la même table doit être dans la même base de données)

### 4. Commencez,
Tous les travaux de préparation préliminaires sont terminés et il est temps de commencer. La séquence de démarrage recommandée est la suivante :
> La méthode de démarrage dans l'exemple suivant correspond aux paramètres de démarrage du fichier de configuration par défaut.

```bash
./nodeserver -i 9001.3.1.1
./serveur de jeux -i 9001.2.1.1
./gateserver -i 9001.1.1.1
./dispatch -i 9001.4.1.1
./multiserveur -i 9001.5.1.1
./muipserver -i 9001.6.1.1
```

## Chaque fonction de service

### serveur de nœuds nodeserver (avec état, non clusterable), découverte de services, gestion de services

### serveur de connexion de répartition (sans état, clusterable)

### serveur passerelle gateserver (avec état, clusterable), seule interface d'interaction entre le réseau interne et le monde extérieur

### serveur logique gameserver (avec état, clusterable), traitement de la logique métier

### multiserveur Le serveur multijoueur (avec état, non clusterisable) n'a aucun service utile

### muipserver est actuellement uniquement responsable de l'API


## Opérations avancées

### Déploiement multi-gateserver, multi-gameserver
En prenant gateserver comme exemple, il n'y a qu'une seule configuration de 9001.1.1.1 par défaut. Vous pouvez ajouter une configuration de 9001.1.1.2. Vous pouvez utiliser le même fichier exécutable au démarrage. Le premier gateserver est démarré avec -i 9001.1.1.1. et le second est démarré avec - i 9001.1.1.2 peut être démarré Si vous êtes sur la même machine, veillez à ne pas entrer en conflit avec les deux ports configurés.

attendez.........

## Remarques
Veuillez bien gérer les réseaux internes et externes et ne pas autoriser le réseau externe à accéder librement au réseau interne du cluster.
Si la bande passante de votre réseau externe est inférieure à 1 Gpbs/s et que le délai n'est pas inférieur à 10 ms, veuillez ne pas utiliser la base de données du réseau externe.

## Vous souhaitez tester mais ne souhaitez pas configurer un environnement complexe

1. Accédez à [Build-dev](https://github.com/gucooing/hkrpg-go/actions/workflows/Build.yml) pour télécharger le fichier exécutable hkrpg-pe