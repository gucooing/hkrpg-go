[EN](./EN.md) | [简中](./zh-CN.md) | [繁中](./zh-TW.md) | [JP](./JP.md) | [RU](./RU.md) | [FR](./FR.md) | [KR](./KR.md) | [VI](./VI.md)

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