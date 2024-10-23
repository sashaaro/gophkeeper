package config

const DefaultServerAddress = "127.0.0.1:9876"

type Client struct {
	ServerAddr string
}

func NewClient() *Client {
	return &Client{
		ServerAddr: DefaultServerAddress,
	}
}
