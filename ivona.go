package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	ivonago "github.com/omie/ivona-go"
)

var (
	client    *ivonago.Ivona
	voicesMap = make(map[string]ivonago.Voice)
)

func initIvona(accessKey, secretKey string) {
	client = ivonago.New(accessKey, secretKey)
}

func loadVoices() error {
	options := ivonago.Voice{}
	listResp, err := client.ListVoices(options)
	if err != nil {
		return err
	}
	for _, v := range listResp.Voices {
		voicesMap[v.Name] = v
	}
	log.Println("loaded voices: ", len(listResp.Voices))
	return nil
}

func GetTTS(text, voice string) (resp []byte, err error) {
	log.Println("--- GetTTS", text, voice)

	v, ok := voicesMap[voice]
	if !ok {
		err = errors.New("Invalid voice name")
		return
	}
	options := ivonago.NewSpeechOptions(text)
	options.Voice = &v //set voice options

	r, err := client.CreateSpeech(options)
	if err != nil {
		log.Println("Error getting response from Ivona: text:", err)
		return
	}

	return r.Audio, err
}

func GetVoices(name, language, gender string) (resp []byte, err error) {
	log.Println("--- GetVoices", name, language, gender)

	options := ivonago.Voice{name, language, gender}
	r, err := client.ListVoices(options)
	if err != nil {
		log.Println("Error getting response from Ivona: text:", err)
		return
	}

	js, err := json.Marshal(r)
	if err != nil {
		return
	}

	return js, err
}

func getTTSHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("--- getTTSHandler")
	r.ParseForm()
	text := r.FormValue("text")
	voice := r.FormValue("voice")

	tts, err := GetTTS(text, voice)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "audio/mpeg3")
	w.Write(tts)
}

func getVoicesHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("--- getVoicesHandler")
	r.ParseForm()
	name := r.FormValue("name")
	language := r.FormValue("language")
	gender := r.FormValue("gender")

	voices, err := GetVoices(name, language, gender)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(voices)
}

func StartHTTPServer(host string, port string) error {
	r := mux.NewRouter()
	r.StrictSlash(true)

	r.HandleFunc("/", getTTSHandler).Methods("POST")
	r.HandleFunc("/voices", getVoicesHandler).Methods("POST")

	http.Handle("/", r)

	bind := fmt.Sprintf("%s:%s", host, port)
	log.Println("listening on:", bind)
	return http.ListenAndServe(bind, nil)
}
