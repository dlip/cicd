package main

import (
	"fmt"
	"io/ioutil"

	api "github.com/gogits/go-gogs-client"
)

func main() {
	c := api.NewClient("http://localhost:3000", "526886d331fde8e06574c9dd006b16756c0e7cf5")
	orgs, _ := c.ListMyOrgs()
	for _, org := range orgs {
		fmt.Println(org.UserName)
		repos, _ := c.ListOrgRepos(org.UserName)
		for _, repo := range repos {
			fmt.Println(repo.Name)
			branches, _ := c.ListRepoBranches(org.UserName, repo.Name)
			for _, branch := range branches {
				fmt.Println(branch.Name)
				file, _ := c.GetFile(org.UserName, repo.Name, branch.Commit.ID, "concourse.yml")
				ioutil.WriteFile(fmt.Sprintf("%s-%s-%s.yml", org.UserName, repo.Name, branch.Name), file, 0664)
			}

		}
	}

}
