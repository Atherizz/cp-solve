import java.util.Scanner;

public class PertahananRobot {
    
    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);
        int x = 0;
        int y = 0;
        String S;

        S = input.next();

        for(int i = 0; i < S.length(); i++) {
            if (S.charAt(i) == 'R') {
                x+= 1;
            } else if (S.charAt(i) == 'L') {
                x-= 1;
            } else if (S.charAt(i) == 'U') {
                y += 1;
            } else if (S.charAt(i) == 'D') {
                y -= 1;
            }
        }

        System.out.println(x + " " + y);




    }
}
