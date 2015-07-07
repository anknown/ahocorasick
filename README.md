## Aho–Corasick algorithm

#### Intro

An implement of *Aho-Corasick algorithm* in golang and supports Both Chinese and English. To improve the performance and reduce memory usage, the program uses *Double Array Trie* instead of common *Linked List Trie*. In the benchmark, `it is 10 times faster than the most popular AC algorithm implement in golang @ github and tenth of its memory usage`. You can find more information in the benchmark parts.

Besides Multi-Pattern Search using AC algorithm, the program also provide exact match search using Double Array Trie

Aho-Corasick algorithm is first presented in the paper below:

> [Efficient string matching: an aid to bibliographic search](http://dl.acm.org/citation.cfm?id=360855)

the wikipedia link is: [aho-corasick algorithm](https://en.wikipedia.org/wiki/Aho%E2%80%93Corasick_algorithm)

#### Usage

**Multi-Pattern Search Example**

	package main

	import (
	    "bufio"
	    "bytes"
	    "fmt"
	    "io"
	    "io/ioutil"
	    "os"
	    "time"	
	 )

	func ReadRunes(filename string) ([][]rune, error) {
	    dict := [][]rune{}

	    f, err := os.OpenFile(filename, os.O_RDONLY, 0660)
	    if err != nil {
	        return nil, err
	    }

	    r := bufio.NewReader(f)
	    for {
	        l, err := r.ReadBytes('\n')
	        if err != nil || err == io.EOF {
	            break
	        }
	        l = bytes.TrimSpace(l)
	        dict = append(dict, bytes.Runes(l))
	    }

	    return dict, nil
	}

	func main() {
	    dict, err := ReadRunes(your_dict_files)
	    if err != nil {
	        fmt.Println(err)
	        return
	    }

	    content, err := ioutil.ReadFile(your_text)
	    if err != nil {
	        fmt.Println(err)
	        return
	    }

	    contentRune := bytes.Runes([]byte(content))
    
	    m := new(goahocorasick.Machine)
	    if err := m.Build(dict); err != nil {
	        fmt.Println(err)
	        return
	    }
	    
	    terms := m.MultiPatternSearch(contentRune)
        for _, t := range terms {
            fmt.Printf("%d %s\n", t.Pos, string(t.Word))
        }
	}

#### Benchmark

**Multi-Pattern Search**

compare with *cloudflare/ahocorasick* which receive most star and forks in the implement which is written in golang

To Run Benchmark, go to test dir 

	go build test.go
	
	./test


* For Chinese Test

*Dictionary* contains `153,151` words, *Text* contains `777,277` words

	====================================================================
								cost(million sec)	memory usage(MBytes)
	cloudflare/ahocorasick		28926				1911M
	anknown/ahocorasick			1814				155M
	====================================================================

* For English Test

*Dictionary* contains `127,141` words, *Text* contains `674,669` words

	====================================================================
								time(million sec)	memory usage(MBytes)
	cloudflare/ahocorasick		19835				1340M
	anknown/ahocorasick			1619				203M
	====================================================================
