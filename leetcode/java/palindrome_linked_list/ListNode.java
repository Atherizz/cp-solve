package palindrome_linked_list;

import java.util.ArrayList;

public class ListNode {
      int val;
      ListNode next;

      ListNode() {}

      ListNode(int val) { 
        this.val = val;
      }

      ListNode(int val, ListNode next) { 
        this.val = val; 
        this.next = next;
       }
}

class MyStack {
    ArrayList<Integer> list = new ArrayList<>(); 
    int top;

    MyStack() {
        top = -1;
    }

    boolean isEmpty() {
        return list.isEmpty();
    }

    void push(int data) {
        top++;
        list.add(top,data);
    }

    int pop() {
        if (!isEmpty()) {
            int data = list.get(top);
            top--;
            return data;     
        } else {
            return 0;
        }
    }

    int peek() {
        return list.get(top);
    }
}

class Solution {
    public boolean isPalindrome(ListNode head) {
        MyStack stack = new MyStack();

        ListNode current = head;
        while (current != null) {
            stack.push(current.val);
            current = current.next;
        }

        current = head;
        while (current != null) {
            int num = stack.pop();
            if (current.val != num) {
                return false;
            }
            current = current.next;
        }

        return true;
    }
}


