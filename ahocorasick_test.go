package goahocorasick

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
)

func Read(filename string) ([][]rune, error) {
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

func TestBuild(t *testing.T) {
	keywords, err := Read("test_keywords_eng")
	if err != nil {
		t.Error(err)
	}

	m := new(Machine)
	m.Build(keywords)
	//m.PrintFailure()
	//m.PrintOutput()
}

func TestMultiPatternSearchEnglish(t *testing.T) {
	fmt.Printf("===> MultiPattern Search For English \n")
	keywords, err := Read("test_keywords_eng")
	if err != nil {
		t.Error(err)
	}
	m := new(Machine)
	m.Build(keywords)
	//m.PrintFailure()
	//m.PrintOutput()

	content := []rune("ushers")
	terms := m.MultiPatternSearch(content, false)
	for _, term := range terms {
		fmt.Printf("find %s @%d in %s\n", string(term.Word), term.Pos, string(content))
	}
	fmt.Printf("\n")
}

func TestMultiPatternSearchChinese(t *testing.T) {
	fmt.Printf("===> MultiPattern Search For Chinese \n")
	keywords, err := Read("test_keywords_chn")
	if err != nil {
		t.Error(err)
	}
	m := new(Machine)
	m.Build(keywords)
	//m.PrintFailure()
	//m.PrintOutput()

	content := []rune("你不会想到阿拉伯人会踢出阿根廷风格的足球更何况是埃及风格")
	terms := m.MultiPatternSearch(content, false)
	for _, term := range terms {
		fmt.Printf("find %s @%d in %s\n", string(term.Word), term.Pos, string(content))
	}
	fmt.Printf("\n")
}

func TestExactSearchEnglish(t *testing.T) {
	fmt.Printf("===> Exact Search For English\n")
	keywords, err := Read("test_keywords_eng")
	if err != nil {
		t.Error(err)
	}
	m := new(Machine)
	m.Build(keywords)

	for _, k := range keywords {
		if m.ExactSearch(k) == nil {
			t.Error("exact search chinese failed")
		}
	}
	fmt.Printf("Test total:%d words\n\n", len(keywords))
}

func TestExactSearchChinese(t *testing.T) {
	fmt.Printf("===> Exact Search For Chinese\n")
	keywords, err := Read("test_keywords_chn")
	if err != nil {
		t.Error(err)
	}
	m := new(Machine)
	m.Build(keywords)

	for _, k := range keywords {
		if m.ExactSearch(k) == nil {
			t.Error("exact search chinese failed")
		}
	}
	fmt.Printf("Test total:%d words\n\n", len(keywords))
}
