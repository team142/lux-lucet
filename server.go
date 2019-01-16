package mortis_lux

//StartHealthServer starts a server with its own state that can be updated and queried safely
func StartHealthServer() *HealthServer {
	resp := &HealthServer{}
	go func() {
		run(resp)
	}()
	return resp
}

func run(channel *HealthServer) {
	for {
		select {
		case <-channel.stop:
			break
		case <-channel.query:
			channel.out <- *channel.state
		case item := <-channel.in:
			channel.state.update(item)
		}
	}
}

type HealthServer struct {
	in    chan *subsystem
	out   chan SystemState
	query chan bool
	stop  chan bool
	state *SystemState
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
