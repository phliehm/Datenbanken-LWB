package main

import "gfx"
import "fmt"

func main () {
	gfx.Fenster(640,480)
	for {
		taste,gedrueckt,tiefe:= gfx.TastaturLesen1()
		fmt.Println (taste, gedrueckt,tiefe)
		if taste == 27 {
			 break
		 }
	}
}
