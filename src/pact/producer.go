package pact

import (
	"github.com/paulbellamy/mango"
	"encoding/json"
	"strings"
)

func Producer(env mango.Env) (status mango.Status, headers mango.Headers, body mango.Body) {
  env.Logger().Println("Got a", env.Request().Method, "request for", env.Request().RequestURI)

	if strings.HasPrefix(env.Request().RequestURI, "/producer.json") {
		status = 200
		headers = mango.Headers{"Content-Type": []string{"application/json;charset=utf-8"}}

		json_data := make(map[string]interface{})
		json_data["test"] = "NO"
		json_data["date"] = "2013-08-16T15:31:20+10:00"
		json_data["count"] = 1000

		json_string, _ := json.Marshal(json_data)
		body = mango.Body(json_string)
	} else {
		status = 404
		body = mango.Body("Invalid path")
	}
	return
}

