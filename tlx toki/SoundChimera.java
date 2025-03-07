import java.util.Scanner;

public class SoundChimera {
    
    public static void main(String[] args) {
        
        Scanner input = new Scanner(System.in);
        int A,B;
        int perubahan = 0;

        A = input.nextInt();
        B = input.nextInt();

        if (A == B) {
            perubahan = 1;
             } 

        System.out.println(perubahan);

    }
}
