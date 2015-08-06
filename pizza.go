package pizza

// StreetAddress represents
type StreetAddress struct {
	Street  string
	Street2 string
	City    string
	State   string
	Zip     string
}

// Store represents
type Store struct {
	ID      string
	Address StreetAddress
}
