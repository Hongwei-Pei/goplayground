/*
Package pipefilter is a simple implementation of pipe-filter design pattern,
which is suited for data processing and analysis. It provides a simple
interface for users to implement their own filters by encapsulating data
processing method(s). The pipefilter struct is also a filter which can be
used to combine filters and transfer data between filters, or buffer data
streams in asynchronous processing.
*/

package pipefilter

var _ Filter = (*PipeFilter)(nil)

type PipeFilter struct {
	filters []Filter
}

func NewPipeFilter() *PipeFilter {
	return &PipeFilter{
		filters: make([]Filter, 0),
	}
}

func (p *PipeFilter) Process(inp Input) (out Output, err error) {
	for _, filter := range p.filters {
		out, err = filter.Process(inp)
		if err != nil {
			return out, err
		}
		inp = out
	}
	return out, nil
}

func (p *PipeFilter) RegisterFilter(filter ...Filter) {
	if p.filters == nil {
		p.filters = make([]Filter, 0)
	}
	p.filters = append(p.filters, filter...)
}
