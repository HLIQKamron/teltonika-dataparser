package main

import (
	"fmt"
	"net/http"

	"github.com/Projects/teltonika-dataparser/functions"
	"github.com/Projects/teltonika-dataparser/models"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "kamron"
	password = "kamron"
	dbname   = "teltonika_test"
)

func main() {
	r := gin.Default()

	// defer db.Close()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello World from container"})
	})
	r.GET("/longlat/:id", func(c *gin.Context) {
		id := c.Param("id")
		fmt.Println("Id :", id)

		data, err := functions.GetDeviceInfoByImei(id)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status":  "error",
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{"lon": data.Lon, "lat": data.Lat})

	})
	r.POST("/longlat", func(c *gin.Context) {

		var req models.Device
		err := c.ShouldBindJSON(&req)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status":  "error",
				"message": err.Error(),
			})
			return
		}
		err = functions.Update(req)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status":  "error",
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "success"})
	})

	// insert
	// hardcoded

	r.Run("0.0.0.0:2222")

}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
