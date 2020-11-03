package nowcoder.hw;

import java.util.*;

public class HJ14 {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);

        int rows = sc.nextInt();
        String[] array = new String[rows];
        for (int i = 0; i <rows; i++) {
            array[i] = sc.next();
        }
        Arrays.sort(array);
        for (String s : array) {
            System.out.println(s);
        }

    }
}
