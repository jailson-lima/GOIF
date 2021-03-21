package component

import (
	"bytes"
	"fmt"
	"github.com/JPereirax/GOIF/extensions"
	"github.com/JPereirax/GOIF/types"
	"io/ioutil"
	"net/http"
	"strings"
)

type HttpComponent struct{}

func init() {
	CreateComponent("http").Processor(HttpComponent{}.run).End()
}

const prefix = "//"
const applicationJson = "application/json"

func (HttpComponent) run(ctx types.ContextRequest) (*types.HttpTransport, error) {
	uri := ctx.Component.Uri
	if !strings.HasPrefix(uri, prefix) {
		return nil, extensions.SendError(ctx.HttpTransport, "Malformed request url: " + uri, http.StatusInternalServerError)
	}

	componentOptions := matchOptions(uri)
	if len(strings.TrimSpace(componentOptions)) == 0 {
		return nil, extensions.SendError(ctx.HttpTransport, "Options not found in request url: " + uri, http.StatusInternalServerError)
	}

	options := getOptions(componentOptions)
	method := strings.ToUpper(options["method"])

	requestBody := bytes.NewBuffer([]byte{})
	if method != http.MethodGet && len(ctx.HttpTransport.Body) > 0 {
		requestBody = bytes.NewBuffer([]byte(ctx.HttpTransport.Body))
	}

	uri = fmt.Sprintf("http:%s", strings.ReplaceAll(uri, fmt.Sprintf("[%s]", componentOptions), ""))
	request, _ := http.NewRequest(method, uri, requestBody)
	request.Header.Set("Accept", applicationJson)
	request.Header.Set("Content-type", applicationJson)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, extensions.SendError(ctx.HttpTransport, "Fail to send request to " + uri, 500)
	}
	defer response.Body.Close()

	responseBody, _ := ioutil.ReadAll(response.Body)
	ctx.HttpTransport.Body = string(responseBody)

	return ctx.HttpTransport, nil
}

func getOptions(componentOptions string) map[string]string {
	options := make(map[string]string)
	matches := strings.Split(componentOptions, "&")
	for _, match := range matches {
		keyValue := strings.Split(match, "=")
		options[keyValue[0]] = keyValue[1]
	}
	return options
}

func matchOptions(s string) string {
	i := strings.Index(s, "[")
	if i >= 0 {
		j := strings.Index(s, "]")
		if j >= 0 {
			return s[i+1 : j]
		}
	}
	return ""
}