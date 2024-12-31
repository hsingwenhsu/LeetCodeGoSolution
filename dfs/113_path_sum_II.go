/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 func helper(root *TreeNode, prev int, current []int, ans *[][]int, targetSum int) {
    if root == nil {
        return;
    }
    current = append(current, root.Val);
    // if leaf node
    if root.Left == nil && root.Right == nil {
        if prev + root.Val == targetSum {
            /**
			 * ChatGPT: 
             * The final paths stored in *ans will not be the state of current at the time they were added.
             * Instead, they will reflect the last state of current after all recursive calls.
             * So we cannot directly append current to ans.
             **/
            pathCopy := make([]int, len(current))
            copy(pathCopy, current)
            *ans = append(*ans, pathCopy);
        }
        return;
    }

    helper(root.Left, prev + root.Val, current, ans, targetSum);
    helper(root.Right, prev + root.Val, current, ans, targetSum);
}

func pathSum(root *TreeNode, targetSum int) [][]int {
    ans := [][]int{};
    current := []int{};
    helper(root, 0, current, &ans, targetSum);
    return ans;
}