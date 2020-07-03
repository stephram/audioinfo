![Go](https://github.com/stephram/audioinfo/workflows/Go/badge.svg)

# audioinfo
Read header information from WAV files and output as JSON or TEXT.

### Usage
```
% build/audioinfo -h
Usage of build/audioinfo:
  -arr
    	output as JSON array
  -fmt string
    	output format 'text' or 'json'. Default 'json' (default "json")
  -hdr
    	print the column header. Only useful when fmt=text
  -prof
    	enable the pprof package. Listening on port 8080 (default true)
  -r	recurse into directories
```

### Running
Output as JSON
```text
audioinfo ~/Dropbox/Audio/Soundscape/*
```
Example output:
```json5
{
  "ID": "01EC95591XQH9G0Q01E2R4CYGE",
  "Name": "ChannelCount-1-32-discrete.wav",
  "Path": "/Users/sg/Dropbox/Audio/",
  "AbsName": "/Users/sg/Dropbox/Audio//ChannelCount-1-32-discrete.wav",
  "Channels": 32,
  "Bits": 16,
  "SBits": 16,
  "BytesPerSec": 3072000,
  "Rate": 48000,
  "Format": 1,
  "Valid": true,
  "Duration": 31677714192,
  "Description": "Format: WAVE - 32 channels @ 48000 / 16 bits - Duration: 31.677714 seconds",
  "Metadata": {
    "SamplerInfo": null,
    "Artist": "",
    "Comments": "",
    "Copyright": "",
    "CreationDate": "2018-03-28 19:05:37 ",
    "Engineer": "Stephen Graham",
    "Technician": "",
    "Genre": "",
    "Keywords": "",
    "Medium": "",
    "Title": "",
    "Product": "",
    "Subject": "",
    "Software": "Sound Forge Pro 3.0.0 (Build 100)",
    "Source": "",
    "Location": "",
    "TrackNbr": "",
    "CuePoints": null
  }
}
```

Output as TEXT
```text
audioinfo -hdr -fmt=text ~/Dropbox/Audio/Soundscape/*
```
Example output:
```text
    AvgBps |       Bits |   NumChans |      SBits | SampleRate |  WavFormat |      Valid | Filename
    176400 |         16 |          2 |         16 |      44100 |          1 |       true | Ballet Exam - Aug 2019.wav
     44100 |         16 |          1 |         16 |      22050 |          1 |       true | CLAP.WAV
     22311 |          4 |          2 |          4 |      22050 |          2 |      false | CancURL.wav
   1536000 |         16 |         16 |         16 |      48000 |          1 |       true | ChannelCount-1-16-discrete.wav
   3072000 |         16 |         32 |         16 |      48000 |          1 |       true | ChannelCount-1-32-discrete.wav
   3072000 |         16 |         32 |         16 |      48000 |          1 |       true | ChannelCount-1-32.wav
   1536000 |         16 |         16 |         16 |      48000 |          1 |       true | ChannelCount-17-32-discrete.wav
    768000 |         16 |          8 |         16 |      48000 |          1 |       true | Magpie-8ch.wav
```

### Quick start

```
make install
make run
```

Should give output similar to...

```$xslt
{"ID":"01D38C1EFBQPZAJSZWFJZ0GX7F","Name":"alert.wav","Channels":2,"Bits":16,"SBits":16,"BytesPerSec":176400,"Rate":44100,"Format":1,"Valid":true}
{"ID":"01D38C1HQ52PF4TYDT1YEP8CZ5","Name":"gran2bubbleLoop.wav","Channels":2,"Bits":16,"SBits":16,"BytesPerSec":176400,"Rate":44100,"Format":1,"Valid":true}
{"ID":"01D38C1HSE4WQZK8ZKAHYMDG6Q","Name":"MurrindindiSong1.wav","Channels":2,"Bits":16,"SBits":16,"BytesPerSec":176400,"Rate":44100,"Format":1,"Valid":true}
{"ID":"01D38C1HTYTHPF3Z7169WDJWSR","Name":"dig13rev.wav","Channels":2,"Bits":16,"SBits":16,"BytesPerSec":176400,"Rate":44100,"Format":1,"Valid":true}
```
## Note
Lots of code tidying required. Particular in relation to command line options and output formatting.