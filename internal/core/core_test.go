package core

import (
	"reflect"
	"sync"
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
		{
			name: "Test for getData",
			args: args{
				agent: fiber.AcquireAgent(),
				url:   "https://api.github.com/search/repositories?q=stm32&sort=updated&order=asc&per_page=1&page=1",
			},
			want: RepositoryData{
				TotalCount: 51277,
				Items: []Items{
					{
						Name: "stm32f10x_stdperiph_lib",
						Owner: Owner{
							Login: "skywolf",
						},
						HTMLURL:         "https://github.com/skywolf/stm32f10x_stdperiph_lib",
						CreatedAt:       time.Date(2009, 9, 13, 10, 07, 22, 0, time.UTC),
						StargazersCount: 1,
					},
				},
			},
		},
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
		agent        *fiber.Agent
		searchString string
	}
	tests := []struct {
		name string
		args args
		want RepositoryData
	}{
		{
			name: "Test for getData",
			args: args{
				agent:        fiber.AcquireAgent(),
				searchString: "stm32",
			},
			want: RepositoryData{
				TotalCount: 51277,
				Items: []Items{
					{
						Name: "stm32f10x_stdperiph_lib",
						Owner: Owner{
							Login: "skywolf",
						},
						HTMLURL:         "https://github.com/skywolf/stm32f10x_stdperiph_lib",
						CreatedAt:       time.Date(2009, 9, 13, 10, 07, 22, 0, time.UTC),
						StargazersCount: 1,
					},
				},
			},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetfirstRepo(tt.args.agent, tt.args.searchString); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetfirstRepo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWorker(t *testing.T) {
	type args struct {
		inputData Items
		wg        *sync.WaitGroup
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test for Worker",
			args: args{
				inputData: Items{
					Name: "stm32f10x_stdperiph_lib",
					Owner: Owner{
						Login: "skywolf",
					},
					HTMLURL:         "https://github.com/skywolf/stm32f10x_stdperiph_lib",
					CreatedAt:       time.Date(2009, 9, 13, 10, 07, 22, 0, time.UTC),
					StargazersCount: 1,
				},
				wg: &sync.WaitGroup{},
			},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Worker(tt.args.inputData, tt.args.wg)
		})
	}
}

func TestGetAllRepo(t *testing.T) {
	type args struct {
		agent        *fiber.Agent
		from         time.Time
		searchString string
	}
	tests := []struct {
		name string
		args args
		want Items
	}{
		{
			name: "Test for GetAllRepo",
			args: args{
				agent:        fiber.AcquireAgent(),
				from:         time.Date(2009, 9, 13, 10, 07, 22, 0, time.UTC),
				searchString: "stm32",
			},
			want: Items{
				Name: "stm32f10x_stdperiph_lib",
				Owner: Owner{
					Login: "skywolf",
				},
				HTMLURL:         "https://github.com/skywolf/stm32f10x_stdperiph_lib",
				CreatedAt:       time.Date(2009, 9, 13, 10, 07, 22, 0, time.UTC),
				StargazersCount: 1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cgot := GetAllRepo(tt.args.agent, tt.args.from, tt.args.searchString)
			var got Items

			for i := range cgot {
				got = i
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllRepo() = %v, want %v", got, tt.want)
			}

			// if got := GetAllRepo(tt.args.agent, tt.args.from, tt.args.searchString); !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("GetAllRepo() = %v, want %v", got, tt.want)
			// }
		})
	}
}
