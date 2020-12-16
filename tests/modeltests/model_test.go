package modeltests

import (
	"fmt"
	"log"
	"os"
	"testing"

	"backend-go/api/controllers"
	"backend-go/api/models"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var server = controllers.Server{}
var userInstance = models.User{}
var bookInstance = models.Book{}

func TestMain(m *testing.M) {
	var err error
	err = godotenv.Load(os.ExpandEnv("../../.env"))
	if err != nil {
		log.Fatalf("Error getting env %v\n", err)
	}
	Database()

	log.Printf("Before calling m.Run() !!!")
	ret := m.Run()
	log.Printf("After calling m.Run() !!!")
	//os.Exit(m.Run())
	os.Exit(ret)
}

func Database() {

	var err error

	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", os.Getenv("TestDbHost"), os.Getenv("TestDbPort"), os.Getenv("TestDbUser"), os.Getenv("TestDbName"), os.Getenv("TestDbPassword"))
	server.DB, err = gorm.Open("postgres", DBURL)
	if err != nil {
		fmt.Printf("Cannot connect to %s database\n", "postgres")
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the %s database\n", "postgres")
	}

}
