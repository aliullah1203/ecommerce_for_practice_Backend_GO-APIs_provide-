package user

import "ecommerce/domain"

type Service interface {
	Find(email, pass string) (*domain.User, error)
	Create(user domain.User) (*domain.User, error)
}
