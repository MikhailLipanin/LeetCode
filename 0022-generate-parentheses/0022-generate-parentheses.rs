impl Solution {
    pub fn generate_parenthesis(n: i32) -> Vec<String> {
        let mut ans = Vec::new();
        let mut s = String::from("");
        Solution::rec(n, n, 0, s, &mut ans);
        ans
    }
    pub fn rec(nx: i32, ny: i32, bal: i32, mut tmp: String, ans: &mut Vec<String>) {
        if bal < 0 {
            return;
        }
        if nx == 0 && ny == 0 {
            if bal == 0 {
                ans.push(tmp);
            }
            return;
        }
        if nx > 0 {
            let next = format!("{}{}", tmp, "(");
            Solution::rec(nx - 1, ny, bal + 1, next, ans);
        }
        if ny > 0 {
            let next = format!("{}{}", tmp, ")");
            Solution::rec(nx, ny - 1, bal - 1, next, ans);
        }
    }
}