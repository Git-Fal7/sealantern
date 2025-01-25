package entitymanager

import "sync/atomic"

var eid atomic.Int32

func NextEID() int32 {
	eid.Add(1)
	return eid.Load()
}