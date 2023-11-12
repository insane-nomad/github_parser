package core

import (
	"encoding/json"
	"fmt"
	"github_parser/internal/config"
	"github_parser/internal/files"
	"github_parser/internal/limit"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
)

type RepositoryData struct {
	TotalCount int     `json:"total_count"`
	Items      []Items `json:"items"`
}

type Owner struct {
	Login string `json:"login"`
}

type Items struct {
	Name            string    `json:"name"`
	Owner           Owner     `json:"owner"`
	HTMLURL         string    `json:"html_url"`
	CreatedAt       time.Time `json:"created_at"`
	StargazersCount int       `json:"stargazers_count"`
}

func getData(agent *fiber.Agent, url string) RepositoryData {
	var repoData RepositoryData
	// добавляем к запросу аутентификационный токен
	agent.Add("Authorization", "Bearer "+config.GithubToken)
	req := agent.Request()
	req.Header.SetMethod(fiber.MethodGet)
	// отправляем запрос
	req.SetRequestURI(url)

	if err := agent.Parse(); err != nil {
		fmt.Println("Parse error: ", err)
	}

	code, response, errs := agent.Bytes()

	if code != fiber.StatusOK {
		fmt.Println("Return code:", code, " from core file")
	}

	if errs != nil {
		fmt.Println(errs)
	}

	if err := json.Unmarshal(response, &repoData); err != nil {
		fmt.Println(err)
	}
	return repoData
}

func GetfirstRepo(agent *fiber.Agent) RepositoryData {
	// "2016-05-08"
	//firstRepoData := getData(agent, "https://api.github.com/search/repositories?q=stm32&sort=updated&order=asc&per_page=1&page=1")
	firstRepoData := getData(agent, "https://api.github.com/search/repositories?q=stm32+created%3A2016-05-18&sort=updated&order=asc&per_page=1&page=1")

	return firstRepoData
}

func Worker(inputData Items, wg *sync.WaitGroup) {
	defer wg.Done()
	starString := ""
	//fmt.Printf("--------%v\n", inputData)
	if inputData.Owner.Login != "" {
		if inputData.StargazersCount != 0 {
			starString = " [s-" + strconv.Itoa(inputData.StargazersCount) + "]"
		} else {
			starString = ""
		}

		//fullName := "files/" + val.Owner.Login + "/" + val.Name + starString + ".zip"
		fullName := "files/" + inputData.Name + " (" + inputData.Owner.Login + ")" + starString + ".zip"
		fileExist, _ := files.Exists(fullName)

		if !fileExist {
			fmt.Printf("Goroutine started downloading file: %v\n", fullName)
			GetFile := files.GetFileFromURL(inputData.HTMLURL + "/archive/refs/heads/master.zip")
			checkFile := strings.Contains(GetFile, `<!DOCTYPE html>`)
			if !checkFile {
				saveZipFile := files.SaveFile(fullName, GetFile)
				if saveZipFile != nil {
					files.SaveTxt("download_error.txt", fullName)
					fmt.Println(saveZipFile)
				}
			} else {
				files.SaveTxt("url_error.txt", inputData.HTMLURL+"/archive/refs/heads/master.zip")
			}
			fmt.Printf("Goroutine comleted downloading file: %v\n", fullName)
		}
	}
}

func GetAllRepo(agent *fiber.Agent, from time.Time) chan Items {
	outputChan := make(chan Items, 100)
	start := time.Now()
	created := "+created%3A"
	date := ""
	getLimits := limit.GetLimit(agent)
	defer close(outputChan)

	fmt.Printf("\nRemaining resources: %#+v, ", getLimits.Resources.Search.Remaining)
	fmt.Printf("Used resources: %#+v\n", getLimits.Resources.Search.Used)

	// Создаем строку из 10 дней
	for i := 0; i < 1; i++ {
		date = date + created + from.Format(time.DateOnly)
	}

	fmt.Printf("From date: %#+v\n", from.Format(time.DateOnly))
	fmt.Printf("Active goroutines: %#+v\n", runtime.NumGoroutine())
	for i := 1; i < 11; i++ {
		allrepos := getData(agent, "https://api.github.com/search/repositories?q=stm32"+date+"&per_page=100&page="+strconv.Itoa(i))
		fmt.Printf("Files num: %#+v\n", allrepos.TotalCount)
		for _, val := range allrepos.Items {
			outputChan <- val
		}

		if allrepos.TotalCount < 101 {
			break
		}
	}

	pause := 61 - int(time.Since(start).Seconds())
	if getLimits.Resources.Search.Used > 26 {
		fmt.Printf("Pause %#+v seconds\n", pause)
		for i := pause; i > 0; i-- {
			fmt.Printf("\rRemaining %2v seconds", i)
			time.Sleep(time.Second * 1)
		}
	}

	return outputChan
}
