# Installation

Installer Docker sur votre machine et vérifier que tout fonctionne avec la commande :

```
sudo docker run hello-world
```

Trouver le moyen de ne pas taper "sudo docker..." mais juste "docker" sans passer root, et vérifer que cela fonctionne avec la commande :

```
docker run hello-world
```

# Dockerfile

Ecrire un premier fichier Dockerfile qui installe Ubuntu 16.04.

Lancez le container et connectez vous dans un "bash".


# Dockerfile votre application

Reprenez votre programme de l'exo 2 "miniserver" et encapsulez le dans un container.

Rendez le port 8080 accessible depuis votre host et vérifier que cela marche en vous connectant sur :

```
http://localhost:8080/aurevoir.html
```

Faites que le port d'accès sur votre host soit le 80 sans changer le port d'écoute de votre programme (8080).
Vérifiez que cela fonctione en allant sur :

```
http://localhost/bonjour.html
```

# Dockerfile MySQL

Installer et exécuter un container avec la dernière version de MySQL.

Vérifiez que vous pouvez vous y connecter en utilisant la client "mysql" (vous avez le droit d'entrer dans le container via un bash avant la connexion au serveur)

Rendez le port 3306 accessible depuis votre host, et écrivez un petit programme en Go qui y accède pour se connecter au serveur mysql.

# Docker Compose

Prenez la main sur l'outil Docker Compose, qui permet de gérer plusieurs containers et les liens entre eux.

Ecrire un docker-compose.yml lançant un MySQL et votre application en Go y accédant, avec tous les liens nécessaires.

Le tout doit fonctionner avec la commande suivante :

```
docker-compose up -d
```

