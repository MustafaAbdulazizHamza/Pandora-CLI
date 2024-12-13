package clientLogic

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func sendHTTPRequest(method, url, username, password string, body string) (response *http.Response, err error) {
	req, err := http.NewRequest(method, url, bytes.NewBufferString(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("username", username)
	req.Header.Set("password", password)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true, // This skips certificate verification
		},
	}

	client := &http.Client{
		Timeout:   time.Second * 10,
		Transport: tr,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil

}

func toJson(i interface{}) ([]byte, error) {
	j, err := json.Marshal(i)
	return j, err
}
func ParseResponse(inp *http.Response) (response Response, err error) {
	defer inp.Body.Close()
	body, err := ioutil.ReadAll(inp.Body)
	if err != nil {
		return response, err
	}
	if string(body) == "404 page not found" {
		body, err = toJson(Response{
			Status: fmt.Sprint(http.StatusNotFound),
			Text:   "Page not found.",
		})
	}
	err = json.Unmarshal(body, &response)

	return response, err
}

func PrintOutResponse(response Response) {
	// Define colors for different statuses
	//green := color.New(color.BgGreen).Add(color.FgBlack).Add(color.Bold)   // Green background
	red := color.New(color.BgRed).Add(color.FgBlack).Add(color.Bold)       // Red background
	yellow := color.New(color.BgYellow).Add(color.FgBlack).Add(color.Bold) // Yellow background
	white := color.New(color.FgWhite)                                      // White for the message text

	// Extract the numeric status code from the status string
	statusCode := strings.Split(response.Status, " ")[0]

	// Handle status and printing
	switch {
	case statusCode == "200":
		fmt.Println(response.Text)
		// Green square for success with white message
		//green.Printf("[\u2588%s] ", statusCode) // Square with status code in the square
		//white.Print(response.Text + "\n")
	case strings.HasPrefix(statusCode, "4"):
		// Red square for client errors (4xx) with white message
		red.Printf("[\u2588%s] ", statusCode)
		white.Print(response.Text + "\n")
	case strings.HasPrefix(statusCode, "5"):
		// Yellow square for server errors (5xx) with white message
		yellow.Printf("[\u2588%s] ", statusCode)
		white.Print(response.Text + "\n")
	default:
		// Default handling for other status codes with white message
		fmt.Printf("[ ] %s\n", response.Text)
	}
}
