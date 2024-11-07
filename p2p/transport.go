package p2p

// Peer: Interface that represents the remote node of the connection
type Peer interface {}

// Tranasport: Interface that represents any handles between connections
// the connection can either be TCP, UDP, Grpc, ...
type Transport interface {}