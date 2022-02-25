package examplefunctions

import "testing"

func Test_sum(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				a: 1,
				b: 1,
			},
			want:    2,
			wantErr: false,
		},
		{
			name: "test2",
			args: args{
				a: 10,
				b: 11,
			},
			want:    21,
			wantErr: false,
		},
		{
			name: "test3",
			args: args{
				a: -1,
				b: 1,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "test4",
			args: args{
				a: 1,
				b: -2,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "test5",
			args: args{
				a: -11,
				b: -10,
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := sum(tt.args.a, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("sum() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
