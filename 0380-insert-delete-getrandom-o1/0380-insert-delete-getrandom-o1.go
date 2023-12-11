type RandomizedSet struct {
    data map[int]struct{}
}


func Constructor() RandomizedSet {
    rand.Seed(time.Now().UnixNano())
    return RandomizedSet{
        data: make(map[int]struct{}),
    }
}


func (this *RandomizedSet) Insert(val int) bool {
    if _, ok := this.data[val]; !ok {
        this.data[val] = struct{}{}
        return true
    }
    return false
}


func (this *RandomizedSet) Remove(val int) bool {
    if _, ok := this.data[val]; ok {
        delete(this.data, val)
        return true
    }
    return false
}


func (this *RandomizedSet) GetRandom() int {
    idx := rand.Intn(len(this.data))
    now := 0
    for val := range this.data {
        if now == idx {
            return val
        }
        now++
    }
    return -1
}


/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */