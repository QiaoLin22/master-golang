package main

import "fmt"

type Server struct {
	config Config
}

func NewServer() (*Server, error) {
	return &Server{
		config: NewConfig(),
	}, nil
}

func NewServerWithConfig(config Config) (*Server, error) {
	return &Server{
		config: config,
	}, nil
}

type Config struct {
	listenAddr string
	id         string
	name       string
}

func (c Config) WithListenAddr(addr string) Config {
	c.listenAddr = addr
	return c
}

func (c Config) WithID(id string) Config {
	c.id = id
	return c
}

func (c Config) WithName(name string) Config {
	c.name = name
	return c
}

func NewConfig() Config {
	return Config{
		id:         "id",
		listenAddr: "addr",
		name:       "name",
	}
}

func main() {
	config := NewConfig().
		WithID("id").
		WithListenAddr("addr").
		WithName("name")
	server, _ := NewServerWithConfig(config)
	fmt.Println(server)
}
