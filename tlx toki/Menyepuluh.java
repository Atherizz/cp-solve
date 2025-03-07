import java.util.Scanner;

public class Menyepuluh {
    public static void main(String[] args) {
        
        Scanner sc = new Scanner(System.in);
        int N = sc.nextInt();

        String angka = Integer.toString(N);
        
        if (angka.contains("0")) {
            System.out.println("YES");
        } else {
            System.out.println("NO");
        }
    }
}
