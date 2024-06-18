package state

type State uint8

const (
	StateAnywhere           State = 0
	StateCSIEntry           State = 1
	StateCSIIgnore          State = 2
	StateCSIIntermediate    State = 3
	StateCSIParam           State = 4
	StateDCSEntry           State = 5
	StateDCSIgnore          State = 6
	StateDCSIntermediate    State = 7
	StateDCSParam           State = 8
	StateDCSPassthrough     State = 9
	StateEscape             State = 10
	StateEscapeIntermediate State = 11
	StateGround             State = 12
	StateOSCString          State = 13
	StateSOSPMAPCString     State = 14
	StateUTF8               State = 15
)

func StateChange(state State, byte2 uint8) (State, Action) {
	change := stateChanges[StateAnywhere][byte2]
	if change == 0 {
		change = stateChanges[state][byte2]
	}
	return unpack(change)
}
