package types

type DiggingStatus uint8

const (
	StartedDigging DiggingStatus = iota
	CancelledDigging
	FinishedDigging
	DropItemStack
	DropItem
	ShootArrowOrFinishEating
)
