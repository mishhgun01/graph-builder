package graph

import (
	"reflect"
	"testing"
)

func TestNewUGraph(t *testing.T) {
	type args struct {
		graph *Weighed
	}
	var tests []struct {
		name string
		args args
		want *Unweighted
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
		graph *Unweighted
	}
	var tests []struct {
		name string
		args args
		want *Weighed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewWGraph(tt.args.graph); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWGraph() = %v, want %v", got, tt.want)
			}
		})
	}
}
