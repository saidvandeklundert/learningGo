# include.mk

# Define a default value for the variable
GREETING ?= World

hello:
	@echo "Hello, $(GREETING)!"

# this is absolutely useless, just showing how to incorporate other repos in the build:
clone:
	rm -rf testing-examples
	git clone https://github.com/saidvandeklundert/testing-examples.git
	rm -rf bin/testing-examples

move:
	mkdir -p bin/testing-examples
	mv testing-examples/* bin/testing-examples/ 2>/dev/null || true
	mv testing-examples/.[!.]* bin/testing-examples/ 2>/dev/null || true
	rm -rf testing-examples