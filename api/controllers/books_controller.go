package controllers

import (
	"backend-go/api/models"
	"backend-go/api/responses"
	"backend-go/api/utils/formaterror"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (server *Server) CreateBook(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	book := models.Book{}
	err = json.Unmarshal(body, &book)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	book.Prepare()
	err = book.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	bookCreated, err := book.SaveBook(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, bookCreated.ID))
	responses.JSON(w, http.StatusCreated, bookCreated)
}

func (server *Server) LendBook(w http.ResponseWriter, r *http.Request) {

	// Read the data
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Start processing the request data
	bookUpdate := models.Book{}
	err = json.Unmarshal(body, &bookUpdate)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	fmt.Print(bookUpdate)

	// Check if the book exist
	book := models.Book{}
	err = server.DB.Debug().Model(models.Book{}).Where("id = ?", body[1]).Take(&book).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Post not found"))
		return
	}

	// If a user attempt to lend a book not belonging to him
	if 2 != book.LoggedUserID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	bookUpdate.Prepare()
	err = bookUpdate.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	bookUpdate.ID = book.ID //this is important to tell the model the book id to update, the other update field are set above

	bookUpdated, err := bookUpdate.UpdateABook(server.DB)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, bookUpdated)
}
