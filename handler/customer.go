package handler

import (
	"bank/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type customerHadler struct {
	custSer service.CustomerService
}

func NewCustomerHandler(custSer service.CustomerService) customerHadler {
	return customerHadler{custSer: custSer}
}

func (h customerHadler) GetCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := h.custSer.GetCustomers()
	if err != nil {
		handleError(w, err)
		return
	}
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

func (h customerHadler) GetCustomer(w http.ResponseWriter, r *http.Request) {
	customerID, _ := strconv.Atoi(mux.Vars(r)["customerID"])
	customer, err := h.custSer.GetCustomer(customerID)
	if err != nil {

		handleError(w, err)
		// appErr, ok := err.(errs.AppError)
		// if ok {
		// 	w.WriteHeader(appErr.Code)
		// 	fmt.Fprintln(w, appErr.Message)
		// 	return
		// }
		// w.WriteHeader(http.StatusInternalServerError)
		// fmt.Fprintln(w, err)
		return
	}
	w.Header().Set("content=type", "application/json")
	json.NewEncoder(w).Encode(customer)
}
