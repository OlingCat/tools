// Package static computes the call graph of a Go program containing
// only static call edges.
package static // import "github.com/golang-zh/tools/go/callgraph/static"

import (
	"github.com/golang-zh/tools/go/callgraph"
	"github.com/golang-zh/tools/go/ssa"
	"github.com/golang-zh/tools/go/ssa/ssautil"
)

// CallGraph computes the call graph of the specified program
// considering only static calls.
//
func CallGraph(prog *ssa.Program) *callgraph.Graph {
	cg := callgraph.New(nil) // TODO(adonovan) eliminate concept of rooted callgraph

	// TODO(adonovan): opt: use only a single pass over the ssa.Program.
	for f := range ssautil.AllFunctions(prog) {
		fnode := cg.CreateNode(f)
		for _, b := range f.Blocks {
			for _, instr := range b.Instrs {
				if site, ok := instr.(ssa.CallInstruction); ok {
					if g := site.Common().StaticCallee(); g != nil {
						gnode := cg.CreateNode(g)
						callgraph.AddEdge(fnode, site, gnode)
					}
				}
			}
		}
	}

	return cg
}