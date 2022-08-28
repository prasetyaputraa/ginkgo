package ginkgo

import (
	conf "github.com/prasetyaputraa/ginkgo/conf"
	providerApplications "github.com/prasetyaputraa/ginkgo/providers/applications"
)

type Providers struct {
	Applications func(conf conf.Configuration)
}

// get all the providers
func GetProviders() Providers {
	return Providers{
		Applications: providerApplications.Provide,
	}
}
