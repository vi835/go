package nowcoder.hw;

import java.util.Scanner;

public class HJ2 {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        String str = sc.nextLine().toUpperCase();
        String tag = sc.nextLine().toUpperCase();
        String temp = str.replace(tag, "");
        System.out.println(str.length() - temp.length());
    }
}
