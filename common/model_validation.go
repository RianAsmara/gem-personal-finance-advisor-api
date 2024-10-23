package common

import (
	"encoding/json"

	"github.com/RianAsmara/personal-finance-advisor-api/exception"
	"gopkg.in/go-playground/validator.v9"
)

func ValidateStruct(data interface{}) {
	validate := validator.New()
	err := validate.Struct(data)
	if err != nil {
		var messages []map[string]interface{}
		for _, err := range err.(validator.ValidationErrors) {
			messages = append(messages, map[string]interface{}{
				"field":   err.Field(),
				"message": err.Tag(),
			})
		}

		jsonMessages, err := json.Marshal(messages)
		exception.PanicLogging(err)

		panic(exception.ValidationError{
			Message: string(jsonMessages),
		})
	}
}
