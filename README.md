# gandidyndns

Simple Gandi DNS client used to dynamically update A record with your current public IP address.

## Install 

`go get github.com/breed808/gandidyndns`

A docker image is also available if that's your preference: `docker pull breed808/gandidyndns:latest`

## Usage

```bash
export GANDI_API_KEY="supersecretkey"

# Set www.example.com and ftp.example.com value to current public IP address.
./gandidyndns example.com www ftp
```
