package gonant

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var sdsymbols = [32]string{"auxdata", "osc1_oct", "osc1_det", "osc1_detune", "osc1_xenv", "osc1_vol",
	"osc1_waveform", "osc2_oct", "osc2_det", "osc2_detune", "osc2_xenv", "osc2_vol", "osc2_waveform",
	"noise_fader", "env_attack", "env_sustain", "env_release", "env_master", "fx_filter", "fx_freq",
	"fx_resonance", "fx_delay_time", "fx_delay_amt", "fx_pan_freq", "fx_pan_amt", "lfo_osc1_freq", 
	"lfo_fx_freq", "lfo_freq", "lfo_amt", "lfo_waveform", "patterns", "columns",}

func fillStructures() {
	fmt.Println(Column{})
	fmt.Println(Song{})
	fmt.Println(Instrument{})
}

func parseSonantOutput(songdata string) {
	songmap := make(map[string]string)
	var symbols []string = strings.Split(songdata, "song_data_")
	
	for i := 0; i < 32; i++ {
		songmap[sdsymbols[i]] = symbols[i]
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


