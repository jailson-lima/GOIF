package component

import (
	"GOIF/extensions"
	"GOIF/types"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type HttpComponent struct{}

func init() {
	CreateComponent("http").Processor(HttpComponent{}.run).End()
}

func (HttpComponent) run(ctx types.ContextRequest) (*types.HttpTransport, error) {
	uri := ctx.Component.Uri
	if !strings.HasPrefix(uri, "//") {
		return nil, extensions.SendError(ctx.HttpTransport, "Malformed request url: " + uri, http.StatusInternalServerError)
	}

	componentOptions := matchOptions(uri)
	if len(strings.TrimSpace(componentOptions)) == 0 {
		return nil, extensions.SendError(ctx.HttpTransport, "Options not found in request url: " + uri, http.StatusInternalServerError)
	}

	options := make(map[string]string)
	matches := strings.Split(componentOptions, "&")
	for _, match := range matches {
		keyValue := strings.Split(match, "=")
		options[keyValue[0]] = keyValue[1]
	}

	uri = fmt.Sprintf("http:%s", strings.ReplaceAll(uri, fmt.Sprintf("[%s]", componentOptions), ""))
	request, _ := http.NewRequest(strings.ToUpper(options["method"]), uri, bytes.NewBuffer([]byte{}))

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, extensions.SendError(ctx.HttpTransport, "Fail to send request to " + uri, 500)
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
	ctx.HttpTransport.Body = string(body)

	return ctx.HttpTransport, nil
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