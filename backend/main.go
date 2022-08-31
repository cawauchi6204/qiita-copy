package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func dbConnect() *gorm.DB {
	dsn := "user:password@tcp(db:3306)/test_db"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}

type User struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	Name           string    `json:"name"`
	Password       string    `json:"password"`
	Email          string    `json:"email"`
	Nickname       string    `json:"nickname"`
	HpUrl          string    `json:"hp_url"`
	Location       string    `json:"location"`
	GithubAccount  string    `json:"github_account_id"`
	OrganizationId int       `json:"organization_id"`
	IsDeleted      int       `json:"is_deleted"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})
	db := dbConnect()
	r.GET("/users", func(c *gin.Context) {
		users := []User{}
		if err := db.Find(&users).Error; err != nil {
			fmt.Println(err)
		}
		c.JSON(http.StatusOK, users)
	})
	deployPort := os.Getenv("PORT")
	if deployPort == "" {
		deployPort = "8080"
	}
	r.Run(":" + deployPort)
}
