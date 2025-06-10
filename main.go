package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func jsondeserialize() (response Response) {
	jsonData := `{
		"took": 18,
		"timed_out": false,
		"_shards": {
			"total": 1,
			"successful": 1,
			"skipped": 0,
			"failed": 0
		},
		"hits": {
			"total": {
				"value": 100,
				"relation": "eq"
			},
			"max_score": 1,
			"hits": [
				{
					"_index": "products-13",
					"_id": "guyONJcBjv4AvNg2XIoG",
					"_score": 1,
					"_source": {
						"webid": 3,
						"title": "Enormous Steel Bench",
						"description": "Built for efficiency, this product balances aesthetics and practical performance seamlessly.",
						"combinedsearchtext": "Enormous Steel Bench Built for efficiency, this product balances aesthetics and practical performance seamlessly. Teal Books",
						"price": 5,
						"categoryName": "Books",
						"stockLevel": 40,
						"color": "Teal",
						"categoryId": 6,
						"string_facet": [
							{
								"facet_name": "Color",
								"facet_value": "Teal"
							},
							{
								"facet_name": "Category",
								"facet_value": "Books"
							}
						]
					}
				}
			]
		},
		"aggregations": {
			"facets": {
				"doc_count": 200,
				"names": {
					"doc_count_error_upper_bound": 0,
					"sum_other_doc_count": 0,
					"buckets": [
						{
							"key": "Category",
							"doc_count": 100,
							"values": {
								"doc_count_error_upper_bound": 0,
								"sum_other_doc_count": 0,
								"buckets": [
									{
										"key": "Bla",
										"doc_count": 23
									}
								]
							}
						}
					]
				}
			}
		}
	}`

	err := json.Unmarshal([]byte(jsonData), &response)
	if err != nil {
		log.Fatalf("Error unmarshaling JSON: %v", err)
	}
	return response

}

func getSinglePlayer(c *gin.Context) {
	resp := jsondeserialize()
	time.Sleep(10 * time.Millisecond) // Simulate a delay - anropa en databas
	playerID := c.Param("id")
	resp.Hits.Hits[0].ID = playerID
	c.JSON(200, resp.Hits.Hits[0])
}

func getAllPlayers(c *gin.Context) {
	resp := jsondeserialize()
	time.Sleep(10 * time.Millisecond) //  Simulate a delay - anropa en databas
	fibonacciIterative(50)
	c.JSON(200, resp)
}

func fibonacciIterative(n int) int {
	if n <= 1 {
		return n
	}
	var n2, n1 = 0, 1
	for i := 2; i <= n; i++ {
		n2, n1 = n1, n1+n2
	}
	return n1
}

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	router.GET("/players/:id", getSinglePlayer)
	router.GET("/players", getAllPlayers)

	router.Run(":8080")

}
