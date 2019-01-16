package molu

type SystemState struct {
	Ok         bool         `json:"ok"`
	Subsystems []*subsystem `json:"subsystems"`
}

func (systemState *SystemState) update(s *subsystem) {
	found := false
	var index int
	for i, item := range systemState.Subsystems {
		if item.Name == s.Name {
			found = true
			index = i
			break
		}
	}
	if found == true {
		systemState.Subsystems[index] = s
	} else {
		systemState.Subsystems = append(systemState.Subsystems, s)
	}
	systemState.check()
}

func (systemState *SystemState) check() {
	ok := true
	for _, s := range systemState.Subsystems {
		if !s.Ok {
			ok = false
		}
	}
	systemState.Ok = ok
}

type subsystem struct {
	Name string `json:"name"`
	Ok   bool   `json:"ok"`
	Msg  string `json:"msg"`
}
