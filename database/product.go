package database

var ProductList []Product

type Product struct{
	ID           int     `json:"id"`
	Title        string  `json:"title"`
	Description  string  `json:"description"`
	Price        float64 `json:"price"`
	IngUrl       string  `json:"imageUrl"`
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

	

	 ProductList = append(ProductList, prd1)
	 ProductList = append(ProductList, prd2)
	 ProductList = append(ProductList, prd3)
	
}
