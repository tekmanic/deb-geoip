#!/bin/zsh
mkdir GeoLite2
wget -c https://download.maxmind.com/app/geoip_download\?edition_id=GeoLite2-City\&license_key=$1\&suffix=tar.gz -O - | tar -C GeoLite2 --strip-components 1 -xz 
mv GeoLite2/GeoLite2-City.mmdb $2/GeoLite2-City.mmdb
rm -rf GeoLite2