package p2p

// Peer is anything that represents the remote node
type Peer interface{}

// Transport is anything that handles the communication
// between the node in the network.
// form (TCP, UDP, websockets, ...)
type Transport interface{}