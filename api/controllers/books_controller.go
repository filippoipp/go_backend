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

	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()

	t := struct {
		LoggedUserID *uint64 `json:"logged_user_id"`
		BookID *uint64 `json:"book_id"`
		ToUserID *uint64 `json:"to_user_id"`
	}{}

	err := d.Decode(&t)
	if err != nil {
		// bad JSON or unrecognized json field
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if the user exist
	user := models.User{}
	err = server.DB.Debug().Model(models.User{}).Where("id = ?", *t.LoggedUserID).Take(&user).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("User not found"))
		return
	}

	// Check if the to user exist
	user = models.User{}
	err = server.DB.Debug().Model(models.User{}).Where("id = ?", *t.ToUserID).Take(&user).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("To User not found"))
		return
	}

	// Check if the book exist
	book := models.Book{}
	err = server.DB.Debug().Model(models.Book{}).Where("id = ?", *t.BookID).Take(&book).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Book not found"))
		return
	}

	// If a user attempt to lend a book not belonging to him
	if *t.LoggedUserID != book.LoggedUserID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	// Book owner can't land book for himself
	if *t.ToUserID == book.LoggedUserID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	//Check if book's already lended
	if book.ToUserID != 0 {
		responses.ERROR(w, http.StatusBadRequest, errors.New("The book's already lended"))
		return
	}


	bookUpdated, err := book.UpdateABook(server.DB, *t.ToUserID, *t.BookID)

	fmt.Println(bookUpdated)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, bookUpdated)
}

func (server *Server) ReturnBook(w http.ResponseWriter, r *http.Request) {

	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()

	t := struct {
		LoggedUserID *uint64 `json:"logged_user_id"`
		BookID *uint64 `json:"book_id"`
	}{}

	err := d.Decode(&t)
	if err != nil {
		// bad JSON or unrecognized json field
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if the book exist
	book := models.Book{}
	err = server.DB.Debug().Model(models.Book{}).Where("id = ?", *t.BookID).Take(&book).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Book not found"))
		return
	}

	// If a user attempt to lend a book not belonging to him
	if *t.LoggedUserID != book.LoggedUserID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	//Check if book's already lended
	if book.ToUserID == 0 {
		responses.ERROR(w, http.StatusBadRequest, errors.New("The book's isn't lended"))
		return
	}

	bookUpdated, err := book.ReturnABook(server.DB, *t.BookID)

	fmt.Println(bookUpdated)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, bookUpdated)
}
