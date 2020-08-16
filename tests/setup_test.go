package tests

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/Lgdev07/superapi/api/controllers"
	"github.com/Lgdev07/superapi/api/models"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var app = controllers.Server{}
var superInstance = models.Super{}

func TestMain(m *testing.M) {
	//Since we add our .env in .gitignore, Circle CI cannot see it, so see the else statement
	if _, err := os.Stat("./../.env"); !os.IsNotExist(err) {
		var err error
		err = godotenv.Load(os.ExpandEnv("./../.env"))
		if err != nil {
			log.Fatalf("Error getting env %v\n", err)
		}
		Database()
	} else {
		CIBuild()
	}
	os.Exit(m.Run())
}

//When using CircleCI
func CIBuild() {
	var err error
	DBURL := fmt.Sprintf(`host=localhost port=5432 user=lgdev07 
	dbname=superapi_test sslmode=disable password=docker`)
	app.DB, err = gorm.Open("postgres", DBURL)
	if err != nil {
		fmt.Printf("Cannot connect to postgres database\n")
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the postgres database\n")
	}
}

func Database() {

	var err error
	DBURL := fmt.Sprintf(`host=%s port=%s user=%s dbname=%s sslmode=disable 
	password=%s`, os.Getenv("TEST_DB_HOST"), os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"), os.Getenv("TEST_DB_NAME"),
		os.Getenv("DB_PASSWORD"))

	app.DB, err = gorm.Open("postgres", DBURL)
	if err != nil {
		fmt.Printf("Cannot connect to postgres database\n")
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the postgres database\n")
	}

}

func RefreshSuperTable() error {
	err := app.DB.DropTableIfExists(&models.Super{}).Error
	if err != nil {
		return err
	}
	err = app.DB.AutoMigrate(&models.Super{}).Error
	if err != nil {
		return err
	}
	return nil
}

func seedOneSuper() (models.Super, error) {
	super := models.Super{
		Name:            "Super Test",
		FullName:        "John Doe",
		Intelligence:    80,
		Power:           50,
		Occupation:      "Student",
		Image:           "https://www.superherodb.com/pictures2/portraits/10/100/1356.jpg",
		Alignment:       "good",
		Groups:          "Group Test, Group Develop, Group Golang",
		NumberOfParents: 5,
	}

	createdSuper, err := super.Save(app.DB)
	if err != nil {
		return models.Super{}, err
	}

	return *createdSuper, nil

}
