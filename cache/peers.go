package cache

import (
	pb "CacheStorm/cachepb"
)

// PeerPicker is the interface that must be implemented to local
// the peer that owns a specific key.
type PeerPicker interface {
	PickPeer(key string) (peer PeerGetter, ok bool)
}

// PeerGetter is the interface that must be implemented by a peer
type PeerGetter interface {
	Get(in *pb.Request, out *pb.Response) error
}
