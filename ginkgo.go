package ginkgo

import (
	conf "github.com/prasetyaputraa/ginkgo/conf"
	providerApplications "github.com/prasetyaputraa/ginkgo/providers/applications"
	providerModels "github.com/prasetyaputraa/ginkgo/providers/models"
)

type Providers struct {
	Applications func(conf conf.Configuration) providerApplications.ApplicationsCollection
	Models       func(conf conf.Configuration) providerModels.ModelsCollection
}

// get all the providers
func GetProviders() Providers {
	return Providers{
		Applications: providerApplications.Provide,
		Models:       providerModels.Provide,
	}
}
