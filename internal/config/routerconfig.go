package config

type RouterConfig struct {
	RoutePrefix string
	Port        string `required:"true" split_words:"true" default:"50051"`
}

type ServerConfig struct {
	ShortLinkName string `required:"true" split_words:"true" default:"innocect"`
}
