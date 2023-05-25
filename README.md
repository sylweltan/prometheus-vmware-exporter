# prometheus-vmware-exporter

Collect metrics ESXi Host

## Badge

[![License](https://img.shields.io/github/license/sylweltan/prometheus-vmware-exporter)](/LICENSE)
[![Release](https://img.shields.io/github/release/sylweltan/prometheus-vmware-exporter.svg)](https://github.com/sylweltan/prometheus-vmware-exporter/releases/latest)
[![GitHub Releases Stats of prometheus-vmware-exporter](https://img.shields.io/github/downloads/sylweltan/prometheus-vmware-exporter/total.svg?logo=github)](https://somsubhra.github.io/github-release-stats/?username=sylweltan&repository=prometheus-vmware-exporter)
[![Go CI](https://github.com/sylweltan/prometheus-vmware-exporter/actions/workflows/ci.yaml/badge.svg?branch=master&event=push)](https://github.com/sylweltan/prometheus-vmware-exporter/actions/workflows/ci.yaml?branch=master&event=push)
[![Frogbot Scan Pull Request](https://github.com/sylweltan/prometheus-vmware-exporter/actions/workflows/frogbot-scan-pr-go.yml/badge.svg)](https://github.com/sylweltan/prometheus-vmware-exporter/actions/workflows/frogbot-scan-pr-go.yml)
[![Scanned by Frogbot](https://raw.github.com/jfrog/frogbot/master/images/frogbot-badge.svg)](https://github.com/jfrog/frogbot#readme)
[![Go Report Card](https://goreportcard.com/badge/github.com/sylweltan/prometheus-vmware-exporter)](https://goreportcard.com/report/github.com/sylweltan/prometheus-vmware-exporter)
![Docker Image Size](https://img.shields.io/docker/image-size/sylweltan/prometheus-vmware-exporter.svg?sort=date)
![Docker Image Version (latest by date):](https://img.shields.io/docker/v/sylweltan/prometheus-vmware-exporter.svg?sort=date)
![Docker Pulls](https://img.shields.io/docker/pulls/sylweltan/prometheus-vmware-exporter.svg)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=prometheus-vmware-exporter&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=prometheus-vmware-exporter)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=prometheus-vmware-exporter&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=prometheus-vmware-exporter)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=prometheus-vmware-exporter&metric=coverage)](https://sonarcloud.io/summary/new_code?id=prometheus-vmware-exporter)

[![Quality gate](https://sonarcloud.io/api/project_badges/quality_gate?project=prometheus-vmware-exporter)](https://sonarcloud.io/summary/new_code?id=prometheus-vmware-exporter)

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
