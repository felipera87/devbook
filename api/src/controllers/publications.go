package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CreatePublication creates a publication on database
func CreatePublication(w http.ResponseWriter, r *http.Request) {
	userID, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var publication models.Publication
	if err = json.Unmarshal(requestBody, &publication); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	publication.AuthorID = userID

	if err = publication.Prepare(); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewPublicationRepository(db)
	publication.ID, err = repository.Create(publication)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, publication)
}

// SearchPublications searches for all publications that shows on user feed
func SearchPublications(w http.ResponseWriter, r *http.Request) {

}

// SearchPublication search for a single an publication
func SearchPublication(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	publicationID, err := strconv.ParseUint(parameters["publicationId"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewPublicationRepository(db)
	publication, err := repository.SearchByID(publicationID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, publication)
}

// UpdatePublication updates a publication
func UpdatePublication(w http.ResponseWriter, r *http.Request) {

}

// DeletePublication deletes a publication
func DeletePublication(w http.ResponseWriter, r *http.Request) {

}