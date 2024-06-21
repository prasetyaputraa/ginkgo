package models

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/prasetyaputraa/ginkgo/conf"
	"github.com/prasetyaputraa/ginkgo/providers"
)

func collectModels(goAppRootPath string, applications []string) ModelsCollection {

	return ModelsCollection{
		CollectionByProvider: providers.CollectionByProvider{
			Provider: "models",
		},
	}
}

type ModelsCollection struct {
	models []func(r *gin.RouterGroup, db *gorm.DB)
	providers.CollectionByProvider
}

func Provide(conf conf.Configuration) ModelsCollection {
	return collectModels(
		os.Getenv("GINKGO_ROOT_PATH"),
		conf.InstalledApplications,
	)
}
