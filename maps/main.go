package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

/*
1. maps are key-value stores
2. unorderd group of elements of one type (element type)
	1. indexed by unique keys (key type)
	2. uninitialized map is nil

*/

func main() {
	// dictionary := "http://www-01.sil.org/linguistics/wordlists/english/wordlist/wordsEn.txt"
	// hashWords(dictionary)
	moby := "http://www.gutenberg.org/files/2701/old/moby10b.txt"
	hashWords(moby)
	// n := hashBucket("Go", 12)
	// fmt.Println(n)
}

func example1() {
	// map[key type]value type
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map (before): ", m)
	delete(m, "k2")
	fmt.Println("map (after): ", m)

	v, ok := m["k2"]
	fmt.Println("ok?:", ok, v)
}

func initialization() {
	var makeGreeting = make(map[string]string)
	makeGreeting["Zac"] = "Good morning"
	makeGreeting["Jacque"] = "Bon jour"
	fmt.Println("makeGreeting:", makeGreeting)

	anotherMap := map[string]string{
		"init": "values",
		"two":  "second",
	}

	fmt.Println("anotherMap: ", anotherMap)

	// single line below narrows scope of val and exist
	// to just this if statement
	if val, exists := makeGreeting["test"]; exists {
		delete(makeGreeting, "test")
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	} else {
		fmt.Println("That value doesn't exist")
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	}

	for key, val := range anotherMap {
		fmt.Println(key, "-", val)
	}
}

func hashWords(url string) {
	res, err := http.Get(url)

	if err != nil {
		log.Fatalln(err)
	}

	words := make(map[int][]string, 12)

	sc := bufio.NewScanner(res.Body)
	defer res.Body.Close()

	sc.Split(bufio.ScanWords)

	for sc.Scan() {
		word := strings.ToLower(string(sc.Text()))
		bucket := hashBucket(word, 12)
		words[bucket] = append(words[bucket], word)
	}

	for k, v := range words {
		fmt.Println(k, " - ", len(v))
	}

	if err := sc.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
}

func hashTable() {

}

func hashBucket(word string, buckets int) int {
	letter := int(word[0])
	return letter % buckets
}
