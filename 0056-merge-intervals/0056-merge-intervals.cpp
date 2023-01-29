class Solution {
public:
    vector<vector<int>> merge(vector<vector<int>>& intervals) {
        vector <vector <int>> res;
        sort(intervals.begin(), intervals.end(), cmp);
        vector <int> cur = intervals[0];
        for (int i = 1; i < intervals.size(); i++) {
            if (intervals[i][0] > cur[1]) {
                res.push_back(cur);
                cur = intervals[i];
            } else {
                cur[1] = max(cur[1], intervals[i][1]);
            }
        }
        res.push_back(cur);
        return res;
    }
    static bool cmp(vector <int> &x, vector <int> &y) {
        return x[0] < y[0];
    }
};