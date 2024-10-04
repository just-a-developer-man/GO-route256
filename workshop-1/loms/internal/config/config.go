package config

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

type Config struct {
	Host string
	Port string
}

func New() *Config {
	config := &Config{
		Host: "localhost",
		Port: "8080",
	}

	if host, ok := os.LookupEnv("HOST"); ok {
		config.Host = host
	}

	if port, ok := os.LookupEnv("PORT"); ok {
		config.Port = port
	}

	return config
}

func checkHost(host string) error {
	if host == "localhost" {
		return nil
	}

	ip := net.ParseIP(host)
	if ip == nil {
		return fmt.Errorf("bad host: %s", host)
	}
	return nil
}

func checkPort(port string) error {
	portInt, err := strconv.Atoi(port)
	if err != nil {
		return fmt.Errorf("bad port: %s", port)
	}

	if portInt < 1 || portInt > 65535 {
		return fmt.Errorf("bad port: %s", port)
	}

	return nil
}
