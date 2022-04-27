package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"log"
	"polsl/tab/estate-agency/internal/model"
	"time"
)

type Database struct {
	user       string
	password   string
	dbName     string
	connection *sql.DB
	log        *logrus.Logger
}

func NewDatabaseConnection(user, password, dbName string, log *logrus.Logger) Database {
	return Database{
		user:     user,
		password: password,
		dbName:   dbName,
		log:      log,
	}
}

func (d *Database) Connect() error {
	dataSourceName := d.user + ":" + d.password + "@/" + d.dbName
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return err
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	d.connection = db
	return nil
}

// ——————————————————————————————————————————————————————————————————————————————————————

// Add
func (d Database) AddApartment(a model.Apartment) *model.Apartment {
	stmt, err := d.connection.Prepare(`INSERT INTO apartments (city, address, area, tenantId, rentPrice)
 									VALUES (?, ?, ?, ?, ?)`)
	if err != nil {
		log.Fatal(err)
		// return false
	}
	defer stmt.Close()

	tId, err := d.GetTenant(a.Tenant.Id)
	if err != nil {
		log.Fatal(err)
		// return false
	}

	if _, err := stmt.Exec(a.City, a.Address, a.Area, tId, a.RentPrice); err != nil {
		log.Fatal(err)
		// return false
	}
	return &a
}

func (d Database) AddTenant(t model.Tenant) *model.Tenant {
	stmt, err := d.connection.Prepare(`INSERT INTO tenants (firstName, lastName, userName, password)
 									VALUES (?, ?, ?, ?)`)
	if err != nil {
		log.Fatal(err)
		// return false
	}
	defer stmt.Close()

	if _, err := stmt.Exec(t.FirstName, t.LastName, t.Login, t.Password); err != nil {
		log.Fatal(err)
		// return false
	}
	return &t
}

func (d Database) AddPayment(c model.ApartmentCost) *model.ApartmentCost {
	stmt, err := d.connection.Prepare(`INSERT INTO payments (price, date, expiry, status, apartmentId)
 									VALUES (?, ?, ?, ?, ?)`)
	if err != nil {
		log.Fatal(err)
		// return false
	}
	defer stmt.Close()

	aId, err := d.GetApartment(c.Apartment.Id)
	if err != nil {
		log.Fatal(err)
		// return false
	}

	if _, err := stmt.Exec(c.Price, c.Date, c.Expiry, c.Status, aId); err != nil {
		log.Fatal(err)
		// return false
	}
	return &c
}

func (d Database) AddTransaction(t model.Transaction) *model.Transaction {
	stmt, err := d.connection.Prepare(`INSERT INTO transactions (price, tenantId, apartmentId, date, status)
 									VALUES (?, ?, ?, ?, ?)`)
	if err != nil {
		log.Fatal(err)
		// return false
	}
	defer stmt.Close()

	tId, err := d.GetTenant(t.Tenant.Id)
	if err != nil {
		log.Fatal(err)
		// return false
	}

	aId, err := d.GetApartment(t.Apartment.Id)
	if err != nil {
		log.Fatal(err)
		// return false
	}

	if _, err := stmt.Exec(t.Price, tId, aId, t.Date, t.Status); err != nil {
		log.Fatal(err)
		// return false
	}
	return &t
}

func (d Database) AddHistory(h model.ApartmentRentalHistory) *model.ApartmentRentalHistory {
	stmt, err := d.connection.Prepare(`INSERT INTO histories (apartmentId, tenantId, rentBegin, rentEnd)
 									VALUES (?, ?, ?, ?)`)
	if err != nil {
		log.Fatal(err)
		// return false
	}
	defer stmt.Close()

	aId, err := d.GetApartment(h.Apartment.Id)
	if err != nil {
		log.Fatal(err)
		// return false
	}
	tId, err := d.GetTenant(h.Tenant.Id)
	if err != nil {
		log.Fatal(err)
		// return false
	}

	if _, err := stmt.Exec(aId, tId, h.RentBegin, h.RentEnd); err != nil {
		log.Fatal(err)
		// return false
	}
	return &h
}

// Get Collection
func (d Database) GetAllApartments() ([]model.Apartment, error) {
	stmt, err := d.connection.Query("SELECT * FROM apartments")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer stmt.Close()

	var as []model.Apartment
	var tId int

	for stmt.Next() {
		var a model.Apartment
		fmt.Println(stmt.Scan(&a.Id,
			&a.City,
			&a.Address,
			&a.Area,
			&tId,
			&a.RentPrice))
		a.Tenant, err = d.GetTenant(tId)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		as = append(as, a)
	}

	return as, nil
}

func (d Database) GetAllTenants() ([]model.Tenant, error) {
	stmt, err := d.connection.Query("SELECT * FROM tenants")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer stmt.Close()

	var ts []model.Tenant

	for stmt.Next() {
		var t model.Tenant
		fmt.Println(stmt.Scan(&t.Id,
			&t.FirstName,
			&t.LastName,
			&t.Login,
			&t.Password))
		ts = append(ts, t)
	}
	return ts, nil
}

func (d Database) GetAllPayments() ([]model.ApartmentCost, error) {
	stmt, err := d.connection.Query("SELECT * FROM payments")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer stmt.Close()

	var cs []model.ApartmentCost
	var aId int

	for stmt.Next() {
		var c model.ApartmentCost
		fmt.Println(stmt.Scan(&c.Id,
			&c.Price,
			&c.Date,
			&c.Expiry,
			&c.Status,
			&aId))
		a, err := d.GetApartment(aId)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		c.Apartment = *a
		cs = append(cs, c)
	}

	return cs, nil
}

func (d Database) GetAllTransactions() ([]model.Transaction, error) {
	stmt, err := d.connection.Query("SELECT * FROM transactions")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer stmt.Close()

	var ts []model.Transaction
	var tId int
	var aId int

	for stmt.Next() {
		var t model.Transaction
		fmt.Println(stmt.Scan(&t.Id,
			&t.Price,
			&tId,
			&aId,
			&t.Date,
			&t.Status))
		ten, err := d.GetTenant(tId)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		a, err := d.GetApartment(aId)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		t.Tenant = *ten
		t.Apartment = *a

		ts = append(ts, t)
	}
	return ts, nil
}

func (d Database) GetAllHistories() ([]model.ApartmentRentalHistory, error) {
	stmt, err := d.connection.Query("SELECT * FROM histories")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer stmt.Close()

	var hs []model.ApartmentRentalHistory
	var aId int
	var tId int

	for stmt.Next() {
		var h model.ApartmentRentalHistory
		fmt.Println(stmt.Scan(&h.Id,
			&aId,
			&tId,
			&h.RentBegin,
			&h.RentEnd))
		a, err := d.GetApartment(aId)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		t, err := d.GetTenant(tId)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		h.Apartment = *a
		h.Tenant = *t

		hs = append(hs, h)
	}
	return hs, nil
}

// ——————————————————————————————————————————————————————————————————————————————————————

// Get Instance
func (d Database) GetApartment(id int) (*model.Apartment, error) {
	stmt, err := d.connection.Prepare("SELECT * FROM apartments WHERE apartmentId = ?")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer stmt.Close()

	var a *model.Apartment = new(model.Apartment)
	var tId int
	err = stmt.QueryRow(id).Scan(&a.Id,
		&a.City,
		&a.Address,
		&a.Area,
		&tId,
		&a.RentPrice)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Fatal("No rows were found.")
			return nil, err
		}
	}

	a.Tenant, err = d.GetTenant(tId)

	return a, nil
}

func (d Database) GetTenant(id int) (*model.Tenant, error) {
	stmt, err := d.connection.Prepare("SELECT * FROM tenants WHERE tenantId = ?")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer stmt.Close()

	var t *model.Tenant = new(model.Tenant)
	err = stmt.QueryRow(id).Scan(&t.Id,
		&t.FirstName,
		&t.LastName,
		&t.Login,
		&t.Password)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Fatal("No rows were found.")
			return nil, err
		}
	}
	return t, nil
}

func (d Database) GetPayment(id int) (*model.ApartmentCost, error) {
	stmt, err := d.connection.Prepare("SELECT * FROM payments WHERE paymentId = ?")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer stmt.Close()

	var c *model.ApartmentCost = new(model.ApartmentCost)
	var aId int
	err = stmt.QueryRow(id).Scan(&c.Id,
		&c.Price,
		&c.Date,
		&c.Expiry,
		&c.Status,
		&aId)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Fatal("No rows were found.")
			return nil, err
		}
	}

	a, err := d.GetApartment(aId)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	c.Apartment = *a

	return c, nil
}

func (d Database) GetTransaction(id int) (*model.Transaction, error) {
	stmt, err := d.connection.Prepare("SELECT * FROM transactions WHERE transactionId = ?")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer stmt.Close()

	var t *model.Transaction = new(model.Transaction)
	var tId int
	var aId int
	err = stmt.QueryRow(id).Scan(&t.Id,
		&t.Price,
		&tId,
		&aId,
		&t.Date,
		&t.Status)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Fatal("No rows were found.")
			return nil, err
		}
	}

	ten, err := d.GetTenant(tId)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	apt, err := d.GetApartment(aId)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	t.Tenant = *ten
	t.Apartment = *apt

	return t, nil
}

func (d Database) GetHistory(id int) (*model.ApartmentRentalHistory, error) {
	stmt, err := d.connection.Prepare("SELECT * FROM histories WHERE historyId = ?")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer stmt.Close()

	var h *model.ApartmentRentalHistory = new(model.ApartmentRentalHistory)
	var aId int
	var tId int
	err = stmt.QueryRow(id).Scan(&h.Id,
		&aId,
		&tId,
		&h.RentBegin,
		&h.RentEnd)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Fatal("No rows were found.")
			return nil, err
		}
	}

	a, err := d.GetApartment(aId)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	t, err := d.GetTenant(tId)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	h.Apartment = *a
	h.Tenant = *t

	return h, nil
}

// ——————————————————————————————————————————————————————————————————————————————————————

// Update
func (d Database) UpdateApartment(id int, a model.Apartment) *model.Apartment {
	stmt, err := d.connection.Prepare(`UPDATE apartments 
 									SET city = ?, address = ?, area = ?, tenantId = ?, rentPrice = ?
 									WHERE apartmentId = ?`)
	if err != nil {
		log.Fatal(err)
		// return false
	}
	defer stmt.Close()

	tId, err := d.GetTenant(a.Tenant.Id)
	if err != nil {
		log.Fatal(err)
		// return false
	}

	if _, err := stmt.Exec(a.City, a.Address, a.Area, tId, a.RentPrice, id); err != nil {
		log.Fatal(err)
		// return false
	}
	return &a
}

func (d Database) UpdateTenant(id int, t model.Tenant) *model.Tenant {
	stmt, err := d.connection.Prepare(`UPDATE tenants 
 									SET firstName = ?, lastName = ?, userName = ?, password = ?
 									WHERE tenantId = ?`)
	if err != nil {
		log.Fatal(err)
		// return false
	}
	defer stmt.Close()

	if _, err := stmt.Exec(t.FirstName, t.LastName, t.Login, t.Password, id); err != nil {
		log.Fatal(err)
		// return false
	}
	return &t
}

func (d Database) UpdatePayment(id int, c model.ApartmentCost) *model.ApartmentCost {
	stmt, err := d.connection.Prepare(`UPDATE payments
									SET price = ?, date = ?, expiry = ?, status = ?, apartmentId = ?
 									WHERE paymentId = ?`)

	if err != nil {
		log.Fatal(err)
		// return false
	}
	defer stmt.Close()

	aId, err := d.GetApartment(c.Apartment.Id)
	if err != nil {
		log.Fatal(err)
		// return false
	}

	if _, err := stmt.Exec(c.Price, c.Date, c.Expiry, c.Status, aId); err != nil {
		log.Fatal(err)
		// return false
	}
	return &c
}

func (d Database) UpdateTransaction(id int, t model.Transaction) *model.Transaction {
	stmt, err := d.connection.Prepare(`UPDATE transactions 
 									SET price = ?, tenantId = ?, apartmentId = ?, date = ?, status = ?
 									WHERE transactionId = ?`)
	if err != nil {
		log.Fatal(err)
		// return false
	}
	defer stmt.Close()

	tId, err := d.GetTenant(t.Tenant.Id)
	if err != nil {
		log.Fatal(err)
		// return false
	}

	aId, err := d.GetApartment(t.Apartment.Id)
	if err != nil {
		log.Fatal(err)
		// return false
	}

	if _, err := stmt.Exec(t.Price, tId, aId, t.Date, t.Status, id); err != nil {
		log.Fatal(err)
		// return false
	}
	return &t
}

func (d Database) UpdateHistory(id int, h model.ApartmentRentalHistory) *model.ApartmentRentalHistory {
	stmt, err := d.connection.Prepare(`UPDATE histories 
 									SET apartmentId = ?, tenantId = ?, rentBegin = ?, rentEnd = ?
 									WHERE historyId = ?`)
	if err != nil {
		log.Fatal(err)
		// return false
	}
	defer stmt.Close()

	aId, err := d.GetApartment(h.Apartment.Id)
	if err != nil {
		log.Fatal(err)
		// return false
	}
	tId, err := d.GetTenant(h.Tenant.Id)
	if err != nil {
		log.Fatal(err)
		// return false
	}

	if _, err := stmt.Exec(aId, tId, h.RentBegin, h.RentEnd, id); err != nil {
		log.Fatal(err)
		// return false
	}
	return &h
}

// ——————————————————————————————————————————————————————————————————————————————————————

// Auxiliary
func (d Database) GetApartmentSearch(search string) (*model.Apartment, error) {
	stmt, err := d.connection.Prepare(`SELECT * FROM apartments
									WHERE apartmentId LIKE '%?%'
										OR city LIKE '%?%'
										OR address LIKE '%?%'
										OR area LIKE '%?%'
										OR tenantId LIKE '%?%'
										OR rentPrice LIKE '%?%';`)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer stmt.Close()

	var a *model.Apartment = new(model.Apartment)
	err = stmt.QueryRow(search).Scan(a.Id,
		a.City,
		a.Address,
		a.Area,
		a.Tenant,
		a.RentPrice)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Fatal("No rows were found.")
			return nil, err
		}
	}
	return a, nil
}
