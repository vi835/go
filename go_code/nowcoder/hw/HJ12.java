package nowcoder.hw;

import java.util.Scanner;

public class HJ12 {
    public static void main(String[] args) {
        Scanner sc=new Scanner(System.in);
        String str=sc.nextLine();
        System.out.println(new StringBuilder(str).reverse().toString());
    }
}
