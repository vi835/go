package nowcoder.hw;

import java.util.Scanner;
import java.util.TreeSet;

public class HJ3 {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        while(sc.hasNext()) {
            int count = sc.nextInt();

            TreeSet<Integer> set = new TreeSet<>();
            for (int i = 0; i < count; i++) {
                set.add(sc.nextInt());
            }

            for (Integer n : set) {
                System.out.println(n);
            }
        }
    }
}