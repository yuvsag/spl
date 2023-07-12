package packet

import "errors"

var ErrNotImplemented = errors.New("pkg packet: method not implemented")
var ErrBytesOverflow = errors.New("pkg packet: field does not fit in the length")
var ErrNotAnOption = errors.New("pkg packet: value given is not an available option")
var ErrIncorrectType = errors.New("pkg packet: value given is of an unexpected type")
var ErrMissingField = errors.New("pkg packet: the packet has a missing field")
var ErrProtocolNotFound = errors.New("pkg packet: no protocol found for this packet")
