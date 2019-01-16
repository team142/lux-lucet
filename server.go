package molu

//StartHealthServer starts a server with its own state that can be updated and queried safely
func StartHealthServer() *HealthServer {
	resp := &HealthServer{
		state: &SystemState{
			Ok: false,
		},
		in:    make(chan *subsystem),
		out:   make(chan SystemState),
		query: make(chan bool),
		stop:  make(chan bool),
	}
	go func() {
		run(resp)
	}()
	return resp
}

func run(channel *HealthServer) {
outer:
	for {
		select {
		case <-channel.stop:
			break outer
		case <-channel.query:
			channel.out <- *channel.state
		case item := <-channel.in:
			channel.state.update(item)
		}
	}
}

//HealthServer holds state for a server and exposes an interface for retrieving and modifying state
type HealthServer struct {
	in    chan *subsystem
	out   chan SystemState
	query chan bool
	stop  chan bool
	state *SystemState
}

//Update is a simple method for marking a subsystem as healthy
func (c *HealthServer) UpdateOk(subSystemName string) {
	c.Update(subSystemName, true, "")
}

//Update passes an update a subsystem and the overall health
func (c *HealthServer) Update(subSystemName string, ok bool, msg string) {
	s := &subsystem{
		Ok:   ok,
		Name: subSystemName,
		Msg:  msg,
	}
	c.in <- s
}

//Query retrieves the latest health of the system
func (c *HealthServer) Query() SystemState {
	go func() {
		c.query <- true
	}()
	i := <-c.out
	return i
}

func (c *HealthServer) Stop() {
	c.stop <- true
}
