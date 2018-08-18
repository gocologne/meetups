package statemachine

type stateMachine struct {
	state   string
	actionc chan func()
	quitc   chan struct{}
}

func (sm *stateMachine) loop() {
	for {
		select {
		case f := <-sm.actionc:
			f()
		case <-sm.quitc:
			return
		}
	}
}

func (sm *stateMachine) foo() int {
	c := make(chan int)
	sm.actionc <- func() {
		if sm.state == "A" {
			sm.state = "B"
		}
		c <- 123
	}
	return <-c
}

func New() *stateMachine {
	sm := &stateMachine{
		state:   "initial",
		actionc: make(chan func()),
		quitc:   make(chan struct{}),
	}
	go sm.loop()
	return sm
}
