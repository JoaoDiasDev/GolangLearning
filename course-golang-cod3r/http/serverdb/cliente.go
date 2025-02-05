package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// Usuario
type Usuario struct {
	Id   int    `json:"id"`
	Nome string `json:"nome"`
}

// UsuarioHandler analisa o request e delega para função adequada
func UsuarioHandler(w http.ResponseWriter, r *http.Request) {
	sid := strings.TrimPrefix(r.URL.Path, "/user/")
	id, _ := strconv.Atoi(sid)

	switch {
	case r.Method == "GET" && id > 0:
		usuarioPorID(w, r, id)
	case r.Method == "GET":
		usuarioTodos(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Desculpa. Nada encontrado.")
	}
}

func usuarioPorID(w http.ResponseWriter, r *http.Request, id int) {
	fmt.Fprintf(w, "Buscando usuario %d", id)
	db, err := sql.Open("mysql", "developer:Senha@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var u Usuario
	db.QueryRow("select id,nome from usuarios where id = ?", id).Scan(&u.Id, &u.Nome)

	if u.Id == 0 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, " but the user %d does not exist :(", id)
		return
	}

	json, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))
}

func usuarioTodos(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Buscando todos os usuarios")
	db, err := sql.Open("mysql", "developer:Senha@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, _ := db.Query("select id,nome from usuarios")
	defer rows.Close()

	var usuarios []Usuario
	for rows.Next() {
		var usuario Usuario
		rows.Scan(&usuario.Id, &usuario.Nome)
		usuarios = append(usuarios, usuario)
	}

	json, _ := json.Marshal(usuarios)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))
}
