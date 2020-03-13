package main
import (
	"bufio"
	"strconv"
    "fmt"
    "log"
    "os"
)

func main() {
	//open input file
	file, err := os.Open("/Users/vipul/workspace/go/src/vipulktiwari/Go-In-Action/wordCount/Data/names.txt")
    if err != nil {
        log.Fatal(err)
    }
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	wordCount := make(map[string]int)
    for scanner.Scan() {
		str := scanner.Text()
		n, ok := wordCount[str]
		if !ok {
			wordCount[str] = 1
		}else{
			wordCount[str] = n + 1
		}
    }
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
	}
	
	//write to output file
	f, err := os.OpenFile("out.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Println(err)
    }
    defer f.Close()
    w := bufio.NewWriter(f)
    defer w.Flush()
    for key, val := range wordCount {
		str := fmt.Sprintf("%s,%s\n", key, strconv.Itoa(val))
		w.WriteString(str)
    }
}