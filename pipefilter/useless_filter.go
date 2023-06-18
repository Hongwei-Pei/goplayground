package pipefilter

import "fmt"

type UselessFilter struct{}

func (n *UselessFilter) Process(data Input) (Output, error) {
	const ctag = "UselessFilter.Process()"
	d, ok := data.(string)
	if !ok {
		return nil, fmt.Errorf("%s can not handle %+v", ctag, data)
	}
	d = d + " is useless"
	return d, nil
}
