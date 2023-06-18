package pipefilter

type Input interface{}

type Output interface{}

// Input --Filter--> Output
type Filter interface {
	Process(data Input) (Output, error)
}
