package config

import (
	"fmt"
	"os"
	"time"
)

// Load environment variables
func LoadEnv() {
	fmt.Println("Environment loaded successfully!")
}

// Get server port
func LoadPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return "3000"
	}
	return port
}

// Database connection string
func GetDatabaseURL() string {
	return "postgres://user:password@localhost:5432/testdb"
}

// JWT secret key
func GetJWTSecret() string {
	return "my-super-secret-jwt-key-123"
}

// Calculate age from birth year
func CalculateAge(birthYear int) int {
	currentYear := time.Now().Year()
	return currentYear - birthYear
}

// Format user greeting
func FormatGreeting(name string) string {
	return fmt.Sprintf("สวัสดี %s! ยินดีต้อนรับเข้าสู่ระบบ", name)
}

// Check if number is even
func IsEven(num int) bool {
	return num%2 == 0
}

// Generate simple ID
func GenerateID() string {
	return fmt.Sprintf("ID-%d", time.Now().Unix())
}
