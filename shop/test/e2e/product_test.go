package e2e

import (
	"fmt"
	"net/http"
	"time"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe("Product Test", ginkgo.Ordered, func() {
	ginkgo.Context("version v1", func() {
		ginkgo.It("create route with service shop v1", func() {
			reqBody := `{
				"uri": "/products/*",
				"upstream": {
					"type":"roundrobin",
					"service_name": "shopv1",
					"discovery_type": "nacos"
				}
			}`
			respBody, status, err := HttpPut(
				fmt.Sprintf("%s/apisix/admin/routes/1", GatewayHost),
				map[string]string{
					"Content-Type": "application/json",
					"X-API-KEY":    "edd1c9f034335f136f87ad84b625c8f1",
				},
				reqBody,
			)
			gomega.Expect(err).ShouldNot(gomega.HaveOccurred(), "create route")
			gomega.Expect(string(respBody)).Should(gomega.ContainSubstring(`"service_name":"shopv1"`))
			gomega.Expect(status).Should(gomega.Or(gomega.Equal(http.StatusCreated), gomega.Equal(http.StatusOK)))
		})

		ginkgo.It("hit the route", func() {
			// sleep 1 second to wait for configuration to take effect
			time.Sleep(1 * time.Second)

			body, status, err := HttpGet(
				GatewayHost+"/products/1",
				map[string]string{
					"Content-Type": "application/json",
				},
			)
			gomega.Expect(err).ShouldNot(gomega.HaveOccurred(), "get product")
			gomega.Expect(status).Should(gomega.Equal(http.StatusOK))
			gomega.Expect(string(body)).Should(gomega.ContainSubstring(`"id":1`))
			gomega.Expect(string(body)).Should(gomega.ContainSubstring(`"version":"v1"`))
		})
	})

	ginkgo.Context("version v2", func() {
		ginkgo.It("create route with service shop v2", func() {
			reqBody := `{
				"uri": "/products",
				"upstream": {
					"type":"roundrobin",
					"service_name": "shopv2",
					"discovery_type": "nacos"
				}
			}`
			respBody, status, err := HttpPut(
				fmt.Sprintf("%s/apisix/admin/routes/2", GatewayHost),
				map[string]string{
					"Content-Type": "application/json",
					"X-API-KEY":    "edd1c9f034335f136f87ad84b625c8f1",
				},
				reqBody,
			)
			gomega.Expect(err).ShouldNot(gomega.HaveOccurred(), "create route")
			gomega.Expect(string(respBody)).Should(gomega.ContainSubstring(`"service_name":"shopv2"`))
			gomega.Expect(status).Should(gomega.Or(gomega.Equal(http.StatusCreated), gomega.Equal(http.StatusOK)))
		})

		ginkgo.It("hit the route", func() {
			// sleep 1 second to wait for configuration to take effect
			time.Sleep(1 * time.Second)

			body, status, err := HttpGet(
				GatewayHost+"/products",
				map[string]string{
					"Content-Type": "application/json",
				},
			)
			gomega.Expect(err).ShouldNot(gomega.HaveOccurred(), "get products")
			gomega.Expect(status).Should(gomega.Equal(http.StatusOK))
			gomega.Expect(string(body)).Should(gomega.ContainSubstring(`"id":1`))
			gomega.Expect(string(body)).Should(gomega.ContainSubstring(`"id":2`))
			gomega.Expect(string(body)).Should(gomega.ContainSubstring(`"version":"v2"`))
		})
	})
})
