package model

import (
	"time"
)

// User model (Landlords & Tenants)
type User struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	Email          string    `gorm:"unique;not null" json:"email"`
	PhoneNumber    string    `gorm:"not null" json:"phone_number"`
	Password       string    `gorm:"not null" json:"password"`
	FirstName      string    `gorm:"not null" json:"first_name"`
	MiddleInitial  string    `gorm:"not null" json:"middle_initial"`
	LastName       string    `gorm:"not null" json:"last_name"`
	Age            int       `gorm:"not null" json:"age"`
	Address        string    `gorm:"not null" json:"address"`
	ValidID        string    `gorm:"not null" json:"valid_id"`
	BusinessName   string    `json:"business_name"`
	BusinessPermit string    `json:"business_permit"`
	AccountStatus  string    `gorm:"not null;default:'Pending'"` // verified / unverified
	UserType       string    `gorm:"not null" json:"user_type"`  // "Landlord" or "Tenant"
	CreatedAt      time.Time `json:"created_at"`
}

// Apartment model
type Apartment struct {
	ID            uint    `gorm:"primaryKey"`
	UserID        uint    `gorm:"not null"`
	PropertyName  string  `gorm:"not null"`
	Address       string  `gorm:"not null"`
	PropertyType  string  `gorm:"not null"` // "Bed Space" or "Apartment"
	RentPrice     float64 `gorm:"not null"`
	LocationLink  string  `gorm:"not null"`
	ContactNumber string  `gorm:"not null"`
	Email         string  `gorm:"not null"`
	Facebook      string
	Status        string `gorm:"not null;default:'Pending'"` // "Pending", "Approved", "Rejected", "Open", "Full"
	CreatedAt     time.Time
}

// Apartment images
type ApartmentImage struct {
	ID          uint   `gorm:"primaryKey"`
	ApartmentID uint   `gorm:"not null"`
	ImageURL    string `gorm:"not null"`
}

// Room model
type Room struct {
	ID          uint   `gorm:"primaryKey"`
	ApartmentID uint   `gorm:"not null"`
	Status      string `gorm:"not null;default:'Vacant'"` // "Vacant" or "Occupied"
}

// Booking model
type Booking struct {
	ID        uint `gorm:"primaryKey"`
	TenantID  uint `gorm:"not null"`
	RoomID    uint `gorm:"not null"`
	StartDate time.Time
	EndDate   time.Time
}

// Amenity model
type Amenity struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"not null;unique"`
}

type ApartmentAmenity struct {
	ID          uint `gorm:"primaryKey"`
	ApartmentID uint `gorm:"not null"`
	AmenityID   uint `gorm:"not null"`
}

// House Rule model
type HouseRule struct {
	ID   uint   `gorm:"primaryKey"`
	Rule string `gorm:"not null;unique"`
}

type ApartmentHouseRule struct {
	ID          uint `gorm:"primaryKey"`
	ApartmentID uint `gorm:"not null"`
	HouseRuleID uint `gorm:"not null"`
}

// Admin model
type Admin struct {
	ID       uint   `gorm:"primaryKey"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}
