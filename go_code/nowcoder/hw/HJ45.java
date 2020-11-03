package nowcoder.hw;

import java.util.*;

public class HJ45 {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        while (sc.hasNext()) {
            int rows = Integer.parseInt(sc.nextLine());
            for (int i = 0; i < rows; i++) {
                String str = sc.nextLine().toLowerCase();
                char[] chars = str.toCharArray();
                int[] arr = new int[26];
                for (char c : chars) {
                    arr[c - 'a'] += 1;
                }
                Arrays.sort(arr);
                int sum = 0;
                int n = 26;
                for (int m = arr.length - 1; m >= 0; m--) {
                    sum += arr[m] * n;
                    n--;
                }
                System.out.println(sum);
            }
        }
    }
}
