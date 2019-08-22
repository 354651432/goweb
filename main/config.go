package main

type Config struct {
	Server struct {
		Listen string
	}
	DB struct {
		Dsn string
	}
}
