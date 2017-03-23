// The MIT License (MIT)
//
// Copyright (c) 2017 aerth <aerth@riseup.net>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
// the Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
// FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
// COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
// IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// playwav plays a .wav file using ALSA
package main

import (
	"io"
	"os"

	"github.com/cocoonlife/goalsa"
	"github.com/cryptix/wav"
	"github.com/fatih/color"
)

func main() {
	if len(os.Args) < 2 {
		color.Red("What .WAV file to play?")
		return
	}

	// file exists
	soundfile, err := os.Open(os.Args[1])
	if err != nil {
		color.Red("open: %s\n", err.Error())
		return
	}

	// stat for size
	sndfileinfo, err := os.Stat(soundfile.Name())
	if err != nil {
		color.Red("stat: %s\n", err.Error())
		return
	}

	// wavReader
	wavReader, err := wav.NewReader(soundfile, sndfileinfo.Size())
	if err != nil {
		color.Red("Wave: %s\n\n", err.Error())
		return
	}

	// require wavReader
	if wavReader == nil {
		color.Red("nil WAV reader")
		return
	}

	// print .WAV info
	color.Cyan(wavReader.String())

	// open default ALSA playback device
	out, err := alsa.NewPlaybackDevice("default", 1, alsa.FormatS16LE, int(wavReader.GetSampleCount()), alsa.BufferParams{})
	if err != nil {
		color.Red("alsa: %s\n\n", err.Error())
		return
	}

	// require ALSA device
	if out == nil {
		color.Red("No device")
		return
	}

	// close device when finished
	defer out.Close()


	for {
		s, err := wavReader.ReadSampleEvery(2, 0)
		var cvert []int16
		for _, b := range s {
			cvert = append(cvert, int16(b))
		}
		if cvert != nil {
			// play!
			out.Write(cvert)
		}
		cvert = []int16{}

		if err == io.EOF {
			break
		} else if err != nil {
			color.Red("Bad:", err.Error())
			return
		}
	}

	// all done
}
