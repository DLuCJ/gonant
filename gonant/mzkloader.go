package gonant

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
	"unsafe"
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

//TODO: dissociate splitSonantOutput interpretation from array idx
type SDTYPE struct {
	ptr unsafe.Pointer
	sdtype int    //AUXDATA..SONG_DATA_COLUMNS
}

func fillStructures(parm_arr [32]unsafe.Pointer) { //Song {
	//var song Song
	fmt.Println(*(*string)(parm_arr[AUXDATA]))

	for i:= SONG_DATA_OSC1_OCT; i < SONG_DATA_PATTERNS ; i++ {

		if i == SONG_DATA_FX_FREQ {
			fmt.Println(*(*[NUM_INSTS]float64)(parm_arr[i]))

		} else {
			fmt.Println(*(*[NUM_INSTS]int64)(parm_arr[i]))
		}
	}

	fmt.Println()
	fmt.Println(*(*string)(parm_arr[SONG_DATA_PATTERNS]))
	fmt.Println()
	fmt.Println(*(*string)(parm_arr[SONG_DATA_COLUMNS]))

	fmt.Println("*****")
	fmt.Println(*(*[NUM_INSTS]int64)(parm_arr[SONG_DATA_OSC2_VOL]))
}

func splitSonantOutput(songdata string) ([32]unsafe.Pointer){

	var parm_arr [32]unsafe.Pointer
	var err error
	var symbols []string = strings.Split(songdata, "song_data_")
	var sdarr [32]string	

	for i, elt := range symbols {
		sdarr[i] = elt
	}

	parm_arr[AUXDATA] = unsafe.Pointer(&sdarr[AUXDATA])
	parm_arr[SONG_DATA_PATTERNS] = unsafe.Pointer(&sdarr[SONG_DATA_PATTERNS])
	parm_arr[SONG_DATA_COLUMNS] = unsafe.Pointer(&sdarr[SONG_DATA_COLUMNS])

	for j := SONG_DATA_OSC1_OCT; j < SONG_DATA_PATTERNS; j++ {
		
		var digarr [NUM_INSTS]int64
		var parm_loc unsafe.Pointer

		tokens := strings.Split(sdarr[j], " ")
		trimmedtok := strings.TrimSpace(tokens[1])
		digits := strings.Split(trimmedtok, ",")

		
		if j == SONG_DATA_FX_FREQ {
			var flotarr [NUM_INSTS]float64

			for idx, elt := range digits {
				flotarr[idx], err = strconv.ParseFloat(elt, 32)
			}
			
			parm_loc = unsafe.Pointer(&flotarr)

		} else if ((j == SONG_DATA_ENV_ATTACK) || (j == SONG_DATA_ENV_SUSTAIN) || (j == SONG_DATA_ENV_RELEASE) ) {
			
			for idx, elt := range digits {
				digarr[idx], err = strconv.ParseInt(elt, 10, 32)
			}

			parm_loc = unsafe.Pointer(&digarr)
				
		} else {
			for idx, elt := range digits {
				digarr[idx], err = strconv.ParseInt(elt, 10, 8)
			}

			parm_loc = unsafe.Pointer(&digarr)
		}

		if err != nil {
			fmt.Println(err)
			panic("ERROR: failed to parse music.inc file")
			
		}

		parm_arr[j] = parm_loc
	}

	return parm_arr
}

//func LoadSongData(filename string) Song{
func LoadSongData(filename string) {
	//TODO: handle errors
	songbytes,_  := ioutil.ReadFile(filename)
	songstr := string(songbytes)

	parm_arr := splitSonantOutput(songstr)
	
	fillStructures(parm_arr)
}


