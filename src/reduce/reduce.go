package reduce

type ReduceInput struct {
}

type Reducer interface {
    Reduce(r ReduceInput)
}
