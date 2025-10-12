package product

type service struct {
	productRepo ProductRepo
}

func NewService(prdct ProductRepo) Service {
	return &service{
		productRepo: prdct,
	}
}
