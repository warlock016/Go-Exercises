# Exercise 09: Testing JSON APIs (🌐 HTTP)

**Learning Goal:** Test JSON API endpoints - parse responses, validate structure, check error cases.

---

## API to Test

```go
func GetUserHandler(w http.ResponseWriter, r *http.Request)
func CreateUserHandler(w http.ResponseWriter, r *http.Request)
```

---

## Testing JSON Responses

```go
func TestGetUserHandler(t *testing.T) {
    req := httptest.NewRequest("GET", "/user/123", nil)
    rec := httptest.NewRecorder()

    GetUserHandler(rec, req)

    var response User
    if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
        t.Fatalf("failed to parse JSON: %v", err)
    }

    if response.ID != 123 {
        t.Errorf("user ID = %d, want 123", response.ID)
    }
}
```

---

## Your Task

Test JSON APIs: parse responses, validate fields, check Content-Type headers.

**Next Exercise:** `10_mocking_http_clients` - Mock external HTTP services
