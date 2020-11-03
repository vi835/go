package nowcoder;

import java.util.Scanner;

public class Main {
//    public static void main(String[] args) {
//        Scanner sc= new Scanner(System.in);
//        while(sc.hasNext()){
//            double num=sc.nextDouble();
//
//            int result= (int)(num+0.5);
//            System.out.println(result);
//        }
//    }

//    public static void main(String[] args) {
//        Scanner sc = new Scanner(System.in);
//        int n = sc.nextInt();
//        if (n == 0) {
//            System.out.println(0);
//            return;
//        }
//        int flag;
//        while (n != 0) {
//            flag = n % 10;
//            n = n / 10;
//            System.out.print(flag);
//        }
//    }

    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        while (sc.hasNext()) {
            int n = sc.nextInt();
            System.out.println((3*n+1)*n/2);
        }
    }
}
