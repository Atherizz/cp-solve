public class MainSolution {
  public static void main(String[] args) {
ListNode head = new ListNode(
    1, new ListNode(
    1, new ListNode(
    2, new ListNode(
    3, new ListNode(
    3, new ListNode(4))))));


    Solution solution = new Solution();
    solution.deleteDuplicates(head);

    ListNode current = head;

    while (current != null) {
        System.out.print(current.val + ", ");
        current = current.next;
    }
  }
}
