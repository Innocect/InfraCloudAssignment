package config

import "net/url"

type RouterConfig struct {
	DependencyURL url.URL
	RoutePrefix   string
}
