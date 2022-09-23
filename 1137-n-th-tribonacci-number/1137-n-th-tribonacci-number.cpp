class Solution {
public:
    int tribonacci(int n) {
        if (n == 0) return 0;
        if (n <= 2) return 1;
        int x = 0, y = 1, z = 1;
        n -= 2;
        while (n--) {
            int temp = z + y + x;
            x = y;
            y = z;
            z = temp;
        }
        return z;
    }
};