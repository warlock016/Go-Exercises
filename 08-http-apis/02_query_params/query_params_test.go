package query_params

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSearchHandler(t *testing.T) {
	tests := []struct {
		name       string
		url        string
		wantStatus int
		wantResult *SearchResult
	}{
		{
			name:       "with all parameters",
			url:        "/search?q=golang&page=2&limit=25",
			wantStatus: http.StatusOK,
			wantResult: &SearchResult{Query: "golang", Page: 2, Limit: 25},
		},
		{
			name:       "with defaults for page and limit",
			url:        "/search?q=golang",
			wantStatus: http.StatusOK,
			wantResult: &SearchResult{Query: "golang", Page: 1, Limit: 10},
		},
		{
			name:       "with custom page only",
			url:        "/search?q=http&page=5",
			wantStatus: http.StatusOK,
			wantResult: &SearchResult{Query: "http", Page: 5, Limit: 10},
		},
		{
			name:       "with custom limit only",
			url:        "/search?q=testing&limit=50",
			wantStatus: http.StatusOK,
			wantResult: &SearchResult{Query: "testing", Page: 1, Limit: 50},
		},
		{
			name:       "missing required q parameter",
			url:        "/search",
			wantStatus: http.StatusBadRequest,
			wantResult: nil,
		},
		{
			name:       "empty q parameter",
			url:        "/search?q=",
			wantStatus: http.StatusBadRequest,
			wantResult: nil,
		},
		{
			name:       "unicode query",
			url:        "/search?q=世界",
			wantStatus: http.StatusOK,
			wantResult: &SearchResult{Query: "世界", Page: 1, Limit: 10},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, tt.url, nil)
			w := httptest.NewRecorder()

			SearchHandler(w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("SearchHandler() status = %d, want %d", w.Code, tt.wantStatus)
			}

			if tt.wantResult != nil {
				var got SearchResult
				if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
					t.Fatalf("failed to decode response: %v", err)
				}

				if got != *tt.wantResult {
					t.Errorf("SearchHandler() = %+v, want %+v", got, *tt.wantResult)
				}

				contentType := w.Header().Get("Content-Type")
				if contentType != "application/json" {
					t.Errorf("Content-Type = %q, want %q", contentType, "application/json")
				}
			}
		})
	}
}

func TestFilterHandler(t *testing.T) {
	tests := []struct {
		name       string
		url        string
		wantStatus int
		wantTags   []string
	}{
		{
			name:       "multiple tags",
			url:        "/filter?tag=golang&tag=http&tag=json",
			wantStatus: http.StatusOK,
			wantTags:   []string{"golang", "http", "json"},
		},
		{
			name:       "single tag",
			url:        "/filter?tag=backend",
			wantStatus: http.StatusOK,
			wantTags:   []string{"backend"},
		},
		{
			name:       "no tags",
			url:        "/filter",
			wantStatus: http.StatusOK,
			wantTags:   []string{},
		},
		{
			name:       "empty tag parameter",
			url:        "/filter?tag=",
			wantStatus: http.StatusOK,
			wantTags:   []string{""},
		},
		{
			name:       "two tags with one empty",
			url:        "/filter?tag=go&tag=",
			wantStatus: http.StatusOK,
			wantTags:   []string{"go", ""},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, tt.url, nil)
			w := httptest.NewRecorder()

			FilterHandler(w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("FilterHandler() status = %d, want %d", w.Code, tt.wantStatus)
			}

			var got FilterResult
			if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
				t.Fatalf("failed to decode response: %v", err)
			}

			// Handle nil vs empty slice
			if got.Tags == nil {
				got.Tags = []string{}
			}
			wantTags := tt.wantTags
			if wantTags == nil {
				wantTags = []string{}
			}

			if len(got.Tags) != len(wantTags) {
				t.Errorf("FilterHandler() tags = %v, want %v", got.Tags, wantTags)
				return
			}

			for i := range got.Tags {
				if got.Tags[i] != wantTags[i] {
					t.Errorf("FilterHandler() tags[%d] = %q, want %q", i, got.Tags[i], wantTags[i])
				}
			}

			contentType := w.Header().Get("Content-Type")
			if contentType != "application/json" {
				t.Errorf("Content-Type = %q, want %q", contentType, "application/json")
			}
		})
	}
}
