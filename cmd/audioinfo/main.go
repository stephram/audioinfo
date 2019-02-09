package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/go-audio/wav"
	log "github.com/sirupsen/logrus"
)

// Constants to avoid duplicate literals.
const (
	TextFormat = "text"
	JSONFormat = "json"
)

// FileMetadata ...
type FileMetadata struct {
	Name        string
	Channels    uint16
	Bits        uint16
	SBits       int32
	BytesPerSec uint32
	Rate        uint32
	Format      uint16
	Valid       bool
}

func main() {
	printHdr := flag.Bool("h", false, "print the column header")
	outFmt := flag.String("ofmt", "json", "output format 'text' or 'json'. Default 'json'")
	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Print("usage: audioinfo <audiofile>\n")
		os.Exit(1)
	}
	args := os.Args[flag.NArg():]

	if *printHdr && strings.EqualFold(*outFmt, TextFormat) {
		displayHeader()
	}

	for i := 0; i < len(args); i++ {
		fileName := args[i]

		f, err := os.Open(fileName) // nolint: gosec
		if err != nil {
			log.WithError(err).Fatalf("failed to open file: %s", fileName)
		}
		defer f.Close()

		d := wav.NewDecoder(f)
		d.ReadMetadata()
		if d.Err() != nil {
			log.WithError(d.Err()).Fatalf("failed to read metadata from file: %s", fileName)
		}
		fileMetadata := FileMetadata{
			Name:        fileName,
			Channels:    d.NumChans,
			Bits:        d.BitDepth,
			SBits:       d.SampleBitDepth(),
			BytesPerSec: d.AvgBytesPerSec,
			Rate:        d.SampleRate,
			Format:      d.WavAudioFormat,
			Valid:       d.IsValidFile(),
		}
		switch *outFmt {
		case TextFormat:
			displayText(fileMetadata)
		case JSONFormat:
			displayJSON(fileMetadata)
		}
	}
	os.Exit(0)
}

func displayHeader() {
	fmt.Printf("%10s | %10s | %10s | %10s | %10s | %10s | %10s | %s\n",
		"AvgBps", "Bits", "NumChans", "SBits", "SampleRate", "WavFormat", "Valid", "Filename")
}

func displayJSON(fileMetadata FileMetadata) {
	jsonArr, jsonErr := json.Marshal(fileMetadata)
	if jsonErr != nil {
		log.WithError(jsonErr).Errorf("failed to marshal JSON for %+v", fileMetadata)
		return
	}
	fmt.Println(string(jsonArr))
}

func displayText(fmd FileMetadata) {
	fmt.Printf("%10d | %10d | %10d | %10d | %10d | %10d | %10t | %s\n",
		fmd.BytesPerSec, fmd.Bits, fmd.Channels, fmd.SBits, fmd.Rate, fmd.Format, fmd.Valid, fmd.Name)
}
