package app

// all the routes in the application are defined here.
import (
	"github.com/shraddhagajul/golang-microservices/mvc/controllers"
)

func mapUrls() {
	// http.HandleFunc("/users", controllers.GetUser)
	router.GET("/users/:user_id",controllers.GetUser)

}