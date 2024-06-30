package types

type ResourcePackResult int

const (
	ResourcePackResultSucessfullyLoaded ResourcePackResult = iota
	ResourcePackResultDeclined
	ResourcePackResultFailedDownload
	ResourcePackResultAccepted
)
