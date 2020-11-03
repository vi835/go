package nowcoder.hw;

import java.util.LinkedHashSet;
import java.util.Scanner;
import java.util.Set;

public class HJ9 {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        int num = sc.nextInt();
        if (num == 0) {
            System.out.println(0);
            return;
        }
        Set<Integer> set = new LinkedHashSet<>();
        while (num != 0) {
            int temp = num % 10;
            num = num / 10;
            set.add(temp);
        }
        for(Integer n:set){
            System.out.print(n);
        }
    }
}
