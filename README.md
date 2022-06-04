# HTTP server "hello world"

## Build Docker image:

```sh
docker build -t my-golang-app .
```

## Run Docker container:

```sh
docker run --env PORT=3000 -p 3000:3000 -it --rm --name my-running-app my-golang-app
```
