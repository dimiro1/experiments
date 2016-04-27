import java.util.PriorityQueue;

public class Main {
    public static void main(String[] args) {
        PriorityQueue<Integer> queue = new PriorityQueue<>((x, y) -> y - x);
        
        queue.add(1);
        queue.add(10);
        queue.add(3);
        queue.add(5);
        queue.add(99);
        queue.add(5);
        queue.add(6);
        
        System.out.println(queue);
    }
}