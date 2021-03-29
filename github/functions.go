package github

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
	"unicode"
)

var (
	deprecatedRepos = [2]string{"https://github.com/go-martini/martini", "https://github.com/pilu/traffic"}
	repos           []Repo
)



func GetAccessToken(tokenFile string) string {
	tokenBytes, err := ioutil.ReadFile(tokenFile)
	if err != nil {
		log.Fatal("Error occurs when read access token file")
	}
	return strings.TrimSpace(string(tokenBytes))
}

func IsDeprecatedRepo(repoUrl string) bool {
	for _, deprecatedRepo := range deprecatedRepos {
		if repoUrl == deprecatedRepo {
			return true
		}
	}
	return false
}

func TrimSpaceAndSlash(r rune) bool {
	return unicode.IsSpace(r) || (r == rune('/'))
}

func SaveRanking(repos []Repo) {
	readme, err := os.OpenFile("../README.md", os.O_RDWR|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer readme.Close()
	_, _ = readme.WriteString(head)
	for _, repo := range repos {
		if IsDeprecatedRepo(repo.URL) {
			repo.Description = warning + repo.Description
		}
		_, _ = readme.WriteString(fmt.Sprintf("| [%s](%s) | %d | %d | %d | %s | %v |\n", repo.Name, repo.URL, repo.Stars, repo.Forks, repo.Issues, repo.Description, repo.LastCommitDate.Format("2006-01-02 15:04:05")))
	}
	_, _ = readme.WriteString(fmt.Sprintf(tail, time.Now().Format(time.RFC3339)))
}
