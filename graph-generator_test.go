package main

import (
	"testing"
)

func TestMinMaxLessThan(t *testing.T) {
	lower, upper := minMax(1, 2)

	if lower != 1 {
		t.Fatalf("Expected lower to be 1, got %v\n", lower)
	}

	if upper != 2 {
		t.Fatalf("Expected upper to be 2, got %v\n", upper)
	}
}

func TestMinMaxGreaterThan(t *testing.T) {
	lower, upper := minMax(2, 1)

	if lower != 1 {
		t.Fatalf("Expected lower to be 1, got %v\n", lower)
	}

	if upper != 2 {
		t.Fatalf("Expected upper to be 2, got %v\n", upper)
	}
}

func TestMinMaxEqual(t *testing.T) {
	lower, upper := minMax(3, 3)

	if lower != 3 {
		t.Fatalf("Expected lower to be 3, got %v\n", lower)
	}

	if upper != 3 {
		t.Fatalf("Expected upper to be 3, got %v\n", upper)
	}
}

func TestHasEdgeEmptyGraph(t *testing.T) {
	g := newGraph(2)

	if g.hasEdge(0, 1) {
		t.Fatal("Expected edge to be absent")
	}

	if g.hasEdge(1, 0) {
		t.Fatal("Expected edge to be absent")
	}
}

func TestAddEdge(t *testing.T) {
	g := newGraph(2)
	g.addEdge(0, 1)

	if !g.hasEdge(0, 1) {
		t.Fatal("Expected edge to exist")
	}

	if !g.hasEdge(1, 0) {
		t.Fatal("Expected edge to exist")
	}
}

func TestAddEdge2(t *testing.T) {
	g := newGraph(3)
	g.addEdge(0, 1)
	g.addEdge(0, 2)

	if !g.hasEdge(0, 1) {
		t.Fatal("Expected edge to exist")
	}

	if !g.hasEdge(0, 2) {
		t.Fatal("Expected edge to exist")
	}

	if g.hasEdge(1, 2) {
		t.Fatal("Expected edge to be absent")
	}
}

func TestCalcGraphSummary1(t *testing.T) {
	g := newGraph(3)

	summary := calcGraphSummary(g)

	if summary.numberVertices != 0 {
		t.Fatalf("Expected 0 vertices, got %v\n", summary.numberVertices)
	}

	if summary.numberEdges != 0 {
		t.Fatalf("Expected 0 edges, got %v\n", summary.numberEdges)
	}
}

func TestCalcGraphSummary2(t *testing.T) {
	g := newGraph(3)
	g.addEdge(0, 2)

	summary := calcGraphSummary(g)

	if summary.numberVertices != 2 {
		t.Fatalf("Expected 2 vertices, got %v\n", summary.numberVertices)
	}

	if summary.numberEdges != 1 {
		t.Fatalf("Expected 1 edges, got %v\n", summary.numberEdges)
	}
}

func TestCalcGraphSummary3(t *testing.T) {
	g := newGraph(3)
	g.addEdge(0, 1)
	g.addEdge(0, 2)

	summary := calcGraphSummary(g)

	if summary.numberVertices != 3 {
		t.Fatalf("Expected 3 vertices, got %v\n", summary.numberVertices)
	}

	if summary.numberEdges != 2 {
		t.Fatalf("Expected 2 edges, got %v\n", summary.numberEdges)
	}
}

func TestBuildRandomGraphNoEdges(t *testing.T) {
	g := buildRandomGraph(2, 0.0)
	summary := calcGraphSummary(g)

	if summary.numberVertices != 0 {
		t.Fatalf("Expected 0 vertices, got %v\n", summary.numberVertices)
	}

	if summary.numberEdges != 0 {
		t.Fatalf("Expected 0 edges, got %v\n", summary.numberEdges)
	}
}

func TestBuildRandomGraphAllEdges(t *testing.T) {
	g := buildRandomGraph(2, 1.0)
	summary := calcGraphSummary(g)

	if summary.numberVertices != 2 {
		t.Fatalf("Expected 2 vertices, got %v\n", summary.numberVertices)
	}

	if summary.numberEdges != 1 {
		t.Fatalf("Expected 1 edges, got %v\n", summary.numberEdges)
	}
}

func TestBuildGraphFixedNumEdges(t *testing.T) {
	g := buildRandomGraphFixedNumEdges(10, 30)
	summary := calcGraphSummary(g)

	if summary.numberEdges != 30 {
		t.Fatalf("Expected 30 edges, got %v\n", summary.numberEdges)
	}
}
