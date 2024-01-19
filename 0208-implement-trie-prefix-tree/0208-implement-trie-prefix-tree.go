type Trie struct {
    ch       [26]*Trie
    wordEnd  bool
    cntWords int
}


func Constructor() Trie {
    return Trie{}
}


func (this *Trie) Insert(word string)  {
    root := this
    for _, c := range word {
        if root.ch[c-'a'] == nil {
            root.ch[c-'a'] = &Trie{}
        }
        root = root.ch[c-'a']
        root.cntWords++
    }
    root.wordEnd = true
}


func (this *Trie) Search(word string) bool {
    root := this
    for _, c := range word {
        if root.ch[c-'a'] == nil {
            return false
        }
        root = root.ch[c-'a']
    }
    return root.wordEnd
}


func (this *Trie) StartsWith(prefix string) bool {
    root := this
    for _, c := range prefix {
        if root.ch[c-'a'] == nil {
            return false
        }
        root = root.ch[c-'a']
    }
    return root.cntWords > 0
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */