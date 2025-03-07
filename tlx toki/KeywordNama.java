import java.util.Scanner;

public class KeywordNama {
    
    public static void main(String[] args) {
        
        Scanner input = new Scanner(System.in);

        char keyword;
        keyword = input.next().charAt(0);
        int count = 0;
        String result = ""; 

        String[] nama = {"agus", "Ammar", "Anthony" , "Ayaz"};

        for (String arrayNama : nama) {
            if (arrayNama.toLowerCase().contains(String.valueOf(keyword))) {
                count++;
                result = arrayNama;
            }
            
        }

        if (count == 0) {
            System.out.println("Tidak ada");
        } else if (count == 1) {
            System.out.println(result);
        } else if (count > 1) {
            System.out.println("Bingung");
        }

    


    }
}
