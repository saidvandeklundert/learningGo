go get github.com/openconfig/ygot/generator
go install github.com/openconfig/ygot/generator
go get github.com/openconfig/ygot/ygot

generator -path="C:\dev\learningGo\yang\ygot_example" -output_file="C:\dev\learningGo\yang\ygot_example\ygot_example\exampleconfig.go" -package_name=ygot_sample example-config.yang'


```
docker pull golang
sudo docker run --name='go_learn' --hostname='go_learn' -p 8080:8080  -di golang:latest /bin/sh
```