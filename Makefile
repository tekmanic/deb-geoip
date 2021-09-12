build: 
	docker build -t geoip:latest .

dockup:
	docker run -d --name geoip -p 3000:3000 geoip:latest

run:
	cd geoip && go run main.go

compile:
	cd geoip && go build main.go
	if test -d binaries; then mv geoip/main binaries/geoip;	else mkdir binaries && mv geoip/main binaries/geoip; fi

test:
	cd geoip && go test -v ./...

download:
	./scripts/mmdb.sh ${MAXMIND_KEY} geoip/internal/handlers

all: download compile build dockup