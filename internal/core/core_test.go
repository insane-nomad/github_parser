package core

import (
	"reflect"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Test_getData(t *testing.T) {
	type args struct {
		agent *fiber.Agent
		url   string
	}
	tests := []struct {
		name string
		args args
		want RepositoryData
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getData(tt.args.agent, tt.args.url); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetfirstRepo(t *testing.T) {
	type args struct {
		agent *fiber.Agent
	}
	tests := []struct {
		name string
		args args
		want RepositoryData
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetfirstRepo(tt.args.agent); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetfirstRepo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAllRepo(t *testing.T) {
	type args struct {
		agent *fiber.Agent
		from  time.Time
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetAllRepo(tt.args.agent, tt.args.from)
		})
	}
}
