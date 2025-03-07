import java.util.Scanner;

public class NyalaRambut {
    
    public static void main(String[] args) {
        int n;
        double d,k;
        Scanner input = new Scanner(System.in);

        n = input.nextInt();
        d = input.nextInt();
        k = input.nextInt();

        double suhu = Math.pow(d, k);


        if (n > suhu) {
            System.out.println("YES");
        } else {
            System.out.println("NO");
        }


    }
}
