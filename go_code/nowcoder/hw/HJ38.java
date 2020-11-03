package nowcoder.hw;

import java.util.Scanner;

public class HJ38 {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        while (sc.hasNext()) {
            int height = sc.nextInt();
            double sum = 0;
            for (int i = 1; i <= 5; i++) {
                sum += height;
                height /= 2;
            }
            System.out.printf("%0.6f\n",sum);
            System.out.printf("%0.6f\n",height);
        }
    }
}
