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

Storing IP addresses in the map as strings is obviously not productive in terms of memory and speed (there can be 2^32 IPv4 addresses in total, one line takes 16 bytes on average, i.e. 2^36 bytes = 64 GB of memory will be required to store such a filled map).

I have written two versions of the solution. \
The first one converts IPv4 addresses to 32-bit numbers (uint32) and stores them into map. It's a little efficiently and works on not too many IP addresses, but grows in memory on a large one. \
The second version initializes an array of 2^32 bytes (4 GB) at the start of the application, the indexes of which are all 32-bit numbers (IPv4 addresses). Then it also converts IP addresses into numbers and use them as indexes in the array. This version saves memory and runs faster than the first one on large amounts of data (about 100 MB on my processor).
