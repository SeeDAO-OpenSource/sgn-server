package apiserver

type ApiServerOptions struct {
	Port int
}

var AsOptions = &ApiServerOptions{
	Port: 5000,
}
