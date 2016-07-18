package gonant

/*
#cgo LDFLAGS: -L. C:/Windows/System32/winmm.dll

#include <windows.h>
#include <mmsystem.h>

static short *wave_buf = 0;

static HWAVEOUT Wave_handle;
static WAVEFORMATEX Wave_format = {};
static WAVEHDR Wave_header = {};

static void Init_SoundBuf(short *buf)
{
    wave_buf = buf;
}

//wrappers for waveOut structure initialization

static void Init_WAVEFORMATEX(int nChannels, int nSamplesPerSec, int nAvgBytesPerSec, 
                              int nBlockAlign, int wBitsPerSample)
{
    Wave_format.wFormatTag = WAVE_FORMAT_PCM;
    Wave_format.nChannels = nChannels;
    Wave_format.nSamplesPerSec = nSamplesPerSec;
    Wave_format.nAvgBytesPerSec = nAvgBytesPerSec;
    Wave_format.nBlockAlign = nBlockAlign;
    Wave_format.wBitsPerSample = wBitsPerSample;
}

static void Init_WAVEHDR(int dwBufferLength)
{
    Wave_header.lpData = (char *)wave_buf;
    Wave_header.dwBufferLength = dwBufferLength;
    Wave_header.dwBytesRecorded = 0;
    Wave_header.dwUser = 0;
    Wave_header.dwFlags = 0;
    Wave_header.dwLoops = 0;
    Wave_header.lpNext = 0;
    Wave_header.reserved = 0;
}

static void Call_waveOutOpen()
{
    waveOutOpen(&Wave_handle,WAVE_MAPPER,&Wave_format,0,0,0);
}

static void Call_waveOutPrepareHeader()
{
    waveOutPrepareHeader(Wave_handle,&Wave_header,sizeof(WAVEHDR));
}

static void Call_waveOutWrite()
{
    waveOutWrite(Wave_handle,&Wave_header,sizeof(WAVEHDR));
}

*/
import "C"
import (
	"math"
	"fmt"
	"unsafe"
)

var WAVE_CHAN int = 2 //channels
var WAVE_SPS int = 8000//44100 // samples per second
var WAVE_BITS int = 16 // bits per sample
var WAVE_ALIGN int = WAVE_CHAN * WAVE_BITS / 8 // bytes per sample
var WAVE_SIZE int = WAVE_CHAN * WAVE_SPS * 240 // buffer size in samples
var AUDIO_CLIPAMP int = 32767 // audio clipping amplitude

var wave_buf = make([]int16, WAVE_SIZE * WAVE_CHAN)

func RenderSine() {
	t := float64(0)

	tone_hz := int16(512)
	tone_vol := int16(3000)
	cur_sample := 0

	waveperiod := WAVE_SPS / int(tone_hz)

	for sample_index := 0; sample_index < WAVE_SIZE ; sample_index++ {

		sineValue := math.Sin(t)
		
		sample_val := int16(sineValue * float64(tone_vol))
		wave_buf[cur_sample] = sample_val
		wave_buf[cur_sample + 1] = sample_val

		t += float64(2.0 * 3.141592653589793 * (1.0 / float64(waveperiod)))
		if t > (2.0 * 3.141592653589793) {
			t -= 2.0 * 3.141592653589793
		}

		cur_sample += 2
	}
	
}

func RenderWurstcapturez() {

	t := 0

	tone_vol := int16(300)
	cur_sample := 0

	for sample_index := 0; sample_index < WAVE_SIZE ; sample_index++ {
		
		yval := int16(t & 4095)

		texpr1 := t >> 12
		texpr2 := 0x9866 >> uint(texpr1 & 12)
		texpr3 := (15 & texpr2)

		xval := int16(t * (texpr3) / 6 & 127)

		yexpr1 := math.Sin(2000.0 / float64(yval))
		yexpr2 := (yexpr1 * 25.0) + (float64(xval * yval) / 10000.0)
		
		texpr4 := (((t >> 6) ^ (t >> 8)) | (t >> 12))
		texpr5 := (int16(texpr4) | xval) & 63
		
		val := int16(yexpr2) + texpr5

		sample_val := val * tone_vol
		wave_buf[cur_sample] = sample_val
		wave_buf[cur_sample + 1] = sample_val
		cur_sample += 2
		
		t++
	}
}

func Gonant_Init() {
	fmt.Println("Starting Point of our MZK")

	C.Init_WAVEFORMATEX(C.int(WAVE_CHAN), C.int(WAVE_SPS), C.int(WAVE_ALIGN) * C.int(WAVE_SPS),
		C.int(WAVE_ALIGN), C.int(WAVE_BITS))
	C.Init_SoundBuf((*C.short)(unsafe.Pointer(&wave_buf[0])))
	C.Init_WAVEHDR(C.int(WAVE_SIZE * WAVE_CHAN * WAVE_BITS / 8))

//	RenderSine()
	RenderWurstcapturez()

	C.Call_waveOutOpen()
	C.Call_waveOutPrepareHeader()
	C.Call_waveOutWrite()
}
