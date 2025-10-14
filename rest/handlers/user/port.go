package user

import "ecommerce/domain"


type Service interface{
	Create(user domain.User) (*domain.User, error)
	Find(email string, pass string) (*domain.User, error)
}