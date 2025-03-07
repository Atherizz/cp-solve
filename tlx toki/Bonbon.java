import java.util.Scanner;

public class Bonbon {
    
    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);

        double X,Y,Z, hargaDiskon, hargaAkhir;

        X = input.nextDouble();
        Y = input.nextDouble();
        Z = input.nextDouble();

        hargaDiskon = X - (X * Y/100);
        hargaAkhir = hargaDiskon + (hargaDiskon * Z/100);
        int intHarga = (int) hargaAkhir;

        System.out.println(intHarga);



    }
}
