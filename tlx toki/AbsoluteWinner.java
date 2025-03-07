import java.util.Scanner;

public class AbsoluteWinner {
    
    public static void main(String[] args) {
        
        Scanner input = new Scanner(System.in);
        int A,B,C, round;

        A = input.nextInt();
        B = input.nextInt();
        C = input.nextInt();
        round = (A + B + C) / 7;

        if (A == 4 * round || B == 4 * round || C == 4 * round) {
            System.out.println("YA");
        } else {
            System.out.println("TIDAK");
        }

    }
}
