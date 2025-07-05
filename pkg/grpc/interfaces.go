package grpc


type Client interface {
	StopClient() error
}

type Server interface {
	StartServer(target string) error
	StopServer() error
}