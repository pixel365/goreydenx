package goreydenx

type PlatformCode string
type LaunchMode string

const (
	BaseUrl = "https://api.reyden-x.com/v1"

	Twitch   PlatformCode = "twitch"
	YouTube  PlatformCode = "youtube"
	GoodGame PlatformCode = "goodgame"
	Trovo    PlatformCode = "trovo"
	VkPlay   PlatformCode = "vkplay"
	Kick     PlatformCode = "kick"

	LaunchModeAuto   LaunchMode = "auto"
	LaunchModeDelay  LaunchMode = "delay"
	LaunchModeManual LaunchMode = "manual"
)
