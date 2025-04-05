package handlers

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/rs/zerolog/log"
	"github.com/shirou/gopsutil/v4/host"
)

func GetHostnameHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	hostStat, err := host.Info()
	if err != nil {
		log.Fatal().Msg("Failed to get host info")
	}
	return mcp.NewToolResultText(hostStat.Hostname), nil
}
