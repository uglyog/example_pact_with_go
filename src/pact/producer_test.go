package pact

import (
	"encoding/json"
	"github.com/bmizerany/assert"
	"github.com/drewolson/testflight"
	"github.com/kr/pretty"
	"github.com/paulbellamy/mango"
	"io/ioutil"
	"pact"
	"testing"
)

func TestProducer(t *testing.T) {
	stack := new(mango.Stack)
	handler := stack.HandlerFunc(pact.Producer)

	testflight.WithServer(handler, func(r *testflight.Requester) {

		pact_str, err := ioutil.ReadFile("../pacts/my_consumer-my_producer.json")
		if err != nil {
			t.Error(err)
		}

		pacts := make(map[string]interface{})
		err = json.Unmarshal(pact_str, &pacts)
		if err != nil {
			t.Error(err)
		}

		for _, i := range pacts["interactions"].([]interface{}) {
			interaction := i.(map[string]interface{})
			t.Logf("Given %s", interaction["producer_state"])
			t.Logf("  %s", interaction["description"])

			request := interaction["request"].(map[string]interface{})
			var actualResponse *testflight.Response
			switch request["method"] {
			case "get":
				actualResponse = r.Get(request["path"].(string) + "?" + request["query"].(string))
			}

			expectedResponse := interaction["response"].(map[string]interface{})

			assert.Equal(t, int(expectedResponse["status"].(float64)), actualResponse.StatusCode)

			for k, v := range expectedResponse["headers"].(map[string]interface{}) {
				assert.Equal(t, v, actualResponse.RawResponse.Header[k][0])
			}

			responseBody := make(map[string]interface{})
			err = json.Unmarshal([]byte(actualResponse.Body), &responseBody)
			if err != nil {
				t.Error(err)
			}
			for _, diff := range pretty.Diff(expectedResponse["body"], responseBody) {
				t.Log(diff)
			}
			assert.Equal(t, expectedResponse["body"], responseBody)
		}
	})
}
