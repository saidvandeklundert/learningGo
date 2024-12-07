# include.mk

# Define a default value for the variable
GREETING ?= World

hello:
	@echo "Hello, $(GREETING)!"