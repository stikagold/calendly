package interfaces

type Jsonable interface {
	ToJson() ([]byte, error)
}
