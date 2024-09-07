[EN](./EN.md) | [简中](./zh-CN.md) | [繁中](./zh-TW.md) | [JP](./JP.md) | [RU](./RU.md) | [FR](./FR.md) | [KR](./KR.md) | [VI](./VI.md)

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