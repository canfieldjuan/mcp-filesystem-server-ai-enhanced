package intelligence

import (
	"context"
	"fmt"
	"time"
	
	"github.com/mark3labs/mcp-go/mcp"
)

// IntelligenceMiddleware wraps existing handlers with AI capabilities
type IntelligenceMiddleware struct {
	engine   *IntelligenceEngine
	enabled  bool
	config   MiddlewareConfig
}

// MiddlewareConfig configuration for intelligence middleware
type MiddlewareConfig struct {
	EnableContentAnalysis bool
	EnableSemanticSearch  bool
	EnablePredictiveCache bool
	EnableSmartSuggestions bool
	AsyncProcessing       bool
	CacheResults          bool
}

// IntelligentReadFileHandler enhances read_file with AI analysis
type IntelligentReadFileHandler struct {
	middleware *IntelligenceMiddleware
}

// IntelligentSearchHandler enhances search with semantic capabilities  
type IntelligentSearchHandler struct {
	middleware *IntelligenceMiddleware
}

// NewIntelligenceMiddleware creates new intelligence middleware
func NewIntelligenceMiddleware(engine *IntelligenceEngine, config MiddlewareConfig) *IntelligenceMiddleware {
	return &IntelligenceMiddleware{
		engine:  engine,
		enabled: true,
		config:  config,
	}
}

// WrapReadFileHandler wraps the read_file handler with intelligence
func (im *IntelligenceMiddleware) WrapReadFileHandler(baseHandler func(context.Context, mcp.CallToolRequest) (*mcp.CallToolResult, error)) func(context.Context, mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		// Execute original handler first
		result, err := baseHandler(ctx, request)
		if err != nil {
			return result, err
		}
		
		// Skip enhancement if disabled
		if !im.enabled || !im.config.EnableContentAnalysis {
			return result, nil
		}
		
		// Extract path and content for analysis
		path, _ := request.RequireString("path")
		
		// Create operation context
		opCtx := &OperationContext{
			Operation:   "read_file",
			Path:        path,
			RequestID:   generateRequestID(),
			Timestamp:   time.Now(),
			Metadata:    make(map[string]interface{}),
		}
		
		// Extract content from result for analysis
		if content, ok := im.extractContentFromResult(result); ok {
			opCtx.Content = content
			
			// Process with intelligence engine (async if configured)
			if im.config.AsyncProcessing {
				go im.processAsync(ctx, opCtx)
			} else {
				enhanced, err := im.engine.Process(ctx, opCtx)
				if err == nil && enhanced != nil {
					// Enhance the result with AI insights
					result = im.enhanceReadFileResult(result, enhanced)
				}
			}
		}
		
		return result, nil
	}
}

// WrapSearchHandler wraps search handlers with semantic search
func (im *IntelligenceMiddleware) WrapSearchHandler(baseHandler func(context.Context, mcp.CallToolRequest) (*mcp.CallToolResult, error)) func(context.Context, mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		// Check if semantic search should be used
		if im.enabled && im.config.EnableSemanticSearch {
			// Try semantic search first
			if semanticResult, err := im.performSemanticSearch(ctx, request); err == nil {
				return semanticResult, nil
			}
		}
		
		// Fallback to original handler
		result, err := baseHandler(ctx, request)
		if err != nil {
			return result, err
		}
		
		// Enhance results with AI insights
		if im.enabled && im.config.EnableSmartSuggestions {
			enhanced := im.enhanceSearchResults(ctx, result, request)
			if enhanced != nil {
				return enhanced, nil
			}
		}
		
		return result, nil
	}
}

// WrapWriteFileHandler wraps write_file with content analysis and indexing
func (im *IntelligenceMiddleware) WrapWriteFileHandler(baseHandler func(context.Context, mcp.CallToolRequest) (*mcp.CallToolResult, error)) func(context.Context, mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		// Execute original handler first
		result, err := baseHandler(ctx, request)
		if err != nil {
			return result, err
		}
		
		// Index content for semantic search (async)
		if im.enabled && im.config.EnableSemanticSearch {
			go im.indexContentAsync(ctx, request)
		}
		
		return result, nil
	}
}

// extractContentFromResult extracts content bytes from MCP result
func (im *IntelligenceMiddleware) extractContentFromResult(result *mcp.CallToolResult) ([]byte, bool) {
	if result == nil || len(result.Content) == 0 {
		return nil, false
	}
	
	// Look for text content
	for _, content := range result.Content {
		if textContent, ok := content.(mcp.TextContent); ok {
			return []byte(textContent.Text), true
		}
	}
	
	return nil, false
}

// enhanceReadFileResult adds AI analysis to read_file results
func (im *IntelligenceMiddleware) enhanceReadFileResult(original *mcp.CallToolResult, enhanced *EnhancedResult) *mcp.CallToolResult {
	if enhanced.Analysis == nil {
		return original
	}
	
	// Create enhanced result with AI insights
	enhancedResult := &mcp.CallToolResult{
		Content: append(original.Content, 
			mcp.TextContent{
				Type: "text",
				Text: fmt.Sprintf("\n--- AI Analysis ---\nContent Type: %s\nLanguage: %s\nSummary: %s\nComplexity: %.2f\nKey Topics: %v\n",
					enhanced.Analysis.ContentType,
					enhanced.Analysis.Language, 
					enhanced.Analysis.Summary,
					enhanced.Analysis.Complexity,
					enhanced.Analysis.KeyTopics,
				),
			},
		),
		IsError: original.IsError,
	}
	
	// Add security warnings if any
	if len(enhanced.Analysis.SecurityIssues) > 0 {
		securityText := "\n--- Security Issues ---\n"
		for _, issue := range enhanced.Analysis.SecurityIssues {
			securityText += fmt.Sprintf("- %s (%s): %s\n", issue.Type, issue.Severity, issue.Description)
		}
		
		enhancedResult.Content = append(enhancedResult.Content,
			mcp.TextContent{
				Type: "text", 
				Text: securityText,
			},
		)
	}
	
	return enhancedResult
}

// performSemanticSearch performs AI-powered semantic search
func (im *IntelligenceMiddleware) performSemanticSearch(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// Extract search parameters
	_, err := request.RequireString("path")
	if err != nil {
		return nil, err
	}
	
	substring, err := request.RequireString("substring")
	if err != nil {
		return nil, err
	}
	
	// Create semantic query
	query := &SemanticQuery{
		Text:    substring,
		Limit:   20,
		Threshold: 0.7,
		Rerank:  true,
		IncludeAnalysis: true,
	}
	
	// Perform semantic search
	if im.engine.semanticSearch == nil {
		return nil, fmt.Errorf("semantic search not available")
	}
	
	results, err := im.engine.semanticSearch.Search(ctx, query)
	if err != nil {
		return nil, err
	}
	
	// Format results
	return im.formatSemanticSearchResults(results, substring), nil
}

// formatSemanticSearchResults formats semantic search results for MCP
func (im *IntelligenceMiddleware) formatSemanticSearchResults(results []SearchResult, query string) *mcp.CallToolResult {
	if len(results) == 0 {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				mcp.TextContent{
					Type: "text",
					Text: fmt.Sprintf("No semantic matches found for '%s'", query),
				},
			},
		}
	}
	
	resultText := fmt.Sprintf("Found %d semantic matches for '%s':\n\n", len(results), query)
	
	for i, result := range results {
		resultText += fmt.Sprintf("%d. %s (Score: %.3f)\n", i+1, result.Path, result.Score)
		if result.Analysis != nil {
			resultText += fmt.Sprintf("   Summary: %s\n", result.Analysis.Summary)
		}
		resultText += "\n"
	}
	
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.TextContent{
				Type: "text",
				Text: resultText,
			},
		},
	}
}

// enhanceSearchResults adds AI-powered suggestions to search results
func (im *IntelligenceMiddleware) enhanceSearchResults(ctx context.Context, result *mcp.CallToolResult, request mcp.CallToolRequest) *mcp.CallToolResult {
	// Add related suggestions, spelling corrections, etc.
	// This is a placeholder for advanced search enhancement logic
	return result
}

// indexContentAsync indexes content for semantic search in background
func (im *IntelligenceMiddleware) indexContentAsync(ctx context.Context, request mcp.CallToolRequest) {
	contentPath, _ := request.RequireString("path")
	content, _ := request.RequireString("content")
	
	if im.engine.semanticSearch != nil {
		// Perform content analysis first
		if im.engine.contentAnalyzer != nil {
			analysis, err := im.engine.contentAnalyzer.AnalyzeFile(ctx, contentPath, []byte(content), "text/plain")
			if err == nil {
				// Index with analysis
				im.engine.semanticSearch.IndexContent(ctx, contentPath, []byte(content), analysis)
			}
		} else {
			// Index without analysis
			im.engine.semanticSearch.IndexContent(ctx, contentPath, []byte(content), nil)
		}
	}
}

// processAsync processes content analysis asynchronously
func (im *IntelligenceMiddleware) processAsync(ctx context.Context, opCtx *OperationContext) {
	_, err := im.engine.Process(ctx, opCtx)
	if err != nil {
		// Log error but don't fail the main operation
		fmt.Printf("Async intelligence processing failed: %v\n", err)
	}
}

// generateRequestID generates a unique request ID
func generateRequestID() string {
	return fmt.Sprintf("req_%d", time.Now().UnixNano())
}
