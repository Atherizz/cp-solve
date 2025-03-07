import java.util.Scanner;

public class PermainanUbin {

    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);
        int N = input.nextInt();
        int M = input.nextInt();
        
        // Jumlah segitiga minimum yang dibutuhkan adalah N * M
        int jumlahSegitiga = N * M;
        
        System.out.println(jumlahSegitiga);
    }
}