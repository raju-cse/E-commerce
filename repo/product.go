package repo


type Product struct{
	ID           int     `json:"id"`
	Title        string  `json:"title"`
	Description  string  `json:"description"`
	Price        float64 `json:"price"`
	IngUrl       string  `json:"imageUrl"`
}


type ProductRepo interface{
	Create(p Product) (*Product, error)
	Get(productID int) (*Product, error)
	List() ([]*Product, error)
	Delete(productID int) error
	Update(p Product) (*Product, error)
}

type productRepo struct{
	productList []*Product
} 

func NewProductRepo() productRepo {
	repo := &productRepo{}

	generateInitialProducts(repo)
	return *repo

}

func (r *productRepo) Create(p Product) (*Product, error){
	p.ID = len(r.productList)+1
	r.productList = append(r.productList, &p)
	return  &p, nil
}
func (r *productRepo) Get(productID int) (*Product, error){
	for _, product := range r.productList{
		if product.ID == productID{
			return product, nil
		}
	}

	return nil , nil

}
func (r *productRepo) List() ([]*Product, error){
	return r.productList, nil

}
func (r *productRepo)	Delete(productID int) error{
		var tempList []*Product

	for _, p := range r.productList{
		if p.ID == productID{
			tempList = append(tempList, p)
		}
	}

	r.productList = tempList
	return nil

}
func (r *productRepo)	Update(product Product) (*Product, error){
		for idx, p := range r.productList{
		if p.ID == product.ID{
			r.productList[idx] = &product
		}
	}
  return &product, nil
}

func generateInitialProducts(r *productRepo){

	prd1 := &Product{
		ID: 1,
		Title: "Orange",
		Description: "I love orange",
		Price: 300,
		IngUrl: "https://www.dole.com/sites/default/files/media/2025-01/oranges.png",
	}

	prd2 := &Product{
		ID: 2,
		Title: "Banana",
		Description: "I love Banana",
		Price: 300,
		IngUrl: "https://encrypted-tbn3.gstatic.com/images?q=tbn:ANd9GcS5i9ZZAnf6oOD7l5oE56xsU7jqpg9TYEGPc5pwXsz0Ly-7RrT7jvCb3Ruh_zajoZdOT-gohM-CGntkupYvTTD8PRp1SCAkGkjBODqaww",
	}

	prd3 := &Product{
		ID: 3,
		Title: "Mango",
		Description: "I love Mango",
		Price: 100,
		IngUrl: "https://www.dole.com/sites/default/files/styles/1024w768h-80/public/media/2025-01/mangos.png?itok=pS1eZJtK-r6z70DGP",
	}

	

	 r.productList = append(r.productList, prd1)
	 r.productList = append(r.productList, prd2)
	 r.productList = append(r.productList, prd3)
	
}
