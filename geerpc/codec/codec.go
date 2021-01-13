package codec

import(
	"io"
)

type Header struce {
	ServiceMethod string // fromat "Service.Method"
	Seq uint64 // sequence number chosen by client
	Error string
}

type Codec interface {
	io.Closer
	ReadHeader(*Header) error
	ReadBody(interface{}) error
	Write(*Header, interface{}) error
}

type NewCodecFunc func(io.ReadWriterCloser) Codec

type Type string

const (
	GlobType Type = "application/gob"
	JsonType Type = "application/json" // not implemneted
)

var NewCodeFuncMap map[Type]NewCodecFunc

func init() {
	NewCodeFuncMap = make(map[Type]NewCodecFunc)
	NewCodeFuncMap[GobType] = NewGobCodec
}
