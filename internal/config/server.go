package config

const DefaultServerListen = ":9876"

type Server struct {
	Listen string
}

func NewServer() *Server {
	return &Server{
		Listen: DefaultServerListen,
	}
}
