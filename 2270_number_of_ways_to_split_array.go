func waysToSplitArray(nums []int) int {
    leftSum := 0;
    rightSum := 0;
    for _, num := range(nums) {
        rightSum+=num;
    }
    ans := 0;
    for i := 0; i < len(nums) - 1; i++ {
        leftSum+=nums[i];
        rightSum-=nums[i];
        if leftSum >= rightSum {
            ans++;
        }
    }
    return ans;
}