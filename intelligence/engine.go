// Package intelligence provides AI-powered enhancements for the MCP filesystem server
package intelligence

import (
	"context"
	"fmt"
	"sync"
)

// IntelligenceEngine is the core AI enhancement engine
type IntelligenceEngine struct {
	contentAnalyzer  *ContentAnalyzer
	semanticSearch   *SemanticSearchEngine
	predictiveCache  *PredictiveCacheLayer
	smartOperations  *SmartOperationFramework
	middleware       []Middleware
	mutex            sync.RWMutex
}

// NewIntelligenceEngine creates a new AI enhancement engine
func NewIntelligenceEngine(config IntelligenceConfig) (*IntelligenceEngine, error) {
	engine := &IntelligenceEngine{
		middleware: make([]Middleware, 0),
	}
	
	if config.ContentAnalysis {
		analyzer, err := NewContentAnalyzer(config)
		if err != nil {
			return nil, fmt.Errorf("failed to create content analyzer: %w", err)
		}
		engine.contentAnalyzer = analyzer
	}
	
	if config.SemanticSearch {
		search, err := NewSemanticSearchEngine(config)
		if err != nil {
			return nil, fmt.Errorf("failed to create semantic search: %w", err)
		}
		engine.semanticSearch = search
	}
	
	if config.PredictiveCache {
		cache, err := NewPredictiveCacheLayer(config)
		if err != nil {
			return nil, fmt.Errorf("failed to create predictive cache: %w", err)
		}
		engine.predictiveCache = cache
	}
	
	if config.SmartOperations {
		operations, err := NewSmartOperationFramework(config)
		if err != nil {
			return nil, fmt.Errorf("failed to create smart operations: %w", err)
		}
		engine.smartOperations = operations
	}
	
	return engine, nil
}

// Process enhances an operation with AI intelligence
func (ie *IntelligenceEngine) Process(ctx context.Context, opCtx *OperationContext) (*EnhancedResult, error) {
	result := &EnhancedResult{
		Metadata: make(map[string]interface{}),
	}
	
	// Apply middleware chain
	for _, middleware := range ie.middleware {
		enhanced, err := middleware.Process(ctx, opCtx)
		if err != nil {
			return nil, fmt.Errorf("middleware %s failed: %w", middleware.Name(), err)
		}
		
		// Merge enhanced results
		if enhanced.Analysis != nil {
			result.Analysis = enhanced.Analysis
		}
		if enhanced.Suggestions != nil {
			result.Suggestions = append(result.Suggestions, enhanced.Suggestions...)
		}
		if enhanced.CacheHint != nil {
			result.CacheHint = enhanced.CacheHint
		}
		
		// Merge metadata
		for k, v := range enhanced.Metadata {
			result.Metadata[k] = v
		}
	}
	
	return result, nil
}

// AddMiddleware adds an enhancement middleware to the processing chain
func (ie *IntelligenceEngine) AddMiddleware(middleware Middleware) {
	ie.mutex.Lock()
	defer ie.mutex.Unlock()
	
	ie.middleware = append(ie.middleware, middleware)
	
	// Sort by priority (higher priority first)
	for i := len(ie.middleware) - 1; i > 0; i-- {
		if ie.middleware[i].Priority() > ie.middleware[i-1].Priority() {
			ie.middleware[i], ie.middleware[i-1] = ie.middleware[i-1], ie.middleware[i]
		} else {
			break
		}
	}
}
