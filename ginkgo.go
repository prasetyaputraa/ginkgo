package ginkgo

import (
	conf "github.com/prasetyaputraa/ginkgo/conf"
	providerApplications "github.com/prasetyaputraa/ginkgo/providers/applications"
)

type Providers struct {
	applications func(conf conf.Configuration)
}

func GetProviders() Providers {
	return Providers{
		applications: providerApplications.Provide,
	}
}
