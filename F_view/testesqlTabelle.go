package main


import (
		"fmt"
		"SQL"
		"../Klassen/sqlTabelle"
		)

func main() {
	// Verbindungsaufbau
	var conn SQL.Verbindung
	conn = SQL.PgSQL ("user=lewein dbname=lewein")
	defer conn.Beenden ()
	fmt.Println("Verbindung hergestellt.\n")
	
	var anfrage string
	
	anfrage = "SELECT * FROM veranstaltungen;"
	sT := sqlTabelle.New(conn,anfrage)
	fmt.Println(sT.GibTabelle())
	
}
