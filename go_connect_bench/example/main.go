/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a server for Greeter service.
package main

import (
	"context"
	"github.com/bufbuild/connect-go"
	_ "go.uber.org/automaxprocs"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	helloworld "local/proto/helloworld"
	helloworldconnect "local/proto/helloworld/helloworldconnect"
	"log"
	"net/http"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	helloworldconnect.UnimplementedGreeterHandler
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *connect.Request[helloworld.HelloRequest]) (*connect.Response[helloworld.HelloReply], error) {
	return connect.NewResponse(&helloworld.HelloReply{Response: in.Msg.GetRequest()}), nil
}

func main() {
	mux := http.NewServeMux()
	mux.Handle(helloworldconnect.NewGreeterHandler(&server{}))
	err := http.ListenAndServe(
		port,
		h2c.NewHandler(mux, &http2.Server{}),
	)
	log.Fatalf("listen failed: %v", err)
}
