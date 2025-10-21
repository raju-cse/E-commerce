package cmd

import (
	"ecommerce/config"
	"ecommerce/infra/db"
	"ecommerce/product"
	"ecommerce/repo"
	"ecommerce/rest"
	prodcthandler "ecommerce/rest/handlers/product"
	usrHandler "ecommerce/rest/handlers/user"
	middleware "ecommerce/rest/middlewares"
	"ecommerce/user"
	"fmt"
	"os"
)

func Serve(){
	cnf := config.GetConfig()

	// fmt.Printf("%+v", cnf.DB)

	dbCon, err := db.NewConnection(cnf.DB)
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}

	err = db.MigrateDB(dbCon, "./migrations")
  if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}

	// repos
	productRepo := repo.NewProductRepo(dbCon)
	userRepo := repo.NewUserRepo(dbCon)
	
	// domains
  userSvc := user.NewService(userRepo)
	prdctSvc := product.NewService(productRepo)


  middlewares :=	middleware.NewMiddlewares(cnf)
  
  //handlers
	productHandler := prodcthandler.NewHandler(middlewares, prdctSvc)
	userHandler := usrHandler.NewHandler(cnf, userSvc)
	

 server	:= rest.NerServer(
	cnf,
	productHandler,
	userHandler, 
)

	server.Start()
 
}