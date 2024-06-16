package main

import (
	"github.com/nexidian/gocliselect"
	"io/ioutil"
	"log"
	"os/exec"
)

func main() {
	files, err := ioutil.ReadDir("videos")
	menu := gocliselect.NewMenu("Pick A File")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		menu.AddItem(file.Name(), file.Name())
	}
	choice := menu.Display()
	app := "mpv"

	arg0 := "-fs"
	arg1 := "videos/"
	arg2 := choice

	cmd := exec.Command(app, arg0, arg1, arg2)
	_, err = cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

}
