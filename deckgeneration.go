package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func generateDeck() {
	files, err := ioutil.ReadDir("./images")
	if err != nil {
		log.Fatal(err)
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(files), func(i, j int) { files[i], files[j] = files[j], files[i] })

	var cardnames []string

	for _, file := range files {
		cardnames = append(cardnames, "./images/"+file.Name())
	}

	fmt.Println(cardnames[:6])

	cmd := exec.Command("ls")
	if runtime.GOOS == "windows" {
		cmd = exec.Command("magick.exe", "montage", cardnames[0], cardnames[1], cardnames[2], cardnames[3], cardnames[4], cardnames[5], "-tile", "6x1", "-geometry", "127x205+2+2", "-shadow", "montage.jpg")
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
}
