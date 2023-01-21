package request

import (
	"encoding/json"
	"net/http"

	validate "github.com/go-playground/validator/v10"
)

func DecodeBody(r *http.Request, body interface{}) error {
	err := json.NewDecoder(r.Body).Decode(body)
	if err != nil {
		return err
	}

	if err := validate.New().Struct(body); err != nil {
		return err
	}

	return nil
}
