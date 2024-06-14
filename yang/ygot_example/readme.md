## YGOT

YGOT is a YANG centric toolkit: https://github.com/openconfig/ygot

One of the things you can do with it is generate Go code from yang models. 

This code does just that:
`main.go`: creates and validates a struct using code that was generated using ygot
`/example/exampleconfig.go`: contains the code that ygot generated
`example-config.yang`: the Yang model used in this example


Some steps to build it yourself (your mileage may vary):


```
docker pull golang
sudo docker run --name='go_learn' --hostname='go_learn' -p 8080:8080  -di golang:latest /bin/sh
sudo docker exec -it -u root go_learn bash
adduser klundert
usermod -aG sudo klundert
su --login klundert
go env -w GOPROXY=direct

go get github.com/openconfig/ygot/generator
go install github.com/openconfig/ygot/generator
go get github.com/openconfig/ygot/ygot

generator -path="/home/klundert/learningGo/yang/ygot_example/example/" -output_file="/home/klundert/learningGo/yang/ygot_example/example/exampleconfig.go" -package_name=example example-config.yang
```





Note:

you can generate it on Windows as well. It will looks something like this:
```
generator -path="C:\dev\learningGo\yang\ygot_example" -output_file="C:\dev\learningGo\yang\ygot_example\ygot_example\exampleconfig.go" -package_name=ygot_sample example-config.yang'
```