#==================================================================
# 查看帮助:
#	make help
#
# Reference:
#	https://shields.io/
#	https://makefiletutorial.com/
#==================================================================
IMAGE_NAME := im_server
TAG := latest

.PHONY : build run

build:
	docker build ./ -t $(IMAGE_NAME):$(TAG)

run:
	docker run  -d --name $(IMAGE_NAME) -p 8080:8080 $(IMAGE_NAME):$(TAG)

stop:
	docker rm -f $(IMAGE_NAME)

