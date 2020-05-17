package handlers

import (
	"encoding/json"
	"github.com/bauidch/hyrt-api/dao"
	"github.com/bauidch/hyrt-api/models"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func Index(w http.ResponseWriter, r *http.Request) {
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func Ping(w http.ResponseWriter, r *http.Request) {
	respondWithJson(w, http.StatusOK, "PONG")
}

func GetAllSeries(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	payload := dao.GetAllSeries()
	json.NewEncoder(w).Encode(payload)
}

func CreateSerie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var task models.Series
	_ = json.NewDecoder(r.Body).Decode(&task)
	log.Println(task, r.Body)
	dao.InsertOneSerie(task)
	respondWithJson(w, http.StatusCreated, map[string]string{"message": "Serie created Successfully"})
}

func GetOneSeries(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	payload, e := dao.GetOneSeries(params["id"])
	if e != nil {
		respondWithJson(w, http.StatusNotFound, map[string]string{"message": "Serie not found"})
		log.Print(e)
	}
	json.NewEncoder(w).Encode(payload)
}

func GetAllSeed(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	payload := dao.GetAllSeed()
	json.NewEncoder(w).Encode(payload)
}

func CreateSeed(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var task models.Seed
	_ = json.NewDecoder(r.Body).Decode(&task)
	log.Println(task, r.Body)
	dao.InsertOneSeed(task)
	respondWithJson(w, http.StatusCreated, map[string]string{"message": "Seed created Successfully"})
}


func GetAllSeedJornal(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	payload := dao.GetAllSeedJournal()
	json.NewEncoder(w).Encode(payload)
}
