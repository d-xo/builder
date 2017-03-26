package docker

import (
	"testing"
)

func Test_volumeArgs(t *testing.T) {
	type args struct {
		volumes map[string]string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{volumes: map[string]string{
				"/host/dir/1": "/guest/dir/1",
				"/host/dir/2": "/guest/dir/2",
			}},
			want: "-v/host/dir/1:/guest/dir/1 -v/host/dir/2:/guest/dir/2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := volumeArgs(tt.args.volumes); got != tt.want {
				t.Errorf("volumesString() = %v, want %v", got, tt.want)
			}
		})
	}
}
