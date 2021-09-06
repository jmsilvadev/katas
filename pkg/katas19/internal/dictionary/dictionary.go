package dictionary

import (
	"bufio"
	"fmt"
	"os"
)

//Load get words with the input length from the dictionary in the especific path in the filesystem
func Load(wordLength int, path string) []string {
	fw, err := os.Open(path)
	if err != nil {
		fmt.Println("An error occurs when load dictionary from storage", err)
		return nil
	}
	defer fw.Close()

	var filteredDic []string

	dic := bufio.NewScanner(fw)
	for dic.Scan() {
		if l := len(dic.Text()); l == wordLength {
			filteredDic = append(filteredDic, dic.Text())
		}
	}
	return filteredDic
}
