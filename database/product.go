package database

var productList []Product

type Product struct{
	ID           int     `json:"id"`
	Title        string  `json:"title"`
	Description  string  `json:"description"`
	Price        float64 `json:"price"`
	IngUrl       string  `json:"imageUrl"`
}

func Store(p Product) Product{
	p.ID = len(productList)+1
	productList = append(productList, p)
	return  p
}

func List() []Product{
	return productList
}

func Get(productID int) *Product{

	for _, product := range productList{
		if product.ID == productID{
			return &product
		}
	}

	return nil 
}

func Update(product Product){
	for idx, p := range productList{
		if p.ID == product.ID{
			productList[idx] = product
		}
	}
}

func Delete(productID int){
	var tempList []Product

	for _, p := range productList{
		if p.ID == productID{
			tempList = append(tempList, p)
		}
	}

	productList = tempList
}

func init(){

	prd1 := Product{
		ID: 1,
		Title: "Orange",
		Description: "I love orange",
		Price: 300,
		IngUrl: "https://www.dole.com/sites/default/files/media/2025-01/oranges.png",
	}

	prd2 := Product{
		ID: 2,
		Title: "Banana",
		Description: "I love Banana",
		Price: 300,
		IngUrl: "https://encrypted-tbn3.gstatic.com/images?q=tbn:ANd9GcS5i9ZZAnf6oOD7l5oE56xsU7jqpg9TYEGPc5pwXsz0Ly-7RrT7jvCb3Ruh_zajoZdOT-gohM-CGntkupYvTTD8PRp1SCAkGkjBODqaww",
	}

	prd3 := Product{
		ID: 3,
		Title: "Mango",
		Description: "I love Mango",
		Price: 100,
		IngUrl: "https://www.dole.com/sites/default/files/styles/1024w768h-80/public/media/2025-01/mangos.png?itok=pS1eZJtK-r6z70DGP",
	}

	

	 productList = append(productList, prd1)
	 productList = append(productList, prd2)
	 productList = append(productList, prd3)
	
}
