package main

type trie_node struct{
	is_leaf bool
	next [26]*trie_node
}

type trie struct{
	head *trie_node
}

func (tr *trie)insert(word string){
	t := tr.head
	for i:=0; i<len(word) ;i++{
		if(t.next[word[i]-'a'] == nil){
			t.next[word[i]-'a'] = new(trie_node)
		}
		t = t.next[word[i]-'a']
	}
	t.is_leaf = true
}

func main(){
	t := trie{}
	t.head = new(trie_node)
	t.insert("manoj")
	t.insert("taraka")
	t.insert("ramarao")
	t.insert("gunuru")
}