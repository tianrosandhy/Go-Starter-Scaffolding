package rest

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type RestClient struct {
	Debug   *bool
	URL     string
	Method  string
	Timeout int
	Headers map[string]string
	Request interface{}
	Log     *logrus.Logger
}

func NewRestClient(log *logrus.Logger) *RestClient {
	return &RestClient{
		Log: log,
	}
}

func (r *RestClient) SetDebug(debug bool) *RestClient {
	r.Debug = &debug
	return r
}

func (r *RestClient) SetURL(url string) *RestClient {
	r.URL = url
	return r
}

func (r *RestClient) SetMethod(method string) *RestClient {
	allowedMethods := []string{"get", "post", "delete", "patch", "put", "options", "head"}
	if _, hasSlice := FindSlice(allowedMethods, strings.ToLower(method)); !hasSlice {
		method = "get"
	}
	r.Method = method
	return r
}

func (r *RestClient) SetTimeout(timeout int) *RestClient {
	if timeout <= 1 {
		timeout = 15 //static default timeout.
	}
	r.Timeout = timeout
	return r
}

func (r *RestClient) SetHeaders(headers map[string]string) *RestClient {
	for hname, hval := range headers {
		r = r.AddHeader(hname, hval)
	}
	return r
}

func (r *RestClient) AddHeader(headerName, headerValue string) *RestClient {
	if len(r.Headers) == 0 {
		r.Headers = make(map[string]string)
	}
	r.Headers[headerName] = headerValue
	return r
}

func (r *RestClient) SetRequest(param interface{}) *RestClient {
	r.Request = param
	return r
}

func (r *RestClient) Execute() (httpBody string, httpStatus int) {
	var restrequest []byte
	restRequestID, _ := uuid.NewRandom()

	if len(r.Method) == 0 {
		r.Method = "GET"
	}
	if r.Timeout == 0 {
		r.Timeout = 15
	}

	var req *http.Request
	var err error

	switch request := r.Request.(type) {
	case string:
		restrequest = []byte(request)
		req, err = http.NewRequest(r.Method, r.URL, strings.NewReader(string(restrequest)))
	case []byte:
		req, err = http.NewRequest(r.Method, r.URL, bytes.NewReader(request))
	case *bytes.Buffer:
		// log
		req, err = http.NewRequest(r.Method, r.URL, request)
	default:
		restrequest, _ = json.Marshal(r.Request)
		req, err = http.NewRequest(r.Method, r.URL, bytes.NewReader(restrequest))
	}

	if err != nil {
		httpStatus = 0
		httpBody = "Error creating request : " + err.Error()
		r.Log.Printf("Error creating request : %s", err.Error())
		return
	}

	for hname, hval := range r.Headers {
		req.Header.Set(hname, hval)
	}
	req.Close = true // this is required to prevent too many files open

	// Create HTTP Connection
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
		Timeout: time.Duration(r.Timeout) * time.Second,
	}

	isDebug := true
	if r.Debug != nil {
		isDebug = *r.Debug
	}

	// Now hit to destionation endpoint
	if isDebug {
		loggedRequest := string(restrequest)
		if len(loggedRequest) > 1024 {
			loggedRequest = loggedRequest[:1024] + "... (log first 1KB only)"
		}

		r.LogStruct(map[string]interface{}{
			"REQUESTID": restRequestID.String(),
			"URL":       r.URL,
			"METHOD":    r.Method,
			"REQUEST":   loggedRequest,
		}, "RESTCLIENT REQUEST LOG")
	}

	res, err := client.Do(req)
	if err != nil {
		r.Log.Printf("Call URL Failed : %s", err.Error())
		httpStatus = 0
		httpBody = "Call URL Failed : " + err.Error()
		return
	}
	defer res.Body.Close()

	buff := new(bytes.Buffer)
	buff.ReadFrom(res.Body)
	httpBody = buff.String()
	httpStatus = res.StatusCode

	if isDebug {
		r.LogStruct(map[string]interface{}{
			"REQUESTID": restRequestID.String(),
			"RESPONSE":  httpBody,
		}, "REST CLIENT RESPONSE LOG")
	}

	return
}

// LogStruct func
func (r *RestClient) LogStruct(data interface{}, message ...string) string {
	byteData, _ := json.Marshal(data)
	var prefix string
	if len(message) > 0 {
		prefix = message[0] + " : "
	}

	printedLog := fmt.Sprintf("%s%s", prefix, byteData)
	r.Log.Println(printedLog)
	return printedLog
}

// FindSlice Find string on slice
func FindSlice(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}
