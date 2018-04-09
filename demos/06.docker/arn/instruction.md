Note de Arnaud sur les dockerfile !!!
===================================

## Ubuntu

A mette dans le docker file : 

>RUN apt-get update;\
apt-get install -y sysvbanner;\


Permet de mettre à jour les packets et de d'intaller banner.
Si on ajoute la ligne <code>banner yoyo</code>, on auras l'affichage de <code>banner yoyo</code> à la fin de la compilation

Pour lancer le bash avec ubuntu 16.04 :

* docker run --rm dockubuntu bash

------------------------------------
## Go
Important, toujours avoir un fichier en .go dans le même répertoire avant de build le dockerfile 

* docker build -t dockergo  .
* docker run --rm --publish 8080:8080 dockergo

<code>-t</code> permet de donner un nom aux dockers, il est obligatoire

<code>--rm</code> permet d'arrêter le docker lorsque l'on fait ^C

<code>--publish portPC:portDOCKER</code> permet de faire de joindre le port à gauche (celui du PC) vers celui de droite (dans le docker)

----------------------------------------

## MySQL
