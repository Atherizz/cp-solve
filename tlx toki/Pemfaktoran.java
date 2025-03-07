import java.util.Scanner;

public class Pemfaktoran {

    public static void main(String[] args) {
        
        Scanner sc = new Scanner(System.in);
        int N = sc.nextInt();
        int countFaktor = 0;

        for (int i = 1; i <= N; i++) {
            if (N % i == 0) {
                countFaktor++;
            }
        }

        if (countFaktor == 5) {
            System.out.println("YES");
        } else {
            System.out.println("NO");
        }
    }
}