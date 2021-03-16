package extensions

import (
	"encoding/json"
	"errors"
	"github.com/JPereirax/GOIF/types"
	"net/http"
)

func SendError(transport *types.HttpTransport, message string, code int) error {
	setDefaultHeaders(transport.Response, code)

	json.NewEncoder(transport.Response).Encode(&types.ErrorResponse{
		Status:  code,
		Message: message,
	})

	return errors.New(message)
}

func SendJson(transport *types.HttpTransport, v interface{}, code int) error {
	setDefaultHeaders(transport.Response, code)

	err := json.Unmarshal([]byte(transport.Body), v)
	if err != nil {
		return SendError(transport, "Fail to return json object.", http.StatusInternalServerError)
	}

	json.NewEncoder(transport.Response).Encode(v)
	return nil
}

func setDefaultHeaders(response http.ResponseWriter, code int) {
	response.Header().Set("Content-Type", "application/json; charset=utf-8")
	response.Header().Set("X-Content-Type-Options", "nosniff")

	response.WriteHeader(code)
}