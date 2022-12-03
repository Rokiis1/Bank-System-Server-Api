package main

type APIServer struct {
	listenAddr string
}

func newAPIServer(listenAddr string) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
	}
}