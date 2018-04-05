Ecrire un programme go qui écoute en http sur le port 8080.

Il affiche le contenu du fichier www/bonjour.html dans le navigateur lorsqu'on tape l'url suivante :

```
http://localhost:8080/bonjour.html
```

Etendre le programme pour appeler dynamiquement un fichier dans le dossier www/

Si on ajoute n'importe quel fichier *filename* dans le sous-dossier www/ on pourra y accéder via l'url http://localhost:8080/filename.

Par exemple,

```
http://localhost:8080/bonjour.html
```
Affichera le contenu du fichier www/bonjour.html

alors que,
```
http://localhost:8080/aurevoir.html
```
Affichera le contenu du fichier www/aurevoir.html

