package user

// for returning user data
type UserDTO struct {
	ID            int    `json:"id"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Email         string `json:"email"`
	PhoneNumber   string `json:"phone_number,omitempty"`
	PhoneVerified bool   `json:"phone_verified"`
	SocialID      string `json:"social_id,omitempty"`
}

// for creating user
type CreateUserDTO struct {
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	Email            string `json:"email"`
	Password         string `json:"password"`
	PhoneNumber      string `json:"phone_number,omitempty"`
	SocialID         string `json:"social_id,omitempty"`
	VerificationCode string `json:"verification_code,omitempty"`
}

// UpdateUserDTO is used for updating user details (for PUT requests)
type UpdateUserDTO struct {
	FirstName   string `json:"first_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	Email       string `json:"email,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	Password    string `json:"password,omitempty"`
	Address     string `json:"address,omitempty"`
}
