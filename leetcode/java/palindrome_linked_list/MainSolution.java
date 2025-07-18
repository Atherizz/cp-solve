package palindrome_linked_list;

public class MainSolution {
    public static void main(String[] args) {
    ListNode head = new ListNode(
    1, new ListNode(
    2, new ListNode(
    2, new ListNode(
    1, new ListNode(
    )))));

    Solution solution = new Solution();
    boolean result = solution.isPalindrome(head);
    System.out.println(result);
    
    }
}
