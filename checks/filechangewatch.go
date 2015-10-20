package checks

import "fmt"
import "os"

func FileChangeWatch(f string) {

	file, err := os.Open(f)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	fileStat, _ := file.Stat()
	fmt.Println(fileStat.ModTime())

}
