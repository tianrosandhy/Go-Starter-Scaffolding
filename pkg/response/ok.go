package response

import (
	"github.com/labstack/echo/v4"
)

// OK is a helper function to return a JSON response with status code 200
// param can be combination of nil, "message string, data", "data, message string", "data" only, or "message string" only
func OK(c echo.Context, param ...interface{}) error {
	if len(param) > 0 {
		message := "OK"
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

		return c.JSON(200, Response{
			Type:    "success",
			Message: message,
			Data:    data,
		})
	}

	return c.JSON(200, Response{
		Type:    "success",
		Message: "OK",
	})
}
