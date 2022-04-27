package model

import "time"

type TransactionStatus int

const (
	Paid TransactionStatus = iota
	Unpaid
	InProgress
	ApartmentCostError
)

type CostType int

const (
	Repair CostType = iota
)

// ApartmentCost defines a struct to store apartment cost (repairs etc)
type ApartmentCost struct {
	Id        int       `json:"apartmentCostId"`
	Price     int       `json:"price"`
	Date      time.Time `json:"date"`
	Expiry    time.Time `json:"expiry"`
	Status    CostType  `json:"status"`
	Apartment Apartment `json:"apartment"`
}

// Apartment defines a struct to store information about apartment
type Apartment struct {
	Id        int     `json:"apartmentId"`
	City      string  `json:"city"`
	Address   string  `json:"address"`
	Area      float32 `json:"area"`
	Tenant    *Tenant `json:"tenant"`
	RentPrice float32 `json:"rentPrice"`
}

// Tenant defines a struct to store information about tenant
type Tenant struct {
	Id        int    `json:"tenantId"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Login     string `json:"login"`
	Password  string `json:"password"`
}

// Transaction defines a struct to store tenant-apartment transaction history
type Transaction struct {
	Id        int               `json:"transactionId"`
	Price     int               `json:"price"`
	Tenant    Tenant            `json:"tenant"`
	Apartment Apartment         `json:"apartment"`
	Date      time.Time         `json:"transactionDate"`
	Status    TransactionStatus `json:"transactionStatus"`
}

// ApartmentRentalHistory defines a struct to store apartment rental history
type ApartmentRentalHistory struct {
	Id        int       `json:"apartmentHistoryId"`
	Apartment Apartment `json:"apartmentId"`
	Tenant    Tenant    `json:"tenant"`
	RentBegin time.Time `json:"rentBegin"`
	RentEnd   time.Time `json:"rentEnd"`
}
