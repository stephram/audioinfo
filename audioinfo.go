package main

import (
	"fmt"
	"github.com/go-audio/wav"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Print("usage: audioinfo <audiofile>\n")
		os.Exit(1)
	}
	args := os.Args

	fmt.Printf("%10s | %10s | %10s | %10s | %10s | %10s | %10s | %s\n",
		"AvgBps", "Bits", "NumChans", "SBits", "SampleRate", "WavFormat", "Valid", "Filename")

	for i := 1; i < len(args); i++ {
		fileName := args[i]

		f, err := os.Open(fileName)
		if err != nil {
			log.WithError(err).Fatalf("failed to open file: %s", fileName)
		}
		defer f.Close()

		d := wav.NewDecoder(f)
		d.ReadMetadata()
		if d.Err() != nil {
			log.WithError(d.Err()).Fatalf("failed to read metadata from file: %s", fileName)
		}
		fmt.Printf("%10d | %10d | %10d | %10d | %10d | %10d | %10t | %s\n",
			d.AvgBytesPerSec, d.BitDepth, d.NumChans, d.SampleBitDepth(), d.SampleRate, d.WavAudioFormat, d.IsValidFile(), fileName)
	}
	os.Exit(0)
}
