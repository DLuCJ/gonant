package gonant

var Endpattern int
var Rowlen int
var Fastforward int = -1

type Column struct {
	notes [32]uint8
}

type Instrument struct {

	//Oscillator 1

	osc1_oct uint8       // Octave knob
	osc1_det uint8       // Detune knob
	osc1_detune uint8    // Actual detune knob
	osc1_xenv uint8      // Multiply freq by envelope
	osc1_vol uint8       // Volume knob
	osc1_waveform uint8  // Wave form
	
	//Oscillator 2

	osc2_oct uint8       // Octave knob
	osc2_det uint8       // Detune knob
	osc2_detune uint8    // Actual detune knob
	osc2_xenv uint8      // Multiply freq by envelope
	osc2_vol uint8       // Volume knob
	osc2_waveform uint8  // Wave form
	
	//Noise oscillator

	noise_fader uint8    // Amount of noise to add
	
	//Envelope

	env_attack uint      // Attack
	env_sustain uint     // Sustain
	env_release uint     // Release
	env_master uint8     // Master volume knob

	//FX

	fx_filter uint8      // Hi/lo/bandpass or notch toggle
	fx_freq float32      // FX Frequency
	fx_resonance uint8   // FX Resonance
	fx_delay_time uint8  // Delay time
	fx_delay_amt uint8   // Delay amount
	fx_pan_freq uint8    // Panning frequency
	fx_pan_amt uint8     // Panning amount
	
	//LFO

	lfo_osc1_freq uint8  // Modify osc1 freq (FM) toggle
	lfo_fx_freq uint8    // Modify fx freq toggle
	lfo_freq uint8       // LFO freq
	lfo_amt uint8        // LFO amount
	lfo_waveform uint8   // LFO waveform

	//Patterns
	pats [48]int8

	cols [10]Column
}

type Song struct {
	Inst [8]Instrument
}
