package response

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/iancoleman/strcase"
)

type ErrorData struct {
	Name      string      `json:"name" example:"Username"`                                 // data name
	Path      string      `json:"path" example:"user.username"`                            // object property path
	Type      string      `json:"type,omitempty" example:"string"`                         // data type
	Value     interface{} `json:"value,omitempty" swaggertype:"string" example:"jane doe"` // value
	Validator string      `json:"validator" example:"required"`                            // validator type, see [more details](https://github.com/go-playground/validator#baked-in-validations)
	Criteria  interface{} `json:"criteria,omitempty" swaggertype:"number" example:"10"`    // criteria, example: if validator is gte (greater than) and criteria is 10, then it means a maximum of 10
	Message   string      `json:"-" example:"invalid value"`                               // Field message
}

var validateMessages = map[string]string{
	"required": "%v is required %v",
	"gte":      "%v must be greater than or equal %v",
	"gt":       "%v must be greater than %v",
	"lte":      "%v must be less than or equal %v",
	"lt":       "%v must be less than %v",
	"unique":   "%v already exists %v",
	"email":    "%v must be a valid email address %v",
	"uuid":     "%v must be a valid UUID %v",
}

// Fail is a helper function to return a JSON response with status code 200
// param can be combination of nil, "message string, data", "data, message string", "data" only, or "message string" only
func Fail(c *fiber.Ctx, httpCode int, param ...interface{}) error {
	c.Status(httpCode)
	if len(param) > 0 {
		message := "Error"
		var data interface{}

		if msg, ok := param[0].(string); ok {
			message = msg
			if len(param) > 1 {
				data = param[1]
			}
		} else if len(param) > 1 {
			data = param[0]
			if msg, ok := param[1].(string); ok {
				message = msg
			}
		} else {
			data = param[0]
		}

		return c.JSON(Response{
			Type:      "error",
			Message:   message,
			ErrorData: data,
		})
	}

	return c.JSON(Response{
		Type:    "error",
		Message: "Something bad occured",
	})
}

func ErrorMessage(c *fiber.Ctx, errorMessage string, httpCode ...int) error {
	if len(httpCode) == 0 {
		httpCode = append(httpCode, fiber.StatusBadRequest)
	}

	return Fail(c, httpCode[0], errorMessage)
}

func Error(c *fiber.Ctx, err error, httpCode ...int) error {
	if len(httpCode) == 0 {
		httpCode = append(httpCode, fiber.StatusBadRequest)
	}

	message := "Something bad occured"
	var errorData interface{}

	if errs, ok := err.(validator.ValidationErrors); ok {
		message = "Request body does not meet the requirements"
		errorDetails := []ErrorData{}
		for _, err := range errs {
			fields := strings.Split(err.StructNamespace(), ".")
			field := strcase.ToSnake(err.Field())
			if len(fields) > 1 {
				names := []string{}
				for _, i := range fields[1:] {
					f := strcase.ToSnake(i)
					names = append(names, f)
				}
				field = strings.Join(names, ".")
			}
			fieldType := strings.ReplaceAll(fmt.Sprint(err.Type()), "*", "")
			regex := regexp.MustCompile(`.*\.([^\.]+)$`)
			fieldType = regex.ReplaceAllString(fieldType, "$1")
			errorData := ErrorData{
				Name:      strcase.ToSnake(err.Field()),
				Path:      field,
				Type:      fieldType,
				Value:     err.Value(),
				Validator: err.Tag(),
			}

			if i, e := strconv.Atoi(err.Param()); nil == e {
				errorData.Criteria = i
			} else {
				errorData.Criteria = err.Param()
			}

			if m, ok := validateMessages[errorData.Validator]; ok {
				errorData.Message = fmt.Sprintf(m, errorData.Name, errorData.Criteria)
			}
			errorDetails = append(errorDetails, errorData)
		}
		errorData = errorDetails
	} else {
		message = err.Error()
	}

	return Fail(c, httpCode[0], message, errorData)
}
