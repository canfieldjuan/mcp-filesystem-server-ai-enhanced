package intelligence

import (
	"context"
	"time"
)

// Supporting types for intelligence components

// Core structure types
type ContentAnalyzer struct {
	analyzer      AIAnalyzer
	cache         AnalysisCache
	config        AnalyzerConfig
	enabled       bool
}

type SemanticSearchEngine struct {
	vectorStore   VectorStore
	embeddings    EmbeddingService
	indexer       ContentIndexer
	config        SearchConfig
	enabled       bool
}

type PredictiveCacheLayer struct {
	patterns      PatternLearner
	cache         PredictiveCache
	analytics     UsageAnalytics
	enabled       bool
}

type SmartOperationFramework struct {
	contextEngine ContextEngine
	suggestions   OperationSuggester
	conflictRes   ConflictResolver
	enabled       bool
}

// Configuration types
type SearchConfig struct {
	VectorDimensions  int
	IndexBatchSize    int
	SearchLimit       int
	SimilarityThreshold float64
	EnableReranking   bool
	ChunkSize         int
	ChunkOverlap      int
}

type VectorStoreStats struct {
	TotalVectors   int     `json:"total_vectors"`
	IndexSize      int64   `json:"index_size"`
	AverageScore   float64 `json:"average_score"`
	LastUpdated    time.Time `json:"last_updated"`
}

type IndexStats struct {
	IndexedFiles   int       `json:"indexed_files"`
	TotalChunks    int       `json:"total_chunks"`
	AverageChunks  float64   `json:"average_chunks"`
	LastIndexed    time.Time `json:"last_indexed"`
}

type SemanticQuery struct {
	Text        string            `json:"text"`
	Filters     map[string]interface{} `json:"filters"`
	Limit       int               `json:"limit"`
	Threshold   float64           `json:"threshold"`
	Rerank      bool              `json:"rerank"`
	IncludeAnalysis bool          `json:"include_analysis"`
}

type SearchResult struct {
	Path        string                 `json:"path"`
	Score       float64                `json:"score"`
	Content     string                 `json:"content"`
	Metadata    map[string]interface{} `json:"metadata"`
	Chunk       *ContentChunk          `json:"chunk,omitempty"`
	Analysis    *ContentAnalysis       `json:"analysis,omitempty"`
}

type ContentChunk struct {
	ID          string    `json:"id"`
	Path        string    `json:"path"`
	Content     string    `json:"content"`
	StartOffset int       `json:"start_offset"`
	EndOffset   int       `json:"end_offset"`
	ChunkIndex  int       `json:"chunk_index"`
	Vector      []float64 `json:"vector,omitempty"`
}

// Core data structures
type OperationContext struct {
	Operation     string
	Path          string
	Content       []byte
	Metadata      map[string]interface{}
	UserContext   UserContext
	RequestID     string
	Timestamp     time.Time
}

type EnhancedResult struct {
	OriginalResult interface{}
	Analysis       *ContentAnalysis
	Suggestions    []OperationSuggestion
	Metadata       map[string]interface{}
	CacheHint      *CacheHint
}

type ContentAnalysis struct {
	ContentType     string            `json:"content_type"`
	Language        string            `json:"language"`
	Summary         string            `json:"summary"`
	KeyTopics       []string          `json:"key_topics"`
	Sentiment       float64           `json:"sentiment"`
	Complexity      float64           `json:"complexity"`
	Dependencies    []string          `json:"dependencies"`
	SecurityIssues  []SecurityIssue   `json:"security_issues"`
	QualityMetrics  QualityMetrics    `json:"quality_metrics"`
	Metadata        map[string]interface{} `json:"metadata"`
	Timestamp       time.Time         `json:"timestamp"`
	CacheKey        string            `json:"cache_key"`
}

type CodeAnalysis struct {
	Language        string            `json:"language"`
	Functions       []FunctionInfo    `json:"functions"`
	Classes         []ClassInfo       `json:"classes"`
	Imports         []string          `json:"imports"`
	Complexity      CyclomaticComplexity `json:"complexity"`
	TestCoverage    float64           `json:"test_coverage"`
	Documentation   float64           `json:"documentation"`
	CodeSmells      []CodeSmell       `json:"code_smells"`
	Patterns        []DesignPattern   `json:"patterns"`
}

type DocumentAnalysis struct {
	DocumentType    string            `json:"document_type"`
	WordCount       int               `json:"word_count"`
	ReadingLevel    string            `json:"reading_level"`
	KeyConcepts     []string          `json:"key_concepts"`
	Outline         []Section         `json:"outline"`
	References      []Reference       `json:"references"`
	Metadata        DocumentMetadata  `json:"metadata"`
}

type IntelligenceConfig struct {
	ContentAnalysis   bool
	SemanticSearch    bool
	PredictiveCache   bool
	SmartOperations   bool
	AnalysisDepth     int
	CacheSize         int
	VectorDimensions  int
}

// Interface for middleware
type Middleware interface {
	Process(ctx context.Context, operation *OperationContext) (*EnhancedResult, error)
	Name() string
	Priority() int
}

// Supporting types for intelligence components

// UserContext contains user-specific context information
type UserContext struct {
	UserID      string            `json:"user_id"`
	Preferences map[string]interface{} `json:"preferences"`
	History     []string          `json:"history"`
}

// OperationSuggestion represents a suggested operation
type OperationSuggestion struct {
	Type        string  `json:"type"`
	Description string  `json:"description"`
	Confidence  float64 `json:"confidence"`
	Action      string  `json:"action"`
}

// CacheHint provides caching recommendations
type CacheHint struct {
	ShouldCache bool          `json:"should_cache"`
	TTL         time.Duration `json:"ttl"`
	Priority    int           `json:"priority"`
}

// SecurityIssue represents a security concern found in content
type SecurityIssue struct {
	Type        string `json:"type"`
	Severity    string `json:"severity"`
	Description string `json:"description"`
	Line        int    `json:"line,omitempty"`
	Column      int    `json:"column,omitempty"`
	Suggestion  string `json:"suggestion"`
}

// QualityMetrics represents code/content quality metrics
type QualityMetrics struct {
	Maintainability float64 `json:"maintainability"`
	Readability     float64 `json:"readability"`
	Reusability     float64 `json:"reusability"`
	Testability     float64 `json:"testability"`
	Overall         float64 `json:"overall"`
}

// FunctionInfo represents information about a function
type FunctionInfo struct {
	Name       string `json:"name"`
	StartLine  int    `json:"start_line"`
	EndLine    int    `json:"end_line"`
	Parameters []string `json:"parameters"`
	ReturnType string `json:"return_type"`
}

// ClassInfo represents information about a class
type ClassInfo struct {
	Name      string         `json:"name"`
	StartLine int           `json:"start_line"`
	EndLine   int           `json:"end_line"`
	Methods   []FunctionInfo `json:"methods"`
	Fields    []string      `json:"fields"`
}

// CyclomaticComplexity represents complexity metrics
type CyclomaticComplexity struct {
	Average float64 `json:"average"`
	Maximum int     `json:"maximum"`
	Total   int     `json:"total"`
}

// CodeSmell represents a code quality issue
type CodeSmell struct {
	Type        string `json:"type"`
	Description string `json:"description"`
	Line        int    `json:"line"`
	Severity    string `json:"severity"`
}

// DesignPattern represents a detected design pattern
type DesignPattern struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Location    string `json:"location"`
	Confidence  float64 `json:"confidence"`
}

// Section represents a document section
type Section struct {
	Title    string    `json:"title"`
	Level    int       `json:"level"`
	Content  string    `json:"content"`
	Children []Section `json:"children,omitempty"`
}

// Reference represents a document reference
type Reference struct {
	Type string `json:"type"`
	URL  string `json:"url"`
	Text string `json:"text"`
}

// DocumentMetadata represents document metadata
type DocumentMetadata struct {
	Author      string    `json:"author"`
	Title       string    `json:"title"`
	CreatedDate time.Time `json:"created_date"`
	ModifiedDate time.Time `json:"modified_date"`
	Keywords    []string  `json:"keywords"`
}

// Interfaces for dependency injection

// AIAnalyzer interface for different AI analysis backends
type AIAnalyzer interface {
	AnalyzeContent(ctx context.Context, content []byte, contentType string) (*ContentAnalysis, error)
	AnalyzeCode(ctx context.Context, code []byte, language string) (*CodeAnalysis, error)
	AnalyzeDocument(ctx context.Context, doc []byte, docType string) (*DocumentAnalysis, error)
}

// AnalysisCache caches analysis results to avoid recomputation
type AnalysisCache interface {
	Get(key string) (*ContentAnalysis, bool)
	Set(key string, analysis *ContentAnalysis, ttl time.Duration)
	Invalidate(key string)
	Clear()
}

// VectorStore interface for storing and querying vectors
type VectorStore interface {
	Store(ctx context.Context, id string, vector []float64, metadata map[string]interface{}) error
	Search(ctx context.Context, queryVector []float64, limit int, threshold float64) ([]SearchResult, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, id string, vector []float64, metadata map[string]interface{}) error
	GetStats() VectorStoreStats
}

// EmbeddingService generates vector embeddings from text
type EmbeddingService interface {
	Embed(ctx context.Context, text string) ([]float64, error)
	EmbedBatch(ctx context.Context, texts []string) ([][]float64, error)
	GetDimensions() int
	GetModel() string
}

// ContentIndexer indexes file content for semantic search
type ContentIndexer interface {
	IndexFile(ctx context.Context, path string, content []byte, analysis *ContentAnalysis) error
	RemoveFile(ctx context.Context, path string) error
	UpdateFile(ctx context.Context, path string, content []byte, analysis *ContentAnalysis) error
	GetIndexStats() IndexStats
}

// PatternLearner learns usage patterns for optimization
type PatternLearner interface {
	LearnPattern(ctx context.Context, pattern UsagePattern) error
	PredictUsage(ctx context.Context, context UsageContext) (*UsagePrediction, error)
	GetPatterns() []UsagePattern
}

// PredictiveCache provides predictive caching capabilities
type PredictiveCache interface {
	Get(key string) (interface{}, bool)
	Set(key string, value interface{}, ttl time.Duration)
	Preload(keys []string) error
	GetStats() CacheStats
}

// UsageAnalytics tracks usage patterns
type UsageAnalytics interface {
	RecordUsage(ctx context.Context, usage UsageEvent) error
	GetAnalytics(ctx context.Context, timeRange TimeRange) (*Analytics, error)
}

// ContextEngine provides context-aware operations
type ContextEngine interface {
	GetContext(ctx context.Context, path string) (*OperationContext, error)
	UpdateContext(ctx context.Context, opCtx *OperationContext) error
}

// OperationSuggester suggests smart operations
type OperationSuggester interface {
	SuggestOperations(ctx context.Context, context *OperationContext) ([]OperationSuggestion, error)
}

// ConflictResolver resolves operation conflicts
type ConflictResolver interface {
	DetectConflicts(ctx context.Context, operations []Operation) ([]Conflict, error)
	ResolveConflicts(ctx context.Context, conflicts []Conflict) ([]Resolution, error)
}

// Supporting data structures

type UsagePattern struct {
	ID          string                 `json:"id"`
	Type        string                 `json:"type"`
	Frequency   float64                `json:"frequency"`
	Context     map[string]interface{} `json:"context"`
	Timestamp   time.Time              `json:"timestamp"`
}

type UsageContext struct {
	Path        string                 `json:"path"`
	Operation   string                 `json:"operation"`
	Time        time.Time              `json:"time"`
	Metadata    map[string]interface{} `json:"metadata"`
}

type UsagePrediction struct {
	Probability float64               `json:"probability"`
	Operations  []string              `json:"operations"`
	Confidence  float64               `json:"confidence"`
	Reasoning   string                `json:"reasoning"`
}

type CacheStats struct {
	HitRate     float64 `json:"hit_rate"`
	MissRate    float64 `json:"miss_rate"`
	Size        int     `json:"size"`
	LastUpdated time.Time `json:"last_updated"`
}

type UsageEvent struct {
	UserID    string    `json:"user_id"`
	Operation string    `json:"operation"`
	Path      string    `json:"path"`
	Duration  time.Duration `json:"duration"`
	Success   bool      `json:"success"`
	Timestamp time.Time `json:"timestamp"`
}

type TimeRange struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}

type Analytics struct {
	TotalOperations int                    `json:"total_operations"`
	PopularPaths    []string               `json:"popular_paths"`
	AverageLatency  time.Duration          `json:"average_latency"`
	ErrorRate       float64                `json:"error_rate"`
	Insights        map[string]interface{} `json:"insights"`
}

type Operation struct {
	ID        string    `json:"id"`
	Type      string    `json:"type"`
	Path      string    `json:"path"`
	Timestamp time.Time `json:"timestamp"`
	UserID    string    `json:"user_id"`
}

type Conflict struct {
	ID          string      `json:"id"`
	Type        string      `json:"type"`
	Operations  []Operation `json:"operations"`
	Description string      `json:"description"`
	Severity    string      `json:"severity"`
}

type Resolution struct {
	ConflictID  string `json:"conflict_id"`
	Strategy    string `json:"strategy"`
	Description string `json:"description"`
	Applied     bool   `json:"applied"`
}
