package dictionary

import (
	"os"
	"bufio"
	"log"
)

var storage map[string]bool = make(map[string]bool)

func loadFile(name string) {
	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		storage[scanner.Text()] = true
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func DoesWordExist(w string) bool {
	if len(storage) == 0 {
		loadFile("/home/ekoira/projects/gotests/src/permutations/dictionary/nounlist.txt")
	}

	return storage[w]
}