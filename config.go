package main

import (
	"flag"
	"fmt"
)

type Config struct {
	homeserver *string
	username   *string
	password   *string
}

type ErrMissingArg string

func (e ErrMissingArg) Error() string {
	return fmt.Sprint(string(e))
}

func parseArgs() (*Config, error) {
	homeserver := flag.String("homeserver", "https://matrix.org", "")
	username := flag.String("username", "", "")
	password := flag.String("password", "", "")

	flag.Parse()

	if *username == "" {
		return nil, ErrMissingArg("username empty")
	}
	if *password == "" {
		return nil, ErrMissingArg("password empty")
	}

	config := Config{
		homeserver: homeserver,
		username:   username,
		password:   password,
	}

	return &config, nil
}
