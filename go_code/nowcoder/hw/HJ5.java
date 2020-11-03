package nowcoder.hw;

import java.util.Scanner;

public class HJ5 {
    public static void main(String[] args) {
        Scanner sc=new Scanner(System.in);
        while(sc.hasNext()){
            String str=sc.nextLine();
            System.out.println(Integer.valueOf(str.substring(2),16).toString());
        }
    }
}
