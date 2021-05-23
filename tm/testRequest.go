package tm

import (
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"go-tester/utils/json"
	"log"
)

type HttpMethod string

const (
	POST HttpMethod = "post"
	GET  HttpMethod = "get"
)

type HttpRequest struct {
	URL    string     `json:"url"`
	Method HttpMethod `json:"method"`
	//RequestParams contains query params for this WorkReq
	// only used if this is a GET WorkReq, ignored otherwise
	QueryParams map[string]string `json:"query_params"`
	PathParams  map[string]string `json:"path_params"`
	//RequestBody contains post body if this is a POST WorkReq
	// ignored otherwise
	RequestBody map[string]interface{} `json:"request_body"`

	// Headers to pass along with the WorkReq
	Headers map[string]string `json:"headers"`
}

type HttpTestCase struct {
	*HttpRequest
	ExpectedStatusCode int16                  `json:"expected_status_code"`
	ExpectedResponse   map[string]interface{} `json:"expected_response"`
	Debug              bool                   `json:"debug"`
	restyRunner        *resty.Client
}

func (tc *HttpTestCase) Run(runner *resty.Client) {

}

func (tc *HttpTestCase) runGet() error {
	if tc.Method != GET {
		return errors.New("Not a Get Request")
	}
	response, err := tc.restyRunner.R().SetHeaders(tc.Headers).SetQueryParams(tc.QueryParams).Get(tc.URL)

	return nil
}

func validateTestCase(testCase HttpTestCase, response *resty.Response, err error) error {
	if testCase.Debug {
		log.Printf("\n\n url = %+v, response = %+v, err = %+v \n\n", testCase.URL, response, err)
	}

	if testCase.ExpectedStatusCode < 400 && (err != nil || response.IsError()) {
		return err
	}

	parsedResponse, err := json.UnmarshalToMap(response.String())
	if err != nil {
		log.Fatal(err)
	}
	if testCase.Debug {
		log.Printf("\n\n parsedResponse = %+v\n\n", parsedResponse)
	}

	return matchResponseWithExpectedTypes(testCase.ExpectedResponse, parsedResponse)
}

func matchResponseWithExpectedTypes(expectedResponse map[string]interface{}, parsedResponse map[string]interface{}) error {
	for key, value := range expectedResponse {
		data := parsedResponse[key]
		switch value.(type) {
		case string:
			{

				if data.(string) != value.(string) {
					return errors.New(fmt.Sprintf("\n\ntype of %+v (%v) does not match. here's what was returned = %v\n\n", key, value.(string), data))
				}
			}

		case map[string]interface{}:
			{
				return matchResponseWithExpectedTypes(value.(map[string]interface{}), data)
			}
		}
	}
	return nil
}
