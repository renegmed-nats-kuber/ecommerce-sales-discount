.PHONY: build
build:
	go build .

.PHONY: docker
docker:
	docker build -t ecommerce-sales-discount .

.PHONY: run
run:
	docker run --name ecommerce-sales-discount -d ecommerce-sales-discount

.PHONY: tag
tag:
	docker tag ecommerce-sales-discount:latest renegmedal/ecommerce-sales-discount:1.0.2

.PHONY: push
push:
	docker push renegmedal/ecommerce-sales-discount:1.0.2

.PHONY: up
up:
	docker-compose up --build -d
