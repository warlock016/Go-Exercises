package method_routing

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestItemStore(t *testing.T) {
	store := NewItemStore()

	// Test Create
	item1 := store.Create("Item 1")
	if item1.ID != 1 || item1.Name != "Item 1" {
		t.Errorf("Create() = %+v, want ID=1, Name='Item 1'", item1)
	}

	item2 := store.Create("Item 2")
	if item2.ID != 2 {
		t.Errorf("Create() second item ID = %d, want 2", item2.ID)
	}

	// Test Get
	got, ok := store.Get(1)
	if !ok || got.Name != "Item 1" {
		t.Errorf("Get(1) = %+v, %v, want Item 1, true", got, ok)
	}

	_, ok = store.Get(999)
	if ok {
		t.Errorf("Get(999) should return false")
	}

	// Test Update
	if !store.Update(1, "Updated Item") {
		t.Errorf("Update(1) = false, want true")
	}

	got, _ = store.Get(1)
	if got.Name != "Updated Item" {
		t.Errorf("After Update, Get(1).Name = %q, want 'Updated Item'", got.Name)
	}

	// Test Delete
	if !store.Delete(1) {
		t.Errorf("Delete(1) = false, want true")
	}

	_, ok = store.Get(1)
	if ok {
		t.Errorf("After Delete, Get(1) should return false")
	}
}

func TestItemHandler(t *testing.T) {
	tests := []struct {
		name       string
		method     string
		url        string
		body       string
		wantStatus int
		wantItem   *Item
	}{
		{
			name:       "POST creates item",
			method:     http.MethodPost,
			url:        "/items",
			body:       `{"name":"Test Item"}`,
			wantStatus: http.StatusCreated,
			wantItem:   &Item{ID: 1, Name: "Test Item"},
		},
		{
			name:       "GET retrieves item",
			method:     http.MethodGet,
			url:        "/items?id=1",
			wantStatus: http.StatusOK,
			wantItem:   &Item{ID: 1, Name: "Test Item"},
		},
		{
			name:       "GET non-existent item",
			method:     http.MethodGet,
			url:        "/items?id=999",
			wantStatus: http.StatusNotFound,
		},
		{
			name:       "PUT updates item",
			method:     http.MethodPut,
			url:        "/items?id=1",
			body:       `{"name":"Updated Item"}`,
			wantStatus: http.StatusOK,
			wantItem:   &Item{ID: 1, Name: "Updated Item"},
		},
		{
			name:       "DELETE removes item",
			method:     http.MethodDelete,
			url:        "/items?id=1",
			wantStatus: http.StatusNoContent,
		},
		{
			name:       "PATCH not allowed",
			method:     http.MethodPatch,
			url:        "/items",
			wantStatus: http.StatusMethodNotAllowed,
		},
	}

	store := NewItemStore()
	handler := ItemHandler(store)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var req *http.Request
			if tt.body != "" {
				req = httptest.NewRequest(tt.method, tt.url, bytes.NewBufferString(tt.body))
				req.Header.Set("Content-Type", "application/json")
			} else {
				req = httptest.NewRequest(tt.method, tt.url, nil)
			}

			w := httptest.NewRecorder()
			handler(w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("ItemHandler() status = %d, want %d", w.Code, tt.wantStatus)
			}

			if tt.wantItem != nil {
				var got Item
				if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
					t.Fatalf("failed to decode response: %v", err)
				}

				if got != *tt.wantItem {
					t.Errorf("ItemHandler() = %+v, want %+v", got, *tt.wantItem)
				}
			}

			if tt.wantStatus == http.StatusMethodNotAllowed {
				allow := w.Header().Get("Allow")
				if allow == "" {
					t.Errorf("ItemHandler() missing Allow header for 405 response")
				}
			}
		})
	}
}
