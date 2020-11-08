package main

import (
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/client/selector"
	"github.com/micro/go-micro/v2/registry"
)

// 过滤器，指定IP
func IPFilter(v string) client.CallOption {
	filter := func(services []*registry.Service) []*registry.Service {
		var filtered []*registry.Service

		for _, service := range services {
			for _, node := range service.Nodes {
				if node.Address == v {
					filtered = append(filtered, service)
				}
			}
		}

		return filtered
	}

	return client.WithSelectOption(selector.WithFilter(filter))
}
