package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type Member struct {
	Email       string    `json:"email"`
	Password    string    `json:"password,omitempty"`
	Name        string    `json:"name"`
	DateOfBirth time.Time `json:"date_of_birth"`
	Gender      string    `json:"gender"`
	Address     string    `json:"address"`
	Subscribed  bool      `json:"subscribed"`
	Age         int       `json:"age,omitempty"`
}

var members = map[string]Member{}

func Register(c *gin.Context) {
	var newMember Member
	if err := c.ShouldBindJSON(&newMember); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := HashPassword(newMember.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	newMember.Password = hashedPassword

	members[newMember.Email] = newMember

	c.JSON(http.StatusOK, gin.H{"message": "Registration successful", "member": newMember})
}

func ViewProfile(c *gin.Context) {
	email := c.Query("email")
	member, exists := members[email]

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Member not found"})
		return
	}

	currentTime := time.Now()
	age := currentTime.Year() - member.DateOfBirth.Year()
	member.Age = age

	c.JSON(http.StatusOK, member)
}

func EditProfile(c *gin.Context) {
	email := c.Query("email")
	member, exists := members[email]

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Member not found"})
		return
	}

	var updatedData struct {
		DateOfBirth time.Time `json:"date_of_birth"`
		Gender      string    `json:"gender"`
		Address     string    `json:"address"`
		Subscribed  bool      `json:"subscribed"`
	}
	if err := c.ShouldBindJSON(&updatedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	member.DateOfBirth = updatedData.DateOfBirth
	member.Gender = updatedData.Gender
	member.Address = updatedData.Address
	member.Subscribed = updatedData.Subscribed

	members[email] = member

	c.JSON(http.StatusOK, gin.H{"message": "Profile updated", "member": member})
}

func DeleteProfile(c *gin.Context) {
	email := c.Query("email")
	if _, exists := members[email]; !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Member not found"})
		return
	}

	delete(members, email)

	c.JSON(http.StatusOK, gin.H{"message": "Member deleted successfully"})
}

func ChangePassword(c *gin.Context) {
	var changePasswordData struct {
		Email           string `json:"email"`
		CurrentPassword string `json:"current_password"`
		NewPassword     string `json:"new_password"`
		ConfirmPassword string `json:"confirm_password"`
	}

	if err := c.ShouldBindJSON(&changePasswordData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	member, exists := members[changePasswordData.Email]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Member not found"})
		return
	}

	if !CheckPasswordHash(member.Password, changePasswordData.CurrentPassword) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Current password is incorrect"})
		return
	}

	if changePasswordData.NewPassword != changePasswordData.ConfirmPassword {
		c.JSON(http.StatusBadRequest, gin.H{"error": "New password and confirm password do not match"})
		return
	}

	hashedNewPassword, err := HashPassword(changePasswordData.NewPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash new password"})
		return
	}

	member.Password = hashedNewPassword
	members[changePasswordData.Email] = member

	c.JSON(http.StatusOK, gin.H{"message": "Password changed successfully"})
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
