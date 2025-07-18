import java.util.List;
import java.util.ArrayList;

public class TreeNode {
      int val;
      TreeNode left;
      TreeNode right;
      TreeNode() {}
      TreeNode(int val) { this.val = val; }
      TreeNode(int val, TreeNode left, TreeNode right) {
          this.val = val;
          this.left = left;
          this.right = right;
      }
  }
 
class Solution {
    public List<Integer> inorderTraversal(TreeNode root) {
        List<Integer> tree = new ArrayList<>();

        if (root == null) {
            return tree;
        }

        List<Integer> leftResult = inorderTraversal(root.left);
        tree.addAll(leftResult);

        tree.add(root.val);

        List<Integer>rightResult = inorderTraversal(root.right);
        tree.addAll(rightResult);
        
        return tree;
    }
}