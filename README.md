# deb-geoip

Golang GeoIP REST API using Fiber, Maxmind, IP-API.com built on Debian Buster Slim. This leverages the example Fiber geoip apps but packages them
into a docker image.

## Installation

Set up the environment `MAXMIND_KEY` is needed to download the Maxmind database.

- Run `export MAXMIND_KEY=<your_key>`
- Run the download script to pull in the mmdb files `make download`
- Run `make build` to build the docker image
- Run `make dockup` to run the docker image
- Run `make test` to run the tests
- Run `make run` to run the app

## Environment Variables

`GEOIP_DIR` - Env variable used mainly by the docker container to find views and static files.

## Application

By default the app will start on `localhost:3000`.  Beyond the web application, there are two GET endpoints: `/geo/:ip` and `/geomm/:ip`. The first returns the IP-API.com result and the second returns the Maxmind result.
