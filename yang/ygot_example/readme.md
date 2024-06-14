go get github.com/openconfig/ygot/generator
go install github.com/openconfig/ygot/generator
go get github.com/openconfig/ygot/ygot

generator -path="/home/klundert/learningGo/yang/ygot_example/example/" -output_file="/home/klundert/learningGo/yang/ygot_example/example/exampleconfig.go" -package_name=example example-config.yang



```
docker pull golang
sudo docker run --name='go_learn' --hostname='go_learn' -p 8080:8080  -di golang:latest /bin/sh
adduser klundert
usermod -aG sudo klundert
su --login klundert
go env -w GOPROXY=direct
```







Note:

you can generate it on Windows as well. It will looks something like this:
```
generator -path="C:\dev\learningGo\yang\ygot_example" -output_file="C:\dev\learningGo\yang\ygot_example\ygot_example\exampleconfig.go" -package_name=ygot_sample example-config.yang'
```