package user

import "ecommerce/domain"

type service struct {
	UserRepo UserRepo
}

func NewService(userRepo UserRepo) Service {
	return &service{
		UserRepo: userRepo,
	}
}
func (svc *service) Create(user domain.User) (*domain.User, error) {
	usr, err := svc.UserRepo.Create(user)
	if err != nil {
		return nil, err
	}
	if usr == nil {
		return nil, nil
	}
	return usr, nil
}

func (svc *service) Find(email, pass string) (*domain.User, error) {
	usr, err := svc.UserRepo.Find(email, pass)
	if err != nil {
		return nil, err
	}
	if usr == nil {
		return nil, nil
	}
	return usr, nil
}
