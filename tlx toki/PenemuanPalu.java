import java.util.Scanner;

public class PenemuanPalu {
   
    public static void main(String[] args) {
        
        Scanner input = new Scanner(System.in);
        int x,y;

        y = input.nextInt();
        x = input.nextInt();

        if (y > x) {
            System.out.println("YES");
        } else {
            System.out.println("NO");
        }




    }
}
