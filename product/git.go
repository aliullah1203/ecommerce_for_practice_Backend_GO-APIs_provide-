package product

import "ecommerce/domain"

func (svc *service) Get(id int) (*domain.Product, error) {
	return svc.productRepo.Get(id)
}
