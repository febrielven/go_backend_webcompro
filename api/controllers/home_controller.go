package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/febrielven/go_backend_webcompro/api/models"
	"github.com/febrielven/go_backend_webcompro/api/repositories"
	"github.com/febrielven/go_backend_webcompro/api/responses"
	"github.com/febrielven/go_backend_webcompro/api/utils/formaterror"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Welcome home api")
	responses.JSON(w, http.StatusOK, "Welcom home api")
}

func (server *Server) GetAll(w http.ResponseWriter, r *http.Request) {
	// user := models.User{}
	centrixStepRepository := repositories.NewStepRepository()
	steps, err := centrixStepRepository.FindAll()

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, steps)

}

func (server *Server) Create(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	step := models.CentrixStep{}

	err = json.Unmarshal(body, &step)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	fmt.Printf("data Name1: %+v\n", step.Name)
	centrixStepRepository := repositories.NewStepRepository()

	stepCreate, err := centrixStepRepository.Save(&step)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, stepCreate.ID))

	responses.JSON(w, http.StatusCreated, stepCreate)

}
