package customer

type Address struct {
	Street     string `json:"street" validate:"required"`
	Number     string `json:"number" validate:"required"`
	Complement string `json:"complement,omitempty"`
	City       string `json:"city" validate:"required"`
	State      string `json:"state" validate:"required"`
	Country    string `json:"country" validate:"required"`
	ZipCode    string `json:"zip" validate:"required"`
}

type Customer struct {
	PersonalID string  `json:"personal_id" validate:"required"`
	Name       string  `json:"name" validate:"required"`
	Email      string  `json:"email" validate:"required,email"`
	Phone      string  `json:"phone" validate:"required"`
	Address    Address `json:"address" validate:"required"`
}
