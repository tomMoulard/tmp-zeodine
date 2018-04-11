# Reponse

## Ubuntu Dockerfile:

```DockerFile
FROM ubuntu:16.04
```

build :
```bash
docker build -t ubuntu-perso .
```

run :
```bash
docker run --rm -ti ubuntu-perso bash
```

## GO Dockerfile

```Dockerfile
FROM golang:onbuild
EXPOSE 9000
```

build :
```bash
docker build -t app-go .
```

run :
```bash
docker run --rm --publish 80:9000 app-go
```
I'm mapping the port 80 on my computer with the port 9000 of the container

# MySQL Docker

```bash
docker run --name mysql-perso -e MYSQL_ROOT_PASSWORD=zeodine --publish 3306:3306 -d mysql
```

To access the server:

```bash
docker exec -it some-mysql bash
```

To access the database:

```bash
docker run -it --link mysql-perso:mysql --rm mysql sh -c 'exec mysql -h"$MYSQL_PORT_3306_TCP_ADDR" -P"$MYSQL_PORT_3306_TCP_PORT" -uroot -p"$MYSQL_ENV_MYSQL_ROOT_PASSWORD"'
```

Create user, a database to access it with the go program

```MySQL
CREATE USER tom@localhost
IDENTIFIED BY 'multipass';
```

then, run the api:

```bash
go run db-api.go

```

## Docker-compose

```bash
docker-compose up -d
```

## Traefik SSl

```bash
docker-compose up -d # to run the service
```

### Only with google-chrome

After, open a web nav with the page : http://web.docker.localhost/
It works.

The open : https://web.docker.localhost/
It doesn't work

(and same with admin-web.docker.localhost )
