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

#API
- Les exemples de paramètres sont tous basés sur uid = 1 | sign_key = 123456

### Paramètres du chemin
- **cmd** : `int16` requis | **Commande d'appel**
- **uid** : `uint32` requis | **UID du joueur**
- **sign_key** : `string` facultatif | **key**

#### Exemple de requête :
```texte brut
OBTENIR : api?cmd=1&uid=1&sign_key=123456
```

#### Analyse de rappel de `json` :
- **code** : statut 0 succès -1 échec
- **msg** : contenu du rappel

___

### Définir le niveau mondial cmd 1001
**paramètre**:
- **world_level** : `uint32` requis | **Niveau mondial à définir**
#### Exemple de requête :
```texte brut
OBTENIR : api?cmd=1001&uid=1&sign_key=123456&world_level=6
```

___

### Récupérer les données du compte cmd 1002
**Paramètres** : Aucun
#### Exemple de requête :
```texte brut
OBTENIR : api?cmd=1002&uid=1&sign_key=123456
```

___

### Obtenir l'état du serveur cmd 1003
**Paramètres** : Aucun
#### Exemple de requête :
```texte brut
OBTENIR : api?cmd=1003&sign_key=123456
```

___

### Obtenez les accessoires cmd 1004
- **tous** : `bool` facultatif | **Si tous les éléments doivent être | 0:false|1:true**
- **id** : `uint32` facultatif | **identifiant de l'élément**
- **num** : `uint32` facultatif | **nombre d'éléments**
#### Exemple de requête :
```texte brut
OBTENIR : api?cmd=1004&uid=1&sign_key=123456&all=0&id=22&num=999
```

___

### Obtenez la relique sacrée cmd 1005
- **tous** : `bool` facultatif | **Si tous les éléments doivent être | 0:false|1:true**
- **id** : `uint32` requis | **Identifiant de la relique sacrée**
- **num** : `uint32` requis | **Nombre de reliques sacrées**
- **main** : `uint32` facultatif | **Spécifiez l'attribut principal de la relique sacrée**
- **sub** : `string` facultatif | **Spécifiez les attributs secondaires de la sainte relique**
#### Exemple de requête :
```texte brut
OBTENIR : api?cmd=1005&uid=1&sign_key=123456&all=0&id=31011&num=1&main=1&sub=[2:10][3:9][4:8]
```

___