package tools

import (
	log "github.com/sirupsen/logrus"
)

// define the type of data our db can hold
type LoginDetails struct {
	AuthToken string //for validating the req
	Username  string
}
type CoinDetails struct {
	Coins    int64 //balance
	Username string
}
type DatabaseInterface interface {
	GetUserLoginDetails(username string) *LoginDetails
	GetUserCoins(username string) *CoinDetails
	SetupDatabase() error
} //allows us to swap out our databases really easily
func NewDatabase() (*DatabaseInterface, error) {
	var database DatabaseInterface = &mockDB{} //set it to a struct that will implement our interface
	var err error = database.SetupDatabase()
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &database, nil
}
