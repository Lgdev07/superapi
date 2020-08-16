package tests

import (
	"log"
	"testing"

	"github.com/Lgdev07/superapi/api/services"
	"github.com/stretchr/testify/assert"
)

func TestApiGetSuperByName(t *testing.T) {
	err := RefreshSuperTable()
	if err != nil {
		log.Fatal(err)
	}

	super, _, err := services.ApiGetSuperByName("Batman")
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, super.Name, "Batman")
	assert.Equal(t, super.Alignment, "good")

	super, _, err = services.ApiGetSuperByName("Not a Super")
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, super.Name, "")
	assert.Equal(t, super.Alignment, "")

}
