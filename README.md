# rhizome the meta cryptocurrency

Rhizome allows you to build your own cryptocurrency without any friction.

[Livre blanc FR](./doc/fr)

# Installation requirements

- [Docker](https://www.docker.com/) >= 2.0.0
- [Golang](https://golang.org/) >= 1.8.0

# Setting up dev environment (regtest network)

Build docker's image and build the containers
```
docker-composer up --build
```

Build multiple nodes where n is the number of nodes
```
docker-compose up --build --scale node=n
```
