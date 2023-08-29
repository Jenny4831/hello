package greeting

import (
	"net/http"
	"testing"
)

// this is generated, need to be filled with actual test
func TestGreeting(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Greeting(tt.args.w, tt.args.r)
		})
	}
}

func TestGreet(t *testing.T) {
	t.Parallel()
	// the ... sppeds up test running by telling it that the size of slice will not change
	tests := [...]struct {
		name    string
		nameArg string
		want    string
	}{
		{
			name:    "Test Greet",
			nameArg: "Jenny",
			want:    "Hello World, Jenny",
		},
	}
	for _, tt := range tests {
		// this is required, due to tests running parallel, making a copy to prevent looping the same item
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := Greet(tt.nameArg); got != tt.want {
				t.Errorf("Greet() = %v, want %v", got, tt.want)
			}
		})
	}
}
