

install:
	go get -u github.com/go-audio/wav
	go get -u github.com/sirupsen/logrus

build:
	go build cmd/audioinfo/main.go

run:
	#./main /**/*.wav
	#./audioinfo /Volumes/SANDISK-64G/RAAF\ Pt\ Cook/SOUNDS/**/*.wav
	find / -name [aA-zZ]*.wav -exec ./main -h {} \;

runJSON:
	#find $(HOME) -name *.wav -exec ./main -ofmt=json {} \;
	find $(HOME) -name [aA-zZ]*.wav -exec ./main {} \;



