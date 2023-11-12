// main.exe | tee mylog.txt
package main

import (
	"github_parser/internal/core"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	agent := fiber.AcquireAgent()
	start := time.Now()
	wg := &sync.WaitGroup{}
	// получаем дату самого первого размещенного репозитория
	getFirstRepoData := core.GetfirstRepo(agent).Items[0].CreatedAt

	for getFirstRepoData.Before(start) {
		// идем от даты размещения первого репозитория по текущую дату
		getFirstRepoData = getFirstRepoData.Add(time.Hour * 24)
		for data := range core.GetAllRepo(agent, getFirstRepoData) {
			wg.Add(1)
			go core.Worker(data, wg)
		}
		wg.Wait()
	}
}
