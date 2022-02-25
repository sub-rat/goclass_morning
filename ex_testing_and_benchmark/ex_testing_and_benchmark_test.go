package extestingandbenchmark

import "testing"

// func TestSumBasic(t *testing.T) {
// 	output, err := Sum(2, 2) // 2
// 	if err != nil {
// 		t.Errorf("want = %d, got =%d", 4, output)
// 	}

// 	if output != 4 {
// 		t.Errorf("want = %d, got =%d", 4, output)
// 	}
// }

func TestSumWithTable(t *testing.T) {
	type args struct {
		a int
		b int
	}

	type tests struct {
		name string
		args args
		want int
		err  bool
	}

	var tes = []tests{
		{
			name: "test1",
			args: args{
				a: 1,
				b: 1,
			},
			want: 2,
			err:  false,
		},
		{
			name: "test2",
			args: args{
				a: 10,
				b: 11,
			},
			want: 21,
			err:  false,
		},
		{
			name: "test3",
			args: args{
				a: -1,
				b: 1,
			},
			want: 0,
			err:  true,
		},
		{
			name: "test4",
			args: args{
				a: 1,
				b: -2,
			},
			want: 0,
			err:  true,
		},
		{
			name: "test5",
			args: args{
				a: -11,
				b: -10,
			},
			want: 0,
			err:  true,
		},
	}
	// error => error ( right)
	// 2 => error (wrong)

	for _, tt := range tes {
		t.Run(tt.name, func(t *testing.T) {
			output, err := Sum(tt.args.a, tt.args.b)
			if err != nil && !tt.err {
				t.Errorf("want = %d, got =%d", tt.want, output)
			}

			if output != tt.want {
				t.Errorf("want = %d, got =%d", tt.want, output)
			}

		})
	}

}

func BenchmarkSum(t *testing.B) {
	for i := 0; i < t.N; i++ {
		Sum(10, 11)
	}
}

func Benchmark_Loop(t *testing.B) {
	for i := 0; i < t.N; i++ {
		Loop()
	}
}
