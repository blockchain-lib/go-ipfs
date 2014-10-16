// Code generated by protoc-gen-gogo.
// source: semver.proto
// DO NOT EDIT!

/*
Package handshake is a generated protocol buffer package.

It is generated from these files:
	semver.proto

It has these top-level messages:
	Handshake1
*/
package handshake

import proto "github.com/jbenet/go-ipfs/Godeps/_workspace/src/code.google.com/p/gogoprotobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type Handshake1 struct {
	// protocolVersion determines compatibility between peers
	ProtocolVersion *string `protobuf:"bytes,1,opt,name=protocolVersion" json:"protocolVersion,omitempty"`
	// agentVersion is like a UserAgent string in browsers, or client version in bittorrent
	// includes the client name and client. e.g.   "go-ipfs/0.1.0"
	AgentVersion     *string `protobuf:"bytes,2,opt,name=agentVersion" json:"agentVersion,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Handshake1) Reset()         { *m = Handshake1{} }
func (m *Handshake1) String() string { return proto.CompactTextString(m) }
func (*Handshake1) ProtoMessage()    {}

func (m *Handshake1) GetProtocolVersion() string {
	if m != nil && m.ProtocolVersion != nil {
		return *m.ProtocolVersion
	}
	return ""
}

func (m *Handshake1) GetAgentVersion() string {
	if m != nil && m.AgentVersion != nil {
		return *m.AgentVersion
	}
	return ""
}

func init() {
}
