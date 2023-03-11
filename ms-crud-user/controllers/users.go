package controllers

import (
	"github.com/gin-gonic/gin"
	"ms-crud-user/database"
	"ms-crud-user/entities"
	"net/http"
)

type CreateUserInput struct {
	Name     string `json:"name" binding:"required"`
	LastName string `json:"lastName" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
}

type UpdateUserInput struct {
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Phone    string `json:"phone"`
}

// FindUsers GET /Users
// Find all Users
// @Summary Get all users
// @Description get all users
// @Tags Users
// @Success 200 {array} entities.User
// @Failure 404 {object} object
// @Router / [get]
func FindUsers(c *gin.Context) {
	var Users []entities.User
	database.Instance.Find(&Users)

	c.JSON(http.StatusOK, gin.H{"data": Users})
}

// FindUser GET /Users/:id
// Find a User
func FindUser(c *gin.Context) {
	// Get model if exist
	var User entities.User
	if err := database.Instance.Where("id = ?", c.Param("id")).First(&User).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": User})
}

type User struct {
	ID int
}

// CreateUser POST /Users
// Create new User
func CreateUser(c *gin.Context) {
	// Validate input
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var currentUser = User{ID: 100}

	// Create will set product's `CreatedBy`, `UpdatedBy` to `currentUser`'s primary key if `audited:current_user` is a valid model
	//db.Set("audited:current_user", currentUser).Create(&product)

	// Create User
	User := entities.User{Name: input.Name, LastName: input.LastName, Phone: input.Phone}
	database.Instance.Set("audited:current_user", currentUser).Create(&User)

	c.JSON(http.StatusCreated, gin.H{"data": User})
}

// UpdateUser PATCH /Users/:id
// Update a User
func UpdateUser(c *gin.Context) {
	// Get model if exist
	var User entities.User

	if err := database.Instance.Where("id = ?", c.Param("id")).First(&User).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.Instance.Model(&User).Set("audited:current_user", User).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": User})
}

// DeleteUser DELETE /Users/:id
// Delete a User
func DeleteUser(c *gin.Context) {
	// Get model if exist
	var User entities.User
	if err := database.Instance.Where("id = ?", c.Param("id")).First(&User).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// set is delete true
	var input = User
	input.IsDelete = true

	database.Instance.Model(&User).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": User})
}
