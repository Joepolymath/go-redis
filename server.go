package main

import (
	"log"
	"net"
)

const (
	defaultListener = ":5324"
)

type Config struct {
	ListenAddr	string
}

type Server struct {
	Config
	peers map[*Peer]bool
	listener	net.Listener
	addPeerCh chan *Peer
}

func NewServer(cfg Config) Server {
	if len(cfg.ListenAddr) == 0 {
		cfg.ListenAddr = defaultListener
	}
	return Server{
		Config: cfg,
		peers: make(map[*Peer]bool),
		addPeerCh: make(chan *Peer),
	}
}

func (s *Server) Start() error {
	listener, err := net.Listen("tcp", s.ListenAddr)
	if err != nil {
		return err
	}
	s.listener = listener
	return s.acceptLoop()
}

func (s *Server) loop() {
	for {
		select {
		case peer := <- s.addPeerCh:
			s.peers[peer] = true
		default:
			log.Println("Error in adding peers")
		}
	}
}

func (s * Server) acceptLoop() error {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			log.Println("accept error", "err", err)
			continue
		}

		go s.handleConn(conn)
		return nil
	}
}

func (s *Server) handleConn(conn net.Conn) {

}