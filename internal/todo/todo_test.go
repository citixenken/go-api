package todo_test

import (
	"context"
	"my_first_api/internal/db"
	"my_first_api/internal/todo"
	"reflect"
	"testing"
)

type MockDB struct {
	items []db.Item
}

func (m *MockDB) InsertItem(_ context.Context, item db.Item) error {
	m.items = append(m.items, item)
	return nil
}

func (m *MockDB) GetAllItems(_ context.Context) ([]db.Item, error) {
	return m.items, nil
}

func TestService_Search(t *testing.T) {

	tests := []struct {
		name           string
		toDosToAdd     []string
		query          string
		expectedResult []string
	}{
		{
			name:           "given a todo of book and a search of bo, I should get book back",
			toDosToAdd:     []string{"book"},
			query:          "bo",
			expectedResult: []string{"book"},
		},
		{
			name:           "still returns book, even if case doesn't match",
			toDosToAdd:     []string{"Booking"},
			query:          "bo",
			expectedResult: []string{"Booking"},
		},
		{
			name:           "spaces",
			toDosToAdd:     []string{"take Book"},
			query:          "take",
			expectedResult: []string{"take Book"},
		},
		{
			name:           "space at start of word",
			toDosToAdd:     []string{" Space at beginning"},
			query:          "space",
			expectedResult: []string{" Space at beginning"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockDB{}
			svc := todo.NewService(m)
			for _, toAdd := range tt.toDosToAdd {
				err := svc.Add(toAdd)
				if err != nil {
					t.Error(err)
				}
			}
			got, err := svc.Search(tt.query)
			if err != nil {
				t.Error(err)
			}
			if !reflect.DeepEqual(got, tt.expectedResult) {
				t.Errorf("Search() = %v, want %v", got, tt.expectedResult)
			}
		})
	}
}

func TestService_GetAll(t *testing.T) {
	tests := []struct {
		name           string
		expectedResult []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockDB{}
			svc := todo.NewService(m)
			got, err := svc.GetAll()
			if err != nil {
				t.Error(err)
			}
			if !reflect.DeepEqual(got, tt.expectedResult) {
				t.Errorf("GetAll() = %v, want %v", got, tt.expectedResult)
			}
		})
	}
}
