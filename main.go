package main

import (
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// =====================
// Request Struct
// =====================

type RegisterRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// =====================
// Regex Validation
// =====================

var (
	emailRegex    = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	letterRegex = regexp.MustCompile(`[A-Za-z]`)
	digitRegex  = regexp.MustCompile(`[0-9]`)
	fullRegex   = regexp.MustCompile(`^[A-Za-z0-9]{6,}$`)
)

func isValidPassword(pw string) bool {
	return fullRegex.MatchString(pw) &&
		letterRegex.MatchString(pw) &&
		digitRegex.MatchString(pw)
}

// =====================
// Main
// =====================

func main() {
	r := gin.Default()

	r.POST("/register", registerHandler)
	r.POST("/login", loginHandler)

	r.Run(":8080")
}

// =====================
// Handlers
// =====================

func registerHandler(c *gin.Context) {
	var req RegisterRequest

	// ShouldBindWith JSON
	if err := c.ShouldBindWith(&req, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	// Regex validation
	if !emailRegex.MatchString(req.Email) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "email format tidak valid",
		})
		return
	}

	if !isValidPassword(req.Password) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "password minimal 6 karakter dan harus mengandung angka",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "register berhasil",
		"data": gin.H{
			"email": req.Email,
			"name":  req.Name,
		},
	})
}

func loginHandler(c *gin.Context) {
	var req LoginRequest

	// ShouldBindWith JSON
	if err := c.ShouldBindWith(&req, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	// Regex validation
	if !emailRegex.MatchString(req.Email) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "email format tidak valid",
		})
		return
	}

	if req.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "password wajib diisi",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "login berhasil",
	})
}
