#!/bin/bash

echo "Skrypt uruchamiajacy skrypt zadanie1.pp via Docker"
echo "Podaj nazwe kontenera Dockera:"
read name
if [ "${#name}" -le 0 ]; then
    echo "Nalezy podac nazwe!"
    exit 1
fi
docker container inspect "$name" > /dev/null 2>&1
if [ "$?" -ne 0 ]; then
    echo "Podany kontener nie istnieje!"
    exit 2
fi
echo "Uruchamiam kontener..."
docker start "$name"
docker cp zadanie1.pp "$name:/home/student/projobj/zadanie1.pp"
docker exec $name fpc zadanie1.pp
docker exec $name ./zadanie1
echo "Wylaczam kontener..."
docker stop "$name"