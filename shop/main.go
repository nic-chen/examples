package main

import (
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Product struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var products = map[int]Product{
	1: {
		ID:   1,
		Name: "Apple",
	},
	2: {
		ID:   2,
		Name: "Pear",
	},
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/products", func(c *gin.Context) {
		version := os.Getenv("VERSION")

		c.JSON(200, gin.H{
			"code":    0,
			"data":    products,
			"version": version,
		})
	})

	r.GET("/products/:id", func(c *gin.Context) {
		productID := c.Param("id")
		if productID == "" {
			c.JSON(400, gin.H{
				"message": "id is required",
			})
			return
		}

		id, err := strconv.Atoi(productID)
		if err != nil {
			c.JSON(400, gin.H{
				"message": "invalid productId",
			})
			return
		}

		version := os.Getenv("VERSION")

		c.JSON(200, gin.H{
			"code":    0,
			"data":    products[id],
			"version": version,
		})
	})

	return r
}

func main() {
	r := setupRouter()
	err := r.Run()
	if err != nil {
		panic(err)
	}
}
