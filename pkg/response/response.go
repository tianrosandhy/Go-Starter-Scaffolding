package response

type Response struct {
	Type      string      `json:"type"`
	Message   string      `json:"message"`
	ErrorData interface{} `json:"error_data,omitempty"`
	Data      interface{} `json:"data,omitempty"`
}
