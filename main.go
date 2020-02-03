package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/goyod/labs/fizzbuzz"
)

func main() {
	r := gin.Default()
	r.GET("/fizzbuzzr", JWTWrapper(fizzbuzzRandomHandler))
	r.GET("/fizzbuzz/:number", fizzbuzzHandler)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func tokenHandler(c *gin.Context) {
	mySigningKey := []byte("AllYourBase")

	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Minute * 5).Unix(),
		Issuer:    "pallat",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.String(http.StatusOK, ss)
}

func JWTWrapper(next gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")[7:]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return []byte("AllYourBase"), nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid {
			_ = claims
			c.JSON(http.StatusUnauthorized, map[string]string{
				"message": err.Error(),
			})
			return
		}

		next(c)
	}
}

func fizzbuzzHandler(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")[7:]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte("AllYourBase"), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid {
		c.String(http.StatusUnauthorized, err.Error())
		_ = claims
		return
	}

	n, err := strconv.Atoi(c.Param("number"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.String(http.StatusOK, fizzbuzz.New(n).String())
}

func fizzbuzzRandomHandler(c *gin.Context) {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	n := r.Intn(100)

	c.JSON(http.StatusOK, FizzbuzzResponse{
		Number:  n,
		Message: fizzbuzz.New(n).String(),
	})
}

type FizzbuzzResponse struct {
	Number  int    `json:"number"`
	Message string `json:"message"`
}
