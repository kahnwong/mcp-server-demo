package main

import (
	"fmt"

	"github.com/kahnwong/mcp-server-demo-go/handlers"
	"github.com/mark3labs/mcp-go/server"
	"github.com/rs/zerolog/log"
)

var s *server.MCPServer

func main() {
	// init
	log.Info().Msg("Starting server")
	s = server.NewMCPServer(
		"Demo",
		"1.0.0",
	)

	// tools
	s.AddTool(handlers.GetHostnameTool(), handlers.GetHostnameHandler)

	// start server
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
