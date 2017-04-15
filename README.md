## Aho–Corasick algorithm

#### Intro

A faster and more effective implement of *Aho-Corasick algorithm* in golang and supports Both Chinese and English. To improve the performance and reduce memory usage, the program uses *Double Array Trie* instead of common *Linked List Trie*. In the benchmark, `it is 10 times faster than the most popular AC algorithm implement in golang @ github and tenth of its memory usage`. You can find more information in the benchmark parts.

This Project is inspired by [hankcs/AhoCorasickDoubleArrayTrie](https://github.com/hankcs/AhoCorasickDoubleArrayTrie)

Besides Multi-Pattern Search using AC algorithm, the program also provides "exact match search" using Double Array Trie

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
	    "os"
	 )

	import (
		"github.com/anknown/ahocorasick"
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
	    dict, err := ReadRunes("your_dict_files")
	    if err != nil {
	        fmt.Println(err)
	        return
	    }
	
	    content := []rune("your text")
	
	    m := new(goahocorasick.Machine)
	    if err := m.Build(dict); err != nil {
	        fmt.Println(err)
	        return
	    }
	
	    terms := m.MultiPatternSearch(content, false)
	    for _, t := range terms {
	        fmt.Printf("%d %s\n", t.Pos, string(t.Word))
	    }
	}

I do not provide read file API because I think your dict may coming form other source

#### Benchmark

**Multi-Pattern Search**

compare with `cloudflare/ahocorasick` who receives most stars and forks in all the implements written in golang

To Run Benchmark, go to test dir 

	go build benchmark.go
	
	./benchmark


* For Chinese Test

*Dictionary* contains `153,151` words, *Text* contains `777,277` words

	====================================================================
				cost(million sec)	memory usage(MBytes)
	cloudflare/ahocorasick		28926				1911
	anknown/ahocorasick		1814				155
	====================================================================

* For English Test

*Dictionary* contains `127,141` words, *Text* contains `674,669` words

	====================================================================
				time(million sec)	memory usage(MBytes)
	cloudflare/ahocorasick		19835				1340
	anknown/ahocorasick		1619				203
	====================================================================

#### License

MIT License
