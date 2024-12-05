# TARGETS

MY_VARIABLE = Make!
all: message build run compile final_message

pre_message:
	echo "building"

message: pre_message # pre_message is called as a pre-requisite here
	echo "a Go project with $(MY_VARIABLE)!"

build:
	go build -o bin/main main.go

run:
	go run main.go

compile:
	echo "compiling for other architectures"
	echo "Compiling for every OS and Platform"
	GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 main.go
	GOOS=linux GOARCH=386 go build -o bin/main-linux-386 main.go
	GOOS=windows GOARCH=386 go build -o bin/main-windows-386 main.go

final_message:
	echo "all done now"
