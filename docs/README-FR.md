![hkrpg-go](https://socialify.git.ci/gucooing/hkrpg-go/image?description=1&font=Inter&forks=1&language=1&name=1&owner=1&pattern=Circuit%20Board&stargazers=1&theme=Auto)

<div align="center">
<table>
<td valign="center"><a href="README.md"><img src="https://github.com/twitter/twemoji/blob/master/assets/svg/1f1fa-1f1f8.svg" width="16"/> English</td>
 
<td valign="center"><a href="docs/README_zh-CN.md"><img src="https://em-content.zobj.net/thumbs/120/twitter/351/flag-china_1f1e8-1f1f3.png" width="16"/> 简中</td>
 
<td valign="center"><a href="docs/README_zh-TW.md"><img src="https://em-content.zobj.net/thumbs/120/twitter/351/flag-china_1f1e8-1f1f3.png" width="16"/> 繁中</td>
 
<td valign="center"><a href="docs/README-JP.md"><img src="https://github.com/twitter/twemoji/blob/master/assets/svg/1f1ef-1f1f5.svg" width="16"/> 日本語</td>
 
<td valign="center"><a href="docs/README-RU.md"><img src="https://github.com/twitter/twemoji/blob/master/assets/svg/1f1f7-1f1fa.svg" width="16"/> Русский</a></td>

<td valign="center"><a href="docs/README-FR.md"><img src="https://em-content.zobj.net/thumbs/160/twitter/154/flag-for-france_1f1eb-1f1f7.png" width="16"/> Français</td>
 
<td valign="center"><a href="docs/README-KR.md"><img src="https://em-content.zobj.net/source/twitter/53/flag-for-south-korea_1f1f0-1f1f7.png" width="16"/> 한국어</td>
 
<td valign="center"><a href="docs/README-VI.md"><img src="https://em-content.zobj.net/thumbs/120/twitter/351/flag-vietnam_1f1fb-1f1f3.png" width="16"/> Tiếng Việt </a>
</td>
</table>
</div>

# **Bienvenue sur notre [Discord](https://discord.gg/222yVp6pUq)**

## Un ami fait son retour, gardez un œil dessus [hk4e-dmca](https://github.com/flswld/hk4e-go)

## Veuillez ne pas l'utiliser dans un environnement de production

### Documentation :
* [easy-tutorial](./docs/tutorial/zh-cn.md)
* [config parsing](./docs/conf/zh-CN.md)
* [api usage](./docs/command/zh-CN.md)
* [Details](./docs/progress/zh-CN.md)

### Notes :
* Si vous souhaitez apporter votre aide à ce projet, n'hésitez pas à le soumettre.

 ### Contenu terminé
- **Sacs à dos**
- **Bataille**
- **Tirage de cartes**
- **Formation**
- **Courrier**
- **Amis**
- **Accessoires de scène/Génération de monstres/PNJ** - **Scénario**
- **Intrigue**
- **La Cour de l'oubli et autres spin-offs**
- **Univers simulé**
- **Univers différentiel**
- **Transmission temporelle régulière (partielle**

### Connexion au client (Fiddler)
1. Installez et exécutez [Fiddler Classic](https://www.telerik.com/fiddler).
2. Configurez Fiddler pour décrypter le trafic https (Outils -> Options -> HTTPS -> Décrypter le trafic HTTPS), assurez-vous que l'option `Ignorer les erreurs de certificat du serveur` est cochée.
3. Copiez et collez le code suivant dans l'onglet Fiddlerscript de Fiddler Classic :

```javascript
import System; import System.
import  System.Windows.Forms;
import Fiddler; import System.
import System.Windows.Forms; import Fiddler; import System.Text.
import System.Text.RegularExpressions; class Handlers
static function OnBeforeRequest(oS: Session) {
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

4. Connectez-vous en utilisant votre nom de compte, le mot de passe peut être défini sur n'importe quelle valeur.