import java.util.Scanner;

public class BungaGabungan {
    
    public static void main(String[] args) {
        
        Scanner input = new Scanner(System.in);
        int p,q, totalBunga;

        p = input.nextInt();
        q = input.nextInt();

        totalBunga = (p*p) + (q*q) + 1;

        if (totalBunga % 4 == 0) {
            System.out.println(totalBunga/4);
        } else {
            System.out.println(-1);
        }

    }
}
