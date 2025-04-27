# Area search

This service accepts the **searched area** (*textual description or address*) and returns the corresponding latitude and longitude coordinates.

It uses the Nominatim [Public API](https://nominatim.org/release-docs/develop/api/Search/#structured-query) for the purpose of converting addresses into coordinates, which uses OpenStreetMap data under the hood for geo-coding.


## Setup Instructions

### Run Server

```bash
go mod tidy
air main.go # or `go run main.go`
```

---

## Testing Instructions

### Run Client

```bash
go run tests/client.go
```