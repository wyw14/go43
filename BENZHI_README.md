# API quota

This repository provides a runnable Go component for api quota, including domain values, storage, policy, lifecycle, export, and batch processing.

## Build and verify

```bash
docker buildx build --platform linux/amd64 -f benzhi.Dockerfile -t quotawindow:amd64 --load .
docker buildx build --platform linux/arm64 -f benzhi.Dockerfile -t quotawindow:arm64 --load .
docker run --rm --platform linux/amd64 quotawindow:amd64 bash -c "go build ./... && go test ./... && go vet ./..."
```

Run the demo with `go run ./cmd/demo`.
