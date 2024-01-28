package reduce

type ReduceInput struct {
}
//This is the interface that the client written reducer should adhere to.
type Reducer interface {
    Reduce(r ReduceInput)
}
