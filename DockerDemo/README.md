# DockerDemo
Demo on Docker

A simple application that helps you learn docker image pulls.
The image is automatically built on https://hub.docker.com/r/linuxsmurfen/dockerdemo


To pull this image:
```
docker pull linuxsmurfen/dockerdemo
```

To run this image:
```
docker run -p 7000:7000 -it linuxsmurfen/dockerdemo
```

Docker stack:
```
docker stack deploy -c docker-stack.yml DockerDemo
```
