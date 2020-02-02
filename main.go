package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/goyod/labs/fizzbuzz"
)

func main() {
	r := gin.Default()
	r.GET("/fizzbuzz/:number", fizzbuzzHandler)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func fizzbuzzHandler(c *gin.Context) {
	n, err := strconv.Atoi(c.Param("number"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.String(http.StatusOK, fizzbuzz.FizzBuzz(n))
}
