package main

import (
		"os"
		"os/exec"
		"path/filepath"
		)

func main(){
	
	wd,_ := os.Getwd()
	relativePath := "../D_sql"
	combinedPath := filepath.Join(wd,relativePath)
	absolutePath,_:= filepath.Abs(combinedPath)			// nimmt .. weg
	os.Chdir(absolutePath)
	cmd := exec.Command("bash", "-c", "psql -U lewein -d lewein -f Install-LWBadventure.sql")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
	os.Chdir(wd)
}
