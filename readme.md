# Random graph generator

This Golang code generates a random graph and saves it as an adjacency file. It was written to stress-test the connected component calculator, which is also written in Golang.

There are two modes of operation:

- Fixed probability of an edge between any given pair of vertices;
- Fixed number of edges.

If a vertex doesn't have any edges, it is not written to the adjacency file.

## Usage

- For usage:

```
./graph-generator.exe -h
```

- To generate a random graph with 1000 vertices a fixed probability of an edge of 0.1:

```
./graph-generator.exe -n 1000 -p 0.1 -output edges.csv
```

- To generate a random graph with 1,000 vertices and 100 edges:

```
./graph-generator.exe -n 1000 -e 100 -output edges.csv
```

- To generate a network with 50,000,000 vertices and 40,000,000 edges:

```
./graph-generator.exe -n 50000000 -e 40000000 -output edges.csv
```
