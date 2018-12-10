# Dockerised Microservices

## Pre-requisites:

* Go 1.11.2
* Docker

## Build

There is a make file that builds all microservices and a `docker-compose.yml` file to bring them all up.

```
make dockerise
```

Then to run the microservices, you can just do:

```
docker-compose up
```

Set the environment variable `ZEEBE_BROKER_ADDRESS` to use a remote broker.