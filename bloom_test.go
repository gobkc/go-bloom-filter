package gobloom

import (
	"testing"
)

func TestBloom_Has(t *testing.T) {
	bloom := NewBloom()
	for i := 0; i < 100000000; i++ {
		bloom.Add(i)
	}
	type args struct {
		dest any
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test has element 1",
			args: args{
				dest: 10001,
			},
			want: true,
		},
		{
			name: "test has element 2",
			args: args{
				dest: 100001,
			},
			want: true,
		},
		{
			name: "test not has element 1",
			args: args{
				dest: -100,
			},
			want: false,
		},
		{
			name: "test not has element 2",
			args: args{
				dest: `one`,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bloom.Has(tt.args.dest); got != tt.want {
				t.Errorf("Has() = %v, want %v", got, tt.want)
			}
		})
	}
}
