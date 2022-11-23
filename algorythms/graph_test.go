package algorythms

import (
	"reflect"
	"testing"
)

func TestNewUGraph(t *testing.T) {
	type args struct {
		graph *WeighedUndirectedGraph
	}
	tests := []struct {
		name string
		args args
		want *UndirectedGraph
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUGraph(tt.args.graph); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUGraph() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewWGraph(t *testing.T) {
	type args struct {
		graph *UndirectedGraph
	}
	tests := []struct {
		name string
		args args
		want *WeighedUndirectedGraph
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewWGraph(tt.args.graph); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWGraph() = %v, want %v", got, tt.want)
			}
		})
	}
}
