package handlers

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/mem"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/rs/zerolog/log"
	"github.com/shirou/gopsutil/v4/host"
)

func AddTool() mcp.Tool {
	return mcp.NewTool("add",
		mcp.WithDescription("Add two numbers"),
		mcp.WithNumber("a",
			mcp.Required(),
			mcp.Description("first number"),
		),
		mcp.WithNumber("b",
			mcp.Required(),
			mcp.Description("second number"),
		),
	)
}
func AddHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	a, ok := request.Params.Arguments["a"].(float64)
	if !ok {
		return nil, errors.New("a must be a number")
	}

	b, ok := request.Params.Arguments["b"].(float64)
	if !ok {
		return nil, errors.New("b must be a number")
	}

	return mcp.NewToolResultText(fmt.Sprintf("%v", a+b)), nil
}

func GetHostnameTool() mcp.Tool {
	return mcp.NewTool("get_hostname",
		mcp.WithDescription("Display hostname"),
	)
}
func GetHostnameHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	hostStat, err := host.Info()
	if err != nil {
		log.Fatal().Msg("Failed to get host info")
	}
	return mcp.NewToolResultText(hostStat.Hostname), nil
}

func GetCpuUtilizationTool() mcp.Tool {
	return mcp.NewTool("get_cpu_utilization",
		mcp.WithDescription("Display current CPU utilization"),
	)
}
func GetCpuUtilizationHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	cpuStat, err := cpu.Percent(time.Second, false)
	if err != nil {
		log.Fatal().Msg("Failed to get cpu info")
	}

	return mcp.NewToolResultText(fmt.Sprintf("%.2f", cpuStat[0])), nil
}

func GetMemoryUtilizationTool() mcp.Tool {
	return mcp.NewTool("get_memory_utilization",
		mcp.WithDescription("Display current memory utilization"),
	)
}
func GetMemoryUtilizationHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	memoryStat, err := mem.VirtualMemory()
	if err != nil {
		log.Fatal().Msg("Failed to get cpu info")
	}

	return mcp.NewToolResultText(fmt.Sprintf("%.2f", memoryStat.UsedPercent)), nil
}
