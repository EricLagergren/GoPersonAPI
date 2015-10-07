package person

// Model is a ...
type Model struct {
	ID           int64  `json:"id"`
	FirstName    string `json:"firstname"`
	LastName     string `json:"lastname"`
	EmailAddress string `json:"emailaddress"`
	Phone        string `json:"phone"`
	Weight       int    `json:"weight"`
}
