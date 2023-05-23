FROM golang:1.20 AS builder
WORKDIR /src/github.com/sylweltan/prometheus-vmware-exporter
COPY ./ /src/github.com/sylweltan/prometheus-vmware-exporter
RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux go build

FROM alpine:3.8
COPY --from=builder /src/github.com/sylweltan/prometheus-vmware-exporter/prometheus-vmware-exporter /usr/bin/prometheus-vmware-exporter
EXPOSE 9512
ENTRYPOINT ["prometheus-vmware-exporter"]
