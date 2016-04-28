import java.util.PriorityQueue;

public class Main {
    public static void main(String[] args) {
        PriorityQueue<Integer> queue = new PriorityQueue<>((x, y) -> y - x);
        
        queue.offer(1);
        queue.offer(10);
        queue.offer(3);
        queue.offer(5);
        queue.offer(99);
        queue.offer(5);
        queue.offer(6);
        
        System.out.println(queue);
        
        while (!queue.isEmpty()) {
            System.out.println(queue.poll());
        }
    }
}