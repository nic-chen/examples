package e2e

import (
	"net/http"
	"time"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe("Product Test", ginkgo.Ordered, func() {
	ginkgo.Context("version v1", func() {
		ginkgo.It("Run container with env v1", func() {
			RunContainer("v1")
			time.Sleep(1 * time.Second)
		})
		ginkgo.It("API test", func() {
			body, status, err := HttpGet(
				ServerHost+"/products/1",
				map[string]string{
					"Content-Type": "application/json",
				},
			)
			gomega.Expect(err).ShouldNot(gomega.HaveOccurred(), "get product")
			gomega.Expect(status).Should(gomega.Equal(http.StatusOK))
			gomega.Expect(string(body)).Should(gomega.ContainSubstring(`"id":1`))
			gomega.Expect(string(body)).Should(gomega.ContainSubstring(`"version":"v1"`))
		})
		ginkgo.It("Stop container", func() {
			StopContainer()
		})
	})

	ginkgo.Context("version v2", func() {
		ginkgo.It("Run container with env v2", func() {
			RunContainer("v2")
			time.Sleep(1 * time.Second)
		})
		ginkgo.It("API test", func() {
			body, status, err := HttpGet(
				ServerHost+"/products/2",
				map[string]string{
					"Content-Type": "application/json",
				},
			)
			gomega.Expect(err).ShouldNot(gomega.HaveOccurred(), "get product")
			gomega.Expect(status).Should(gomega.Equal(http.StatusOK))
			gomega.Expect(string(body)).Should(gomega.ContainSubstring(`"id":2`))
			gomega.Expect(string(body)).Should(gomega.ContainSubstring(`"version":"v2"`))
		})
		ginkgo.It("Stop container", func() {
			StopContainer()
		})
	})
})
