package bc

type Block struct {
	*BlockHeader
	ID          Hash
	Transaction []*Tx
}

type BlockHeader struct {
	Version         uint64
	Height          uint64
	PreviousBlockId *Hash
	Timestamp       uint64
}
