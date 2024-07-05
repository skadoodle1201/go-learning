package tools

import "time"

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"tanish": {
		AuthToken: "Test123",
		Username:  "coolBoii",
	},
	"tanu": {
		AuthToken: "Test124",
		Username:  "coolGirl",
	},
	"rajeshwari": {
		AuthToken: "Test143",
		Username:  "Bhendi",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"tanish": {
		Coins:    400,
		Username: "coolBoii",
	},
	"tanu": {
		Coins:    600,
		Username: "coolGirl",
	},
	"rajeshwari": {
		Coins:    3,
		Username: "Bhendi",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(time.Second * 1)
	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	time.Sleep(time.Second * 1)
	var coinDataForUser = CoinDetails{}
	coinDataForUser, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}

	return &coinDataForUser
}

func (d *mockDB) SetupDatabase() error {
	return nil
}
