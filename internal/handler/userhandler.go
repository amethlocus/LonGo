package handler

import (
	"LonGo/internal/database"
	"LonGo/internal/models"
	"LonGo/internal/views"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *gin.Context) {
	var input models.User

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error encriptando la contraseña"})
		return
	}
	input.Password = string(hashedPassword)

	if err := database.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	input.Password = "" // no enviar la contraseña en la respuesta
	c.JSON(http.StatusCreated, gin.H{"message": "Usuario creado", "user": input})
}

func CreateUserForm(c *gin.Context) {
	c.HTML(http.StatusOK, "", views.UserForm())
}
