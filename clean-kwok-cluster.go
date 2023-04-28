package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func fetchClusterList() []string {
	cmd := exec.Command("kwokctl", "get", "clusters")
	data, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	allClusters := string(data)
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
		fmt.Println("Deleting cluster ", c)
		removeCluster(c)
	}
	fmt.Println("successfully removed all kwok clusters")
}
