package applications

import (
	"fmt"
	"os"

	"github.com/prasetyaputraa/ginkgo/conf"
)

func collectApplications(goAppRootPath string, applications []string) {
	for application := range applications {
		fmt.Println(application)
	}
}

func Provide(conf conf.Configuration) {
	collectApplications(
		os.Getenv("GINKGO_ROOT_PATH"),
		conf.InstalledApplications,
	)
}
