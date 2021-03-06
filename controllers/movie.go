package controllers
import (
	"fmt"
	"github.com/shotu/gin-boilerplate/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)
//ArticleController ...
type MovieController struct{}

var movieModel = new(models.MovieModel)

//Create ...
func (ctrl MovieController) Create(c *gin.Context) {
	// userID := getUserID(c)
	// if userID == 0 {
	// 	c.JSON(403, gin.H{"message": "Please login first"})
	// 	c.Abort()
	// 	return
	// }

	var movie models.Movie
	db, err := gorm.Open("postgres", "dbname=gorm sslmode=disable")
	if err != nil {
		panic(err.Error())
 	}
	c.BindJSON(&movie)
	db.Create(&movie)
	c.JSON(200, movie)
}

func (ctrl MovieController) All(c *gin.Context) {
	db, err := gorm.Open("postgres", "dbname=gorm sslmode=disable")
	if err != nil {
		panic(err.Error())
	 }

	 var movies []models.Movie
	//  var movie []models.Movie
	 if err := db.Find(&movies).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
			c.JSON(200, movies)
	}
}

//All ...
// func (ctrl ArticleController) All(c *gin.Context) {
// 	userID := getUserID(c)

// 	if userID == 0 {
// 		c.JSON(403, gin.H{"message": "Please login first"})
// 		c.Abort()
// 		return
// 	}

// 	data, err := articleModel.All(userID)

// 	if err != nil {
// 		c.JSON(406, gin.H{"Message": "Could not get the articles", "error": err.Error()})
// 		c.Abort()
// 		return
// 	}

// 	c.JSON(200, gin.H{"data": data})
// }

//One ...
// func (ctrl ArticleController) One(c *gin.Context) {
// 	userID := getUserID(c)

// 	if userID == 0 {
// 		c.JSON(403, gin.H{"message": "Please login first"})
// 		c.Abort()
// 		return
// 	}

// 	id := c.Param("id")

// 	if id, err := strconv.ParseInt(id, 10, 64); err == nil {

// 		data, err := articleModel.One(userID, id)
// 		if err != nil {
// 			c.JSON(404, gin.H{"Message": "Article not found", "error": err.Error()})
// 			c.Abort()
// 			return
// 		}
// 		c.JSON(200, gin.H{"data": data})
// 	} else {
// 		c.JSON(404, gin.H{"Message": "Invalid parameter"})
// 	}
// }

//Update ...
// func (ctrl ArticleController) Update(c *gin.Context) {
// 	userID := getUserID(c)

// 	if userID == 0 {
// 		c.JSON(403, gin.H{"message": "Please login first"})
// 		c.Abort()
// 		return
// 	}

// 	id := c.Param("id")
// 	if id, err := strconv.ParseInt(id, 10, 64); err == nil {

// 		var articleForm forms.ArticleForm

// 		if c.BindJSON(&articleForm) != nil {
// 			c.JSON(406, gin.H{"message": "Invalid parameters", "form": articleForm})
// 			c.Abort()
// 			return
// 		}

// 		err := articleModel.Update(userID, id, articleForm)
// 		if err != nil {
// 			c.JSON(406, gin.H{"Message": "Article could not be updated", "error": err.Error()})
// 			c.Abort()
// 			return
// 		}
// 		c.JSON(200, gin.H{"message": "Article updated"})
// 	} else {
// 		c.JSON(404, gin.H{"Message": "Invalid parameter", "error": err.Error()})
// 	}
// }

//Delete ...
// func (ctrl ArticleController) Delete(c *gin.Context) {
// 	userID := getUserID(c)

// 	if userID == 0 {
// 		c.JSON(403, gin.H{"message": "Please login first"})
// 		c.Abort()
// 		return
// 	}

// 	id := c.Param("id")
// 	if id, err := strconv.ParseInt(id, 10, 64); err == nil {

// 		err := articleModel.Delete(userID, id)
// 		if err != nil {
// 			c.JSON(406, gin.H{"Message": "Article could not be deleted", "error": err.Error()})
// 			c.Abort()
// 			return
// 		}
// 		c.JSON(200, gin.H{"message": "Article deleted"})
// 	} else {
// 		c.JSON(404, gin.H{"Message": "Invalid parameter"})
// 	}
// }
