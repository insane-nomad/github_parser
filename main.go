// main.exe | tee mylog.txt
package main

import (
	"fmt"
	"github_parser/internal/core"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	var searchString string
	fmt.Println("Введите поисковый запрос: ")
	fmt.Scanf("%s\n", &searchString)

	agent := fiber.AcquireAgent()
	start := time.Now()
	wg := &sync.WaitGroup{}
	// 2018-06-25
	// получаем дату самого первого размещенного репозитория
	getFirstRepoData := core.GetfirstRepo(agent, searchString).Items[0].CreatedAt

	for getFirstRepoData.Before(start) {
		// идем от даты размещения первого репозитория по текущую дату
		for data := range core.GetAllRepo(agent, getFirstRepoData, searchString) {
			wg.Add(1)
			go core.Worker(data, wg)
		}
		getFirstRepoData = getFirstRepoData.Add(time.Hour * 24)
		wg.Wait()
	}
}
