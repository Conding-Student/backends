package controller

import (
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"

	"intern_template_v1/middleware"
	"intern_template_v1/model"
	"intern_template_v1/model/response" // Replace with actual module path
)

// LoginRequest struct allows login with either Email or Phone Number
type LoginRequest struct {
	Identifier string `json:"identifier"` // Can be Email or Phone Number
	Password   string `json:"password"`
}

// LoginUser authenticates a user and returns a JWT token
func LoginUser(c *fiber.Ctx) error {
	// Parse and validate request body
	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		log.Println("[ERROR] Failed to parse request body:", err)
		return c.Status(fiber.StatusBadRequest).JSON(response.ResponseModel{
			RetCode: "400",
			Message: "Invalid request format",
			Data:    nil,
		})
	}

	// Trim inputs to avoid accidental spaces
	req.Identifier = strings.TrimSpace(req.Identifier)
	req.Password = strings.TrimSpace(req.Password)

	// Validate input fields
	if req.Identifier == "" || req.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(response.ResponseModel{
			RetCode: "400",
			Message: "Email/Phone and password are required",
			Data:    nil,
		})
	}

	var user model.User

	// Query the database to find the user by email or phone number
	result := middleware.DBConn.Table("users").
		Where("email = ? OR phone_number = ?", req.Identifier, req.Identifier).
		First(&user)

	if result.Error != nil {
		log.Println("[ERROR] Database query error:", result.Error)
		return c.Status(fiber.StatusInternalServerError).JSON(response.ResponseModel{
			RetCode: "500",
			Message: "Server error. Please try again later.",
			Data:    nil,
		})
	}

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(response.ResponseModel{
			RetCode: "404",
			Message: "Account not found",
			Data:    nil,
		})
	}

	// Debugging: Print fetched user details (remove in production)
	log.Println("[DEBUG] User found:", user.Email, "| Middle Initial:", user.MiddleInitial)

	// Compare the provided password with the stored hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		log.Println("[WARNING] Incorrect password for user:", user.Email)
		return c.Status(fiber.StatusUnauthorized).JSON(response.ResponseModel{
			RetCode: "401",
			Message: "Incorrect password",
			Data:    nil,
		})
	}

	// ✅ Generate JWT token
	token, err := middleware.GenerateJWT(user.ID, user.Email, user.UserType)
	if err != nil {
		log.Println("[ERROR] Failed to generate JWT token:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(response.ResponseModel{
			RetCode: "500",
			Message: "Could not generate authentication token",
			Data:    nil,
		})
	}

	// ✅ Successful login response with JWT token
	return c.Status(fiber.StatusOK).JSON(response.ResponseModel{
		RetCode: "200",
		Message: "Login successful",
		Data: fiber.Map{
			"id":             user.ID,
			"first_name":     user.FirstName,
			"middle_initial": user.MiddleInitial, // ✅ Middle Initial now included
			"last_name":      user.LastName,
			"email":          user.Email,
			"phone_number":   user.PhoneNumber,
			"address":        user.Address,
			"user_type":      user.UserType,
			"token":          token, // ✅ JWT token for authentication
		},
	})
}
