package main

import (
	"errors"
	"fmt"
)

// Option 1: config struct
type Config struct {
	Port int
}

// Option 2: builder patter
type ConfigBuilderConfig struct {
	Port int
	Name string
}

type ConfigBuilder struct {
	port *int
	name *string
}

func (b *ConfigBuilder) Port(port int) *ConfigBuilder {
	b.port = &port
	return b
}

func (b *ConfigBuilder) Name(name string) *ConfigBuilder {
	b.name = &name
	return b
}
func (b *ConfigBuilder) Build() (ConfigBuilderConfig, error) {
	cfg := ConfigBuilderConfig{}
	if b.name == nil {
		cfg.Name = "default"
	} else {
		cfg.Name = *b.name
	}
	if b.port == nil {
		cfg.Port = 80
	} else {
		if *b.port < 0 {
			return ConfigBuilderConfig{}, errors.New("port should be positive")
		} else {
			cfg.Port = *b.port
		}
	}
	return cfg, nil
}

/*
Option 3: Functional options pattern

The idea here is as follows:
- An unexported struct holds the configuration: options.
- Each option is a function that returns the same type:

'type Option func(options *options) error'

For example, WithPort accepts an int argument that represents the port and returns an Option type that represents how
to update the options struct.
*/

const defaultHTTPPort = 8080

type Server struct {
	Port int
	Name string
}
type options struct {
	port *int
	name *string
}

type Option func(options *options) error

func WithPort(port int) Option {
	return func(options *options) error {
		if port < 0 {
			return errors.New("port should be positive")
		}
		options.port = &port
		return nil
	}
}

func WithName(name string) Option {
	return func(options *options) error {

		options.name = &name
		return nil
	}
}

func NewServer(opts ...Option) (*Server, error) {
	var options options

	for _, opt := range opts {
		err := opt(&options)
		if err != nil {
			return nil, err
		}
	}

	// At this stage, the options struct is built and contains the config
	// Therefore, we can implement our logic related to port configuration
	var port int
	var name string
	if options.port == nil {
		port = defaultHTTPPort
	} else {
		if *options.port == 0 {
			port = randomPort()
		} else {
			port = *options.port
		}
	}

	_ = port

	if options.name != nil {
		name = *options.name
	} else {
		name = "default"
	}
	return &Server{Port: port, Name: name}, nil
}

func randomPort() int {
	return 4 // Chosen by fair dice roll, guaranteed to be random.
}
func main() {
	fmt.Println("option 1: config struct")
	config := Config{Port: 1}
	fmt.Printf("%+v\n", config)
	fmt.Println("option 2: builder pattern")
	config_builder := ConfigBuilder{}
	config_builder.Port(8080)
	config_builder.Name("something")
	configuration, err := config_builder.Build()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", configuration)
	fmt.Println("option 3: functional options pattern")

	server, _ := NewServer(WithPort(80), WithName("non-default-name"))
	fmt.Printf("%+v\n", server)
	default_server, _ := NewServer()
	fmt.Printf("%+v\n", default_server)
}
