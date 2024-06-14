## Some udemy course

### Building a module in Go:
go version
go version go1.20.2 windows/amd64
### make a mod in the toolkit folder
cd .\toolkit
go mod init github.com/saidvandeklundert/toolkit
cd ..
### make a mod in the app folder
cd .\app\
go mod init myapp
### anchor the workspace
go work init toolkit app 

### put in some code and then from app, which contains main.go:
go run .


### run tests:
go test
go test . -v

### add something to workspace (from the workspace root)
go work use app-upload