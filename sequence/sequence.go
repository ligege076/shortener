package sequence

type Sequence interface {
	Next() (seq uint64, err error)
}
