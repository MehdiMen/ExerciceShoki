# ExerciceShoki

Je pars du principe où tous les outils et logiciel sont installé sur la machine.
On se place dans le répertoire de travail où se trouve tous les dossiers envoyés.
Pour la partie go les chemins relatifs nécessaires sont:
    "github.com/httprouter"
    "github.com/gocolly/colly"
    
Pour commencer, on lance 2 terminaux de commandes (sur windows) pour le front et le back:

Dans le 1er, en prenant comme point de départ le répertoire de travail, on effectue les commandes suivantes:
    >cd react
    >cd src
    >yarn start
On se rend, si ce n'est pas fait automatiquement, à l'adresse : http://localhost:3000/ (adresse sur le screen front)
On peut voir l'interface de notre webapp.

Dans le 2eme, en prenant toujours comme point de départ le répertoire de travail, on tape:
    >cd go
    >cd src
    >cd App
    >App
On se rend à l'adresse: http://localhost:8083/?url=https://www.tripadvisor.fr/Restaurant_Review-g35805-d1894904-Reviews-The_Kerryman-Chicago_Illinois.html (adresse sur le screen back)
Si il n'y avais pas de problème, l'adresse serai: http://localhost:8083/
On voit ici les résultat de la recherche sur le back. Les données ont été récupérées à l'aide de Colly et se trouve au emplacement indiqués par le screen site.


Comme on peut le voir il manque la liaison entre le front et le back. Le boutton a pour but d'envoyer l'URL au back directement sans passer par un paramètre qui doit toujours être indiqué dans la barre de recherche pour que cela fonctionne.
Inversement le back doit envoyer au front les données récupérées et les implémenter dans la classe du front pour que les changements soient pris en compte à l'affichage de la page.
