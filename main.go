package main

import (
	"github.com/sirupsen/logrus"

	"github.com/ChenYuTingJerry/gin-bdd-godog/api"
)

func main() {
	r := api.SetRouter()
	logrus.Info("listen and serve on 0.0.0.0:8080 ")
	r.Run("0.0.0.0:8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
