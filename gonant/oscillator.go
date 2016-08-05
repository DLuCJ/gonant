package gonant

//note: https://en.wikibooks.org/wiki/Sound_Synthesis_Theory/Oscillators_and_Wavetables

import (
	"math"
)

type oscfn func(float64) float64

var oscarr = []oscfn{oscSin, oscSquare, oscSaw, oscTri}

func getOscOutput(nwaveform uint8, fvalue float64) float64 {
	return oscarr[nwaveform](fvalue)
}

func oscSin(value float64) float64 {
	return math.Sin(2.0 * 3.141592653589793 * value)
}

func oscSquare(value float64) float64 {
	if oscSin(value) < 0 {	return -1.0 } 
	return 1.0
}

func oscSaw(value float64) float64 {
/*	
  "fld1;"          //push 1.0 onto reg stack
  "fld %1;"        //push value onto reg stack

  //ST(0) : value
  //ST(1) : 1.0

  "fprem;"         //computes partial remainder of st(0) / st(1)
  "fstp %%st(1);"  //copy ST(0) to ST(1) & pop
  "fstp %0;"       //copy ST(0) to result & pop, so clean fp stack?
  : "=m" (result)
  : "m" (value)
*/
	return math.Mod(value, 1.0) - 0.5	
}

func oscTri(value float64) float64 {
	var v2 float64 = (oscSaw(value) + 0.5) * 4.0
	if v2 < 2.0 { return v2 - 1.0 }
	return 3.0 - v2
}

