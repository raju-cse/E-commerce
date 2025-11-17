package repo

import (
	"database/sql"
	"ecommerce/domain"
	"ecommerce/product"
	"fmt"

	"github.com/jmoiron/sqlx"
)

// type Product struct{
// 	ID           int     `json:"id" db:"id"`
// 	Title        string  `json:"title" db:"title"`
// 	Description  string  `json:"description" db:"description"`
// 	Price        float64 `json:"price" db:"price"`
// 	IngUrl       string  `json:"imageUrl" db:"img_url"`
// }

type ProductRepo interface{
	product.ProductRepo
}

type productRepo struct{
	db *sqlx.DB
}

func NewProductRepo(db *sqlx.DB) ProductRepo {
    return &productRepo{
        db: db,
    }
}

func (r *productRepo) Create(p domain.Product) (*domain.Product, error){

	query := `
		INSERT INTO products (
		 title,
		 description,
		 price,
		 img_url
		)
		VALUES (
		 $1,
	   $2,
	   $3,
	   $4
		)
		RETURNING id
	`
  row := r.db.QueryRow(query, p.Title, p.Description, p.Price, p.IngUrl)
  err := row.Scan(&p.ID)
  if err != nil{
		fmt.Println(err)
		return nil, err
	}
	return &p, nil
}

func (r *productRepo) Get(id int) (*domain.Product, error){

	var prd domain.Product

	query := `
	  SELECT
		  id,
			title,
			description,
			price,
			img_url
		from products
		where id = $1
	`
  err := r.db.Get(&prd, query, id)
  if err != nil{
		if err == sql.ErrNoRows{
			return nil, nil
		}
		return nil, err
	}

	return &prd, nil
}

func (r *productRepo) List(page, limit int64) ([]*domain.Product, error){
  
	offset := ((page - 1) * limit) + 1

	var prdList []*domain.Product

	query := `
	  SELECT
		  id,
			title,
			description,
			price,
			img_url
		from products
		LIMIT $1
		OFFSET $2;
	`
  err := r.db.Select(&prdList, query, limit, offset)
  if err != nil{
		return nil, err
	}

	return prdList, nil

}

func (r *productRepo) Count() (int64, error) {
    query := `
        SELECT COUNT(*) 
        FROM products
    `

    var count int64
    err := r.db.QueryRow(query).Scan(&count)   
    if err != nil {
        return 0, err
    }

    return count, nil
}

func (r *productRepo)	Delete(id int) error{
	 query := `
	  DELETE FROM products WHERE id = $1
	 `
	_, err := r.db.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}

func (r *productRepo)	Update(p domain.Product) (*domain.Product, error){
	 query := `
	  UPDATE products
	  SET title=$1, description=$2, price=$3, img_url=$4
	  WHERE id = $5
	 `
	 row := r.db.QueryRow(query, p.Title, p.Description, p.Price, p.IngUrl)
	 err := row.Err()
	 if err != nil{
		return nil, err
	 }
	  return &p, nil
	}

// func generateInitialProducts(r *productRepo){

// 	prd1 := &Product{
// 		ID: 1,
// 		Title: "Orange",
// 		Description: "I love orange",
// 		Price: 300,
// 		IngUrl: "https://www.dole.com/sites/default/files/media/2025-01/oranges.png",
// 	}

// 	prd2 := &Product{
// 		ID: 2,
// 		Title: "Banana",
// 		Description: "I love Banana",
// 		Price: 300,
// 		IngUrl: "https://encrypted-tbn3.gstatic.com/images?q=tbn:ANd9GcS5i9ZZAnf6oOD7l5oE56xsU7jqpg9TYEGPc5pwXsz0Ly-7RrT7jvCb3Ruh_zajoZdOT-gohM-CGntkupYvTTD8PRp1SCAkGkjBODqaww",
// 	}

// 	prd3 := &Product{
// 		ID: 3,
// 		Title: "Mango",
// 		Description: "I love Mango",
// 		Price: 100,
// 		IngUrl: "https://www.dole.com/sites/default/files/styles/1024w768h-80/public/media/2025-01/mangos.png?itok=pS1eZJtK-r6z70DGP",
// 	}

// 	 r.productList = append(r.productList, prd1)
// 	 r.productList = append(r.productList, prd2)
// 	 r.productList = append(r.productList, prd3)

// }

// package repo

// import (
// 	"database/sql"
// 	"fmt"

// 	"github.com/jmoiron/sqlx"
// )

// type Product struct {
// 	ID          int     `json:"id" db:"id"`
// 	Title       string  `json:"title" db:"title"`
// 	Description string  `json:"description" db:"description"`
// 	Price       float64 `json:"price" db:"price"`
// 	IngUrl      string  `json:"imageUrl" db:"img_url"`
// }

// type ProductRepo interface {
// 	Create(p Product) (*Product, error)
// 	Get(productID int) (*Product, error)
// 	List() ([]*Product, error)
// 	Delete(productID int) error
// 	Update(p Product) (*Product, error)
// }

// type productRepo struct {
// 	db *sqlx.DB
// }

// func NewProductRepo(db *sqlx.DB) ProductRepo {
// 	return &productRepo{
// 		db: db,
// 	}
// }

// func (r *productRepo) Create(p Product) (*Product, error) {
// 	query := `
// 		INSERT INTO products (title, description, price, img_url)
// 		VALUES ($1, $2, $3, $4)
// 		RETURNING id
// 	`
// 	row := r.db.QueryRow(query, p.Title, p.Description, p.Price, p.IngUrl)
// 	err := row.Scan(&p.ID)
// 	if err != nil {
// 		fmt.Println(err)
// 		return nil, err
// 	}
// 	return &p, nil
// }

// func (r *productRepo) Get(id int) (*Product, error) {
// 	var prd Product
// 	query := `
// 		SELECT id, title, description, price, img_url
// 		FROM products
// 		WHERE id = $1
// 	`
// 	err := r.db.Get(&prd, query, id)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			return nil, nil
// 		}
// 		return nil, err
// 	}
// 	return &prd, nil
// }

// func (r *productRepo) List() ([]*Product, error) {
// 	var prdList []*Product
// 	query := `
// 		SELECT id, title, description, price, img_url
// 		FROM products
// 	`
// 	err := r.db.Select(&prdList, query)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return prdList, nil
// }

// func (r *productRepo) Delete(id int) error {
// 	query := `DELETE FROM products WHERE id = $1`
// 	_, err := r.db.Exec(query, id)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (r *productRepo) Update(p Product) (*Product, error) {
// 	query := `
// 		UPDATE products
// 		SET title=$1, description=$2, price=$3, img_url=$4
// 		WHERE id=$5
// 	`
// 	_, err := r.db.Exec(query, p.Title, p.Description, p.Price, p.IngUrl, p.ID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &p, nil
// }
