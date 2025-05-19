package tool

import "testing"

// TestTernary
//
//	@Description: 测试Ternary
//	@param t
func TestTernary(t *testing.T) {
	type args struct {
		condition  bool
		trueValue  any
		falseValue any
	}

	tests := []struct {
		name string
		args args
		want any
	}{
		{
			name: "test1",
			args: args{
				condition:  true,
				trueValue:  1,
				falseValue: 2,
			},
			want: 1,
		},
		{
			name: "test2",
			args: args{
				condition:  false,
				trueValue:  1,
				falseValue: 2,
			},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Ternary(tt.args.condition, tt.args.trueValue, tt.args.falseValue); got != tt.want {
				t.Errorf("Ternary() = %v, want %v", got, tt.want)
			}
		})
	}
}
