package main

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ydhnwb/golang_api/config"
	"github.com/ydhnwb/golang_api/controller"
	"github.com/ydhnwb/golang_api/entity"
	"github.com/ydhnwb/golang_api/middleware"
	"github.com/ydhnwb/golang_api/repository"
	"github.com/ydhnwb/golang_api/service"
	"gorm.io/gorm"
)

var (
	db                *gorm.DB                     = config.SetupDatabaseConnection()
	userRepository    repository.UserRepository    = repository.NewUserRepository(db)
	ratingRepository  repository.RatingRepository  = repository.NewRatingRepository(db)
	productRepository repository.ProductRepository = repository.NewProductRepository(db)
	jwtService        service.JWTService           = service.NewJWTService()
	userService       service.UserService          = service.NewUserService(userRepository)
	authService       service.AuthService          = service.NewAuthService(userRepository)
	ratingService     service.RatingService        = service.NewRatingService(ratingRepository)
	productService    service.ProductService       = service.NewProductService(productRepository)
	authController    controller.AuthController    = controller.NewAuthController(authService, jwtService)
	userController    controller.UserController    = controller.NewUserController(userService, jwtService, authService)
	ratingController  controller.RatingController  = controller.NewRatingController(ratingService)
	productController controller.ProductController = controller.NewProductController(productService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	fmt.Println("Running")
	db.AutoMigrate(&entity.Rating{})
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"}
	config.AddAllowHeaders("Content-Type")
	config.AddAllowHeaders("Authorization")
	r.Use(cors.New(config))
	//Serving Static Files Like Images.
	r.Static("/public/images", "./public/images")
	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	userRoutes := r.Group("api/user", middleware.AuthorizeUser(jwtService, authService))
	{
		userRoutes.GET("/profile", userController.Profile)
		userRoutes.PUT("/profile", userController.Update)

	}

	userAdminRoutes := r.Group("api/user", middleware.AuthorizeAdmin(jwtService, authService))
	{

		userAdminRoutes.POST("/all", userController.AllUsers)
		userAdminRoutes.DELETE("/delete/:id", userController.DeleteUser)
		userAdminRoutes.POST("/update", userController.Update)
		userAdminRoutes.POST("/create", authController.Register)
	}

	ratingRoutes := r.Group("api/rating", middleware.AuthorizeUser(jwtService, authService))
	{
		ratingRoutes.POST("/", ratingController.Insert)
	}

	ratingAdminRoutes := r.Group("api/rating", middleware.AuthorizeAdmin(jwtService, authService))
	{
		ratingAdminRoutes.POST("/all", ratingController.AllRatings)
	}

	productRoutes := r.Group("api/product", middleware.AuthorizeUser(jwtService, authService))
	{
		productRoutes.POST("/all", productController.GetAllProducts)
		productRoutes.GET("/:id", productController.GetProductByID)
	}
	productAdminRoutes := r.Group("api/product", middleware.AuthorizeAdmin(jwtService, authService))
	{
		productAdminRoutes.POST("/import", productController.ImportExcel)
		productAdminRoutes.POST("/delete", productController.DeleteProducts)
		productAdminRoutes.POST("/save", productController.SaveProduct)
	}
	r.Run()
}
