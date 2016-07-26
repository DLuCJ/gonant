package gonant

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

const NUM_INSTS = 8

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

func fillStructures(auxdata string, parms map[int][NUM_INSTS]int64, fxfreq [NUM_INSTS]float64, patterns string, columns string) { //Song {
	//var song Song
	fmt.Println(auxdata)
	fmt.Println()
	fmt.Println(parms)
	fmt.Println()
	fmt.Println(fxfreq)
	fmt.Println()
	fmt.Println(patterns)
	fmt.Println()
	fmt.Println(columns)


	fmt.Println("*****")
	fmt.Println(parms[SONG_DATA_OSC2_VOL])
}

func splitSonantOutput(songdata string) (string, map[int][NUM_INSTS]int64, [NUM_INSTS]float64, string, string){
	var err error

	var symbols []string = strings.Split(songdata, "song_data_")

	var sdarr [32]string	
	for i, elt := range symbols {
		sdarr[i] = elt
	}

	var fx_filter_params [NUM_INSTS]float64
	parm_map := make(map[int][NUM_INSTS]int64)

	j := 1

	for i := 0; i < 29; i++ {
		
		var digarr [NUM_INSTS]int64
		
		tokens := strings.Split(sdarr[j], " ")
		trimmedtok := strings.TrimSpace(tokens[1])
		digits := strings.Split(trimmedtok, ",")

		if j == SONG_DATA_FX_FREQ {
			var flotarr [NUM_INSTS]float64

			for idx, elt := range digits {
				flotarr[idx], err = strconv.ParseFloat(elt, 32)
			}
			
			fx_filter_params = flotarr

		} else if ((j == SONG_DATA_ENV_ATTACK) || (j == SONG_DATA_ENV_SUSTAIN) || (j == SONG_DATA_ENV_RELEASE) ) {
			
			for idx, elt := range digits {
				digarr[idx], err = strconv.ParseInt(elt, 10, 32)
			}
				
		} else {
			for idx, elt := range digits {
				digarr[idx], err = strconv.ParseInt(elt, 10, 8)
			}
		}

		if err != nil {
			fmt.Println(err)
			panic("ERROR: failed to parse music.inc file")
			
		}

		if j == SONG_DATA_FX_FREQ {
			j++
			continue
		}

		parm_map[j] = digarr
		j++		
	}

	return sdarr[0], parm_map, fx_filter_params, sdarr[30], sdarr[31]
}

//func LoadSongData(filename string) Song{
func LoadSongData(filename string) {
	//TODO: handle errors
	songbytes,_  := ioutil.ReadFile(filename)
	songstr := string(songbytes)

	auxdata, parms, fxfreq, patterns, columns := splitSonantOutput(songstr)
	
	fillStructures(auxdata, parms, fxfreq, patterns, columns)
}


