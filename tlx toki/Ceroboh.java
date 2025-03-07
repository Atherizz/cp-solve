import java.util.Scanner;

public class Ceroboh {

    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);
    
        int N,M,X,Y, nilai, nilaiMaksimal;

        N = input.nextInt();
        M = input.nextInt();
        X = input.nextInt();
        Y = input.nextInt();

        nilai = (N - X) * 1 + (M - Y) * 2;
        nilaiMaksimal = (20 * 1) + (15 * 2);
    

        if (nilaiMaksimal / 2 < nilai) {
            System.out.println("LOLOS");
        } else {
            System.out.println("TIDAK LOLOS");
        }
        }



    }
    

