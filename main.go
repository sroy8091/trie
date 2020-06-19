// TODO: use file for word dictionary
// TODO: use goroutines for filing up the trie ds
// TODO: add test suite for unit test

package main

import "fmt"

type trie struct {
	character rune
	childrens []*trie
	is_end    bool
}

var found_in_child bool
var node *trie
var d *trie

func create_new_node(c rune) *trie {
	return &trie{c, nil, false}
}

func insert(word string, root *trie) {
	node = root
	for _, c := range word {
		//fmt.Printf(" => %c", c)
		found_in_child = false
		for _, d := range node.childrens {
			if d.character == c {
				found_in_child = true
				node = d
				break
			}
		}
		if !found_in_child {
			new_node := create_new_node(c)
			//fmt.Printf(" => %c", c)
			node.childrens = append(node.childrens, new_node)
			node = new_node
		}

	}
	node.is_end = true
}

func autocomplete_util(node *trie, prefix string, res *[]string) {
	if node.is_end{
		*res = append(*res, prefix)
	}
	for _, i := range node.childrens {
		autocomplete_util(i, prefix+string(i.character), res)
	}
}


func autocomplete(root *trie, prefix string) {
	node = root
	res := make([]string, 0)
	for _, j := range prefix {
		found := false
		for _, d = range node.childrens {
			if j == d.character {
				found = true
				break
			}
		}
		if !found {
			fmt.Println("Prefix not found")
			return
		}
		node = d
	}
	autocomplete_util(node, prefix, &res)
	fmt.Println(res)
}

func main() {
	fmt.Println("Hello, World!")
	root := create_new_node('*')
	insert("dog", root)
	insert("dear", root)
	insert("deal", root)
	insert("abstract", root)
	insert("deer", root)
	insert("absurd", root)
	insert("anatomy", root)
	autocomplete(root, "de")
	autocomplete(root, "abs")
	autocomplete(root, "dea")
	autocomplete(root, "an")
	autocomplete(root, "a")
	autocomplete(root, "ca")
	fmt.Println("And yes I didn't removed the hello world")
}
