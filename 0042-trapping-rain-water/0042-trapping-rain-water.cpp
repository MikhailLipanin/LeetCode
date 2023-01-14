class Solution {
public:
    int trap(vector<int>& height) {
        int n = height.size();
        vector <int> lef(n, -1);
        vector <int> rig(n, -1);
        for (int i = 0; i < n; i++) {
            lef[i] = height[i];
            if (i) lef[i] = max(lef[i], lef[i - 1]);
        }
        for (int i = n - 1; i >= 0; i--) {
            rig[i] = height[i];
            if (i + 1 < n) rig[i] = max(rig[i], rig[i + 1]);
        }
        int ans = 0;
        for (int i = 0; i < n; i++) {
            ans += min(lef[i], rig[i]) - height[i];
        }
        return ans;
    }
};