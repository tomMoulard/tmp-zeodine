Note de Arnaud sur les dockerfile !!!
===================================

## Ubuntu

A mettre dans le docker file : 

>RUN apt-get update;\
apt-get install -y sysvbanner;\

Note sur les tag : pensez à aller sur https://hub.docker.com pour verifier les tags dispo 

Permet de mettre à jour les packets et de d'intaller banner.
Si on ajoute la ligne <code>banner yoyo</code>, on auras l'affichage de <code>banner yoyo</code> à la fin de la compilation

Pour lancer le bash avec ubuntu 16.04 :

* docker run --rm dockubuntu bash

------------------------------------
## Go
Important, toujours avoir un fichier en .go dans le même répertoire avant de build le dockerfile 

* docker build -t dockergo  .
* docker run --rm --publish 8080:8080 dockergo

<code>-t name</code> permet de donner un nom aux dockers, il est obligatoire

<code>--rm</code> permet d'arrêter le docker lorsque l'on fait ^C

<code>--publish portPC:portDOCKER</code> permet de faire de joindre le port à gauche (celui du PC) vers celui de droite (dans le docker)

----------------------------------------

## MySQL

>ENV MYSQL_ROOT_PASSWORD=123

Permet de modifier une variable d'environement et évite de le définir manuelement lors de l'execution du docker avec <code>docker run --name sql dkmysql</code>

On se connecte au serveur mySQL avec :
* <code>docker run -it --link sql:mysql --rm mysql sh -c 'exec mysql -h"$MYSQL_PORT_3306_TCP_ADDR" -P"$MYSQL_PORT_3306_TCP_PORT" -uroot -p"$MYSQL_ENV_MYSQL_ROOT_PASSWORD"'</code>

<code>docker build -t sql_go .</code>

<code>docker run --rm sql_go</code>