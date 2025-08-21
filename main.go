package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func hellohHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Hello Bangladesh")
}

func aboutHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "I m jahid , i'm Youtuber & software enginner")
}

type Product struct{
	ID           int     `json:"id"`
	Title        string  `json:"title"`
	Description  string  `json:"description"`
	Price        float64 `json:"price"`
	IngUrl       string  `json:"imageUrl"`
}

var productList []Product

func getProducts(w http.ResponseWriter, r *http.Request){

  handleCors(w)
  handlePreflightReg(w, r) 
	
	if r.Method != "GET"{
	http.Error(w, "give me GET request", 400)
	return
 }
 sendData(w, productList, 200)
//  encoder := json.NewEncoder(w)
//  encoder.Encode(productList)

}

func createProduct(w http.ResponseWriter, r *http.Request){
	
  handleCors(w)
  handlePreflightReg(w, r)

  if r.Method != "POST"{
		http.Error(w, "Plz give me post requst", 400)
		return
	}

	var newProduct Product  //Create struct instance 

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)

	if err != nil{
		fmt.Println(err)
		http.Error(w, "plz give me valid json", 400)
		return
	}

	newProduct.ID = len(productList)+1
	productList = append(productList, newProduct)

	sendData(w, newProduct, 201)
	// w.WriteHeader(201)
	// encoder := json.NewEncoder(w)
	// encoder.Encode(newProduct)

}

func handleCors(w http.ResponseWriter){

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	

}

func handlePreflightReg(w http.ResponseWriter, r *http.Request){

	 if r.Method == "OPTIONS"{
		w.WriteHeader(200)
	}
}

func sendData(w http.ResponseWriter, data interface{}, statusCode int){

	w.WriteHeader(statusCode)
	encode := json.NewEncoder(w)
	encode.Encode(data)
}

func main(){

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hellohHandler)
 	mux.HandleFunc("/about", aboutHandler)
	mux.HandleFunc("/products", getProducts)
	mux.HandleFunc("/create-products", createProduct)


	fmt.Println("Server running on : 8080")
	err := http.ListenAndServe(":8080", mux)

	if err != nil{
		fmt.Println("Error startin the Servr :", err)
	}

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

	// prd4 := Product{
	// 	ID: 4,
	// 	Title: "Apple",
	// 	Description: "I love Apple",
	// 	Price: 6200,
	// 	IngUrl: "https://static.vecteezy.com/system/resources/thumbnails/012/086/172/small_2x/green-apple-with-green-leaf-isolated-on-white-background-vector.jpg",
	// }

	// prd5 := Product{
	// 	ID: 5,
	// 	Title: "Anar",
	// 	Description: "I love Anar",
	// 	Price: 6000,
	// 	IngUrl: "https://sevasahayog.org/wp-content/uploads/2024/06/Anar.jpeg",
	// }

	// prd6 := Product{
	// 	ID: 6,
	// 	Title: "Pape",
	// 	Description: "I love Pape",
	// 	Price: 300,
	// 	IngUrl: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQ1OQ3WHKGIyLXoJBlH184WL-eswOup8tlZmKV_G90C0Q0DFQTsQ1QiCyQlLs4iMMl43Uw&usqp=CAU",
	// }

	 productList = append(productList, prd1)
	 productList = append(productList, prd2)
	 productList = append(productList, prd3)
	//  productList = append(productList, prd4)
	//  productList = append(productList, prd5)
	//  productList = append(productList, prd6)
}