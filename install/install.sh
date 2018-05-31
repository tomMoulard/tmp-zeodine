#!/bin/bash

START=$(date +%s)

if [ `whoami` != root ]; then
  	echo Please run this script as root or using sudo
    exit
fi

# Going to home
cd $HOME

# Installing dep
apt-get update
for x in docker docker-compose git curl ; do
	if [ ! -x /usr/bin/$x ]; then
		echo $x is not instaled
		apt-get install $x -y
	fi;
done

if [ ! -x /usr/lib/go-1.10/bin ]; then
	echo go-1.10 is not instaled
	add-apt-repository ppa:gophers/archive -y
	apt-get update
	apt-get install golang-1.10-go -y						
fi;

# Get public ip

PUBLIC_IP=`curl http://ipecho.net/plain`

# Setup ssh

if [ ! -x $HOME/.ssh ]; then # No ssh key generated
	mkdir $HOME/.ssh
	ssh-keygen -t rsa -b 4096 -C "server@$PUBLIC_IP" -N "" -f $HOME/.ssh/id_rsa
	echo Pub key to add to the git
	echo ~~~~~~~~~~~~~~~~~~~~~~~
	cat $HOME/.ssh/id_rsa.pub
	echo ~~~~~~~~~~~~~~~~~~~~~~~

	read -p "Ok ? [Y/n] " OK
	if [[ $OK == "n" ]]; then
		exit
	fi
fi

echo Cloning repop

git clone ssh://git@git.softinnov.fr:2222/zeodine/first.git

echo Staring services

# Starting the server
cd first/server
docker-compose up -d --build
docker exec -ti go bash -c "go test ./..."

# Starting the updater && make it a permanet process
cd ../updater
go build
cp updater.service /lib/systemd/system/
chmod 755 /lib/systemd/system/updater.service
systemctl enable upgrader.service
systemctl start updater.service

echo "Took" $(( $(date +%s) - $START )) "seconds"
