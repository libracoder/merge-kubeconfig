package main

import (
	"fmt"
	"io/fs"
	"os"
)

func main() {
	// Check if two kubeconfig files are provided as arguments
	if len(os.Args) != 3 {
		fmt.Println("Usage: merge_kubeconfig <kubeconfig_file1> <kubeconfig_file2>")
		os.Exit(1)
	}

	// Read kubeconfig file contents
	kubeconfigFile1, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("Error reading kubeconfig file '%s': %s\n", os.Args[1], err)
		os.Exit(1)
	}

	kubeconfigFile2, err := os.ReadFile(os.Args[2])
	if err != nil {
		fmt.Printf("Error reading kubeconfig file '%s': %s\n", os.Args[2], err)
		os.Exit(1)
	}

	// Merge the contents of kubeconfigFile2 into kubeconfigFile1
	mergedKubeconfig := string(kubeconfigFile1) + string(kubeconfigFile2)

	// Write the merged kubeconfig to a new file
	mergedFile := "merged_kubeconfig"
	err = os.WriteFile(mergedFile, []byte(mergedKubeconfig), fs.FileMode(0644))
	if err != nil {
		fmt.Printf("Error writing merged kubeconfig to file '%s': %s\n", mergedFile, err)
		os.Exit(1)
	}

	fmt.Printf("Merged kubeconfig file created: %s\n", mergedFile)
}
