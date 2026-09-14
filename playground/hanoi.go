package main

// TODO: Continue

type CannotAcceptErr struct {
	s string
}

func (err *CannotAcceptErr) Error() string {
	return err.s
}

type PegMember int
type Peg struct {
	members []PegMember
}

func (p Peg) receiveMember(newMember PegMember) error {
	if p.canAcceptMember(newMember) {
		p.members = append(p.members, newMember)
		return nil
	} else {
		return &CannotAcceptErr{"CANNOT_ACCEPT"}
	}
}

func (p Peg) canAcceptMember(newMember PegMember) bool {
	return p.members[len(p.members)-1] < newMember
}

func (p Peg) isEmpty() bool {
	return len(p.members) == 0
}

func movesNeededForTotalTransfer(pegSize, numberOfPegs int) int {
	movesNeeded := 0
	var pegs []Peg
	// load pegZero with pegSize sorted members
	for i := 0; i < pegSize; i++ {
		pegs[0].members = append(pegs[0].members, PegMember(i))
	}

	for !isFullyTransferred(pegs) {
	}

	return movesNeeded
}

// All the members should be on one peg
// It shouldn't be the peg[0]
// Only one none-empty peg should exist
func isFullyTransferred(pegs []Peg) bool {
	countNoneEmpties := 0
	for i := 0; i < len(pegs); i++ {
		if !pegs[i].isEmpty() {
			countNoneEmpties += 1
		}
	}

	return countNoneEmpties == 1
}
