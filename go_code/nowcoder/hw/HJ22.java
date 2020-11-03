package nowcoder.hw;

import java.util.Scanner;

public class HJ22 {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        while (sc.hasNext()) {
            int empty = sc.nextInt();
            if (empty == 0) {
                return;
            }
            System.out.println(empty / 2);
        }
    }
}
