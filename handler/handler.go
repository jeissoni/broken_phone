package handler

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"telefono_roto/chat"
	"text/template"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type ErrorData struct {
	ErrorMessage string
}

func Index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "method not found", http.StatusNotFound)
		return
	}

	//tmpl, err := template.New("index").ParseFiles("templates/index.html")

	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	//leer el contenido del archivo hmtl
	indexHTML, _ := os.ReadFile("templates/registerUser.html")

	err := tmpl.Execute(w, string(indexHTML))

	if err != nil {
		panic(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)

	//http.ServeFile(w, r, "templates/index.html")
}

func UserRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "method not found", http.StatusNotFound)
		return
	}

	// Parsear los datos del formulario
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error al parsear los datos del formulario", http.StatusInternalServerError)
		return
	}

	//validar que no se envie un campo vacio
	if r.Form.Get("username") == "" {
		http.Error(w, "El campo username es requerido", http.StatusBadRequest)
		return
	}

	//normalizar el campo username
	username := strings.ToLower(r.Form.Get("username"))

	//crear el usuario
	u, err := chat.CreateUser(username)

	if err != nil {
		//tmpl := template.Must(template.ParseFiles("templates/error.html"))

		/*
			data := ErrorData{
				ErrorMessage: err.Error(),
			}
		*/
		http.Error(w, err.Error(), http.StatusBadRequest)

		//tmpl.Execute(w, data)
		return

	}

	//get room empty
	roomId := chat.GetRoomEmpty()

	if roomId == 0 {
		roomId = chat.CreateRoom("Room")
	}

	//add user to room
	chat.AddUserToRoom(roomId, &u)

	fmt.Println("RoomId: ", roomId)
	fmt.Println(chat.RoomsMap[roomId].Users[0])
	http.ServeFile(w, r, "templates/register.html")

}
