package intelligence

import (
	"context"
	"crypto/sha256"
	"fmt"
	"path/filepath"
	"strings"
	"time"
)

// AnalyzerConfig configuration for content analysis
type AnalyzerConfig struct {
	MaxFileSize       int64
	CacheSize         int
	CacheTTL          time.Duration
	AnalysisDepth     AnalysisDepth
	EnableSentiment   bool
	EnableSecurity    bool
	EnableQuality     bool
	SupportedTypes    []string
}

// AnalysisDepth defines how deep the analysis should be
type AnalysisDepth int

const (
	ShallowAnalysis AnalysisDepth = iota
	MediumAnalysis
	DeepAnalysis
)

// NewContentAnalyzer creates a new content analyzer
func NewContentAnalyzer(config IntelligenceConfig) (*ContentAnalyzer, error) {
	analyzerConfig := AnalyzerConfig{
		MaxFileSize:    10 * 1024 * 1024, // 10MB
		CacheSize:      1000,
		CacheTTL:       time.Hour * 24,
		AnalysisDepth:  MediumAnalysis,
		EnableSentiment: true,
		EnableSecurity:  true,
		EnableQuality:   true,
		SupportedTypes: []string{
			"text/plain", "application/json", "text/x-python",
			"text/x-go", "text/javascript", "text/x-java",
			"text/x-c", "text/x-cpp", "text/html", "text/css",
		},
	}
	
	cache := NewMemoryAnalysisCache(analyzerConfig.CacheSize)
	analyzer := NewLocalAIAnalyzer(analyzerConfig)
	
	return &ContentAnalyzer{
		analyzer: analyzer,
		cache:    cache,
		config:   analyzerConfig,
		enabled:  true,
	}, nil
}

// AnalyzeFile performs comprehensive analysis of a file
func (ca *ContentAnalyzer) AnalyzeFile(ctx context.Context, path string, content []byte, contentType string) (*ContentAnalysis, error) {
	if !ca.enabled {
		return nil, fmt.Errorf("content analyzer is disabled")
	}
	
	// Check file size limits
	if int64(len(content)) > ca.config.MaxFileSize {
		return nil, fmt.Errorf("file too large for analysis: %d bytes", len(content))
	}
	
	// Generate cache key
	cacheKey := ca.generateCacheKey(path, content, contentType)
	
	// Check cache first
	if cached, found := ca.cache.Get(cacheKey); found {
		return cached, nil
	}
	
	// Perform analysis based on content type
	var analysis *ContentAnalysis
	var err error
	
	if ca.isCodeFile(contentType) {
		language := ca.detectLanguage(path, contentType)
		codeAnalysis, err := ca.analyzer.AnalyzeCode(ctx, content, language)
		if err != nil {
			return nil, fmt.Errorf("code analysis failed: %w", err)
		}
		analysis = ca.convertCodeAnalysis(codeAnalysis, contentType)
	} else if ca.isDocumentFile(contentType) {
		docAnalysis, err := ca.analyzer.AnalyzeDocument(ctx, content, contentType)
		if err != nil {
			return nil, fmt.Errorf("document analysis failed: %w", err)
		}
		analysis = ca.convertDocumentAnalysis(docAnalysis, contentType)
	} else {
		analysis, err = ca.analyzer.AnalyzeContent(ctx, content, contentType)
		if err != nil {
			return nil, fmt.Errorf("content analysis failed: %w", err)
		}
	}
	
	// Enhance with additional metadata
	analysis.CacheKey = cacheKey
	analysis.Timestamp = time.Now()
	analysis.ContentType = contentType
	
	// Cache the result
	ca.cache.Set(cacheKey, analysis, ca.config.CacheTTL)
	
	return analysis, nil
}

// generateCacheKey creates a unique cache key for the content
func (ca *ContentAnalyzer) generateCacheKey(path string, content []byte, contentType string) string {
	return fmt.Sprintf("%s:%s:%x", path, contentType, sha256.Sum256(content))
}

// isCodeFile determines if the content type represents code
func (ca *ContentAnalyzer) isCodeFile(contentType string) bool {
	codeTypes := []string{
		"text/x-python", "text/x-go", "text/javascript",
		"text/x-java", "text/x-c", "text/x-cpp",
		"text/x-csharp", "text/x-ruby", "text/x-php",
	}
	
	for _, codeType := range codeTypes {
		if strings.Contains(contentType, codeType) {
			return true
		}
	}
	return false
}

// isDocumentFile determines if the content type represents a document
func (ca *ContentAnalyzer) isDocumentFile(contentType string) bool {
	docTypes := []string{
		"text/markdown", "text/html", "application/pdf",
		"application/msword", "text/rtf",
	}
	
	for _, docType := range docTypes {
		if strings.Contains(contentType, docType) {
			return true
		}
	}
	return false
}

// detectLanguage attempts to detect programming language from file extension
func (ca *ContentAnalyzer) detectLanguage(path, contentType string) string {
	ext := strings.ToLower(filepath.Ext(path))
	
	langMap := map[string]string{
		".go":   "go",
		".py":   "python",
		".js":   "javascript",
		".ts":   "typescript",
		".java": "java",
		".c":    "c",
		".cpp":  "cpp",
		".cs":   "csharp",
		".rb":   "ruby",
		".php":  "php",
		".rs":   "rust",
	}
	
	if lang, found := langMap[ext]; found {
		return lang
	}
	
	// Fallback to content type analysis
	if strings.Contains(contentType, "python") {
		return "python"
	} else if strings.Contains(contentType, "javascript") {
		return "javascript"
	}
	
	return "unknown"
}

// Placeholder methods for missing dependencies
func NewMemoryAnalysisCache(size int) AnalysisCache {
	// Placeholder implementation
	return nil
}

func NewLocalAIAnalyzer(config AnalyzerConfig) AIAnalyzer {
	// Placeholder implementation
	return nil
}

func (ca *ContentAnalyzer) convertCodeAnalysis(codeAnalysis *CodeAnalysis, contentType string) *ContentAnalysis {
	// Placeholder implementation
	return &ContentAnalysis{
		ContentType: contentType,
		Language:    codeAnalysis.Language,
		Summary:     "Code analysis placeholder",
		Timestamp:   time.Now(),
	}
}

func (ca *ContentAnalyzer) convertDocumentAnalysis(docAnalysis *DocumentAnalysis, contentType string) *ContentAnalysis {
	// Placeholder implementation
	return &ContentAnalysis{
		ContentType: contentType,
		Summary:     "Document analysis placeholder",
		Timestamp:   time.Now(),
	}
}
