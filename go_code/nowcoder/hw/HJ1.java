package nowcoder.hw;

import java.util.Scanner;

public class HJ1 {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        String str = sc.nextLine();

        String[] arr=str.split(" ");
        System.out.println(arr[arr.length-1].length());
    }
}
