package gonant

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const (
	AUXDATA = iota
	SONG_DATA_OSC1_OCT = iota
	SONG_DATA_OSC1_DET = iota
	SONG_DATA_OSC1_DETUNE = iota
	SONG_DATA_OSC1_XENV = iota
	SONG_DATA_OSC1_VOL = iota
	SONG_DATA_OSC1_WAVEFORM = iota
	SONG_DATA_OSC2_OCT = iota
	SONG_DATA_OSC2_DET = iota
	SONG_DATA_OSC2_DETUNE = iota
	SONG_DATA_OSC2_XENV = iota
	SONG_DATA_OSC2_VOL = iota
	SONG_DATA_OSC2_WAVEFORM = iota
	SONG_DATA_NOISE_FADER = iota
	SONG_DATA_ENV_ATTACK = iota
	SONG_DATA_ENV_SUSTAIN = iota
	SONG_DATA_ENV_RELEASE = iota
	SONG_DATA_ENV_MASTER = iota
	SONG_DATA_FX_FILTER = iota
	SONG_DATA_FX_FREQ = iota
	SONG_DATA_FX_RESONANCE = iota
	SONG_DATA_FX_DELAY_TIME = iota
	SONG_DATA_FX_DELAY_AMT = iota
	SONG_DATA_FX_PAN_FREQ = iota
	SONG_DATA_FX_PAN_AMT = iota
	SONG_DATA_LFO_OSC1_FREQ = iota
	SONG_DATA_LFO_FX_FREQ = iota
	SONG_DATA_LFO_FREQ = iota
	SONG_DATA_LFO_AMT = iota
	SONG_DATA_LFO_WAVEFORM = iota
	SONG_DATA_PATTERNS = iota
	SONG_DATA_COLUMNS = iota
)

func fillStructures() {
	fmt.Println(Column{})
	fmt.Println(Song{})
	fmt.Println(Instrument{})
}

func parseSonantOutput(songdata string) {
	songmap := make(map[int]string)
	var symbols []string = strings.Split(songdata, "song_data_")
	
	for i := 0; i < 32; i++ {
		songmap[i] = symbols[i]
	}

	fmt.Println(songmap)
	fmt.Printf("Endpattern is: %d\n", Endpattern)
}

//func LoadSongData(filename string) Song{
func LoadSongData(filename string) {
	//var song Song

	//TODO: handle errors
	songbytes,_  := ioutil.ReadFile(filename)
	songstr := string(songbytes)

	parseSonantOutput(songstr)
	
	//fillStructures()
	//return song
}


