package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/cucumber/godog"
	"github.com/gin-gonic/gin"
	"github.com/onsi/gomega"
	"github.com/sirupsen/logrus"
)

type apiFeature struct {
	router   *gin.Engine
	recorder *httptest.ResponseRecorder
}

func (a *apiFeature) reset() {
	a.router = SetRouter()
	a.recorder = httptest.NewRecorder()
}

func (a *apiFeature) iSendGETRequestToVersion() error {
	req, err := http.NewRequest("GET", "/version", nil)
	if err != nil {
		return err
	}

	a.router.ServeHTTP(a.recorder, req)
	return nil
}

func (a *apiFeature) iSendGETRequestToHealth() error {
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		return err
	}

	a.router.ServeHTTP(a.recorder, req)
	return nil
}

func (a *apiFeature) theResponseCodeShouldBe(code int) error {
	if code != a.recorder.Code {
		return fmt.Errorf("expected response code to be: %d, but actual is: %d", code, a.recorder.Code)
	}
	return nil
}

func (a *apiFeature) theResponseShouldMatchJson(body *godog.DocString) (err error) {
	defer failHandler(&err)
	want := map[string]interface{}{}
	actual := map[string]interface{}{}
	json.Unmarshal([]byte(body.Content), &want)
	json.Unmarshal(a.recorder.Body.Bytes(), &actual)
	gomega.Ω(actual).Should(gomega.Equal(want))
	return err
}

func failHandler(err *error) {
	if r := recover(); r != nil {
		*err = fmt.Errorf("%s", r)
	}
}
func InitializeScenario(ctx *godog.ScenarioContext) {
	api := &apiFeature{}

	ctx.Before(func(c context.Context, s *godog.Scenario) (context.Context, error) {
		logrus.Info("before")
		api.reset()
		return c, nil
	})

	ctx.After(func(c context.Context, s *godog.Scenario, err error) (context.Context, error) {
		return c, err
	})

	gomega.RegisterFailHandler(func(message string, _ ...int) {
		panic(message)
	})

	ctx.Step(`^I send GET request to \/version$`, api.iSendGETRequestToVersion)
	ctx.Step(`^I send GET request to \/health$`, api.iSendGETRequestToHealth)

	ctx.Step(`^the response code should be (\d+)$`, api.theResponseCodeShouldBe)
	ctx.Step(`^the response should match json:$`, api.theResponseShouldMatchJson)

}
