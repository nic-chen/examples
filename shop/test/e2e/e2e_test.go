package e2e

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"os/exec"
	"testing"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var (
	ServerHost = "http://127.0.0.1:8080"
)

// TestShop is the entry point for the E2E test.
func TestShop(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)

	ginkgo.RunSpecs(t, "Shop Test Suites")
}

// RunCommand runs a command in a shell and returns the output
func RunCommand(command string) (string, error) {
	cmd := exec.Command("/bin/sh", "-c", command)
	res, err := cmd.CombinedOutput()
	return string(res), err
}

func RunContainer(version string) {
	RunCommand("docker run --name=shop-container -d -p 8080:8080 -e VERSION=" + version + " shop:test")
}

func StopContainer() {
	RunCommand("docker stop shop-container && docker rm shop-container")
}

func HttpGet(url string, headers map[string]string) ([]byte, int, error) {
	return httpRequest(http.MethodGet, url, headers, "")
}

func httpRequest(method, url string, headers map[string]string, reqBody string) ([]byte, int, error) {
	var requestBody = new(bytes.Buffer)
	if reqBody != "" {
		requestBody = bytes.NewBuffer([]byte(reqBody))
	}
	req, err := http.NewRequest(method, url, requestBody)

	req.Close = true
	if err != nil {
		return nil, 0, err
	}

	// set header
	for key, val := range headers {
		if key == "Host" {
			req.Host = val
		}
		req.Header.Add(key, val)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, err
	}

	return body, resp.StatusCode, nil
}
