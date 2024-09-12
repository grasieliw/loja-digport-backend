package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/grasieliw/loja-digport-backend/model"
)

func CriaUsuariosHandler(w http.ResponseWriter, r *http.Request) {
	var usuario model.Usuario
	json.NewDecoder(r.Body).Decode(&usuario)

	error := model.CriaUsuario(usuario)

	if error != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(error.Error())
	}

	w.WriteHeader(http.StatusCreated)
}
