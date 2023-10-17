package interfaces

type Owner interface {
	ID() uint
	Own(asset Asset) bool
}

type Asset interface {
	SetOwner(owner Owner)
	Owner() Owner
}
