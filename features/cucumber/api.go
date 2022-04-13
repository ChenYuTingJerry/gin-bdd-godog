package cucumber

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"

	"github.com/ChenYuTingJerry/gin-bdd-godog/api"
)

var opts = godog.Options{Output: colors.Colored(os.Stdout)}

func init() {
	godog.BindCommandLineFlags("godog.", &opts)
}

func TestMain(m *testing.M) {
	pflag.Parse()
	opts.Paths = pflag.Args()

	status := godog.TestSuite{
		Name:                 "godogs",
		TestSuiteInitializer: InitializeTestSuite,
		ScenarioInitializer:  InitializeScenario,
		Options:              &opts,
	}.Run()

	os.Exit(status)
}

type apiFeature struct {
	router   *gin.Engine
	recorder *httptest.ResponseRecorder
	testing  testing.T
	assert   *assert.Assertions
	reqBody  map[string]interface{}
}

func (a *apiFeature) reset() {
	a.router = api.SetRouter()
	a.recorder = httptest.NewRecorder()
	a.assert = assert.New(&a.testing)
}

func (a *apiFeature) clientSendRequestTo(method, url string) error {
	strBody, _ := json.Marshal(a.reqBody)
	req, err := http.NewRequest(method, url, strings.NewReader(string(strBody)))
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

func (a *apiFeature) clientGiveARequestBody(body *godog.DocString) error {
	json.Unmarshal([]byte(body.Content), &a.reqBody)
	return nil
}

func InitializeTestSuite(ctx *godog.TestSuiteContext) {
	ctx.BeforeSuite(func() {})
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
	ctx.Step(`^client give a request body:$`, api.clientGiveARequestBody)

}
