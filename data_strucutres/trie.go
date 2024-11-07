package main

import "fmt"

// note that this trie implementation supports only characters a-z or 1-26

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

func (tr *trie)search(pattern string) bool{
	t:= tr.head
	for i:=0; i<len(pattern); i++{
		if(t.next[pattern[i]-'a'] == nil){
			return false;
		}
		t = t.next[pattern[i]-'a']
	}
	return true;
}

func (tr * trie)PrintDictWords(word string)bool{
	t:= tr.head
	for i:=0; i<len(word); i++{
		if(t.next[word[i]-'a'] == nil){
			return false;
		}
		t = t.next[word[i]-'a']
	}
	if(t.is_leaf){
		return true
	}
	return false
}

func main(){
	t := trie{}
	t.head = new(trie_node)
	t.insert("manoj")
	t.insert("taraka")
	t.insert("ramarao")
	t.insert("gunuru")
	t.insert("rajni")

	lt := [6]string{"man","tara","rama","rao","ka","gun",}
	for _,v := range(lt){
		if(t.search(v)){
			fmt.Println("YES PATTERN IS THERE")
		}else{
			fmt.Println("PATTERN IS NOT THERE")
		}
	}
	lt1 := [3]string{"manoj","rajni","kamal",}
	for _,v:= range lt1{
		if(t.PrintDictWords(v)){
			fmt.Println(v)
		}else{
			fmt.Println("WORD IS NOT THERE")
		}
	}
}