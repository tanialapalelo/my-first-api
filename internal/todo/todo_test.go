package todo_test

import (
	"reflect"
	"testing"

	"my-first-api/internal/todo"
)

func TestService_Search(t *testing.T) {
	tests := []struct {
		name           string
		toDosToAdd     []string
		query          string
		expectedResult []string
	}{
		// test cases
		{
			name:           "given a todo of shop and a search of sh, i should get shop back",
			toDosToAdd:     []string{"shop"},
			query:          "sh",
			expectedResult: []string{"shop"},
		},
		{
			name:           "still returns shop, even if the case does not match",
			toDosToAdd:     []string{"Shop"},
			query:          "sh",
			expectedResult: []string{"Shop"},
		},
		{
			name:           "spaces",
			toDosToAdd:     []string{"go shopping"},
			query:          "go",
			expectedResult: []string{"go shopping"},
		},
		{
			name:           "space at the start of the word",
			toDosToAdd:     []string{" Space at the start"},
			query:          "space",
			expectedResult: []string{" Space at the start"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := todo.NewService()
			for _, toAdd := range tt.toDosToAdd {
				err := svc.Add(toAdd)
				if err != nil {
					t.Error(err)
				}
			}
			if got := svc.Search(tt.query); !reflect.DeepEqual(got, tt.expectedResult) {
				t.Errorf("Search() = %v, want %v", got, tt.expectedResult)
			}
		})
	}
}
