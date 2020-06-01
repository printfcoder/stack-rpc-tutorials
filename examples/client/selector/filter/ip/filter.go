package main

import (
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/v2/client/selector"
	"github.com/micro/go-micro/v2/registry"
)

// 过滤器，指定IP
// 也可改成传入数字，低于或大于某版本号的不调用，可灵活配置
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
