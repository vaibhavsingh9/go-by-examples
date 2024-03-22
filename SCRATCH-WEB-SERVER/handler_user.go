package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"scratch-web-server/internal/database"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct{
		Name string 'name'
	}
	decoder:=json.NewDecoder(r.Body)
	params:=parameters{}
	err:=decoder.Decode(&params)
	if err != nil{
		respondWithError(w, 400, fmt.Sprintf("error parsing json: ", err))
		return
	}
	apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{

	})

	respondWithJSON(w, 200, struct{}{})
}
