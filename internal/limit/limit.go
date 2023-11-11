package limit

import (
	"encoding/json"
	"fmt"
	"github_parser/internal/config"

	"github.com/gofiber/fiber/v2"
)

type Limits struct {
	Resources Resources `json:"resources"`
	Rate      Rate      `json:"rate"`
}
type Core struct {
	Limit     int `json:"limit"`
	Used      int `json:"used"`
	Remaining int `json:"remaining"`
	Reset     int `json:"reset"`
}
type Search struct {
	Limit     int `json:"limit"`
	Used      int `json:"used"`
	Remaining int `json:"remaining"`
	Reset     int `json:"reset"`
}

type Scim struct {
	Limit     int `json:"limit"`
	Used      int `json:"used"`
	Remaining int `json:"remaining"`
	Reset     int `json:"reset"`
}

type Resources struct {
	Core   Core   `json:"core"`
	Search Search `json:"search"`
	Scim   Scim   `json:"scim"`
}

type Rate struct {
	Limit     int `json:"limit"`
	Used      int `json:"used"`
	Remaining int `json:"remaining"`
	Reset     int `json:"reset"`
}

func GetLimit(agent *fiber.Agent) Limits {
	var limitData Limits

	agent.Add("Authorization", "Bearer "+config.GithubToken)
	req := agent.Request()
	req.Header.SetMethod(fiber.MethodGet)
	req.SetRequestURI("https://api.github.com/rate_limit")

	if err := agent.Parse(); err != nil {
		fmt.Println("Parse error: ", err)
	}

	code, response, errs := agent.Bytes()

	if code != fiber.StatusOK {
		fmt.Println("Return code:", code, " from limit file")
	}

	if errs != nil {
		fmt.Println(errs)
	}

	if err := json.Unmarshal(response, &limitData); err != nil {
		fmt.Println(err)
	}

	return limitData

}
