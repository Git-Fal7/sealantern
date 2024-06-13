package player

import (
	"sync"

	"github.com/git-fal7/sealantern/minecraft/player/connplayer"

	"github.com/google/uuid"
)

type PlayerRegistry struct {
	playersCount int
	playersMutex *sync.RWMutex
	playerIDs    map[uuid.UUID]*connplayer.ConnectedPlayer // uuids map
}

func NewPlayerRegistry() *PlayerRegistry {
	return &PlayerRegistry{
		playersCount: 0,
		playerIDs:    map[uuid.UUID]*connplayer.ConnectedPlayer{},
		playersMutex: &sync.RWMutex{},
	}
}

// TODO: Make the Core kick already existing players with same username
func (r *PlayerRegistry) RegisterPlayer(p *connplayer.ConnectedPlayer) bool {
	r.playersMutex.Lock()

	if _, ok := r.playerIDs[p.UUID()]; ok {
		return false
	}

	r.playerIDs[p.UUID()] = p
	r.playersMutex.Unlock()
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
	player, ok := r.playerIDs[uuid]
	r.playersMutex.RUnlock()
	if ok {
		return player
	}
	return nil
}

func (r *PlayerRegistry) GetPlayerFromEID(eid uint16) *connplayer.ConnectedPlayer {
	for _, player := range r.GetPlayers() {
		if player.ID() == eid {
			return player
		}
	}
	return nil
}

func (r *PlayerRegistry) GetPlayers() []*connplayer.ConnectedPlayer {
	r.playersMutex.RLock()
	players := make([]*connplayer.ConnectedPlayer, 0, len(r.playerIDs))
	for _, player := range r.playerIDs {
		players = append(players, player)
	}
	r.playersMutex.RUnlock()
	return players
}
