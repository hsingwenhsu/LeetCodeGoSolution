func isPrefixAndSuffix(str1 string, str2 string) bool {
    l1 := len(str1);
    l2 := len(str2);
    if l1 > l2 {
        return false;
    }

    if l1 == l2 {
        return str1 == str2;
    } 
    for i := 0; i < l1; i++ {
        if str1[i] != str2[i] {
            return false;
        }
        if str1[l1 - 1 - i] != str2[l2 - 1 - i] {
            return false;
        }
    }
    return true;
}
func countPrefixSuffixPairs(words []string) int {
    ans := 0;
    for i := 0; i < len(words); i++ {
        for j := i + 1; j < len(words); j++ {
            if isPrefixAndSuffix(words[i], words[j]) {
                ans++;
            }
        }
    }
    return ans;
}