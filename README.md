# docker-Dasboard
Simple Dashboard written in go

![Screenshot](Screenshot.png)

docker build -t dasboard .

docker run --rm -it -v /var/run/docker.sock:/var/run/docker.sock:ro -p 8088:8080 dasboard
