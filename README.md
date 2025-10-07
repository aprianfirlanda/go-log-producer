# README

## Build

download all dependency
```shell
go mod tidy
```

install ko for build go application as container
https://ko.build/install/

set env variable to put value docker repository
```shell
export KO_DOCKER_REPO=docker.io/aprianfirlanda/go-log-producer
```

build and push image
```shell
ko build --bare --tags 0.1.0 --platform=linux/amd64,linux/arm64 .
```

run as container
```shell
docker login docker.io
docker pull docker.io/aprianfirlanda/go-log-producer:0.1.0
docker run -d --name go-log-producer -p 8080:8080 docker.io/aprianfirlanda/go-log-producer:0.1.0
```
