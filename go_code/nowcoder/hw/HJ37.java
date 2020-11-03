package nowcoder.hw;

import java.util.Scanner;

public class HJ37 {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        while(sc.hasNext()){
            int n=sc.nextInt();
            System.out.println(rabbits(n));
        }
    }

    private static int rabbits(int months) {
        if (months == 1 || months == 2) {
            return 1;
        }
        return rabbits(months-1)+rabbits(months-2);
    }
}

