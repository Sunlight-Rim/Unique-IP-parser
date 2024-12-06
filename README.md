# Unique IP parser

Counts unique IPs from file in such format:
```
145.67.23.4
8.34.5.23
89.54.3.124
89.54.3.124
3.45.71.5
...
```

### Run

Generate random IPs:
```shell
IPS_COUNT=10000000
for ((i=0; i<$IPS_COUNT; i++)); do printf "%d.%d.%d.%d\n" "$((RANDOM % 256))" "$((RANDOM % 256))" "$((RANDOM % 256))" "$((RANDOM % 256))" >> ip_addresses.txt; done
```

Run version working with map:
```shell
make map
```

Or run version working with array:
```shell
make array
```

View memory usage:
```shell
make profile
```

### Conclusions

Map version is better for small amount of data, while array version is better for performing large file.