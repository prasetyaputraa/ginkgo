package applications

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"plugin"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/prasetyaputraa/ginkgo/conf"
	"github.com/prasetyaputraa/ginkgo/providers"
)

func collectApplications(goAppRootPath string, applications []string) ApplicationsCollection {
	var pluginsCollection []string

	for index := range applications {
		fmt.Println("building " + applications[index] + "...")
		applicationPath := filepath.Join(goAppRootPath, applications[index])

		err := exec.Command("go build -buildmode=plugin ./" + applicationPath).Run()

		if err != nil {
			log.Fatal("failed to build " + applications[index])
			continue
		}

		pluginsCollection = append(pluginsCollection, applications[index]+".so")
	}

	applicationsCollection := ApplicationsCollection{
		apps: []func(r *gin.RouterGroup, db *gorm.DB){},

		CollectionByProvider: providers.CollectionByProvider{
			Provider: "applications",
		},
	}

	for index := range pluginsCollection {
		fmt.Println("loading " + pluginsCollection[index] + "...")
		p, err := plugin.Open(pluginsCollection[index])

		if err != nil {
			log.Fatal("failed to load " + pluginsCollection[index])
		}

		App, err := p.Lookup("App")

		if err != nil {
			panic("application must have function App")
		}

		applicationsCollection.appendApp(App.(func(r *gin.RouterGroup, db *gorm.DB)))
	}

	return applicationsCollection
}

type ApplicationsCollection struct {
	apps []func(r *gin.RouterGroup, db *gorm.DB)
	providers.CollectionByProvider
}

func (appsCollection ApplicationsCollection) appendApp(app func(r *gin.RouterGroup, db *gorm.DB)) {
	appsCollection.apps = append(appsCollection.apps, app)
}

func (appsCollection ApplicationsCollection) GetApps() []func(r *gin.RouterGroup, db *gorm.DB) {
	return appsCollection.apps
}

func Provide(conf conf.Configuration) ApplicationsCollection {
	return collectApplications(
		os.Getenv("GINKGO_ROOT_PATH"),
		conf.InstalledApplications,
	)
}
