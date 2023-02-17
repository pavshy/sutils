package sutils

func NewLink(isInternalSSP, isInternalDSP bool) Link {
	switch {
	case isInternalSSP && isInternalDSP:
		return linkIxI
	case !isInternalSSP && isInternalDSP:
		return linkExI
	case isInternalSSP && !isInternalDSP:
		return linkIxE
	case !isInternalSSP && !isInternalDSP:
		return linkExE
	}
	return linkDisconnected
}

// Link shows type of current connection between SSP and DSP
// First parameter sets SSP type, second sets DSP type, 'I' means Internal, 'E' means External
type Link uint

const (
	linkDisconnected = iota
	linkIxI
	linkExI
	linkIxE
	linkExE
)

func (l Link) Switch(IxI, ExI, IxE, ExE func()) {
	var f func()
	switch l {
	case linkIxI:
		f = IxI
	case linkExI:
		f = ExI
	case linkIxE:
		f = IxE
	case linkExE:
		f = ExE
	}
	if f != nil {
		f()
	}
}
