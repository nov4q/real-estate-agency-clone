package api

import (
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"polsl/tab/estate-agency/internal/model"
)

type DBManager interface {
	AddApartment(a model.Apartment) *model.Apartment
	AddTenant(t model.Tenant) *model.Tenant
	AddPayment(c model.ApartmentCost) *model.ApartmentCost
	AddTransaction(t model.Transaction) *model.Transaction
	AddHistory(h model.ApartmentRentalHistory) *model.ApartmentRentalHistory
	GetAllApartments() ([]model.Apartment, error)
	GetAllTenants() ([]model.Tenant, error)
	GetAllPayments() ([]model.ApartmentCost, error)
	GetAllTransactions() ([]model.Transaction, error)
	GetAllHistories() ([]model.ApartmentRentalHistory, error)
	GetApartment(id int) (*model.Apartment, error)
	GetTenant(id int) (*model.Tenant, error)
	GetPayment(id int) (*model.ApartmentCost, error)
	GetTransaction(id int) (*model.Transaction, error)
	GetHistory(id int) (*model.ApartmentRentalHistory, error)
	UpdateApartment(id int, a model.Apartment) *model.Apartment
	UpdateTenant(id int, t model.Tenant) *model.Tenant
	UpdatePayment(id int, c model.ApartmentCost) *model.ApartmentCost
	UpdateTransaction(id int, t model.Transaction) *model.Transaction
	UpdateHistory(id int, h model.ApartmentRentalHistory) *model.ApartmentRentalHistory
	GetApartmentSearch(search string) (*model.Apartment, error)
}

type Handler struct {
	dbManager    DBManager
	log          *log.Logger
	getApartment string
	getRenter    string
}

func NewHandler(dbManager DBManager, log *log.Logger) Handler {
	return Handler{
		dbManager:    dbManager,
		log:          log,
		getApartment: "/v1/apartment/{id}",
		getRenter:    "/v1/renter/{id}",
	}
}

func (h Handler) InitializeEndpoints(mux *mux.Router) {
	mux.HandleFunc(h.getApartment, h.GetApartment).Methods("GET")
	mux.HandleFunc(h.getRenter, h.GetRenter).Methods("GET")
}

func (h Handler) GetApartment(w http.ResponseWriter, r *http.Request) {
	/*params := mux.Vars(r)
	key := params["id"]
	url := strings.Replace(h.getApartment, "{id}", key, 1)
	if r.URL.Path != url {
		h.handleError(w, http.StatusNotFound, "%s: 404 not found", h.getFuncName())
		return
	}/*/

}

func (h Handler) GetRenter(w http.ResponseWriter, r *http.Request) {
}
