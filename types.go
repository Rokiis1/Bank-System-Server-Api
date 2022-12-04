package main

type Account struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Number    int64  `json:"number"`
	Balance   int64  `json:"balance"`
}

func NewAccount(firstName, lastName string) *Account {
	return &Account{
		// rand.Intn(10000)
		ID: 1,
		FirstName: firstName,
		LastName: lastName,
		// int64(rand.Intn(1000000))
		Number: 1,
	}
}