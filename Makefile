all: download compile build slim dockup

build: 
	docker build -t deb-geoip:latest .

dockup:
	docker run -d --name deb-geoip -p 3000:3000 deb-geoip:slim

run:
	cd geoip && go run main.go

compile:
	cd geoip && go build main.go
	if test -d bin; then mv geoip/main bin/geoip;	else mkdir bin && mv geoip/main bin/geoip; fi

test:
	cd geoip && go test -v ./...

download:
	./scripts/mmdb.sh ${MAXMIND_KEY} geoip/internal/handlers

slim:
	docker-slim build --dockerfile Dockerfile --tag deb-geoip:slim .

clean:
	docker kill deb-geoip || true
	docker rm -f deb-geoip
