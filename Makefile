

install:
	go get -u github.com/go-audio/wav
	go get -u github.com/sirupsen/logrus

build:
	go build cmd/audioinfo/main.go

run:
	#./main /**/*.wav
	#./audioinfo /Volumes/SANDISK-64G/RAAF\ Pt\ Cook/SOUNDS/**/*.wav
	#find / -name *.wav* -exec ./main -h {} \;


