package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ydhnwb/golang_api/config"
	"github.com/ydhnwb/golang_api/controller"
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
	userController    controller.UserController    = controller.NewUserController(userService, jwtService)
	ratingController  controller.RatingController  = controller.NewRatingController(ratingService)
	productController controller.ProductController = controller.NewProductController(productService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"}
	config.AddAllowHeaders("Content-Type")
	config.AddAllowHeaders("Authorization")
	r.Use(cors.New(config))
	// r.Use(cors.Default())
	//Serving Static Files Like Images.
	r.Static("/public/images", "./public/images")
	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	userRoutes := r.Group("api/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/profile", userController.Profile)
		userRoutes.PUT("/profile", userController.Update)
		userRoutes.POST("/all", userController.AllUsers)
		userRoutes.POST("/update", userController.Update)
		userRoutes.DELETE("/delete/:id", userController.DeleteUser)
	}

	ratingRoutes := r.Group("api/rating", middleware.AuthorizeJWT(jwtService))
	{
		ratingRoutes.POST("/all", ratingController.AllRatings)
		ratingRoutes.POST("/", ratingController.Insert)
	}

	productRoutes := r.Group("api/product", middleware.AuthorizeJWT(jwtService))
	{
		productRoutes.POST("/all", productController.GetAllProducts)
		productRoutes.POST("/import", productController.ImportExcel)
		productRoutes.POST("/save", productController.SaveProduct)
		productRoutes.GET("/:id", productController.GetProductByID)
		productRoutes.POST("/delete", productController.DeleteProducts)
	}
	r.Run()
}
