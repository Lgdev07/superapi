package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/Lgdev07/superapi/api/models"
	"github.com/Lgdev07/superapi/api/utils"
)

func (s *Server) CreateSuper(w http.ResponseWriter, r *http.Request) {
	apiToken := os.Getenv("SUPERHERO_API_TOKEN")

	var responseInterface map[string]interface{}

	err := json.NewDecoder(r.Body).Decode(&responseInterface)
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}

	url := fmt.Sprintf("https://superheroapi.com/api/%s/search/%s", apiToken, responseInterface["name"])

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	var responseInterface2 map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&responseInterface2)
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}

	results := responseInterface2["results"].([]interface{})

	for i := 0; i < len(results); i++ {
		name := results[i].(map[string]interface{})["name"]

		if name != responseInterface["name"] {
			continue
		}

		fullName := results[i].(map[string]interface{})["biography"].(map[string]interface{})["full-name"]
		intelligence := results[i].(map[string]interface{})["powerstats"].(map[string]interface{})["intelligence"]
		power := results[i].(map[string]interface{})["powerstats"].(map[string]interface{})["power"]
		occupation := results[i].(map[string]interface{})["work"].(map[string]interface{})["occupation"]
		image := results[i].(map[string]interface{})["image"].(map[string]interface{})["url"]
		alignment := results[i].(map[string]interface{})["biography"].(map[string]interface{})["alignment"]
		groups := results[i].(map[string]interface{})["connections"].(map[string]interface{})["group-affiliation"]

		intelligenceInt, _ := strconv.Atoi(intelligence.(string))
		powerInt, _ := strconv.Atoi(power.(string))

		super := &models.Super{
			Name:         name.(string),
			FullName:     fullName.(string),
			Intelligence: intelligenceInt,
			Power:        powerInt,
			Occupation:   occupation.(string),
			Image:        image.(string),
			Alignment:    alignment.(string),
			Groups:       groups.(string),
		}

		// Name            string    `gorm:"size:100;not null" json:"name"`
		// FullName        string    `gorm:"not null" json:"full_name"`
		// Intelligence    string       `gorm:"not null" json:"intelligence"`
		// Power           string       `gorm:"not null" json:"power"`
		// Occupation      string    `gorm:"not null" json:"occupation"`
		// Image           string    `gorm:"not null" json:"image"`
		// alignment       string    `gorm:"not null" json:"alignment"`
		// Groups          string    `json:"groups"`
		// NumberOfParents int       `json:"number_of_parents"`

		createdSuper, err := super.CreateSuper(s.DB)
		if err != nil {
			log.Fatal(err)
		}

		utils.JSON(w, http.StatusOK, createdSuper)
		return
	}

	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Println(string(body))

	// req, err := http.NewRequest("GET", url, nil)
	// if err != nil {
	// 	log.Print(err)
	// 	os.Exit(1)
	// }

	// q := req.URL.Query()
	// req.URL.RawQuery = q.Encode()

	// fmt.Println(req.URL.String())

	var responseFinal = map[string]interface{}{"status": "success", "message": "Author successfully created"}
	utils.JSON(w, http.StatusOK, responseFinal)
	return
}

func (s *Server) ListSupers(w http.ResponseWriter, r *http.Request) {
	return
}

func (s *Server) DeleteSuper(w http.ResponseWriter, r *http.Request) {
	return
}

func (s *Server) Index(w http.ResponseWriter, r *http.Request) {
	var resp = map[string]interface{}{"status": "success", "message": "Author successfully created"}
	utils.JSON(w, http.StatusOK, resp)
	return
}

// Cadastrar um Super/Vilão
// Listar todos os Super's cadastrados
// Listar apenas os Super Heróis
// Listar apenas os Super Vilões
// Buscar por nome
// Buscar por 'uuid'
// Remover o Super
