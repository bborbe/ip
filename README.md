# Ip

## Install

`go get github.com/bborbe/ip/cmd/ip-server`

## Run

```
ip-server \
-logtostderr \
-v=2 \
-port=8080
```

## Docker

```
docker run \
--env PORT=8080 \
--publish 8080:8080 \
docker.io/bborbe/ip:1.1.0 \
-logtostderr \
-v=1
```

```
curl http://localhost:8080
```
