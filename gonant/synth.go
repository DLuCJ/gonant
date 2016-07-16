package gonant

/*
#include <windows.h>
#include <mmsystem.h>

// Sound-Specific
#define WAVE_CHAN 2 // channels
#define WAVE_SPS 44100 // samples per second
#define WAVE_BITS 16 // bits per sample
#define WAVE_ALIGN WAVE_CHAN * WAVE_BITS / 8 // bytes per sample
#define WAVE_SIZE WAVE_CHAN * WAVE_SPS * 240 // buffer size in samples
#define AUDIO_CLIPAMP 32767 // audio clipping amplitude

static short wave_buffer[WAVE_SIZE * WAVE_CHAN];

static HWAVEOUT wave_handle;

static WAVEFORMATEX wave_format = {
  WAVE_FORMAT_PCM, //wFormatTag
  WAVE_CHAN, // nChannels
  WAVE_SPS, // nSamplesPerSec
  WAVE_ALIGN * WAVE_SPS, // nAvgBytesPerSec
  WAVE_ALIGN, // nBlockAlign
  WAVE_BITS // wBitsPerSample
};

static WAVEHDR wave_header = {
  (char *) wave_buffer, // lpData
  WAVE_SIZE * WAVE_CHAN * WAVE_BITS / 8, // dwBufferLength
  0, // dwBytesRecorded
  0, // dwUser
  0, // dwFlags
  0, // dwLoops
  0, // lpNext
  0 // reserved
};

//wrappers for waveOut calls here

*/
import "C"
import "fmt"

func Gonant_Init() {
	fmt.Println("Starting Point of our MZK")
}
