package method_routing

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// Item represents an item in the store
type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// ItemStore manages items in memory
type ItemStore struct {
	items  map[int]Item
	nextID int
}

// NewItemStore creates a new ItemStore
func NewItemStore() *ItemStore {
	// TODO(human): Implement
	return &ItemStore{
		items:  make(map[int]Item),
		nextID: 0,
	}
}

// Get retrieves an item by ID
func (s *ItemStore) Get(id int) (Item, bool) {
	// TODO(human): Implement
	item, ok := s.items[id]
	if !ok {
		return Item{}, false
	}
	return item, true
}

// Create adds a new item and returns it with assigned ID
func (s *ItemStore) Create(name string) Item {
	// TODO(human): Implement
	s.nextID++
	newItem := Item{
		Name: name,
		ID:   s.nextID,
	}
	s.items[newItem.ID] = newItem
	return newItem
}

// Update modifies an existing item
func (s *ItemStore) Update(id int, name string) (Item, bool) {
	// TODO(human): Implement
	if _, ok := s.items[id]; !ok {
		return Item{}, false
	}

	item := s.items[id]
	item.Name = name
	s.items[id] = item
	return item, true
}

// Delete removes an item by ID
func (s *ItemStore) Delete(id int) bool {
	// TODO(human): Implement
	if _, ok := s.items[id]; ok {
		delete(s.items, id)
		return true
	}
	return false
}

// ItemHandler handles CRUD operations on items
func ItemHandler(store *ItemStore) http.HandlerFunc {
	// TODO(human): Implement
	var allowedMethods = map[string]bool{"GET": true, "POST": true, "PUT": true, "DELETE": true}

	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		if !allowedMethods[r.Method] {
			w.Header().Set("Allow", "GET, POST, PUT, DELETE")
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		var input Item

		if r.Method == "POST" {
			if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			input = store.Create(input.Name)
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(input)
			return
		}

		params := r.URL.Query()
		id := params.Get("id")
		val, err := strconv.ParseInt(id, 10, 64)
		if err != nil || val == 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		switch r.Method {
		case "GET":
			res, ok := store.Get(int(val))
			if !ok {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(res)
			return

		case "PUT":
			if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			item, ok := store.Update(int(val), input.Name)
			if !ok {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(item)
			return

		case "DELETE":
			if ok := store.Delete(int(val)); !ok {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
}
