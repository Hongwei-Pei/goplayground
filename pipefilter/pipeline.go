package pipefilter

import (
	"fmt"
)

const (
	channalBuffSize = 1
)

type PipeLine struct {
	InpChan  *chan interface{}
	Filters  []Filter
	Channals []chan interface{}
	OutChan  *chan interface{}
}

func NewPipeLine(inp *chan interface{}, out *chan interface{}) *PipeLine {
	p := &PipeLine{
		Filters:  make([]Filter, 0),
		Channals: make([]chan interface{}, 0),
	}
	p.init(inp, out)
	return p
}

func (p *PipeLine) RegisterFilter(filter ...Filter) {
	if p.Filters == nil {
		p.Filters = make([]Filter, 0)
	}
	if p.Channals == nil {
		p.Channals = make([]chan interface{}, 0)
	}
	if len(p.Filters) == 0 {
		p.Filters = append(p.Filters, filter...)
	} else {
		p.Filters = append(p.Filters, filter...)
		p.Channals = append(p.Channals, make(chan interface{}, channalBuffSize))
	}
}

func (p *PipeLine) Run() {
	p.RunWithBuffChan()
}

func (p *PipeLine) init(inp *chan interface{}, out *chan interface{}) error {
	if *inp == nil || *out == nil {
		return fmt.Errorf("channals is nil")
	}
	p.InpChan = inp
	p.OutChan = out
	if p.Channals == nil || len(p.Channals) == 0 {
		return fmt.Errorf("channals is nil")
	}
	return nil
}

func (p *PipeLine) ProcessWithFilter(filter Filter, inc *chan interface{},
	outc *chan interface{}) {
	data, ok := <-*inc
	if ok {
		out, err := filter.Process(data)
		if err != nil {
			fmt.Println(err)
		}
		*outc <- out
	} else {
		fmt.Println("input chan is closed")
	}
}

func (p *PipeLine) RunWithBuffChan() {
	for {
		var data interface{}
		var ok bool
		for i, filter := range p.Filters {
			if i == 0 {
				data, ok = <-*p.InpChan
			} else {
				data, ok = <-p.Channals[i-1]
			}
			if !ok {
				fmt.Println("input chan is closed")
			}
			out, err := filter.Process(data)
			if err != nil {
				fmt.Println(err)
				continue
			}
			if i == len(p.Filters)-1 {
				*p.OutChan <- out
			} else {
				p.Channals[i] <- out
			}
		}
	}
}
