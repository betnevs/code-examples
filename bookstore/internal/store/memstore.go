package store

import (
	"sync"

	"github.com/betNevS/code-examples/bookstore/store/factory"

	"github.com/betNevS/code-examples/bookstore/store"
)

type MemStore struct {
	sync.RWMutex
	books map[string]*store.Book
}

func (m *MemStore) Create(book *store.Book) error {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.books[book.Id]; ok {
		return store.ErrExist
	}

	nBook := *book
	m.books[book.Id] = &nBook

	return nil
}

func (m *MemStore) Update(book *store.Book) error {
	m.Lock()
	defer m.Unlock()

	oldBook, ok := m.books[book.Id]
	if !ok {
		return store.ErrNotFound
	}

	nBook := *oldBook

	if book.Name != "" {
		nBook.Name = book.Name
	}

	if book.Authors != nil {
		nBook.Authors = book.Authors
	}

	if book.Press != "" {
		nBook.Press = book.Press
	}

	m.books[book.Id] = &nBook

	return nil
}

func (m *MemStore) Get(id string) (store.Book, error) {
	m.RLock()
	defer m.RUnlock()

	if t, ok := m.books[id]; ok {
		return *t, nil
	}
	return store.Book{}, store.ErrNotFound
}

func (m *MemStore) GetAll() ([]store.Book, error) {
	m.RLock()
	defer m.RUnlock()

	allBooks := make([]store.Book, 0, len(m.books))
	for _, b := range m.books {
		allBooks = append(allBooks, *b)
	}
	return allBooks, nil
}

func (m *MemStore) Delete(id string) error {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.books[id]; !ok {
		return store.ErrNotFound
	}

	delete(m.books, id)

	return nil
}

func init() {
	factory.Register("mem", &MemStore{
		books: make(map[string]*store.Book),
	})
}
