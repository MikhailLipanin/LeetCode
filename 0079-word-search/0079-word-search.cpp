class Solution {
private:
    int dx[4] = {-1, 0, 0, 1};
    int dy[4] = {0, -1, 1, 0};
public:
    bool exist(vector<vector<char>>& board, string word) {
        queue <pair <pair <int, int>, pair <int, long>>> q;
        int n = board.size(), m = board[0].size();
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < m; j++) {
                if (board[i][j] != word[0]) continue;
                long id = (1LL << (m * i + j));
                q.push({{i, j}, {1, id}});
            }
        }
        while (!q.empty()) {
            auto now = q.front();
            q.pop();
            int x = now.first.first, y = now.first.second;
            int cnt = now.second.first;
            long id = now.second.second;
            if (cnt == word.size()) {
                return true;
            }
            for (int i = 0; i < 4; i++) {
                int nx = x + dx[i];
                int ny = y + dy[i];
                if (!(0 <= nx && nx < n)) continue;
                if (!(0 <= ny && ny < m)) continue;
                if (board[nx][ny] != word[cnt]) continue;
                long cur_id = (1LL << (m * nx + ny));
                if (id & cur_id) continue;
                cur_id = id | cur_id;
                q.push({{nx, ny}, {cnt + 1, cur_id}});
            }
        }
        return false;
    }
};