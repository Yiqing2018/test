package main

// the node of Trie
type TrieNode struct {
	children map[string]*TrieNode
	endInfo    string
	matchPriority int
}

// Trie
type Trie struct {
    root *TrieNode
}

var matchAll string = "*"

func buildTire(inputs [][]string) *Trie {
	trie := NewTrie()
	for _ , input := range inputs{
		trie.insert(input)
	}
	return trie
}

func (this *Trie) search(input []string) string {
	var nodes []*TrieNode
	nodes = append(nodes, this.root)
	for _, word := range input {
		var matches []*TrieNode
		for _, node := range nodes {
			for _, item := range []string{word, matchAll} {
				_, ok := node.children[item]
				if ok {
					matches = append(matches, node.children[item])
				}
			}
		}
		nodes = matches
	}
	priority := -1
	result := "no avaliable result"
	for _, node := range nodes{
		end := node.endInfo
		if node.matchPriority > priority{
			result = end
			priority = node.matchPriority
		}

	}
	return result
}

func (this *Trie) insert(input []string){
	node := this.root
	counter := 0
	for i := 0; i < inputLength; i++ {
		word := input[i]
		_, ok := node.children[word]
		if !ok {
			node.children[word] = newTrieNode()
		}
		node = node.children[word]
		if word != matchAll{
			counter++
		}
	}
	node.endInfo = input[inputLength]
	node.matchPriority = counter

}

func newTrieNode() *TrieNode{
	return &TrieNode{
		children: make(map[string]*TrieNode), 
		endInfo: "",
		matchPriority: -1}
}

func NewTrie() *Trie {
    return &Trie{root: newTrieNode()}
}