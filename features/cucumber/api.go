package cucumber

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cucumber/godog"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/ChenYuTingJerry/gin-bdd-godog/api"
)

type apiFeature struct {
	router   *gin.Engine
	recorder *httptest.ResponseRecorder
	testing  testing.T
	assert   *assert.Assertions
}

func (a *apiFeature) reset() {
	a.router = api.SetRouter()
	a.recorder = httptest.NewRecorder()
	a.assert = assert.New(&a.testing)
}

func (a *apiFeature) clientSendRequestTo(method, url string) error {
	req, err := http.NewRequest(method, url, nil)
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

func (a *apiFeature) theResponseShouldMatchJson(body *godog.DocString) error {
	want := map[string]interface{}{}
	actual := map[string]interface{}{}
	json.Unmarshal([]byte(body.Content), &want)
	json.Unmarshal(a.recorder.Body.Bytes(), &actual)
	if !a.assert.Equal(want, actual) {
		return fmt.Errorf("response not match. want: %s, actual: %s", want, actual)
	}
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	api := &apiFeature{}

	ctx.Before(func(c context.Context, s *godog.Scenario) (context.Context, error) {
		api.reset()
		return c, nil
	})

	ctx.After(func(c context.Context, s *godog.Scenario, err error) (context.Context, error) {
		return c, err
	})

	ctx.Step(`^client send "([^"]*)" request to "([^"]*)"$`, api.clientSendRequestTo)
	ctx.Step(`^the response code should be (\d+)$`, api.theResponseCodeShouldBe)
	ctx.Step(`^the response should match json:$`, api.theResponseShouldMatchJson)

}
