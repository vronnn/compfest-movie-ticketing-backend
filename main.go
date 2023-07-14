package main

import (
	"log"
	"os"

	"github.com/Caknoooo/golang-clean_template/config"
	"github.com/Caknoooo/golang-clean_template/controller"
	"github.com/Caknoooo/golang-clean_template/middleware"
	"github.com/Caknoooo/golang-clean_template/migration"
	"github.com/Caknoooo/golang-clean_template/repository"
	"github.com/Caknoooo/golang-clean_template/routes"
	"github.com/Caknoooo/golang-clean_template/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {	
	var (
		db             *gorm.DB                  = config.SetUpDatabaseConnection()
		jwtService     services.JWTService       = services.NewJWTService()
		userRepository repository.UserRepository = repository.NewUserRepository(db)
		userService    services.UserService      = services.NewUserService(userRepository)
		userController controller.UserController = controller.NewUserController(userService, jwtService)
		seederRepository repository.SeederRepository = repository.NewSeederRepository(db)
		seederService    services.SeederService      = services.NewSeederService(seederRepository)
		seederController controller.SeederController = controller.NewSeederController(seederService)
		topupRepository repository.TopupRepository = repository.NewTopupRepository(db)
		topupService    services.TopupService      = services.NewTopupService(topupRepository)
		topupController controller.TopupController = controller.NewTopupController(topupService, jwtService)
		movierepository repository.MovieRepository = repository.NewMoviesRepository(db)
		movieService    services.MovieService      = services.NewMovieService(movierepository)
		movieController controller.MovieController = controller.NewMovieController(movieService, jwtService, userService)
		ticketRepository repository.TicketRepository = repository.NewTicketRepository(db)
		ticketService services.TicketService = services.NewTicketService(ticketRepository)
		ticketController controller.TicketController = controller.NewTicketController(ticketService, jwtService)
		withDrawalRepository repository.WithDrawalRepository = repository.NewWithDrawalRepository(db)
		withDrawalService services.WithDrawalService = services.NewWithDrawalService(withDrawalRepository)
		withDrawalController controller.WithDrawalController = controller.NewWithDrawalController(withDrawalService, jwtService)
	)

	server := gin.Default()
	server.Use(middleware.CORSMiddleware())
	routes.Router(server, userController, seederController, topupController, movieController, ticketController, withDrawalController, jwtService)

	if err := migration.Seeder(db); err != nil {
		log.Fatalf("error seeding database: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}
	server.Run(":" + port)
}
