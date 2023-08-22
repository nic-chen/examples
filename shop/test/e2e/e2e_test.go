package e2e

import (
	"testing"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var (
	GatewayHost = "http://127.0.0.1:9080"
)

// TestShop is the entry point for the E2E test.
func TestShop(t *testing.T) {
	res := WaitForHttpService("http://127.0.0.1:18080", 100)
	if !res {
		t.Fatal("failed to wait for etcd to start")
	}
	res = WaitForHttpService("http://127.0.0.1:2379", 100)
	if !res {
		t.Fatal("failed to wait for etcd to start")
	}
	res = WaitForHttpService(GatewayHost, 100)
	if !res {
		t.Fatal("failed to wait for Gateway to start")
	}

	gomega.RegisterFailHandler(ginkgo.Fail)

	ginkgo.RunSpecs(t, "Shop Test Suites")
}

var failed = false
var _ = ginkgo.AfterEach(func() {
	if ginkgo.CurrentSpecReport().Failed() {
		failed = true
	}
	DumpErrorLogs()
})

var _ = ginkgo.AfterSuite(func() {
	if !failed {
		RunCommand("cd .. && docker-compose -f ./docker/docker-compose.yaml down")
	}
})
