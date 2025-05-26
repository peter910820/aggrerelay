package base

var (
	Ch                  = make(chan int)
	LineShutdownChan    = make(chan struct{})
	DiscordShutdownChan = make(chan struct{})
)
