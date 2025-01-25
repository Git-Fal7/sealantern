package player

import (
	"sync"

	"github.com/git-fal7/sealantern/minecraft/player/connplayer"

	"github.com/google/uuid"
)

type PlayerRegistry struct {
	playersMutex sync.RWMutex
	playerIDs    map[uuid.UUID]*connplayer.ConnectedPlayer // uuids map
}

func NewPlayerRegistry() *PlayerRegistry {
	return &PlayerRegistry{
		playerIDs: map[uuid.UUID]*connplayer.ConnectedPlayer{},
	}
}

func (r *PlayerRegistry) RegisterPlayer(p *connplayer.ConnectedPlayer) bool {
	r.playersMutex.Lock()
	defer r.playersMutex.Unlock()

	if _, ok := r.playerIDs[p.UUID()]; ok {
		return false
	}
	r.playerIDs[p.UUID()] = p
	return true
}

func (r *PlayerRegistry) UnregisterPlayer(p *connplayer.ConnectedPlayer) bool {
	if p == nil {
		return false
	}
	r.playersMutex.Lock()
	defer r.playersMutex.Unlock()
	_, ok := r.playerIDs[p.UUID()]
	delete(r.playerIDs, p.UUID())
	return ok
}

func (r *PlayerRegistry) GetPlayerFromUUID(uuid uuid.UUID) *connplayer.ConnectedPlayer {
	r.playersMutex.RLock()
	defer r.playersMutex.RUnlock()
	player, ok := r.playerIDs[uuid]
	if ok {
		return player
	}
	return nil
}

func (r *PlayerRegistry) GetPlayerFromEID(eid int32) *connplayer.ConnectedPlayer {
	for _, player := range r.GetPlayers() {
		if player.ID() == eid {
			return player
		}
	}
	return nil
}

func (r *PlayerRegistry) GetPlayers() []*connplayer.ConnectedPlayer {
	r.playersMutex.RLock()
	defer r.playersMutex.RUnlock()
	players := make([]*connplayer.ConnectedPlayer, 0, len(r.playerIDs))
	for _, player := range r.playerIDs {
		if player == nil {
			continue
		}
		players = append(players, player)
	}
	return players
}
