package model

import (
	"time"
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type User struct {
	ID        string    `json:"id" valid:"uuid"`
	Name      string    `json:"name" valid:"notnull"`
	Email     string    `json:"email" valid:"notnull"`
	CreatedAt time.Time `json:"created_at" valid:"-"`
	UpdatedAt time.Time `json:"updated_at" valid:"-"`
}

func (user *User) isValid() error {
	_, err := govalidator.ValidateStruct(user)

	if err != nil {
		return err
	}

	return nil
}

func NewUser(name string, email string) (*User, error) {
	user := User{
		ID:        uuid.NewV4().String(),
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
	}

	err := user.isValid()
	if err != nil {
		return nil, err
	}

	return &user, nil
}
