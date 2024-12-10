# TARGETS
# make all MY_VARIABLE=SOMETHINGELSE

MY_VARIABLE = Make!

define macro_example
	@echo "+++ invoked by $@ +++"
	echo "macro called"

endef

all: message build run compile invoke_include final_message

pre_message:
	echo "building"
	./make_support.sh
	# uncomment the following to write to STDERR and exit make
	#echo "Some error message" >&2;
	#exit 1

message: pre_message # pre_message is called as a pre-requisite here
	echo "a Go project with $(MY_VARIABLE)!"

build:
	go build -o bin/main main.go

run:
	go run main.go
	$(call macro_example, $@)


compile:
	echo "compiling for other architectures"
	echo "Compiling for every OS and Platform"
	GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 main.go
	GOOS=linux GOARCH=386 go build -o bin/main-linux-386 main.go
	GOOS=windows GOARCH=386 go build -o bin/main-windows-386 main.go

invoke_include:
	echo "invoke include"
	@echo "Invoking include.mk"
	@$(MAKE) -f include.mk hello GREETING="from the first Makefile"
	@$(MAKE) -f include.mk clone
	@$(MAKE) -f include.mk move

final_message:
	echo "all done now"
