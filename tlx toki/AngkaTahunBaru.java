    import java.util.Scanner;

    public class AngkaTahunBaru {

        public static void main(String[] args) {
            Scanner input = new Scanner(System.in);
            int angka, n;
            n = input.nextInt();
            
            
            if (isPrime(n) && n < 7) {
                System.out.println("YES");
            } else {
                System.out.println("NO");
            }
        }

        private static boolean isPrime (int angka) {
            if (angka < 2 || angka > 100) {
                return false;
            }

            for (int i = 2; i <= Math.sqrt(angka); i++) {
                if (angka % i == 0) {
                    return false;
                }
            }

            return true;
        }
    }