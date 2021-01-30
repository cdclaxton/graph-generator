# Random graph generator

This Golang code generates a random graph and saves it as an adjacency file. It was written to stress-test the connected component calculator, which was also written in Golang.

There are two modes of operation:

- Fixed probability of an edge between any given pair of vertices;
- Fixed number of edges.

If a vertex doesn't have any edges, it is not written to the adjacency file.

## Usage

- For usage help:

```
./graph-generator.exe -h
```

- To generate a random graph with a fixed probability of an edge:

```
./graph-generator.exe -n 1000 -p 0.1 -output edges.csv
```

- To generate a random graph with a fixed number of edges:

```
./graph-generator.exe -n 1000 -e 100 -output edges.csv
```
