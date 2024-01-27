package mapper

type MapInput struct {
}
//This is the interface intended to shape the function that client written Map functions should adhere to.
type Mapper interface {
    Map(m MapInput)
}
//TODO: Figure out how to get the Emit to emit to the coordinator
