package dataSignature

import (
	"time"
)

type SigninData struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// POST

type PostNormalAccount struct {
	Name        string    `json:"name" binding:"required"`
	LastName    string    `json:"lastName" binding:"required"`
	UserName    string    `json:"username" binding:"required"`
	PhoneNumber string    `json:"phoneNumber" binding:"required"`
	Password    string    `json:"password" binding:"required"`
	Email       string    `json:"email" binding:"required"`
	Gender      string    `json:"gender" binding:"required"`
	JoinDate    time.Time `json:"JoinDate" binding:"required"`
}

type PostVetAccount struct {
	PostNormalAccount
	ExpertiseLevel        string    `json:"expertiseLevel" binding:"required"`
	ExpertiseArea         string    `json:"expertiseArea" binding:"required"`
	WorkingHoursStartDate string    `json:"workingHoursStartDate" binding:"required"`
	WorkingHoursEndDate   string    `json:"workingHoursEndDate" binding:"required"`
	WorkingHoursStartHour time.Time `json:"workingHoursStartHour" binding:"required"`
	WorkingHoursEndHour   time.Time `json:"workingHoursEndHour" binding:"required"`
}

type PostPetLoverAccount struct {
	PostNormalAccount
	FreeTimeStartDate string    `json:"freeTimeStartDate" binding:"required"`
	FreeTimeEndDate   string    `json:"freeTimeEndDate" binding:"required"`
	FreeTimeStartHour time.Time `json:"freeTimeStartHour" binding:"required"`
	FreeTimeEndHour   time.Time `json:"freeTimeEndHour" binding:"required"`
}

// GET

type GetNormalAccount struct {
	Id          uint64    `json:"id" binding:"required"`
	Name        string    `json:"name" binding:"required"`
	LastName    string    `json:"lastName" binding:"required"`
	UserName    string    `json:"username" binding:"required"`
	PhoneNumber string    `json:"phoneNumber" binding:"required"`
	Password    string    `json:"password" binding:"required"`
	Email       string    `json:"email" binding:"required"`
	Gender      string    `json:"gender" binding:"required"`
	JoinDate    time.Time `json:"JoinDate" binding:"required"`
}

type GetVetAccount struct {
	PostNormalAccount
	ExpertiseLevel        string    `json:"expertiseLevel" binding:"required"`
	ExpertiseArea         string    `json:"expertiseArea" binding:"required"`
	WorkingHoursStartDate string    `json:"workingHoursStartDate" binding:"required"`
	WorkingHoursEndDate   string    `json:"workingHoursEndDate" binding:"required"`
	WorkingHoursStartHour time.Time `json:"workingHoursStartHour" binding:"required"`
	WorkingHoursEndHour   time.Time `json:"workingHoursEndHour" binding:"required"`
}

type GetPetLoverAccount struct {
	PostNormalAccount
	FreeTimeStartDate string    `json:"freeTimeStartDate" binding:"required"`
	FreeTimeEndDate   string    `json:"freeTimeEndDate" binding:"required"`
	FreeTimeStartHour time.Time `json:"freeTimeStartHour" binding:"required"`
	FreeTimeEndHour   time.Time `json:"freeTimeEndHour" binding:"required"`
}

// UPDATE

type UpdateNormalAccount struct {
	Name        string `json:"name" binding:"required"`
	LastName    string `json:"lastName" binding:"required"`
	UserName    string `json:"username" binding:"required"`
	PhoneNumber string `json:"phoneNumber" binding:"required"`
	Password    string `json:"password" binding:"required"`
	Email       string `json:"email" binding:"required"`
	Gender      string `json:"gender" binding:"required"`
}

type UpdateVetAccount struct {
	PostNormalAccount
	ExpertiseLevel        string    `json:"expertiseLevel" binding:"required"`
	ExpertiseArea         string    `json:"expertiseArea" binding:"required"`
	WorkingHoursStartDate string    `json:"workingHoursStartDate" binding:"required"`
	WorkingHoursEndDate   string    `json:"workingHoursEndDate" binding:"required"`
	WorkingHoursStartHour time.Time `json:"workingHoursStartHour" binding:"required"`
	WorkingHoursEndHour   time.Time `json:"workingHoursEndHour" binding:"required"`
}

type UpdatePetLoverAccount struct {
	PostNormalAccount
	FreeTimeStartDate string    `json:"freeTimeStartDate" binding:"required"`
	FreeTimeEndDate   string    `json:"freeTimeEndDate" binding:"required"`
	FreeTimeStartHour time.Time `json:"freeTimeStartHour" binding:"required"`
	FreeTimeEndHour   time.Time `json:"freeTimeEndHour" binding:"required"`
}
