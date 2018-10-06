cmd1: go build
cmd2: go run main.go
cmd3: docker build -t golang_webservice/web_service .
cmd4: docker run --publish 9092:9092 -t golang_webservice/web_service
cmd5: docker push golang_webservice/web_service

docker tag golang_webservice/web_service joelvivek/golang_webservice/web_service
