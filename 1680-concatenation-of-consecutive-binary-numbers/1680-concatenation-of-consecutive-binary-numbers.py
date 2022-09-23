class Solution:
    def concatenatedBinary(self, n: int) -> int:
        num = ""
        for i in range(1, n + 1):
            num += bin(i)[2:]
        ret = int(num, 2)
        ret %= 1000000007
        return ret
        