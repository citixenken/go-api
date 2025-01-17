package transport

import (
	"github.com/citixenken/go-api.git/internal/todo"
	"reflect"
	"testing"
)

func TestNewServer(t *testing.T) {
	type args struct {
		todoSvc *todo.Service
	}
	tests := []struct {
		name string
		args args
		want *Server
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewServer(tt.args.todoSvc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewServer() = %v, want %v", got, tt.want)
			}
		})
	}
}
