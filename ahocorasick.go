package goahocorasick

import (
	"fmt"

	godarts "github.com/anknown/darts"
)

const FAIL_STATE = -1
const ROOT_STATE = 1

type Machine struct {
	trie    *godarts.DoubleArrayTrie
	failure []int
	output  map[int]([][]rune)
}

type Term struct {
	Pos  int
	Word []rune
}

func (m *Machine) Build(keywords [][]rune) (err error) {
	if len(keywords) == 0 {
		return fmt.Errorf("empty keywords")
	}

	d := new(godarts.Darts)

	trie := new(godarts.LinkedListTrie)
	m.trie, trie, err = d.Build(keywords)
	if err != nil {
		return err
	}

	m.output = make(map[int]([][]rune), 0)
	for idx, val := range d.Output {
		m.output[idx] = append(m.output[idx], val)
	}

	queue := make([](*godarts.LinkedListTrieNode), 0)
	m.failure = make([]int, len(m.trie.Base))
	for _, c := range trie.Root.Children {
		m.failure[c.Base] = godarts.ROOT_NODE_BASE
	}
	queue = append(queue, trie.Root.Children...)

	for {
		if len(queue) == 0 {
			break
		}

		node := queue[0]
		for _, n := range node.Children {
			if n.Base == godarts.END_NODE_BASE {
				continue
			}
			inState := m.f(node.Base)
		set_state:
			outState := m.g(inState, n.Code-godarts.ROOT_NODE_BASE)
			if outState == FAIL_STATE {
				inState = m.f(inState)
				goto set_state
			}
			if _, ok := m.output[outState]; ok != false {
				copyOutState := make([][]rune, 0)
				for _, o := range m.output[outState] {
					copyOutState = append(copyOutState, o)
				}
				m.output[n.Base] = append(copyOutState, m.output[n.Base]...)
			}
			m.setF(n.Base, outState)
		}
		queue = append(queue, node.Children...)
		queue = queue[1:]
	}

	return nil
}

func (m *Machine) PrintFailure() {
	fmt.Printf("+-----+-----+\n")
	fmt.Printf("|%5s|%5s|\n", "index", "value")
	fmt.Printf("+-----+-----+\n")
	for i, v := range m.failure {
		fmt.Printf("|%5d|%5d|\n", i, v)
	}
	fmt.Printf("+-----+-----+\n")
}

func (m *Machine) PrintOutput() {
	fmt.Printf("+-----+----------+\n")
	fmt.Printf("|%5s|%10s|\n", "index", "value")
	fmt.Printf("+-----+----------+\n")
	for i, v := range m.output {
		var val string
		for _, o := range v {
			val = val + " " + string(o)
		}
		fmt.Printf("|%5d|%10s|\n", i, val)
	}
	fmt.Printf("+-----+----------+\n")
}

func (m *Machine) g(inState int, input rune) (outState int) {
	if inState == FAIL_STATE {
		return ROOT_STATE
	}

	t := inState + int(input) + godarts.ROOT_NODE_BASE
	if t >= len(m.trie.Base) {
		if inState == ROOT_STATE {
			return ROOT_STATE
		}
		return FAIL_STATE
	}
	if inState == m.trie.Check[t] {
		return m.trie.Base[t]
	}

	if inState == ROOT_STATE {
		return ROOT_STATE
	}

	return FAIL_STATE
}

func (m *Machine) f(index int) (state int) {
	return m.failure[index]
}

func (m *Machine) setF(inState, outState int) {
	m.failure[inState] = outState
}

func (m *Machine) MultiPatternSearch(content []rune, returnImmediately bool, n_noncontinue_chars int) [](*Term) {
	terms := make([](*Term), 0)

	state := ROOT_STATE
	prev_state := state
	noncontion_char_size := 0
	for pos, c := range content {
	start:
		if m.g(state, c) == FAIL_STATE {
			if noncontion_char_size < n_noncontinue_chars {
				noncontion_char_size += 1
				state = prev_state
				continue
			}
			state = m.f(state)
			goto start
		} else {
			state = m.g(state, c)
			prev_state = state
			noncontion_char_size = 0
			if val, ok := m.output[state]; ok != false {
				for _, word := range val {
					term := new(Term)
					term.Pos = pos - len(word) + 1 - noncontion_char_size
					term.Word = word
					terms = append(terms, term)
					if returnImmediately {
						return terms
					}
				}
				state = m.f(state)
				goto start
			}
		}
	}

	return terms
}

func (m *Machine) ExactSearch(content []rune) [](*Term) {
	if m.trie.ExactMatchSearch(content, 0) {
		t := new(Term)
		t.Word = content
		t.Pos = 0
		return [](*Term){t}
	}

	return nil
}
