/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func helper(root *TreeNode, ans *[]int) {
    if root == nil {
        return;
    }

    if root.Left != nil {
        helper(root.Left, ans);
    }
    *ans = append(*ans, root.Val);
    if root.Right != nil {
        helper(root.Right, ans)
    }
}

func inorderTraversal(root *TreeNode) []int {        
    ans := []int{};
    helper(root, &ans);
    return ans;
}