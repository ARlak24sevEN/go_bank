package service

import (
	"bank/errs"
	"bank/logs"
	"bank/repository"
	"database/sql"
	"net/http"
)

type customerService struct {
	custRepo repository.CustomerRepository
}

func NewCustomerService(custRepo repository.CustomerRepository) CustomerService {
	return customerService{custRepo: custRepo}
}

func (s customerService) GetCustomers() ([]CustomerResponse, error) {
	customers, err := s.custRepo.GetAll()
	if err != nil {
		// log.Println(err)
		logs.Error(err)
		return nil, err
	}
	custResponses := []CustomerResponse{}
	for _, customer := range customers {
		custResponse := CustomerResponse{CustomerID: customer.CustomerID,
			Name: customer.Name, Status: customer.Status,
		}
		custResponses = append(custResponses, custResponse)
	}
	return custResponses, nil
}

func (s customerService) GetCustomer(id int) (*CustomerResponse, error) {
	customer, err := s.custRepo.GetById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.AppError{
				Code:    http.StatusNotFound,
				Message: "customer not found ",
			}
		}
		// log.Println(err)
		logs.Error(err)
		return nil, errs.AppError{
			Code:    http.StatusInternalServerError,
			Message: "unexpected error",
		}
	}
	custResponse := CustomerResponse{
		CustomerID: customer.CustomerID, Name: customer.Name, Status: customer.Status,
	}
	return &custResponse, nil
}
