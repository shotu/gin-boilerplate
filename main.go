package main
import (
	"fmt"
	// "./greeting"
	// "net/http"
	// "runtime"
	"github.com/gin-gonic/gin"
	"github.com/shotu/gin-boilerplate/controllers"
	// "github.com/shotu/gin-boilerplate/db"
	// "github.com/gin-gonic/contrib/sessions"
	"github.com/shotu/gin-boilerplate/models"
	_"github.com/lib/pq"
	"github.com/jinzhu/gorm"
)



//CORSMiddleware ...
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}


func main() {
	
	db, err := gorm.Open("postgres", "dbname=gorm sslmode=disable")
	// // db, err := gorm.Open("postgres", "user=gorm password=gorm dbname=gorm sslmode=disable")
	if err != nil {
		 panic(err.Error())
	}
	defer db.Close()
	println("Connection to databse established")
	db.AutoMigrate(&models.Movie{})
	db.AutoMigrate(&models.User{})
	
	println("Done")
	
	r := gin.Default()

	// store, _ := sessions.NewRedisStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	// r.Use(sessions.Sessions("gin-boilerplate-session", store))

	r.Use(CORSMiddleware())

	// db.Init()

	v1 := r.Group("/v1")
	// {
	// 	/*** START USER ***/
	// 	user := new(controllers.UserController)

	// 	v1.POST("/user/signin", user.Signin)
	// user := new(controllers.UserController)
	// v1.POST("/user/signup", user.Signup)
	// 	v1.GET("/user/signout", user.Signout)
	// 	/*** START Article ***/
	movie := new(controllers.MovieController)
	v1.POST("/movie", movie.Create)
	v1.GET("/movie", movie.All)
	// 	v1.GET("/articles", article.All)
	// 	v1.GET("/article/:id", article.One)
	// 	v1.PUT("/article/:id", article.Update)
	// 	v1.DELETE("/article/:id", article.Delete)
	// }

	// r.LoadHTMLGlob("./public/html/*")

	// r.Static("/public", "./public")

	// r.GET("/", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "index.html", gin.H{
	// 		"ginBoilerplateVersion": "v0.03",
	// 		"goVersion":             runtime.Version(),
	// 	})
	// })

	// r.NoRoute(func(c *gin.Context) {
	// 	c.HTML(404, "404.html", gin.H{})
	// })

	r.Run(":9000")
}

type User struct {
	ID uint
	UserName string
	FirstName string
	LastName string
}

var users []User = []User{
	User{UserName:"manish", FirstName: "Manish", LastName: "Atri"},
	User{UserName:"anish", FirstName: "Anish", LastName: "Atri"},
	User{UserName:"lokesh", FirstName: "Lokesh", LastName: "Sharma"},
}