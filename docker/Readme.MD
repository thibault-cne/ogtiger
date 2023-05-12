# ARMv7 Testing docker

## Required

- docker
- docker-buildx
- qemu-user-static
- qemu-user-binfmt

## Build and run the docker on linux

`sudo docker buildx create --name mybuilder`
`sudo docker buildx use mybuilder`
`sudo docker buildx inspect --bootstrap`
`sudo docker run --rm --privileged multiarch/qemu-user-static --reset -p yes`

To check the installation: `docker buildx ls`

To build : `sudo docker buildx build -t arm --output type=docker --platform linux/arm/v7 .`

In order to run your armv7 script you need to put in in the `test.s` file inside this folder.

To run : `sudo docker run --platform linux/arm/v7 --rm -it arm /bin/sh`
Or : `sudo docker run --platform linux/arm/v7 --rm -it arm /prgm`

## Build and run the docker on macos

`sudo docker buildx create --name mybuilder`
`sudo docker buildx use mybuilder`

To check the installation: `docker buildx ls`

To build : `sudo docker buildx build -t arm --output type=docker --platform linux/arm/v7 .`

In order to run your armv7 script you need to put in in the `test.s` file inside this folder.

To run : `sudo docker run --platform linux/arm/v7 --rm -it arm /bin/sh`
Or : `sudo docker run --platform linux/arm/v7 --rm -it arm /prgm`
