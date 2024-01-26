package mapper

type MapInput struct {
}

type Mapper interface {
    Map(m MapInput)
}
//TODO: Figure out how to get the Emit to emit to the coordinator
