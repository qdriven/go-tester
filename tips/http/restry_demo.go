package http

import (
	"bytes"
	"fmt"
	"github.com/go-resty/resty/v2"
	"io/ioutil"
	"strconv"
	"time"
)

func DemoRestry(){
	// Create a Resty Client
	client := resty.New()

	resp, err := client.R().
		EnableTrace().
		Get("https://httpbin.org/get")

	// Explore response object
	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  Time       :", resp.Time())
	fmt.Println("  Received At:", resp.ReceivedAt())
	fmt.Println("  Body       :\n", resp)
	fmt.Println()

	// Explore trace info
	fmt.Println("Request Trace Info:")
	ti := resp.Request.TraceInfo()
	fmt.Println("  DNSLookup     :", ti.DNSLookup)
	fmt.Println("  ConnTime      :", ti.ConnTime)
	fmt.Println("  TCPConnTime   :", ti.TCPConnTime)
	fmt.Println("  TLSHandshake  :", ti.TLSHandshake)
	fmt.Println("  ServerTime    :", ti.ServerTime)
	fmt.Println("  ResponseTime  :", ti.ResponseTime)
	fmt.Println("  TotalTime     :", ti.TotalTime)
	fmt.Println("  IsConnReused  :", ti.IsConnReused)
	fmt.Println("  IsConnWasIdle :", ti.IsConnWasIdle)
	fmt.Println("  ConnIdleTime  :", ti.ConnIdleTime)
	fmt.Println("  RequestAttempt:", ti.RequestAttempt)
	fmt.Println("  RemoteAddr    :", ti.RemoteAddr.String())

	resp, err = client.R().
		SetQueryParams(map[string]string{
			"page_no": "1",
			"limit": "20",
			"sort":"name",
			"order": "asc",
			"random":strconv.FormatInt(time.Now().Unix(), 10),
		}).
		SetHeader("Accept", "application/json").
		SetAuthToken("BC594900518B4F7EAC75BD37F019E08FBC594900518B4F7EAC75BD37F019E08F").
		Get("/search_result")


	// Sample of using Request.SetQueryString method
	resp, err = client.R().
		SetQueryString("productId=232&template=fresh-sample&cat=resty&source=google&kw=buy a lot more").
		SetHeader("Accept", "application/json").
		SetAuthToken("BC594900518B4F7EAC75BD37F019E08FBC594900518B4F7EAC75BD37F019E08F").
		Get("/show_product")


	// If necessary, you can force response content type to tell Resty to parse a JSON response into your struct
	resp, err = client.R().
		SetResult(&AuthSuccess{}).
		ForceContentType("application/json").
		Get("v2/alpine/manifests/latest")
}

type AuthSuccess struct {

}

type AuthError struct {

}
type User struct {
	Username string
	Password string
}

func DemoPost(){
	// Create a Resty Client
	client := resty.New()

	// POST JSON string
	// No need to set content type, if you have client level setting
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(`{"username":"testuser", "password":"testpass"}`).
		SetResult(&AuthSuccess{}).    // or SetResult(AuthSuccess{}).
		Post("https://myapp.com/login")

	// POST []byte array
	// No need to set content type, if you have client level setting
	resp, err = client.R().
		SetHeader("Content-Type", "application/json").
		SetBody([]byte(`{"username":"testuser", "password":"testpass"}`)).
		SetResult(&AuthSuccess{}).    // or SetResult(AuthSuccess{}).
		Post("https://myapp.com/login")

	// POST Struct, default is JSON content type. No need to set one
	resp, err = client.R().
		SetBody(User{Username: "testuser", Password: "testpass"}).
		SetResult(&AuthSuccess{}).    // or SetResult(AuthSuccess{}).
		SetError(&AuthError{}).       // or SetError(AuthError{}).
		Post("https://myapp.com/login")

	// POST Map, default is JSON content type. No need to set one
	resp, err = client.R().
		SetBody(map[string]interface{}{"username": "testuser", "password": "testpass"}).
		SetResult(&AuthSuccess{}).    // or SetResult(AuthSuccess{}).
		SetError(&AuthError{}).       // or SetError(AuthError{}).
		Post("https://myapp.com/login")

	// POST of raw bytes for file upload. For example: upload file to Dropbox
	fileBytes, _ := ioutil.ReadFile("/Users/jeeva/mydocument.pdf")

	// See we are not setting content-type header, since go-resty automatically detects Content-Type for you
	resp, err = client.R().
		SetBody(fileBytes).
		SetContentLength(true).          // Dropbox expects this value
		SetAuthToken("<your-auth-token>").
		SetError(&AuthError{}).       // or SetError(DropboxError{}).
		Post("https://content.dropboxapi.com/1/files_put/auto/resty/mydocument.pdf") // for upload Dropbox supports PUT too

	// Note: resty detects Content-Type for request body/payload if content type header is not set.
	//   * For struct and map data type defaults to 'application/json'
	//   * Fallback is plain text content type
	fmt.Println(resp)
	fmt.Println(err)
}

func DemoPostFile(){
	profileImgBytes, _ := ioutil.ReadFile("/Users/jeeva/test-img.png")
	notesBytes, _ := ioutil.ReadFile("/Users/jeeva/text-file.txt")

	// Create a Resty Client
	client := resty.New()

	resp, err := client.R().
		SetFileReader("profile_img", "test-img.png", bytes.NewReader(profileImgBytes)).
		SetFileReader("notes", "text-file.txt", bytes.NewReader(notesBytes)).
		SetFormData(map[string]string{
			"first_name": "Jeevanandam",
			"last_name": "M",
		}).
		Post("http://myapp.com/upload")

	// Single file scenario
	resp, err = client.R().
		SetFile("profile_img", "/Users/jeeva/test-img.png").
		Post("http://myapp.com/upload")

	// Multiple files scenario
	resp, err = client.R().
		SetFiles(map[string]string{
			"profile_img": "/Users/jeeva/test-img.png",
			"notes": "/Users/jeeva/text-file.txt",
		}).
		Post("http://myapp.com/upload")

	// Multipart of form fields and files
	resp, err = client.R().
		SetFiles(map[string]string{
			"profile_img": "/Users/jeeva/test-img.png",
			"notes": "/Users/jeeva/text-file.txt",
		}).
		SetFormData(map[string]string{
			"first_name": "Jeevanandam",
			"last_name": "M",
			"zip_code": "00001",
			"city": "my city",
			"access_token": "C6A79608-782F-4ED0-A11D-BD82FAD829CD",
		}).
		Post("http://myapp.com/profile")
	fmt.Println(resp)
	fmt.Println(resp)
	fmt.Println(err)

}
