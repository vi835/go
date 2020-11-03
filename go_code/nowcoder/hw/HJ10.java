package nowcoder.hw;

import java.util.HashSet;
import java.util.Scanner;
import java.util.Set;

public class HJ10 {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        String str = sc.nextLine();
        char[] chars = str.toCharArray();

        Set set=new HashSet<>();
        for(char c:chars){
            set.add(c);
        }
        System.out.println(set.size());
    }
}
