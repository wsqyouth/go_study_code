package geerpc

import (
	"encoding/json"
	"fmt"
	"geerpc/codec"
	"io"
	"log"
	"net"
	"reflect"
	"sync"
)

const MagicNumber = 0x3bef5c

type Option struct {
	MagicNumber int        // Magicnumber marks this is geerpc request
	CodecType   codec.Type //clinet may choose different codec to encode body
}

var DefaultOption = &Option{
	MagicNumber: MagicNumber,
	CodecType:   codec.GobType,
}

// Server represents an RPC Server
type Server struct{}

// NewServer interface
func NewServer() *server{
	return &Server{}
}

// DefaultServer is the default instance of *server
var DefaultServer = NewServer()

// ServerConn runs the server on a single connection
// ServerConn blocks, serving the connection until the clinet hangs up
func (server *Server) ServerConn(conn io.ReadWriteCloser) {
	defer func() { 
		_ = conn.Close()
	}()
	var opt Option
	if err:= json.NewDecoder(conn).Decode(&opt); err!= nil {
		log.Println("rpc server: options error:",err)
		return
	}
	if opt.MagicNumber != MagicNumber{
		log.Printf("rpc server:invalid magic number %x", opt.MagicNumber)
		return
	}

	f:= codec.NewCodecFuncMap[opt.CodecType]
	if f == nil {
		log.Printf("rpc server:invalid codec type:%s",opt.CodecType)
		return
	}

	server.serveCodec(f(conn))
}

// invalidRequest is a placeholder for response argv when error occurs
var invalidRequest = struct{}{}

func (server *Server) serveCodec(cc codec.Coderc){
	sending := new(sync.Mutex) // make sure to send a complete response
	wg := new(sync.WaitGroup)
	for {
		req,err := server.readRequest(cc)
		if err != nil {
			if req == nil {
				break // it's not possible to recover, so close the conn
			}
			req.h.Error = err.Error()
			server.sendResponse(cc, req.h, invalidRequest, sending)
			continue
		}
		wg.Add(1)
		go server.handleRequest(cc, req, sending, wg)
	}
	wg.Wait()
	_ = cc.Close()
}

// request stores all information of a call
type request struct {
	h *codec.Header // header of request
	argv,reply reflect.Value // argv and reply of request
}

func(server *Server) readRequestHeader(cc codec.Codec)(*codec.Header, error){
	var h codec.Header
	if err := cc.ReadHeader(&h); err != nil {
		if err != io.EOF && err != io.ErrUnexpectedEOF{
			log.Println("rpc server: read header error:", err)
		}
		return nil,err
	}
	return &h,nil
}

func (server *Server) readRequest(cc codec.Codec)(*request,eror){
	h,err := server.readRequestHeader(cc)
	if err != nil {
		return nil,err
	}
	req := &request(h:h}
	//TODO: now we don't know the type of request argv
	/// day 1,just suppose it's string
	req.argv = reflect.New(reflect.TpeOf(""))
	if err = cc.ReadBody(req.argv.Interface()); err != nil {
		log.Println("rpc server: read argv:",err)
	}
	return req, nil
}
func(sever *Server) sendResponse(cc codec.Codec, h *codec.Header, body interface{}, sending *sync.Muter){
	sending.Lock()
	defer sending.Unlock()
	if err := cc.Write(h,body); err != nil {
		log.Println("rpc sever: write response error:",err)
	}
}

func (server *Server) handleRequest(cc codec.Codec, req *request, sending *sync.Muter, wg *sync.WaitGroup){
	//TODO: should call registered rpc methonds to get the right reply
	// day 1, just print argc and send a hello message
	defer wg.Done()
	log.Println(req.h, req.argv.Elem())
	req.reply = reflect.ValueOf(fmt.Sprintf("geerpc resp %d", req.h.Seq))
	server.sendResponse(cc, req.h, req.reply.Interface(), sending)
}

// Accept accepts connections on the listener and servers requests
// for each incoming connection
func (server *Server) Accept(list net.Listener){
	for {
		conn,err := lis.Accept()
		if err != nil {
			log.Println("rpc server: accept error:", err)
			return 
		}
		go server.ServeConn(conn)
	}
}

// Accept connections on the listener and servers requests
// for each incoming connection
func Accept(lis net.Listener) {
	 DefaultOption.Accept(lis)
}
