package nowcoder.hw;

import java.util.Arrays;
import java.util.Scanner;

public class HJ101 {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        while (sc.hasNext()) {
            int count = sc.nextInt();
            int[] arr = new int[count];
            for (int i = 0; i < count; i++) {
                arr[i] = sc.nextInt();
            }
            int flag = sc.nextInt();
            Arrays.sort(arr);
            if (flag == 0) {
                for (int i = 0; i < arr.length; i++) {
                    System.out.print(arr[i] + " ");
                }
            } else {
                for (int i = arr.length - 1; i >= 0; i--) {
                    System.out.print(arr[i] + " ");
                }
            }
            System.out.println();
        }
    }
}
