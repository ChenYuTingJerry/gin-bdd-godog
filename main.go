package main

import (
	"github.com/ChenYuTingJerry/gin-bdd-godog/api"
	"github.com/sirupsen/logrus"
)

func main() {
	r := api.SetRouter()
	logrus.Info("listen and serve on 0.0.0.0:8080 ")
	r.Run("0.0.0.0:8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
