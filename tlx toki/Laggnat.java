import java.util.Scanner;

public class Laggnat {
    
    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);
         int A,B;
         int C;
      
         A = input.nextInt();
         B = input.nextInt();

        //  C = A;
        //  A = B;
        //  B = C;

         System.out.println(A);
         System.out.println(B);

         if (B >= 1 || B<=31 && A == 1) {
            System.out.println("YES");
         } else if (B >= 1 || B<=28 && A == 2) {
            System.out.println("YES");
         }else if (B >= 1 || B<=31 && A == 3) {
            System.out.println("YES");
         } else  if (B >= 1 || B<=30 && A == 4) {
            System.out.println("YES");
         } else  if (B >= 1 || B<=31 && A == 5) {
            System.out.println("YES");
         } else  if (B >= 1 || B<=30 && A == 6) {
            System.out.println("YES");
         } else  if (B >= 1 || B<=31 && A == 7) {
            System.out.println("YES");
         } else  if (B >= 1 || B<=31 && A == 8) {
            System.out.println("YES");
         } else  if (B >= 1 || B<=30 && A == 9) {
            System.out.println("YES");
         } else  if (B >= 1 || B<=31 && A == 10) {
            System.out.println("YES");
         } else  if (B >= 1 || B<=30 && A == 11) {
            System.out.println("YES");
         } else if (B >= 1 || B<=31 && A == 12) {
            System.out.println("YES");
         } else {
            System.out.println("NO");
         }



    }
}
