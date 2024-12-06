map:
	@go build -o ip_parser cmd/map/main.go
	@time ./ip_parser ip_addresses.txt

array:
	@go build -o ip_parser cmd/array/main.go
	@time ./ip_parser ip_addresses.txt

profile:
	@go tool pprof -web profile.prof