docker-build:
	docker build --no-cache -t go-webservice:latest .
docker-deploy:
	docker run --name gowebservice -d -p 10000:10000 go-webservice:latest
docker-stop:
	docker stop gowebservice
local-build:
	go build .
