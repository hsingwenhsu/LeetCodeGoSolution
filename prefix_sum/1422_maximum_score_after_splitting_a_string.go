func maxScore(s string) int {
    zeroCnt := 0;
    oneCnt := strings.Count(s, "1");
    ans := 0;
    for i := 0; i < len(s) - 1; i++ {
        if (s[i] == '0') {
            zeroCnt++;
        } else {
            oneCnt--;
        }
        ans = max(ans, zeroCnt + oneCnt);
    }
    return ans;
}