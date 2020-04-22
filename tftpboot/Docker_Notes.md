# Docker Notes

## Install
INFO: [Install Docker Engine on Ubuntu](https://docs.docker.com/engine/install/ubuntu/)
[Post-installation steps for Linux](https://docs.docker.com/engine/install/linux-postinstall/)

### Problems

#### cannot import name '_gi'
```
$ sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/
...
ImportError: cannot import name '_gi' from 'gi' (/usr/lib/python3/dist-packages/gi/__init__.py)
```

change default python to Ubuntu default 3.6 
```
sudo update-alternatives --config python3                                                                         [20:52:24] There are 2 choices for the alternative python3 (providing /usr/bin/python3).

  Selection    Path                Priority   Status
------------------------------------------------------------
* 0            /usr/bin/python3.7   2         auto mode
  1            /usr/bin/python3.6   1         manual mode
  2            /usr/bin/python3.7   2         manual mode
```

#### start engine on WSL
`sudo service docker start`

#### Notes for alpine
[Installation of Docker on Alpine Linux](https://docs.genesys.com/Documentation/System/8.5.x/DDG/InstallationofDockeronAlpineLinux)
```
apk add --no-cache  --repository http://dl-cdn.alpinelinux.org/alpine/edge/main --repository  http://dl-cdn.alpinelinux.org/alpine/edge/community docker
docker run hello-world
rc-update add docker boot
service docker start
docker run hello-world
```

### test
`docker run hello-world`

## Docker notes
[INFO:](https://docker-curriculum.com/)
`docker ps -a` process
`docker rm 305297d7a235 ff0a5c3750b9`
`docker rm $(docker ps -a -q -f status=exited)` or `docker container prune` rm stopped
`docker images` available images
`docker pull ubuntu:18.04`
`docker start` - start already created container
`docker exec` - run smth. in running container
`docker attach` - attach to process which start a container
*Note: `Ctr+C` can kill container or any command to `exit`, to exit without killing container `Ctr+P`+`Ctr+Q`*
`-v host_path:conteiner_mount_point` - voluems
`-p host_port:container_port`  

### commit
`docker container ls`
`docker commit id_or_name repo/name_in_repo:tag` 
`docker login` 
`docker push repo/name_in_repo`

### [Dockerfile reference](https://docs.docker.com/engine/reference/builder/)
```
FROM
VOLUME - /mount_point
WORKDIR - /wordir for RUN, CMD....
COPY ./ /where_to_copy - normaly VOLUME, copy all from curret dir
RUN
CMD
EXPOSE 80/udp - it really not publish, it's just an NOTE
```
`docker build -t repo/name_in_repo:tag .` - DOckerfile in current dir or -f





## Dockerfile
```
FROM python:3

# set a directory for the app
WORKDIR /usr/src/app

# copy all the files to the container
COPY . .

# install dependencies
RUN pip install --no-cache-dir -r requirements.txt

# define the port number the container should expose
EXPOSE 5000

# run the command
CMD ["python", "./app.py"]
```
`docker build -t yourusername/catnip .`  username should be the same one you created when you registered on Docker hub
`docker push yourusername/catnip`
`docker run -p 8888:5000 yourusername/catnip`  external_port:internal_port

```
FROM alpine:latest

RUN apk update
RUN apk upgrade
RUN apk add dnsmasq
```



