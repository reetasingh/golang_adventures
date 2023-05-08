package main

import (
	"fmt"
	"os/exec"
	"strings"
)

const NO_CLUSTERS_FOUND = "No clusters found"

func fetchClusterList() []string {
	cmd := exec.Command("kwokctl", "get", "clusters")
	data, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	allClusters := string(data)
	if allClusters == NO_CLUSTERS_FOUND {
		return []string{}
	}
	clusters := strings.Split(allClusters, "\n")
	return clusters
}

func removeCluster(name string) {
	cmd := exec.Command("kwokctl", "delete", "cluster", "--name", name)
	_, err := cmd.Output()
	if err != nil {
		panic(err)
	}
}

func main() {
	clusterList := fetchClusterList()
	for _, c := range clusterList {
		if len(c) == 0 {
			continue
		}
		fmt.Println("Deleting cluster ", c)
		removeCluster(c)
	}
	fmt.Println("successfully removed all kwok clusters")
}
