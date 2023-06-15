package main

import (
		"gfx"
		"fmt"
		"./textboxTabelle"
		"./sqlTabelle"
		"SQL"
		)
		
func main() {
	gfx.Fenster(1200,700)
	gfx.SetzeFont("Schriftarten/terminus-font/Terminus-Bold.ttf",20)
	// Verbindungsaufbau
	var conn SQL.Verbindung
	conn = SQL.PgSQL ("user=lewein dbname=lewein")
	defer conn.Beenden ()
	fmt.Println("Verbindung hergestellt.\n")
	
	var anfrage string
	
	anfrage = "SELECT * FROM spielstaende ORDER BY punkte DESC LIMIT 10;"
	sT := sqlTabelle.New(conn,anfrage)
	fmt.Println(sT.GibTabelle())
	
	// Textbox Tabelle
	tbT := textboxTabelle.New(sT.GibTabelle(),sT.GibKopf(),50,50)
	tbT.SetzeFarbeTabelle(255,0,0)
	tbT.SetzeZeilenAbstand(20)
	tbT.SetzeSchriftgrößeTabelle(20)
	tbT.SetzeSchriftgrößeKopf(30)
	tbT.SetzeFarbeKopf(0,0,255)
	tbT.SetzeFontKopf("Schriftarten/terminus-font/TerminusTTF-Bold-4.49.2.ttf")
	tbT.SetzeFontTabelle("Schriftarten/terminus-font/TerminusTTF-Bold-4.49.2.ttf")
	tbT.Zeichne()
	gfx.TastaturLesen1()
	
}
