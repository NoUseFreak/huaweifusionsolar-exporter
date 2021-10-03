# Huawei Fusion Solar

This project pulls information from the huawei fusion solar API and exposes the values
for Prometheus to scrape.

## Requirements

In order to extract the metrics from the fusion solar api, you will need to request credentials by mail.

```
export HFSE_USERNAME=<username>
export HFSE_PASSWORD=<password>
```

## Usage

For debugging purposes, you can run `go run main.go get KpiStationMonth` to validate if it all works.
To expose the metrics endpoint, run `go run main.go metrics`. This will expose `http://0.0.0.0:2112/metrics`.

## Configuration

The application takes it's configuration from environment variables, the following are available.

| Name | Description | Default |
| --- | --- | --- |
| HFSE_USERNAME | Your api username | "" |
| HFSE_PASSWORD | Your api password | "" |
| HFSE_API_ENDPOINT | The api endpoint | "https://eu5.fusionsolar.huawei.com/thirdData" |