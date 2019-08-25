package nna

// NNA - Nerv network address
// NNA v1 format:
// protocol-version / NodeType / ID / ConnectionType / ConnectionProtocolType / Address

func ParseString(data string) (result NNA, err error) {
	return result, err
}

type NNA struct {
	Version int
	Type    NodeType
	ID      string
	Ctype   ConnectionType
	CPtype  ConnectionProtocolType
	Address string
}

type NodeType struct {
	Byte   byte
	String string
}

type ConnectionType string

type ConnectionProtocolType string

func (addr NNA) Network() string {
	return string(addr.CPtype)
}

func (addr NNA) String() string {
	return addr.Address
}

type NNAFormat struct {
}
