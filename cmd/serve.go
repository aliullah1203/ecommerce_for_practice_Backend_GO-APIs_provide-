package cmd

import (
	"ecommerce/config"
	"ecommerce/infra/db"
	"ecommerce/product"
	"ecommerce/repo"
	"ecommerce/rest"
	producthandler "ecommerce/rest/handlers/product"
	userhandler "ecommerce/rest/handlers/user"
	middleware "ecommerce/rest/middlewares"
	"ecommerce/user"
	"fmt"
	"os"
)

func Serve() {
	cnf := config.GetConfig()

	dbCon, err := db.NewConnection(cnf.DB)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := db.MigrateDB(dbCon, "./migrations"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// repositories
	productRepo := repo.NewProductRepo(dbCon)
	userRepo := repo.NewUserRepo(dbCon)

	// domain services
	userSvc := user.NewService(userRepo)
	productSvc := product.NewService(productRepo)

	middlewares := middleware.NewMiddlewares(cnf)

	// handlers
	productHandler := producthandler.NewHandler(middlewares, productSvc)
	userHandler := userhandler.NewHandler(cnf, userSvc)

	server := rest.NewServer(
		cnf,
		productHandler,
		userHandler,
	)
	server.Start()
}
