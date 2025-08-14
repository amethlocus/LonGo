package handler

import (
	"LonGo/internal/views"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	fmt.Println("Login successful")
	token, err := CreateJWT("pablo")
	if err != nil {
		fmt.Printf("something went wrong")
	}

	c.HTML(http.StatusOK, "", views.Home(token))

}
