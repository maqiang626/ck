

# 1009-Docker

## 构建本地镜像

### Makefile (github.com/maqiang626/ck/httpserver/Makefile)

```makefile
export tag=v1.0
root:
	export ROOT=github.com/maqiang626/ck

build:
	echo "building httpserver binary"
	mkdir -p bin/amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/amd64 .

release: build
	echo "building httpserver container"
	docker build -t maqiang626/httpserver:${tag} .

push: release
	echo "pushing maqiang626/httpserver"
	docker push maqiang626/httpserver:${tag}

```





## 编写 Dockerfile 将练习 2.2 编写的 httpserver 容器化（请思考有哪些最佳实践可以引入到 Dockerfile 中来）

### Dockerfile (github.com/maqiang626/ck/httpserver/Dockerfile)

```dockerfile
FROM centos:centos7.9.2009

MAINTAINER maqiang

ENV MY_SERVICE_PORT=9001

LABEL multi.label1="value1" multi.label2="value2" other="value3"

ADD bin/amd64/httpserver /httpserver

# httpserver listen port
EXPOSE 9001

ENTRYPOINT /httpserver

```





## 将镜像推送至 Docker 官方镜像仓库

```shell
# 将镜像推送至 Docker 官方镜像仓库

[root@d1 httpserver]# 
[root@d1 httpserver]# docker login
Login with your Docker ID to push and pull images from Docker Hub. If you don't have a Docker ID, head over to https://hub.docker.com to create one.
Username: maqiang626
Password: 
WARNING! Your password will be stored unencrypted in /root/.docker/config.json.
Configure a credential helper to remove this warning. See
https://docs.docker.com/engine/reference/commandline/login/#credentials-store

Login Succeeded
[root@d1 httpserver]# echo $?
0
[root@d1 httpserver]# 

```



### make push 流程

```shell
# make push 流程

[root@d1 httpserver]# 
[root@d1 httpserver]# make push
echo "building httpserver binary"
building httpserver binary
mkdir -p bin/amd64
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/amd64 .
echo "building httpserver container"
building httpserver container
docker build -t maqiang626/httpserver:v1.0 .
Sending build context to Docker daemon  12.22MB
Step 1/7 : FROM centos:centos7.9.2009
 ---> eeb6ee3f44bd
Step 2/7 : MAINTAINER maqiang
 ---> Using cache
 ---> 09c2e63842cd
Step 3/7 : ENV MY_SERVICE_PORT=9001
 ---> Using cache
 ---> 962b47e2cb97
Step 4/7 : LABEL multi.label1="value1" multi.label2="value2" other="value3"
 ---> Using cache
 ---> 560393ce4e00
Step 5/7 : ADD bin/amd64/httpserver /httpserver
 ---> Using cache
 ---> faacd4964813
Step 6/7 : EXPOSE 9001
 ---> Using cache
 ---> 868062777223
Step 7/7 : ENTRYPOINT /httpserver
 ---> Using cache
 ---> f24ced221246
Successfully built f24ced221246
Successfully tagged maqiang626/httpserver:v1.0
echo "pushing maqiang626/httpserver"
pushing maqiang626/httpserver
docker push maqiang626/httpserver:v1.0
The push refers to repository [docker.io/maqiang626/httpserver]
4ab5e8e0c2ab: Pushed 
174f56854903: Mounted from library/centos 
v1.0: digest: sha256:a5cb1ec8c96a065b5c495ff1ad9718de118cdd041ff761a40bc5ef57d9cd8353 size: 740
[root@d1 httpserver]# echo $?
0
[root@d1 httpserver]# 
[root@d1 httpserver]# 

```





## 通过 Docker 命令本地启动 httpserver

```shell
# 通过 Docker 命令本地启动 httpserver

[root@d1 httpserver]# 
[root@d1 httpserver]# docker images
REPOSITORY              TAG              IMAGE ID       CREATED             SIZE
maqiang626/httpserver   v1.0             f24ced221246   About an hour ago   210MB
hello-world             latest           feb5d9fea6a5   2 weeks ago         13.3kB
centos                  centos7.9.2009   eeb6ee3f44bd   4 weeks ago         204MB
[root@d1 httpserver]# 
[root@d1 httpserver]# docker run -d maqiang626/httpserver:v1.0
dff271546295598846998af64a86348cf18f1af7ea6bc1977c7ad03b67977bf5
[root@d1 httpserver]# 
[root@d1 httpserver]# docker ps
CONTAINER ID   IMAGE                        COMMAND                  CREATED          STATUS          PORTS      NAMES
dff271546295   maqiang626/httpserver:v1.0   "/bin/sh -c /httpser…"   14 seconds ago   Up 12 seconds   9001/tcp   adoring_taussig
[root@d1 httpserver]# 

```





## 通过 nsenter 进入容器查看 IP 配置

```shell
# 通过 nsenter 进入容器查看 IP 配置

[root@d1 httpserver]# 
[root@d1 httpserver]# docker images
REPOSITORY              TAG              IMAGE ID       CREATED             SIZE
maqiang626/httpserver   v1.0             f24ced221246   About an hour ago   210MB
hello-world             latest           feb5d9fea6a5   2 weeks ago         13.3kB
centos                  centos7.9.2009   eeb6ee3f44bd   4 weeks ago         204MB
[root@d1 httpserver]# 
[root@d1 httpserver]# docker run -d maqiang626/httpserver:v1.0
dff271546295598846998af64a86348cf18f1af7ea6bc1977c7ad03b67977bf5
[root@d1 httpserver]# 
[root@d1 httpserver]# docker ps
CONTAINER ID   IMAGE                        COMMAND                  CREATED          STATUS          PORTS      NAMES
dff271546295   maqiang626/httpserver:v1.0   "/bin/sh -c /httpser…"   14 seconds ago   Up 12 seconds   9001/tcp   adoring_taussig
[root@d1 httpserver]# 
[root@d1 httpserver]# docker inspect -f {{.State.Pid}} dff271546295
18524
[root@d1 httpserver]# 
[root@d1 httpserver]# nsenter -t 18524 -n ip a
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
6: eth0@if7: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default 
    link/ether 02:42:ac:11:00:02 brd ff:ff:ff:ff:ff:ff link-netnsid 0
    inet 172.17.0.2/16 brd 172.17.255.255 scope global eth0
       valid_lft forever preferred_lft forever
[root@d1 httpserver]# 
[root@d1 httpserver]# 

```



