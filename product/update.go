package product

import "ecommerce/domain"

func (svc *service) Update(prdct domain.Product) (*domain.Product, error) {
	return svc.productRepo.Update(prdct)
}
