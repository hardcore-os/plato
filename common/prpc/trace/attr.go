package trace

import (
	"go.opentelemetry.io/otel/attribute"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	gcodes "google.golang.org/grpc/codes"
)

const (
	// GRPCStatusCodeKey is convention for numeric status code of a gRPC request.
	GRPCStatusCodeKey = attribute.Key("rpc.grpc.status_code")
	// RPCNameKey is the name of msg transmitted or received.
	RPCNameKey = attribute.Key("name")
	// RPCMessageTypeKey is the type of msg transmitted or received.
	RPCMessageTypeKey = attribute.Key("msg.type")
	// RPCMessageIDKey is the identifier of msg transmitted or received.
	RPCMessageIDKey = attribute.Key("msg.id")
	// RPCMessageCompressedSizeKey is the compressed size of the msg transmitted or received in bytes.
	RPCMessageCompressedSizeKey = attribute.Key("msg.compressed_size")
	// RPCMessageUncompressedSizeKey is the uncompressed size of the msg
	// transmitted or received in bytes.
	RPCMessageUncompressedSizeKey = attribute.Key("msg.uncompressed_size")
	// ServerEnvironment ...
	ServerEnvironment = attribute.Key("environment")
)

// Semantic conventions for common RPC attributes.
var (
	// RPCSystemGRPC is the semantic convention for gRPC as the remoting system.
	RPCSystemGRPC = semconv.RPCSystemKey.String("grpc")
	// RPCNameMessage is the semantic convention for a msg named msg.
	RPCNameMessage = RPCNameKey.String("msg")
	// RPCMessageTypeSent is the semantic conventions for sent RPC msg types.
	RPCMessageTypeSent = RPCMessageTypeKey.String("SENT")
	// RPCMessageTypeReceived is the semantic conventions for the received RPC msg types.
	RPCMessageTypeReceived = RPCMessageTypeKey.String("RECEIVED")
)

// StatusCodeAttr returns an attribute.KeyValue that represents the give c.
func StatusCodeAttr(c gcodes.Code) attribute.KeyValue {
	return GRPCStatusCodeKey.Int64(int64(c))
}
