package repo

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImagUrl     string  `json:"imageUrl"`
}

type ProductRepo interface {
	Create(p Product) (*Product, error)
	Get(productId int) (*Product, error)
	List() ([]*Product, error)
	Update(product Product) (*Product, error)
	Delete(productId int) error
}

type productRepo struct {
	productList []*Product
}

// constructor or constructor function
func NewProductRepo() ProductRepo {
	repo := &productRepo{}

	generateInitialProduct(repo)
	return repo
}

func (r *productRepo) Create(p Product) (*Product, error) {
	p.ID = len(r.productList) + 1
	r.productList = append(r.productList, &p)
	return &p, nil
}

func (r *productRepo) Get(productId int) (*Product, error) {
	for _, product := range r.productList {
		if product.ID == productId {
			return product, nil
		}
	}

	return nil, nil
}
func (r *productRepo) List() ([]*Product, error) {
	return r.productList, nil
}
func (r *productRepo) Update(product Product) (*Product, error) {
	for idx, p := range r.productList {
		if p.ID == product.ID {
			r.productList[idx] = &product
		}
	}
	return &product, nil
}
func (r *productRepo) Delete(productId int) error {
	var tempList []*Product
	for _, p := range r.productList {
		if p.ID == productId {
			tempList = append(tempList, p)
		}
	}
	r.productList = tempList
	return nil
}

func generateInitialProduct(r *productRepo) {
	prd1 := &Product{
		ID:          1,
		Title:       "Orange",
		Description: "Orange is red, I love orange",
		Price:       100.00,
		ImagUrl:     "https://upload.wikimedia.org/wikipedia/commons/e/e3/Oranges_-_whole-halved-segment.jpg",
	}
	prd2 := &Product{
		ID:          2,
		Title:       "Apple",
		Description: "Apple is green, I hate apple",
		Price:       80.00,
		ImagUrl:     "https://www.collinsdictionary.com/images/full/apple_158989157.jpg",
	}
	prd3 := &Product{
		ID:          3,
		Title:       "Banana",
		Description: "Banana is yellow, I love banana",
		Price:       15.00,
		ImagUrl:     "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQMFueRcp7oghbY4-YTxJQmhfLrQpiD0Q_DLw&s",
	}
	prd4 := &Product{
		ID:          4,
		Title:       "Grapes",
		Description: "grapes is bichi, hahahha",
		Price:       100.00,
		ImagUrl:     "https://upload.wikimedia.org/wikipedia/commons/e/e3/Oranges_-_whole-halved-segment.jpg",
	}
	prd5 := &Product{
		ID:          5,
		Title:       "Watermelon",
		Description: "Watermelon is pure red, I love Watermelon too much",
		Price:       500.00,
		ImagUrl:     "https://cdn.mos.cms.futurecdn.net/v2/t:0,l:200,cw:1200,ch:1200,q:80,w:1200/XenViG9cC4EdGupeibtKa5.jpg",
	}
	prd6 := &Product{
		ID:          6,
		Title:       "Mango",
		Description: "Mango is my favourite, I love mango",
		Price:       500.00,
		ImagUrl:     "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSe4OzPMUZFtQ-TJxNf5k37h07WhwkJjgkezQ&s",
	}
	r.productList = append(r.productList, prd1)
	r.productList = append(r.productList, prd2)
	r.productList = append(r.productList, prd3)
	r.productList = append(r.productList, prd4)
	r.productList = append(r.productList, prd5)
	r.productList = append(r.productList, prd6)
}
