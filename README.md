# prometheus-vmware-exporter

Collect metrics ESXi Host

## Badge

[![Scanned by Frogbot](https://raw.github.com/jfrog/frogbot/master/images/frogbot-badge.svg)](https://github.com/jfrog/frogbot#readme)
[![Go Report Card](https://goreportcard.com/badge/github.com/sylweltan/prometheus-vmware-exporter)](https://goreportcard.com/report/github.com/sylweltan/prometheus-vmware-exporter)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=prometheus-vmware-exporter&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=prometheus-vmware-exporter)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=prometheus-vmware-exporter&metric=coverage)](https://sonarcloud.io/summary/new_code?id=prometheus-vmware-exporter)

## Build

```sh
docker build -t prometheus-vmware-exporter .
```

## Run

```sh
docker run -d \
  --restart=always \
  --name=prometheus-vmware-exporter \
  --env=ESX_HOST esx.domain.local \
  --env=ESX_USERNAME user \
  --env=ESX_PASSWORD password \
  --env=ESX_LOG debug \
  prometheus-vmware-exporter 
```
