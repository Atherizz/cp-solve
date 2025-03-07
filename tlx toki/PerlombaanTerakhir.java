import java.util.Scanner;

public class PerlombaanTerakhir {
    
public static void main(String[] args) {
    String namaLomba;
    char karakter = 'O';
    int countO = 0;

    Scanner S = new Scanner(System.in); 
    namaLomba = S.next();

    for(int i = 0; i < namaLomba.length();i++) {
        if(namaLomba.charAt(i) == 'O') {
            countO++;
        }
        }
    
    if (countO == 1) { 
        System.out.println("Ya");
    } else {
        System.out.println("Tidak");
    }

}
}
