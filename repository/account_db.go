package repository

import "github.com/jmoiron/sqlx"

type accountRepositoryDb struct {
	db *sqlx.DB
}

func NewAccountRepositoryDb(db *sqlx.DB) AccountRepository {
	return accountRepositoryDb{db: db}
}

func (r accountRepositoryDb) Create(acc Account) (*Account, error) {

	query := "INSERT INTO accounts ( customer_id, opening_date, account_type, amount, status) VALUES(?,?,?,?,?)"
	result, err := r.db.Exec(query, acc.CustomerID, acc.OpeningDate, acc.AccountType, acc.Amount, acc.Status)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	acc.AccountID = int(id)

	return &acc, nil
}

func (r accountRepositoryDb) GetAll(customerID int) ([]Account, error) {
	accounts := []Account{}

	query := "SELECT account_id, customer_id, opening_date, account_type, amount, status FROM bank.accounts where customer_id = ?"
	err := r.db.Select(&accounts, query, customerID)
	if err != nil {
		return nil, err
	}
	return accounts, err
}
