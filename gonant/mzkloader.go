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

func parseAuxData(dat string) {
	var err error

	lines := strings.Split(dat, "\n")

	for _, line := range lines {
		tokens := strings.Split(line, " ")
		
		if tokens[0] != "%define" {
			continue
		}

		if tokens[1] == "_4K_SONANT_ROWLEN_" {
			Rowlen, err = strconv.Atoi(tokens[2])
		} else if tokens[1] == "_4K_SONANT_ENDPATTERN_" {
			Endpattern, err = strconv.Atoi(tokens[2])
		} else if tokens[1] == "_4K_SONANT_FASTFORWARD_" {
			Fastforward = 0
		}	
	}

	if err != nil {
		fmt.Println(err)
		panic("ERROR: failed to parse music.inc file")
		
	}	
} 

func fillStructures(parm_arr [32]interface{}) Song{ 
	var song Song
	parseAuxData(*parm_arr[AUXDATA].(*string))

	for i := 0; i < 8; i++ {
		song.Inst[i].osc1_oct       = uint8((*parm_arr[SONG_DATA_OSC1_OCT].(*[NUM_INSTS]int64))[i])
		song.Inst[i].osc1_det       = uint8((*parm_arr[SONG_DATA_OSC1_DET].(*[NUM_INSTS]int64))[i])
		song.Inst[i].osc1_detune    = uint8((*parm_arr[SONG_DATA_OSC1_DETUNE].(*[NUM_INSTS]int64))[i])
		song.Inst[i].osc1_xenv 	    = uint8((*parm_arr[SONG_DATA_OSC1_XENV].(*[NUM_INSTS]int64))[i])
		song.Inst[i].osc1_vol 	    = uint8((*parm_arr[SONG_DATA_OSC1_VOL].(*[NUM_INSTS]int64))[i])
		song.Inst[i].osc1_waveform  = uint8((*parm_arr[SONG_DATA_OSC1_WAVEFORM].(*[NUM_INSTS]int64))[i])
		song.Inst[i].osc2_oct       = uint8((*parm_arr[SONG_DATA_OSC2_OCT].(*[NUM_INSTS]int64))[i])
		song.Inst[i].osc2_det       = uint8((*parm_arr[SONG_DATA_OSC2_DET].(*[NUM_INSTS]int64))[i])
		song.Inst[i].osc2_detune    = uint8((*parm_arr[SONG_DATA_OSC2_DETUNE].(*[NUM_INSTS]int64))[i])
		song.Inst[i].osc2_xenv      = uint8((*parm_arr[SONG_DATA_OSC2_XENV].(*[NUM_INSTS]int64))[i])
		song.Inst[i].osc2_vol       = uint8((*parm_arr[SONG_DATA_OSC2_VOL].(*[NUM_INSTS]int64))[i])
		song.Inst[i].osc2_waveform  = uint8((*parm_arr[SONG_DATA_OSC2_WAVEFORM].(*[NUM_INSTS]int64))[i])
		song.Inst[i].noise_fader    = uint8((*parm_arr[SONG_DATA_NOISE_FADER].(*[NUM_INSTS]int64))[i])
		song.Inst[i].env_attack     = uint((*parm_arr[SONG_DATA_ENV_ATTACK].(*[NUM_INSTS]int64))[i])
		song.Inst[i].env_sustain    = uint((*parm_arr[SONG_DATA_ENV_SUSTAIN].(*[NUM_INSTS]int64))[i])
		song.Inst[i].env_release    = uint((*parm_arr[SONG_DATA_ENV_RELEASE].(*[NUM_INSTS]int64))[i])
		song.Inst[i].env_master     = uint8((*parm_arr[SONG_DATA_ENV_MASTER].(*[NUM_INSTS]int64))[i])
		song.Inst[i].fx_filter      = uint8((*parm_arr[SONG_DATA_FX_FILTER].(*[NUM_INSTS]int64))[i])
		song.Inst[i].fx_freq        = float32((*parm_arr[SONG_DATA_FX_FREQ].(*[NUM_INSTS]float64))[i])
		song.Inst[i].fx_resonance   = uint8((*parm_arr[SONG_DATA_FX_RESONANCE].(*[NUM_INSTS]int64))[i])
		song.Inst[i].fx_delay_time  = uint8((*parm_arr[SONG_DATA_FX_DELAY_TIME].(*[NUM_INSTS]int64))[i])
		song.Inst[i].fx_delay_amt   = uint8((*parm_arr[SONG_DATA_FX_DELAY_AMT].(*[NUM_INSTS]int64))[i])
		song.Inst[i].fx_pan_freq    = uint8((*parm_arr[SONG_DATA_FX_PAN_FREQ].(*[NUM_INSTS]int64))[i])
		song.Inst[i].fx_pan_amt     = uint8((*parm_arr[SONG_DATA_FX_PAN_AMT].(*[NUM_INSTS]int64))[i])
		song.Inst[i].lfo_osc1_freq  = uint8((*parm_arr[SONG_DATA_LFO_OSC1_FREQ].(*[NUM_INSTS]int64))[i])
		song.Inst[i].lfo_fx_freq    = uint8((*parm_arr[SONG_DATA_LFO_FX_FREQ].(*[NUM_INSTS]int64))[i])
		song.Inst[i].lfo_freq       = uint8((*parm_arr[SONG_DATA_LFO_FREQ].(*[NUM_INSTS]int64))[i])
		song.Inst[i].lfo_amt        = uint8((*parm_arr[SONG_DATA_LFO_AMT].(*[NUM_INSTS]int64))[i])
		song.Inst[i].lfo_waveform   = uint8((*parm_arr[SONG_DATA_LFO_WAVEFORM].(*[NUM_INSTS]int64))[i])

		var patarr [48]int8 
		
		for j, elt := range (*parm_arr[SONG_DATA_PATTERNS].(*[NUM_INSTS][48]int64))[i] {
			patarr[j] = int8(elt)
		}
		
		song.Inst[i].pats = patarr
		
		for j := 0; j < 10; j++ {

			var notearr [32]uint8
			for k, elt := range (*parm_arr[SONG_DATA_COLUMNS].(*[NUM_INSTS * 10][32]int64))[i * 10 + j] {
				notearr[k] = uint8(elt)
			}
			
			song.Inst[i].cols[j].notes = notearr
		}
	}

	return song
}

func splitToken(token string) []string {
	trimmedtok := strings.TrimSpace(token)
	digits := strings.Split(trimmedtok, ",")
	return digits[:]
}

func splitSonantOutput(songdata string) [32]interface{} {
	var parm_arr [32]interface{}
	var err error
	var symbols []string = strings.Split(songdata, "song_data_")
	var sdarr [32]string	

	for i, elt := range symbols {
		sdarr[i] = elt
	}

	parm_arr[AUXDATA] = &sdarr[AUXDATA]

	for j := SONG_DATA_OSC1_OCT; j <= SONG_DATA_COLUMNS; j++ {
		var parm_loc interface{}
		var delim string
		
		if j == SONG_DATA_PATTERNS || j == SONG_DATA_COLUMNS {
			delim = "db"
		} else {
			delim = " "
		}
		
		tokens := strings.Split(sdarr[j], delim)
		digits := splitToken(tokens[1])
	
		if j == SONG_DATA_PATTERNS {
			var sdpatarrarr [NUM_INSTS][48]int64

			i := 0
			for idx, elt := range tokens {
				if idx == 0 {
					continue
				}
				
				digits := splitToken(elt)

				for digidx, digelt := range digits {
					sdpatarrarr[i][digidx], err = strconv.ParseInt(digelt, 10, 8)
				}
				
				i++
			}
			
			parm_loc = &sdpatarrarr

		} else if j == SONG_DATA_COLUMNS {

			var sdcolarrarr [NUM_INSTS * 10][32]int64

			i := 0
			for idx, elt := range tokens {
				if idx == 0 {
					continue
				}
				
				digits := splitToken(elt)
				
				for digidx, digelt := range digits {
					sdcolarrarr[i][digidx], err = strconv.ParseInt(digelt, 10, 8)
				}

				i++
			}
			
			parm_loc = &sdcolarrarr
		
		} else if j == SONG_DATA_FX_FREQ {
			var flotarr [NUM_INSTS]float64

			for idx, elt := range digits {
				flotarr[idx], err = strconv.ParseFloat(elt, 32)
			}
			
			parm_loc = &flotarr

		} else if (j == SONG_DATA_ENV_ATTACK || j == SONG_DATA_ENV_SUSTAIN || j == SONG_DATA_ENV_RELEASE ) {
			var digarr [NUM_INSTS]int64
			
			for idx, elt := range digits {
				digarr[idx], err = strconv.ParseInt(elt, 10, 32)
			}

			parm_loc = &digarr
				
		} else {
			var digarr [NUM_INSTS]int64
			
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

func LoadSongData(filename string) Song{
	//TODO: handle errors
	songbytes,_  := ioutil.ReadFile(filename)
	songstr := string(songbytes)

	parm_arr := splitSonantOutput(songstr)	
	song := fillStructures(parm_arr)
	return song
}


