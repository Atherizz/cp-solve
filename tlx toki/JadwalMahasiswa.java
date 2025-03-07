import java.util.Scanner;

public class JadwalMahasiswa {
   
    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);
        int jamNow, menitNow, menitKuliah, jamSelesai, menitSelesai, menitFix;
        String waktuSaatIni;
        

        
        System.out.println("---------------------------------------");
        System.out.println("------ JADWAL MAHASISWA POLINEMA ------");
        System.out.println("---------------------------------------");
        

        System.out.print("Masukkan jam saat ini : ");
        jamNow = input.nextInt();
        
        System.out.print("Masukkan menit saat ini : ");
        menitNow = input.nextInt();


        if (menitNow == 0) {
        waktuSaatIni = jamNow + "." + menitNow + "0";
        } else {
            waktuSaatIni = jamNow + "." + menitNow;
        }

        System.out.println("jam saat ini : " + waktuSaatIni);

        System.out.print("Masukkan waktu kuliah dalam menit : ");
        menitKuliah = input.nextInt();

        menitSelesai = menitNow + menitKuliah;
        menitFix = menitSelesai % 60;
        jamSelesai = jamNow + menitKuliah / 60;

        if (menitFix == 0) {
            System.out.println("kuliah berakhir pada : " + jamSelesai + "." + menitFix + "0");
        } else {
            System.out.println("kuliah berakhir pada : " + jamSelesai + "." + menitFix);
        }

   
        
    }
}
