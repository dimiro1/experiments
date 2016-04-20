# Building

```sh
$ docker build -t my-golang-with-docker .
```

# Development

```sh
$ docker run -v /$(pwd):/go/src/github.com/experiments/docker -p 8080:8080 -it --rm my-golang-with-docker go run main.go
```

# Deployment

```sh
$ docker run -p 8080:8080 -it --rm my-golang-with-docker
```