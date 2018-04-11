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

* docker run -it --rm dockubuntu bash

<code>-it</code> permet de dire de lancer le bash dans le ubuntu du docker

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

On se connecte à la base de donnée avec :
```bash 
 docker run -it --link sql:mysql --rm mysql sh -c 'exec mysql -h"$MYSQL_PORT_3306_TCP_ADDR" -P"$MYSQL_PORT_3306_TCP_PORT" -uroot -p"$MYSQL_ENV_MYSQL_ROOT_PASSWORD"'</code>
```

```bash
docker build -t sql_go .
docker run --rm sql_go
```

----------------------------------------

## SSL

Instruction :

A faire juste une fois :

```bash 
git pull
docker network create webapp
```
Ensuite :
```bash 
cd ~/tmp-zeodine/demos/06.docker/arn/sql_go

docker run -d -v /var/run/docker.sock:/var/run/docker.sock -v $PWD/traefik.toml:/traefik.toml -v $PWD/acme.json:/acme.json -p 80:80 -p 443:443 -l traefik.frontend.rule=Host:monitor.example.com -l traefik.port=8080 --network webapp --name traefik traefik:1.3.6-alpine --docker

docker-compose build; docker-compose up -d
```

Ensuite aller sur https://test.fr

### V2

Instruction :

A faire juste une fois :

```bash 
git pull
docker network create webapp
```
Ensuite :
```bash 
cd ~/tmp-zeodine/demos/06.docker/arn/sql_go
echo "">acme.json
docker-compose build; docker-compose up -d
```

Ensuite aller sur https://test.fr
