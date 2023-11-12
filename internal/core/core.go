package core

import (
	"encoding/json"
	"fmt"
	"github_parser/internal/config"
	"github_parser/internal/files"
	"github_parser/internal/limit"
	"os"
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

func makeDir(dirname string) {
	_, err := os.Stat("files/" + dirname)
	if os.IsNotExist(err) {
		if err := os.Mkdir("files/"+dirname, os.ModePerm); err != nil {
			//log.Fatal(err)
			fmt.Println(err)
		}
	}
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
	firstRepoData := getData(agent, "https://api.github.com/search/repositories?q=stm32&sort=updated&order=asc&per_page=1&page=1")
	//firstRepoData := getData(agent, "https://api.github.com/search/repositories?q=stm32+created%3A2015-04-25&sort=updated&order=asc&per_page=100&page=1")

	return firstRepoData
}

func GetAllRepo(agent *fiber.Agent, from time.Time) {
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	start := time.Now()
	starString := ""
	created := "+created%3A"

	for from.Before(start) {

		date := ""
		// из за лимита запросов (30 поисковых запросов в минуту) спим полторы секунды после каждого запроса

		getLimits := limit.GetLimit(agent)

		fmt.Printf("\nRemaining resources: %#+v, ", getLimits.Resources.Search.Remaining)
		fmt.Printf("Used resources: %#+v\n", getLimits.Resources.Search.Used)

		// Создаем строку из 10 дней
		for i := 0; i < 20; i++ {
			date = date + created + from.Format(time.DateOnly)
			from = from.Add(time.Hour * 24)
		}

		fmt.Printf("From date: %#+v\n", from.Format(time.DateOnly))
		//fmt.Printf("From date: %#+v\n", date)
		fmt.Printf("Active goroutines: %#+v\n", runtime.NumGoroutine())
		for i := 1; i < 11; i++ {
			allrepos := getData(agent, "https://api.github.com/search/repositories?q=stm32"+date+"&per_page=100&page="+strconv.Itoa(i))
			fmt.Printf("Files num: %#+v\n", allrepos.TotalCount)

			for key, val := range allrepos.Items {
				wg.Add(1)
				go func(key int, val Items) {
					mu.Lock()
					//makeDir(val.Owner.Login)
					if val.StargazersCount != 0 {
						starString = " [s-" + strconv.Itoa(val.StargazersCount) + "]"
					} else {
						starString = ""
					}

					//fullName := "files/" + val.Owner.Login + "/" + val.Name + starString + ".zip"
					fullName := "files/" + val.Name + " (" + val.Owner.Login + ")" + starString + ".zip"
					fileExist, _ := files.Exists(fullName)
					mu.Unlock()
					if !fileExist {
						fmt.Printf("Goroutine %v started downloading file: %v\n", key, fullName)

						GetFile := files.GetFileFromURL(val.HTMLURL + "/archive/refs/heads/master.zip")
						checkFile := strings.Contains(GetFile, `<!DOCTYPE html>`)
						if !checkFile {
							saveZipFile := files.SaveFile(fullName, GetFile)
							if saveZipFile != nil {
								files.SaveTxt("download_error.txt", fullName)
								fmt.Println(saveZipFile)
							}
						}
					}

					wg.Done()
				}(key, val)
			}

			wg.Wait()
			pause := 61 - int(time.Since(start).Seconds())
			if getLimits.Resources.Search.Used > 26 {
				fmt.Printf("Pause %#+v seconds\n", pause)
				for i := pause; i > 0; i-- {

					fmt.Printf("\rRemaining %2v seconds", i)
					time.Sleep(time.Second * 1)
				}
				start = time.Now()
			}

			if allrepos.TotalCount < 101 {
				break
			}
		}

	}
}
