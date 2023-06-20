package e2e

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"time"

	"github.com/onsi/ginkgo/v2"
)

func WaitForHttpService(url string, tries int) bool {
	var (
		i int
		s int

		client = http.Client{
			Timeout: 500 * time.Millisecond,
		}
	)

	for {
		i++
		res, err := client.Get(url)
		if err == nil && res.StatusCode != 0 {
			s++
			time.Sleep(500 * time.Millisecond)
			// considered success when consecutive success twice
			if s >= 2 {
				fmt.Println("Service is ready, status:", res.StatusCode)
				return true
			}
		} else {
			fmt.Print(".")
			s = 0
		}

		if i > tries {
			fmt.Print("Service is not ready", err)
			return false
		}

		time.Sleep(500 * time.Millisecond)
	}
}

// RunCommand runs a command in a shell and returns the output
func RunCommand(command string) (string, error) {
	cmd := exec.Command("/bin/sh", "-c", command)
	res, err := cmd.CombinedOutput()
	return string(res), err
}

func HttpGet(url string, headers map[string]string) ([]byte, int, error) {
	return httpRequest(http.MethodGet, url, headers, "")
}

func HttpPut(url string, headers map[string]string, reqBody string) ([]byte, int, error) {
	return httpRequest(http.MethodPut, url, headers, reqBody)
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

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, err
	}

	return body, resp.StatusCode, nil
}

// DumpErrorLogs dump the error logs of the test.
func DumpErrorLogs() {
	if ginkgo.CurrentSpecReport().Failed() {
		ginkgo.GinkgoWriter.Println("--------------  dump logs  --------------")
		res, err := RunCommand("docker logs apisix -n 100")
		ginkgo.GinkgoWriter.Println("server logs:", res, " cmd err:", err)
	}
}
