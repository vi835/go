package nowcoder.hw;

import java.util.Scanner;

public class HJ40 {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        while (sc.hasNext()) {
            String str = sc.nextLine();
            char[] chars = str.toCharArray();
            int letters = 0, whiteSpaces = 0, digits = 0, others = 0;
            for (char c : chars) {
                if (Character.isDigit(c)) {
                    digits++;
                } else if (Character.isLetter(c)) {
                    letters++;
                } else if (Character.isWhitespace(c)) {
                    ++whiteSpaces;
                } else {
                    ++others;
                }
            }
            System.out.println(letters);
            System.out.println(whiteSpaces);
            System.out.println(digits);
            System.out.println(others);
        }
    }
}
