package tick

import (
	"github.com/influxdata/kapacitor/pipeline"
	"github.com/influxdata/kapacitor/tick/ast"
)

// From converts the from pipeline node into the TICKScript AST
type From struct {
	Function
}

// NewFrom creates a Flatten function builder
func NewFrom(parents []ast.Node) *From {
	return &From{
		Function{
			Parents: parents,
		},
	}
}

// Build creates a from ast.Node
func (n *From) Build(f *pipeline.FromNode) (ast.Node, error) {
	n.Pipe("from").
		Dot("database", f.Database).
		Dot("retentionPolicy", f.RetentionPolicy).
		Dot("measurement", f.Measurement).
		DotIf("groupByMeasurement", f.GroupByMeasurementFlag).
		Dot("round", f.Round).
		Dot("truncate", f.Truncate).
		Dot("where", f.Lambda)
	if len(f.Dimensions) > 0 {
		n.Dot("groupBy", f.Dimensions...)
	}
	return n.prev, n.err
}
