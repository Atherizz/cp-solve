import java.util.Scanner;

public class TahunLompat {
    
    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);

        int N, A, B, C;

        N = input.nextInt();
        A = input.nextInt();
        B = input.nextInt();
        C = input.nextInt();

        if (N % A == 0 && N % B != 0 || N % C == 0) {
            System.out.println("YES");
        } else {
           System.out.println("NO"); 
        }

    }
}
