package main

type Config struct {
	Server struct {
		Listen     string
		PublicPath string
	}
	DB struct {
		Dsn string
	}
}
