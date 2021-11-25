package healthcheck

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/valyala/fasthttp"
)

// Todo: user interface as argument

// HealthCheck : Check api health
func HealthCheck(ctx *fasthttp.RequestCtx) {
	json := jsoniter.ConfigCompatibleWithStandardLibrary

	r := struct {
		Name string `json:"route_name"`
	}{
		Name: "health check",
	}
	response, err := json.Marshal(&r)

	// Todo: move this code to a separated function specialized on building the final response
	ctx.Response.Header.Set("Content-Type", "application/json")
	ctx.Response.SetStatusCode(fasthttp.StatusOK)
	if err != nil {
		return
	}
	_, err = ctx.Write(response)
	if err != nil {
		return
	}
}
