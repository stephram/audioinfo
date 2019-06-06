package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	_ "net/http/pprof"

	"github.com/stephram/audioinfo/pkg/app"

	"github.com/stephram/audioinfo/internal/utils"
	"github.com/stephram/audioinfo/internal/utils/ulid"

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
	ID          string
	Name        string
	Channels    uint16
	Bits        uint16
	SBits       int32
	BytesPerSec uint32
	Rate        uint32
	Format      uint16
	Valid       bool
}

var (
	prtHdr *bool
	outFmt *string
)

func init() {
	utils.GetLogger().Infof("initialised logger")
}

func main() {
	prtHdr = flag.Bool("hdr", false, "print the column header. Only useful when fmt=text")
	outFmt = flag.String("fmt", "json", "output format 'text' or 'json'. Default 'json'")

	recurse := flag.Bool("r", false, "recurse into directories")

	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Println(app.New())
		fmt.Print("usage: audioinfo [options] <audiofile>\n")
		os.Exit(1)
	}
	// Get the args after the flags.
	args := os.Args[flag.NFlag()+1:]

	if *prtHdr && strings.EqualFold(*outFmt, TextFormat) {
		displayHeader()
	}

	go func() {
		_ = http.ListenAndServe("localhost:8080", nil)
	}()

	processFiles(args, *recurse)

	os.Exit(0)
}

func processFiles(fileNames []string, recurse bool) {
	for i := 0; i < len(fileNames); i++ {
		fileName := fileNames[i]

		if isDirectory(fileName) {
			if recurse {
				filesInfo, infoErr := ioutil.ReadDir(fileName)
				if infoErr != nil {
					log.WithError(infoErr).Errorf("failed to read directory: %s", fileName)
				}
				processFiles(readFilenames(filesInfo, fileName, recurse), recurse)
			}
			continue
		}

		if isValidWaveFilename(fileName) {
			processWaveFile(fileName)
		}
	}
}

func isValidWaveFilename(fileName string) bool {
	if strings.HasPrefix(fileName, ".") {
		return false
	}

	if !strings.HasSuffix(strings.ToLower(fileName), "wav") {
		return false
	}
	return true
}

func readFilenames(filesInfo []os.FileInfo, filesPath string, recurse bool) []string {
	var fileNames []string

	for i := 0; i < len(filesInfo); i++ {
		fileName := filesInfo[i].Name()
		fileNames = append(fileNames, fmt.Sprintf("%s%c%s", filesPath, filepath.Separator, fileName))
	}

	return fileNames
}

func isDirectory(fileName string) bool {
	fi, fiErr := os.Stat(fileName)
	if fiErr != nil {
		log.WithError(fiErr).Errorf("failed to stat filename: %s", fileName)
		return false
	}
	return fi.Mode().IsDir()
}

func processWaveFile(fileName string) {
	f, err := os.Open(fileName) // nolint: gosec
	if err != nil {
		log.WithError(err).Fatalf("failed to open file: %s", fileName)
	}
	defer f.Close() // nolint: errcheck

	d := wav.NewDecoder(f)
	d.ReadMetadata()

	if d.Err() != nil {
		log.WithError(d.Err()).Fatalf("failed to read metadata from file: %s", fileName)
		return
	}
	outputMetadata(fileName, d, *outFmt)
}

func outputMetadata(fileName string, d *wav.Decoder, outFmt string) {
	fileMetadata := FileMetadata{
		ID:          ulid.New(),
		Name:        fileName,
		Channels:    d.NumChans,
		Bits:        d.BitDepth,
		SBits:       d.SampleBitDepth(),
		BytesPerSec: d.AvgBytesPerSec,
		Rate:        d.SampleRate,
		Format:      d.WavAudioFormat,
		Valid:       d.IsValidFile(),
	}
	switch outFmt {
	case TextFormat:
		displayText(fileMetadata)
	case JSONFormat:
		displayJSON(fileMetadata)
	}
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
