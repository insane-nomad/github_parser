// main.exe | tee mylog.txt
package main

import (
	"github_parser/internal/core"

	"github.com/gofiber/fiber/v2"
)

func main() {
	agent := fiber.AcquireAgent()
	// получаем дату самого первого размещенного репозитория
	getFirstRepoData := core.GetfirstRepo(agent).Items[0].CreatedAt
	// получаем остальные репозитории от даты первого репозитория
	core.GetAllRepo(agent, getFirstRepoData)

}
