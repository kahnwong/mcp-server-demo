use log::info;
use mcpr::error::MCPError;
use mcpr::schema::ToolInputSchema;
use mcpr::{
    Tool,
    server::{Server, ServerConfig},
    transport::stdio::StdioTransport,
};
use serde_json::Value;
use std::error::Error;

fn main() -> Result<(), Box<dyn Error>> {
    // Configure the server
    let server_config = ServerConfig::new()
        .with_name("My MCP Server")
        .with_version("1.0.0")
        .with_tool(Tool {
            name: "hello".to_string(),
            description: Some("A simple hello world tool".to_string()),
            input_schema: ToolInputSchema {
                r#type: "".to_string(),
                properties: None,
                required: None,
            },
        });
    // Create the server
    let mut server = Server::new(server_config);

    // Register tool handlers
    server.register_tool_handler("hello", |params: Value| {
        // Parse parameters
        let name = params
            .get("name")
            .and_then(|v| v.as_str())
            .ok_or_else(|| MCPError::Protocol("Missing name parameter".to_string()))?;

        info!("Handling hello tool call for name: {}", name);

        // Generate response
        let response = serde_json::json!({
            "message": format!("Hello, {}!", name)
        });

        Ok(response)
    })?;

    // Start the server with stdio transport
    let transport = StdioTransport::new();
    server.start(transport)?;

    Ok(())
}
