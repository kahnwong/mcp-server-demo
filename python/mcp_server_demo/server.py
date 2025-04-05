import logging as logger
import platform

import psutil  # type: ignore
from mcp.server.fastmcp import FastMCP

mcp = FastMCP("Demo")


@mcp.tool()
def add(a: int, b: int) -> int:
    """Add two numbers"""
    return a + b


@mcp.tool()
def get_hostname() -> str:
    """Display hostname"""
    return platform.node()


@mcp.tool()
def get_current_cpu_utilization() -> float:
    "Display current CPU utilization"
    return psutil.cpu_percent()


@mcp.tool()
def get_current_memory_utilization() -> float:
    "Display current memory utilization"
    return psutil.virtual_memory().percent


# entrypoint
def run() -> None:
    logger.info("MCP server is running")
    mcp.run()
