package implement_stack_queue;

import java.util.LinkedList;
import java.util.Queue;

class MyStack {
    Queue<Integer> q = new LinkedList<>();

    MyStack() {
    }

    boolean empty() {
        return q.isEmpty();
    }

        void push(int data) {
            q.add(data);

            int size = q.size();

            for (int i = 0; i < size - 1; i++) {
                int num = q.poll();
                q.add(num);
            }
            
        }

    int pop() {
        return q.poll();
    }

    int peek() {
        return q.peek();
    }
}
