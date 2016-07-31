package main

import (
	"fmt"
	"flag"

	"github.com/DLuCJ/gonant/gonant"
)

func main() {
	fmt.Println("Gonant start")
	var song gonant.Song	
	
	var bp = flag.Bool("load", false, "load song data if songdata.go has not been generated using sdgen.py")
	flag.Parse()

	if *bp {
		song = gonant.LoadSongData("gonantdata/music.inc")
	} else {
		song = gonant.SongData
		gonant.Rowlen = gonant.SONANT_ROWLEN_
		gonant.Endpattern = gonant.SONANT_ENDPATTERN_
	}
        	
	fmt.Println(song)
	fmt.Println(gonant.Rowlen)
	fmt.Println(gonant.Endpattern)

	//gonant.Init(song)

	gonant.Init()

	for ;; {
		// For ever and ever and ever and e
	}
}
