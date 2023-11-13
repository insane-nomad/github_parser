package limit

import (
	"reflect"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestGetLimit(t *testing.T) {
	type args struct {
		agent *fiber.Agent
	}
	tests := []struct {
		name string
		args args
		want Limits
	}{
		{
			name: "GetLimit test",
			args: args{
				agent: fiber.AcquireAgent(),
			},
			want: Limits{
				Resources: Resources{
					Search: Search{
						Limit:     30,
						Used:      0,
						Remaining: 30,
					},
				},
			},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetLimit(tt.args.agent); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLimit() = %v, want %v", got, tt.want)
			}
		})
	}
}
