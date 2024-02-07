package main

import (
	"database/sql"
	"fmt"
	"net/http"

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

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	defer db.Close()
	r.GET("/longlat", func(c *gin.Context) {
		data, err := db.Query("select * from longlat")
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		defer data.Close()
		for data.Next() {
			var lon, lat string
			err = data.Scan(&lon, &lat)
			if err != nil {
				c.String(http.StatusInternalServerError, err.Error())
				return
			}
			c.JSON(http.StatusOK, gin.H{"lon": lon, "lat": lat})
		}
	})
	r.POST("/longlat", func(c *gin.Context) {

		lat := c.Query("lat")
		lon := c.Query("lon")

		// updateStmt := fmt.Sprintf(`insert into "longlat"("longitude", "latitude") values('%s', '%s')`, lon, lat)
		updateStmt := fmt.Sprintf(`update "longlat" set "longitude" = '%s', "latitude" = '%s'`, lon, lat)
		_, err := db.Exec(updateStmt)
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

	r.Run("0.0.0.0:4444")

}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
