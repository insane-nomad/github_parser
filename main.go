package main

import (
	"fmt"
	"github_parser/internal/limit"

	"github.com/gofiber/fiber/v2"
)

func main() {
	agent := fiber.AcquireAgent()
	//https://api.github.com/search/repositories?q=stm32+created%3A2022-03-01+created%3A2022-03-02+created%3A2022-03-03+created%3A2022-03-04+created%3A2022-03-05&per_page=100&page=1
	getLimits := limit.GetLimit(agent)

	fmt.Printf("%#+v", getLimits)
}
