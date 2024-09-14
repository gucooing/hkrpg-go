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

# Explication détaillée des paramètres de conf

### LogLevel : niveau de journalisation

### MaxPlayer : Le nombre maximum de joueurs sur ce serveur de jeu

### AutoCreate : s'il faut s'inscrire automatiquement

###Liste d'applications :

 Format : appid[configuration]
 port_player : utilisé pour gateserver, représentant le port de connexion externe, c'est-à-dire le port kcp
 port_gt : utilisé pour gameserver, représentant le port utilisé pour se connecter au gateserver
 port_service : représente le port sur lequel ce service accepte les connexions d'autres services.
 port_http : démarre un serveur http sur ce port

### Application :
 Port : port d'écoute
 InnerAddr : adresse externe
 OuterAddr : adresse d'écoute

### NetConf :
 Format : service[adresse]
 Nœud : représente l'adresse de connexion du serveur de nœuds

### MysqlConf :
 Format : tableau[adresse]
 Chaque adresse de connexion MySQL requise

### RedisConf :
 Format : tableau [configuration de la connexion]
 Chaque configuration de connexion Redis requise

### Expédition : (expédition uniquement)
 Format : [] Adresse du serveur de chaque nœud
 Lorsque plusieurs serveurs ne communiquent pas entre eux, modifiez cette configuration pour vous connecter

### GameDataConfigPath : (pour le serveur de jeux uniquement) renseignez le chemin de la table de configuration (en fonction du répertoire de démarrage du programme)