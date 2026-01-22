class RandomizedSet {
public:
    set <int> storage;
    RandomizedSet() {
        
    }
    
    bool insert(int val) {
        if (storage.find(val) != storage.end()) {
            return false;
        }
        storage.insert(val);
        return true;
    }
    
    bool remove(int val) {
        if (storage.find(val) == storage.end()) {
            return false;
        }
        storage.erase(val);
        return true;
    }
    
    int getRandom() {
        int max = storage.size()-1;
        int min = 0;
        int range = max - min + 1;
        int num = rand() % range + min;
        return *next(storage.begin(), num);
    }
};

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * RandomizedSet* obj = new RandomizedSet();
 * bool param_1 = obj->insert(val);
 * bool param_2 = obj->remove(val);
 * int param_3 = obj->getRandom();
 */