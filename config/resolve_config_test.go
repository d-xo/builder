package config

import (
	"testing"
)

func Test_pathExists(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty path -> false",
			args: args{path: ""},
			want: false,
		},
		{
			name: "root path -> true",
			args: args{path: "/"},
			want: true,
		},
		{
			name: "only separators -> true",
			args: args{path: "////////"},
			want: true,
		},
		{
			name: "non existant path -> false",
			args: args{path: "/asdkucxosasd/-_-_a/df-i/rlly-hope-nobody-ever-has-a-path-like-this"},
			want: false,
		},
		{
			name: "invalid path -> false",
			args: args{path: "*&*&%(%*&@_!_&@!(*@^@!&^)*$%*&#)!_@&*$#@*$%))"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pathExists(tt.args.path); got != tt.want {
				t.Errorf("pathExists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPathRoot(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "root-path *nix -> true",
			args: args{path: "/"},
			want: true,
		},
		{
			name: "non-root path *nix -> false",
			args: args{path: "/abc/"},
			want: false,
		},
		{
			name: "root-path win -> true",
			args: args{path: "C:\\"},
			want: true,
		},
		{
			name: "non-root path win -> false",
			args: args{path: "C:\\ahsdf\\asldkf"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPathRoot(tt.args.path); got != tt.want {
				t.Errorf("isPathRoot() = %v, want %v", got, tt.want)
			}
		})
	}
}
