package model

// init channel
var (
	RelayChan           = make(chan int)
	LineShutdownChan    = make(chan struct{})
	DiscordShutdownChan = make(chan struct{})
)
