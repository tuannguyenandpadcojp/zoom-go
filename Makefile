example-video:
	export $(shell cat .env | xargs) && go run example/video/main.go

host-token:
	curl --location --request POST 'http://localhost:8080/sessions'
