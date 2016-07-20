package gonant

import (
	"fmt"
	"io/ioutil"
)

type stringcture struct {
	sonant_rowlen string
	song_data_osc1_oct string
	song_data_osc1_det string
	song_data_osc1_detune string
	song_data_osc1_xenv string
	song_data_osc1_vol string
	song_data_osc1_waveform string

	song_data_osc2_oct string
	song_data_osc2_det string
	song_data_osc2_detune string
	song_data_osc2_xenv string
	song_data_osc2_vol string
	song_data_osc2_waveform string

	song_data_noise_fader string

	song_data_env_attack string
	song_data_env_sustain string
	song_data_env_release string
	song_data_env_master string

	song_data_fx_filter string
	song_data_fx_freq string
	song_data_fx_resonance string
	song_data_fx_delay_time string
	song_data_fx_delay_amt string
	song_data_fx_pan_freq string
	song_data_fx_pan_amt string

	song_data_lfo_osc1_freq string
	song_data_lfo_fx_freq string
	song_data_lfo_freq string
	song_data_lfo_amt string
	song_data_lfo_waveform string

	song_data_patterns string

	song_data_columns string
}

func fillStructures() {
	fmt.Println(Column{})
	fmt.Println(Song{})
	fmt.Println(Instrument{})
}

func parseSonantOutput(songdata string) stringcture {
	var strctr stringcture
	return strctr
}

func LoadSongData(filename string) {
	//var songdata stringcture
	
	//TODO: handle errors
	songbytes,_  := ioutil.ReadFile(filename)
	songstr := string(songbytes)

	//songdata = parseSonantOutput(songstr)

	fmt.Println(songstr)
	fillStructures()
}


