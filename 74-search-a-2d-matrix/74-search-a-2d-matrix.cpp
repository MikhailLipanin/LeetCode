class Solution {
public:
    bool searchMatrix(vector<vector<int>>& matrix, int target) {
        int n = matrix.size(), m = matrix[0].size();
        int lRow = 0, rRow = n;
        while (rRow - lRow > 1) {
            int mRow = (lRow + rRow) >> 1;
            int lCol = 0, rCol = m;
            while (rCol - lCol > 1) {
                int mCol = (lCol + rCol) >> 1;
                if (matrix[mRow][mCol] <= target) lCol = mCol;
                else rCol = mCol;
            }
            if (matrix[mRow][lCol] == target) {
                return true;
            }
            else if (matrix[mRow][lCol] < target) lRow = mRow;
            else rRow = mRow;
        }
        int lCol = 0, rCol = m;
        while (rCol - lCol > 1) {
            int mCol = (lCol + rCol) >> 1;
            if (matrix[lRow][mCol] <= target) lCol = mCol;
            else rCol = mCol;
        }
        if (matrix[lRow][lCol] == target)
            return true;
        else
            return false;
    }
};