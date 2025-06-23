golang make it easy to build apps for different os (window, unix, macod.. this is called cross complicatuon and one of the best superpower or Go)

Go uses environment variable to control the  build
Goos (go-os)= target operating system
Goarch (go-Arch)= target Architecture like amd64, arm64 etc

GOOS=windows GOARCH=amd64 go build -o app.exe main.go
GOOS=linux GOARCH=amd64 go build -o app main.go
GOOS=darwin GOARCH=amd64 go build -o app main.go  // for mac


check what os/arch you are using : 

go env Goos
go env goarch