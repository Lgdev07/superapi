package tests

import (
	"log"
	"testing"

	"github.com/Lgdev07/superapi/api/models"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestSaveSuper(t *testing.T) {
	err := RefreshSuperTable()
	if err != nil {
		log.Fatal(err)
	}

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
		log.Fatal(err)
	}

	assert.Equal(t, createdSuper.Name, "Super Test")
	assert.Equal(t, createdSuper.FullName, "John Doe")
	assert.Equal(t, createdSuper.Intelligence, 80)
	assert.Equal(t, createdSuper.Power, 50)
	assert.Equal(t, createdSuper.Occupation, "Student")
	assert.Equal(t, createdSuper.Image, "https://www.superherodb.com/pictures2/portraits/10/100/1356.jpg")
	assert.Equal(t, createdSuper.Alignment, "good")
	assert.Equal(t, createdSuper.Groups, "Group Test, Group Develop, Group Golang")
	assert.Equal(t, createdSuper.NumberOfParents, 5)
}

func TestDeleteSuperByID(t *testing.T) {
	err := RefreshSuperTable()
	if err != nil {
		log.Fatal(err)
	}

	super, _ := seedOneSuper()

	supers, err := models.FindSupers(app.DB, make(map[string]string))
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, len(*supers), 1)
	assert.Equal(t, (*supers)[0].Name, "Super Test")
	assert.Equal(t, (*supers)[0].Alignment, "good")

	if err := models.DeleteSuperByID(app.DB, super.ID.String()); err != nil {
		log.Fatal(err)
	}

	supers, err = models.FindSupers(app.DB, make(map[string]string))
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, len(*supers), 0)
}

func TestGetSuperByID(t *testing.T) {
	err := RefreshSuperTable()
	if err != nil {
		log.Fatal(err)
	}

	super, _ := seedOneSuper()

	superFound, err := models.GetSuperByID(app.DB, super.ID.String())
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, superFound.Name, "Super Test")
	assert.Equal(t, superFound.Alignment, "good")

	randomUUID := uuid.NewV4()

	superFound, _ = models.GetSuperByID(app.DB, randomUUID.String())

	assert.Equal(t, superFound.Name, "")
	assert.Equal(t, superFound.Alignment, "")

}

func TestFindSupers(t *testing.T) {
	err := RefreshSuperTable()
	if err != nil {
		log.Fatal(err)
	}

	supers, err := models.FindSupers(app.DB, make(map[string]string))
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, len(*supers), 0)

	_, _ = seedOneSuper()

	supers, err = models.FindSupers(app.DB, make(map[string]string))
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, len(*supers), 1)
	assert.Equal(t, (*supers)[0].Name, "Super Test")
	assert.Equal(t, (*supers)[0].Alignment, "good")

}
