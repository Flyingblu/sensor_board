package main

import (
	"github.com/flyingblu/sensor_board/influxclient"
	"github.com/gin-gonic/gin"
	"strconv"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)
	server := gin.Default()
	ifClient := influxclient.InfluxClient{InfluxUrl: influxclient.HostAddr, Token: influxclient.ClientToken, Database: influxclient.Database}

	server.POST("/write_data", func(c *gin.Context) {
		value, err := strconv.ParseFloat(c.PostForm("value"), 64)

		if err != nil {
			c.JSON(200, gin.H{
				"status": "failed",
				"error":  err.Error(),
			})
			return
		}
		err2 := ifClient.ReceiveData(c.PostForm("token"), c.PostForm("device"), c.PostForm("measurement"), value)

		if err2 != nil {
			c.JSON(200, gin.H{
				"status": "failed",
				"error":  err2.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"status": "success",
			"error":  "nil",
		})
	})
	server.Run()
}
