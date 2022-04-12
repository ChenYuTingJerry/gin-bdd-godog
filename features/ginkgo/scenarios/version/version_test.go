package version_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ChenYuTingJerry/gin-bdd-godog/api"
)

var _ = Describe("Version", func() {
	var request *http.Request
	var recorder *httptest.ResponseRecorder
	var router = api.SetRouter()

	BeforeEach(func() {
		recorder = httptest.NewRecorder()
		request, _ = http.NewRequest("GET", "/version", nil)
	})

	// Scenario
	Describe("should get version number", func() {
		//When
		Context("client send GET request to /version", Label("get"), func() {
			BeforeEach(func() {
				router.ServeHTTP(recorder, request)
			})

			//Then
			It("the response code should be 200", func() {
				Expect(recorder.Code).To(Equal(200))
			})

			//And
			It("the response should match json", func() {
				expectedJsonStr := `{ "version": "v1.0.0" }`
				want := map[string]interface{}{}
				actual := map[string]interface{}{}

				json.Unmarshal([]byte(expectedJsonStr), &want)
				json.Unmarshal(recorder.Body.Bytes(), &actual)
				Expect(actual).To(Equal(want))
			})
		})
	})
})
