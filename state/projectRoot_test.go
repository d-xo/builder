package state

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
			name: "non-root path -> false",
			args: args{path: "/abc/"},
			want: false,
		},
		{
			name: "root-path *nix -> true",
			args: args{path: "/"},
			want: true,
		},
		{
			name: "root-path win -> true",
			args: args{path: "C:\\"},
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

func Test_hash(t *testing.T) {
	type args struct {
		bytes []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "correct hash is computed",
			args: args{bytes: []byte{10, 5, 7, 8, 14}},
			want: "8541cb67c204e32e4f39ea9f7132b8a37e369aa5",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hash(tt.args.bytes); got != tt.want {
				t.Errorf("hash() = %v, want %v", got, tt.want)
			}
		})
	}
}
