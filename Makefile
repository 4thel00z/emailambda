build:
	GOOS=linux
	GOARCH=amd64
	GO111MODULE=on
	go build -o ${PWD}/functions/emailambda cmd/emaillambda/main.go
clean:
	rm -rf functions/*
