package main

import (
	"flag"
	"fmt"
)

type Config struct {
	Homeserver string
	Username   string
	Password   string
}

type ErrMissingArg string

func (e ErrMissingArg) Error() string {
	return fmt.Sprint(string(e))
}

func parseArgs() (Config, error) {
	homeserver := flag.String("homeserver", "https://matrix.org", "")
	username := flag.String("username", "", "")
	password := flag.String("password", "", "")

	flag.Parse()

	if *username == "" {
		return Config{}, ErrMissingArg("username empty")
	}
	if *password == "" {
		return Config{}, ErrMissingArg("password empty")
	}

	config := Config{
		Homeserver: *homeserver,
		Username:   *username,
		Password:   *password,
	}

	return config, nil
}
