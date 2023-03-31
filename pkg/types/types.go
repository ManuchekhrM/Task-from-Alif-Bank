package types

type Money int64

// Phone is the method from which the account is registered
type Phone string

// Some user who has account in shop
type Account struct {
	ID      int64
	Phone   Phone
	Balance Money
}

// ProductCategory is list of all category we have
type ProductCategory string

// Our period for instalment
const (
	ThreeMonth      string = "3"
	SixMonth        string = "6"
	NineMonth       string = "9"
	TwelveMonth     string = "12"
	EighteenMonth   string = "18"
	TwentyFourMonth string = "24"
)

// Our Category
const (
	Smartphone string = "smartphone"
	Laptop     string = "laptop"
	TV         string = "tv"
)

// Show us information about instalment
type Instalment struct {
	ID              string
	AccountID       int64
	ProductCategory ProductCategory
	AccountPhone    string
	Amount          Money
	RangInstalment  int
}
