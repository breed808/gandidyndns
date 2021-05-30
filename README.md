# gandidyndns

Simple Gandi DNS client used to dynamically update A record with your current public IP address.

## Install 

`go get github.com/breed808/gandidyndns`

## Usage

```bash
export GANDI_API_KEY="supersecretkey"

# Set www.example.com and ftp.example.com value to current public IP address.
./gandidyndns example.com www ftp
```
