package graph

//
//import (
//	"reflect"
//	"testing"
//)
//
//func TestConvWeightedToUnweighted(t *testing.T) {
//	type args struct {
//		graph *Abstract
//	}
//	var tests []struct {
//		name string
//		args args
//		want *Unweighted
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := ConvWeightedToUnweighted(tt.args.graph); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("ConvWeightedToUnweighted() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestUnweightedToWeighted(t *testing.T) {
//	type args struct {
//		graph *Unweighted
//	}
//	var tests []struct {
//		name string
//		args args
//		want *Abstract
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := ConvUnweightedToWeighted(tt.args.graph); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("ConvUnweightedToWeighted() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
