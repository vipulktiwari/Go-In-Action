package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
)

const (
   //ALBHABET_SIZE total characters in english alphabet
    ALBHABET_SIZE = 26
)

type trieNode struct {
    childrens [ALBHABET_SIZE]*trieNode
	isWordEnd bool
	count int
}

type trie struct {
    root *trieNode
}

func initTrie() *trie {
    return &trie{
        root: &trieNode{},
    }
}

func (t *trie) insert(word string) {
    wordLength := len(word)
    current := t.root
    for i := 0; i < wordLength; i++ {
        index := word[i] - 'a'
        if current.childrens[index] == nil {
            current.childrens[index] = &trieNode{}
        }
        current = current.childrens[index]
    }
	current.isWordEnd = true
	current.count += 1
}

func display(root *trieNode, str []rune,w *bufio.Writer){
    if(root.isWordEnd ==true){
		s := fmt.Sprintf("%s,%d\n",string(str), root.count)
		w.WriteString(s)
	}
    
    for i := 0; i < ALBHABET_SIZE; i++ {
        if (root.childrens[i] != nil){
			str = append(str,rune(i+'a'))
			display(root.childrens[i],str,w)
		}
    }
}

func main() {
	trie := initTrie()
	file, err := os.Open("/Users/vipul/workspace/go/src/vipulktiwari/Go-In-Action/pubmatic/Data/names.txt")
    if err != nil {
        log.Fatal(err)
    }
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
    for scanner.Scan() {
		str := scanner.Text()
		trie.insert(str)
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
	// write to output file using display utility
	var str []rune
	display(trie.root,str, w)

}