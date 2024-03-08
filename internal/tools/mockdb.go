package tools

import (
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"alex": {
		AuthToken: "123ABC",
		Username:  "alex",
	},
	"jason": {
		AuthToken: "456DEF",
		Username:  "jason",
	},
	"marie": {
		AuthToken: "789GHI",
		Username:  "marie",
	},
}
var mockCoinDetails = map[string]CoinDetails{
	"alex": {
		Coins:    100,
		Username: "alex",
	},
	"jason": {
		Coins:    100,
		Username: "jason",
	},
	"marie": {
		Coins:    100,
		Username: "marie",
	},
} //in a real world app this would be in a DB
func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(time.Second + 1)
	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}
func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	//simulate db call
	time.Sleep(time.Second + 1)
	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username] //function looks up data from our map
	if !ok {
		return nil
	}
	return &clientData
}
func (d *mockDB) SetupDatabase() error {
	return nil //does nothing ?
}
