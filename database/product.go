package database

var ProductList []Product

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImagUrl     string  `json:"imageUrl"`
}

func init() {
	prd1 := Product{
		ID:          1,
		Title:       "Orange",
		Description: "Orange is red, I love orange",
		Price:       100.00,
		ImagUrl:     "https://upload.wikimedia.org/wikipedia/commons/e/e3/Oranges_-_whole-halved-segment.jpg",
	}
	prd2 := Product{
		ID:          2,
		Title:       "Apple",
		Description: "Apple is green, I hate apple",
		Price:       80.00,
		ImagUrl:     "https://www.collinsdictionary.com/images/full/apple_158989157.jpg",
	}
	prd3 := Product{
		ID:          3,
		Title:       "Banana",
		Description: "Banana is yellow, I love banana",
		Price:       15.00,
		ImagUrl:     "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQMFueRcp7oghbY4-YTxJQmhfLrQpiD0Q_DLw&s",
	}
	prd4 := Product{
		ID:          4,
		Title:       "Grapes",
		Description: "grapes is bichi, hahahha",
		Price:       100.00,
		ImagUrl:     "https://upload.wikimedia.org/wikipedia/commons/e/e3/Oranges_-_whole-halved-segment.jpg",
	}
	prd5 := Product{
		ID:          5,
		Title:       "Watermelon",
		Description: "Watermelon is pure red, I love Watermelon too much",
		Price:       500.00,
		ImagUrl:     "https://cdn.mos.cms.futurecdn.net/v2/t:0,l:200,cw:1200,ch:1200,q:80,w:1200/XenViG9cC4EdGupeibtKa5.jpg",
	}
	prd6 := Product{
		ID:          6,
		Title:       "Mango",
		Description: "Mango is my favourite, I love mango",
		Price:       500.00,
		ImagUrl:     "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSe4OzPMUZFtQ-TJxNf5k37h07WhwkJjgkezQ&s",
	}
	ProductList = append(ProductList, prd1)
	ProductList = append(ProductList, prd2)
	ProductList = append(ProductList, prd3)
	ProductList = append(ProductList, prd4)
	ProductList = append(ProductList, prd5)
	ProductList = append(ProductList, prd6)
}
