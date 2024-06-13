package types

type TeamMode int8

const (
	TeamModeCreate TeamMode = iota
	TeamModeRemove
	TeamModeUpdate
	TeamModeAddPlayer
	TeamModeRemovePlayer
)

type TeamFriendlyFire int8

const (
	TeamFriendlyFireOff TeamFriendlyFire = iota
	TeamFriendlyFireOn
	TeamFriendlyFireSeeInvisibles TeamFriendlyFire = 3
)

type TeamNameTagVisibility string

const (
	TeamNameTagVisibilityAlways           TeamNameTagVisibility = "always"
	TeamNameTagVisibilityHideForOtherTeam TeamNameTagVisibility = "hideForOtherTeam"
	TeamNameTagVisibilityHideForOwnTeam   TeamNameTagVisibility = "hideForOwnTeam"
	TeamNameTagVisibilityNever            TeamNameTagVisibility = "never"
)
