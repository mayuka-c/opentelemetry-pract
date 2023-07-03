# opentelemetry-pract
Study and practice opentelemtry

# Setup
## Jaeger Install
```bash
docker run -d --name jaeger -p 14268:14268 -p 16686:16686 jaegertracing/all-in-one:1.24
```

## Running the service
```bash
go run main.go
```