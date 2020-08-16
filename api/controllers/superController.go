package controllers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Lgdev07/superapi/api/models"
	"github.com/Lgdev07/superapi/api/services"
	"github.com/Lgdev07/superapi/api/utils"
	"github.com/gorilla/mux"
)

// ListSupers returns a formated response with the supers list
func (s *Server) ListSupers(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	params := map[string]string{
		"name":      strings.Join(query["name"], ""),
		"alignment": strings.Join(query["alignment"], ""),
	}

	supers, err := models.FindSupers(s.DB, params)
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}

	utils.JSON(w, http.StatusOK, supers)
	return
}

// DeleteSuper deletes a super and returns the deleted id
func (s *Server) DeleteSuper(w http.ResponseWriter, r *http.Request) {
	resp := map[string]interface{}{"status": "success"}
	params := mux.Vars(r)

	err := models.DeleteSuperByID(s.DB, params["id"])
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}

	resp["message"] = "Deleted Super ID " + params["id"]
	utils.JSON(w, http.StatusOK, resp)
	return

}

// ShowSuper returns a formated response with the super
func (s *Server) ShowSuper(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	super, err := models.GetSuperByID(s.DB, params["id"])
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}

	utils.JSON(w, http.StatusOK, super)
	return

}

// CreateSuper creates a new super with the given name in body
func (s *Server) CreateSuper(w http.ResponseWriter, r *http.Request) {
	var responseInterface map[string]string

	err := json.NewDecoder(r.Body).Decode(&responseInterface)
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}

	nameTitle := strings.Title(responseInterface["name"])
	supers, err := models.FindSupers(s.DB, map[string]string{"name": nameTitle})
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	if len(*supers) != 0 {
		resp := map[string]string{"error": "There is Already a Super With That Same Name Registered"}
		utils.JSON(w, http.StatusBadRequest, resp)
		return
	}

	super, status, err := services.ApiGetSuperByName(nameTitle)
	if err != nil {
		utils.ERROR(w, status, err)
		return
	}

	if super.Name == "" {
		resp := map[string]string{"error": "No Super found with this name"}
		utils.JSON(w, http.StatusBadRequest, resp)
		return
	}

	createdSuper, err := super.Save(s.DB)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	utils.JSON(w, http.StatusCreated, createdSuper)
	return
}
