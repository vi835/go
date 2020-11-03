package nowcoder.hw;

import java.util.Scanner;

public class HJ13 {
    public static void main(String[] args) {
        Scanner sc=new Scanner(System.in);
        String str=sc.nextLine();
        String[] s = str.split(" ");
        for(int i=s.length-1;i>=0;i--){
            System.out.print(s[i]+" ");
        }
    }
}
