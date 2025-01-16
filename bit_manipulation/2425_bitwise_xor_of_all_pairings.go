func xorAllNums(nums1 []int, nums2 []int) int {
    n := len(nums1) % 2;
    m := len(nums2) % 2;
    x := 0;
    if (m == 1) {
        for _, num := range(nums1) {
            x ^= num;
        }
    }
    if (n == 1) {
        for _, num := range(nums2) {
            x ^= num;
        }
    }
    return x;
}