package test

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
)

// IsEqualString fails test if got and want are not identical
func IsEqualString(t *testing.T, got, want string) {
	if got != want {
		t.Errorf("Assertion failed, got: %s, want: %s.", got, want)
	}
}

// IsEqualBool fails test if got and want are not identical
func IsEqualBool(t *testing.T, got, want bool) {
	if got != want {
		t.Errorf("Assertion failed, got: %t, want: %t.", got, want)
	}
}

// IsEqualInt fails test if got and want are not identical
func IsEqualInt(t *testing.T, got, want int) {
	if got != want {
		t.Errorf("Assertion failed, got: %d, want: %d.", got, want)
	}
}

// IsNotEmpty fails test if string is empty
func IsNotEmpty(t *testing.T, s string) {
	if s == "" {
		t.Errorf("Assertion failed, got: %s, want: empty.", s)
	}
}

// IsEmpty fails test if string is not empty
func IsEmpty(t *testing.T, s string) {
	if s != "" {
		t.Errorf("Assertion failed, got: %s, want: empty.", s)
	}
}

// IsNil fails test if error not nil
func IsNil(t *testing.T, got error) {
	if got != nil {
		t.Errorf("Assertion failed, got: %s, want: nil.", got.Error())
	}
}

// HttpPageResult tests if a http server is outputting the correct result
func HttpPageResult(t *testing.T, config HttpTestConfig) []*http.Cookie {
	config.init()
	client := &http.Client{}

	data := url.Values{}
	for _, value := range config.PostValues {
		data.Add(value.Key, value.Value)
	}

	req, err := http.NewRequest(config.Method, config.Url, strings.NewReader(data.Encode()))
	IsNil(t, err)

	for _, cookie := range config.Cookies {
		req.Header.Set("Cookie", cookie.toString())
	}
	if len(config.PostValues) > 0 {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	}
	resp, err := client.Do(req)
	IsNil(t, err)

	if resp.StatusCode != 200 {
		t.Errorf("Status %d != 200", resp.StatusCode)
	}
	content, err := ioutil.ReadAll(resp.Body)
	IsNil(t, err)
	if config.IsHtml && !bytes.Contains(content, []byte("</html>")) {
		t.Error(config.Url + ": Incorrect response")
	}
	for _, requiredString := range config.RequiredContent {
		if !bytes.Contains(content, []byte(requiredString)) {
			t.Error(config.Url + ": Incorrect response. Got:\n" + string(content))
		}
	}
	for _, excludedString := range config.ExcludedContent {
		if bytes.Contains(content, []byte(excludedString)) {
			t.Error(config.Url + ": Incorrect response. Got:\n" + string(content))
		}
	}
	resp.Body.Close()
	return resp.Cookies()
}

// HttpTestConfig is a struct for http test init
type HttpTestConfig struct {
	Url             string
	RequiredContent []string
	ExcludedContent []string
	IsHtml          bool
	Method          string
	PostValues      []PostBody
	Cookies         []Cookie
	UploadFileName  string
	UploadFieldName string
}

func (c *HttpTestConfig) init() {
	if c.Url == "" {
		log.Fatalln("No url passed!")
	}
	if c.Method == "" {
		c.Method = "GET"
	}
}

// Cookie is a simple struct to pass cookie values for testing
type Cookie struct {
	Name  string
	Value string
}

func (c *Cookie) toString() string {
	return c.Name + "=" + c.Value
}

// PostBody contains mock key/value post data
type PostBody struct {
	Key   string
	Value string
}

// HttpPostRequest sends a post request
func HttpPostRequest(t *testing.T, config HttpTestConfig) {
	file, err := os.Open(config.UploadFileName)
	IsNil(t, err)
	defer file.Close()
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(config.UploadFieldName, filepath.Base(file.Name()))
	IsNil(t, err)

	io.Copy(part, file)
	writer.Close()
	request, err := http.NewRequest("POST", config.Url, body)
	IsNil(t, err)

	for _, cookie := range config.Cookies {
		request.Header.Set("Cookie", cookie.toString())
	}
	request.Header.Add("Content-Type", writer.FormDataContentType())
	client := &http.Client{}

	response, err := client.Do(request)
	IsNil(t, err)
	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	IsNil(t, err)

	for _, requiredString := range config.RequiredContent {
		if !bytes.Contains(content, []byte(requiredString)) {
			t.Error(config.Url + ": Incorrect response. Got:\n" + string(content))
		}
	}
	for _, excludedString := range config.ExcludedContent {
		if bytes.Contains(content, []byte(excludedString)) {
			t.Error(config.Url + ": Incorrect response. Got:\n" + string(content))
		}
	}
}
