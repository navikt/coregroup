go-build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o coregroups .

docker-build:
	docker build -t coregroups .

docker-run:
	docker run --rm -it -p 8080:80 coregroups
