package app

// every time we import a package eg: app ,
// "go" will go into the init method of the package and initialize it the first time
import (
	"github.com/gin-gonic/gin"
)

var(
//router - pointer of gin.Engine
	router *gin.Engine
)

func init(){
// Default returns an Engine instance with the Logger and Recovery middleware already attached.
	router = gin.Default()
}

func StartApp() {

	mapUrls()
	// if err := http.ListenAndServe(":8080",nil);err != nil{
	// 	panic(err)
	// }
	if err := router.Run(":8080"); err != nil{
		panic(err)
	}
}