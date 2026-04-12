// Copyright (c) 2026, Peter Ohler, All rights reserved.

package fhir_test

import (
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/ohler55/slip"
)

func availablePort() int {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		panic(err)
	}
	var listener *net.TCPListener
	if listener, err = net.ListenTCP("tcp", addr); err != nil {
		panic(err)
	}
	defer func() { _ = listener.Close() }()

	return listener.Addr().(*net.TCPAddr).Port
}

func startMockServer(handler func(w http.ResponseWriter, r *http.Request)) (string, *http.Server) {
	port := availablePort()
	hs := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: http.HandlerFunc(handler),
	}
	go func() { _ = hs.ListenAndServe() }()

	su := fmt.Sprintf("http://localhost:%d", port)
	start := time.Now()
	for time.Since(start) < time.Second*2 {
		time.Sleep(time.Millisecond * 50)
		if resp, err := http.Get(su); err == nil {
			_ = resp.Body.Close()
			break
		}
	}
	return su, &hs
}

func classMemberP(list slip.List, name string) bool {
	for _, v := range list {
		if class, _ := v.(slip.Class); class != nil && strings.EqualFold(class.Name(), name) {
			return true
		}
	}
	return false
}
