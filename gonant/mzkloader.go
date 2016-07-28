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

//TODO: dissociate splitSonantOutput interpretation from array idx
type sdtype struct {
	kind int      //AUXDATA..SONG_DATA_COLUMNS
	name string   //float32, uint, uint8, etc
	size uint     //num bits representation
}

//TODO: may already be taken care of with type assertions?
//TODO: but doesn't indicate what kind of data contained in arrays
type sddata struct {
	ptr interface{}
	sdt sdtype    
}

func fillStructures(parm_arr [32]interface{}) { //Song {
	//var song Song
	fmt.Println(*parm_arr[AUXDATA].(*string))

	for i:= SONG_DATA_OSC1_OCT; i <= SONG_DATA_COLUMNS ; i++ {

		if i == SONG_DATA_FX_FREQ {
			fmt.Println(*parm_arr[i].(*[NUM_INSTS]float64))
		} else if i == SONG_DATA_PATTERNS {
			fmt.Println(*parm_arr[i].(*[NUM_INSTS][48]int64))
		} else if i == SONG_DATA_COLUMNS {
			fmt.Println(*parm_arr[i].(*[NUM_INSTS * 10][32]int64))
		} else {
			fmt.Println(*parm_arr[i].(*[NUM_INSTS]int64))
		}
	}

	fmt.Println()

	fmt.Println("*****")
	fmt.Println(*parm_arr[SONG_DATA_OSC2_VOL].(*[NUM_INSTS]int64))
}

/* TODO: could use some big time cleaning up.. */
func splitSonantOutput(songdata string) ([32]interface{}){

	var parm_arr [32]interface{}
	var err error
	var symbols []string = strings.Split(songdata, "song_data_")
	var sdarr [32]string	

	for i, elt := range symbols {
		sdarr[i] = elt
	}

	parm_arr[AUXDATA] = &sdarr[AUXDATA]

	for j := SONG_DATA_OSC1_OCT; j <= SONG_DATA_COLUMNS; j++ {
		
		var digarr [NUM_INSTS]int64
		var parm_loc interface{}

		tokens := strings.Split(sdarr[j], " ")
		trimmedtok := strings.TrimSpace(tokens[1])
		digits := strings.Split(trimmedtok, ",")
		
		if j == SONG_DATA_PATTERNS {
			var sdpatarr [NUM_INSTS][48]int64
			loctokens := strings.Split(sdarr[j], "db")
			
			i := 0
			for idx, elt := range loctokens {
				if idx == 0 {
					continue
				}
				
				var locdigarr [48]int64
				loctrimmedtok := strings.TrimSpace(elt)
				digits := strings.Split(loctrimmedtok, ",")

				for digidx, digelt := range digits {
					locdigarr[digidx], err = strconv.ParseInt(digelt, 10, 8)
				}

				sdpatarr[i] = locdigarr
				i++
			}
			
			parm_loc = &sdpatarr

		} else if j == SONG_DATA_COLUMNS {

			var sdcolarr [NUM_INSTS * 10][32]int64
			loctokens := strings.Split(sdarr[j], "db")
			
			i := 0
			for idx, elt := range loctokens {
				if idx == 0 {
					continue
				}
				
				var locdigarr [32]int64
				loctrimmedtok := strings.TrimSpace(elt)
				digits := strings.Split(loctrimmedtok, ",")

				for digidx, digelt := range digits {
					locdigarr[digidx], err = strconv.ParseInt(digelt, 10, 8)
				}

				sdcolarr[i] = locdigarr
				i++
			}
			
			parm_loc = &sdcolarr
		
		} else if j == SONG_DATA_FX_FREQ {
			var flotarr [NUM_INSTS]float64

			for idx, elt := range digits {
				flotarr[idx], err = strconv.ParseFloat(elt, 32)
			}
			
			parm_loc = &flotarr

		} else if ((j == SONG_DATA_ENV_ATTACK) || (j == SONG_DATA_ENV_SUSTAIN) || (j == SONG_DATA_ENV_RELEASE) ) {
			
			for idx, elt := range digits {
				digarr[idx], err = strconv.ParseInt(elt, 10, 32)
			}

			parm_loc = &digarr
				
		} else {
			for idx, elt := range digits {
				digarr[idx], err = strconv.ParseInt(elt, 10, 8)
			}

			parm_loc = &digarr
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


