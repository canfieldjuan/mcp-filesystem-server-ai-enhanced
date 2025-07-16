package filesystemserver

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/mark3labs/mcp-filesystem-server/intelligence"
	"github.com/mark3labs/mcp-filesystem-server/filesystemserver/handler"
)

// EnhancedServer wraps the basic filesystem server with AI intelligence
type EnhancedServer struct {
	*server.MCPServer
	intelligence *intelligence.IntelligenceEngine
	handler      *handler.FilesystemHandler
}

// NewEnhancedFilesystemServer creates a new enhanced filesystem server with AI capabilities
func NewEnhancedFilesystemServer(allowedDirs []string) (*EnhancedServer, error) {
	// Create the base handler
	h, err := handler.NewFilesystemHandler(allowedDirs)
	if err != nil {
		return nil, err
	}

	// Create the base MCP server
	baseServer := server.NewMCPServer(
		"enhanced-filesystem-server",
		Version,
		server.WithResourceCapabilities(true, true),
	)
	
	// Initialize intelligence engine
	intelligenceConfig := intelligence.IntelligenceConfig{
		ContentAnalysis:   true,
		SemanticSearch:    true,
		PredictiveCache:   false, // Disabled for now
		SmartOperations:   false, // Disabled for now
		AnalysisDepth:     2,
		CacheSize:         1000,
		VectorDimensions:  384,
	}
	
	engine, err := intelligence.NewIntelligenceEngine(intelligenceConfig)
	if err != nil {
		return nil, err
	}
	
	enhanced := &EnhancedServer{
		MCPServer:    baseServer,
		intelligence: engine,
		handler:      h,
	}
	
	// Add all the standard filesystem tools
	enhanced.addStandardTools()
	
	// Add AI-enhanced tools
	enhanced.addIntelligenceTools()
	
	return enhanced, nil
}

// addStandardTools adds the standard filesystem tools (same as the base server)
func (es *EnhancedServer) addStandardTools() {
	// Register resource handlers (same as base server)
	es.AddResource(mcp.NewResource(
		"file://",
		"File System",
		mcp.WithResourceDescription("Access to files and directories on the local file system"),
	), es.handler.HandleReadResource)

	// List directory contents
	es.AddTool(mcp.NewTool(
		"list_directory",
		mcp.WithDescription("List contents of a directory"),
		mcp.WithString("path",
			mcp.Description("Path to the directory to list"),
			mcp.Required(),
		),
	), es.handler.HandleListDirectory)

	// Read file contents
	es.AddTool(mcp.NewTool(
		"read_file",
		mcp.WithDescription("Read the complete contents of a file"),
		mcp.WithString("path",
			mcp.Description("Path to the file to read"),
			mcp.Required(),
		),
	), es.handler.HandleReadFile)

	// Write file contents
	es.AddTool(mcp.NewTool(
		"write_file",
		mcp.WithDescription("Create a new file or completely overwrite an existing file with new content"),
		mcp.WithString("path",
			mcp.Description("Path where the file should be created or overwritten"),
			mcp.Required(),
		),
		mcp.WithString("content",
			mcp.Description("Complete content to write to the file"),
			mcp.Required(),
		),
	), es.handler.HandleWriteFile)

	// Search files
	es.AddTool(mcp.NewTool(
		"search_files",
		mcp.WithDescription("Search for files and directories"),
		mcp.WithString("path",
			mcp.Description("Starting directory for the search"),
			mcp.Required(),
		),
		mcp.WithString("pattern",
			mcp.Description("Search pattern or filename to look for"),
			mcp.Required(),
		),
	), es.handler.HandleSearchFiles)
}

// addIntelligenceTools adds AI-enhanced tools to the server
func (es *EnhancedServer) addIntelligenceTools() {
	// Add semantic search tool
	es.AddTool(mcp.NewTool(
		"semantic_search",
		mcp.WithDescription("Search files using AI-powered semantic understanding"),
		mcp.WithString("query",
			mcp.Description("Natural language search query"),
			mcp.Required(),
		),
		mcp.WithString("path",
			mcp.Description("Root path to search in (optional)"),
		),
		mcp.WithNumber("limit",
			mcp.Description("Maximum number of results (default: 10)"),
		),
	), es.handleSemanticSearch)
	
	// Add content analysis tool
	es.AddTool(mcp.NewTool(
		"analyze_content",
		mcp.WithDescription("Perform AI analysis of file content"),
		mcp.WithString("path",
			mcp.Description("Path to the file to analyze"),
			mcp.Required(),
		),
	), es.handleContentAnalysis)
}

// Placeholder handlers - these would be implemented with actual logic
func (es *EnhancedServer) handleSemanticSearch(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.NewTextContent("Semantic search functionality is available but not yet fully implemented"),
		},
	}, nil
}

func (es *EnhancedServer) handleContentAnalysis(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.NewTextContent("Content analysis functionality is available but not yet fully implemented"),
		},
	}, nil
}
