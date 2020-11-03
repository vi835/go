package nowcoder.hw;

import java.util.Scanner;

public class HJ105 {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        float posiCount = 0;
        int negaCount = 0;
        float sum = 0;
        while (sc.hasNext()) {
            int num = sc.nextInt();
            if (num < 0) {
                ++negaCount;
            } else {
                sum += num;
                ++posiCount;
            }
        }

        System.out.println(negaCount);
        if (posiCount == 0) {
            System.out.println(0.0);
        } else {
            System.out.printf(".1f\n", sum / posiCount);
        }
    }

//    public static void main(String[] args) {
//        Scanner sc = new Scanner(System.in);
//        float count1 = 0;
//        int count2 = 0;
//        float sum = 0;
//        while (sc.hasNextInt()) {
//            int cur = sc.nextInt();
//            if (cur >= 0) {
//                count1++;
//                sum += cur;
//            } else {
//                count2++;
//            }
//        }
//        System.out.println(count2);
//        if (count1 == 0) {
//            System.out.println(0.0);
//        } else {
//            System.out.printf("%.1f\n", sum / count1);
//        }
//    }
}
