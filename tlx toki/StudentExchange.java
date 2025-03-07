import java.util.Scanner;

public class StudentExchange {

    public static void main(String[] args) {
    
        Scanner input = new Scanner(System.in);
        int tanggalBerangkat, tahunBerangkat, tanggalPulang, tahunPulang, 
        studentExchangeHari, studentExchangeBulan, studentExchangeTahun, bulanBerangkat, bulanPulang;
        String keberangkatan, kepulangan;


        System.out.print("Masukkan tanggal berangkat : ");
        tanggalBerangkat = input.nextInt();

        System.out.print("Masukkan bulan berangkat : ");
        bulanBerangkat = input.nextInt();

        System.out.print("Masukkan tahun berangkat : ");
        tahunBerangkat = input.nextInt();

        keberangkatan = tanggalBerangkat + " - " + bulanBerangkat + " - " + tahunBerangkat;
        System.out.println(keberangkatan);

        System.out.print("Masukkan tanggal pulang : ");
        tanggalPulang = input.nextInt(); 

        System.out.print("Masukkan bulan pulang : ");
        bulanPulang = input.nextInt();
        
        System.out.print("Masukkan tahun pulang : ");
        tahunPulang = input.nextInt();

        kepulangan = tanggalPulang + " - " + bulanPulang + " - " + tahunPulang;
        System.out.println(kepulangan);

        studentExchangeTahun = tahunPulang - tahunBerangkat;
        studentExchangeBulan = bulanPulang - bulanBerangkat;
        studentExchangeHari = (tanggalPulang - tanggalBerangkat);

        studentExchangeBulan = studentExchangeBulan + ((studentExchangeTahun * 12) % 12) * -1;
        studentExchangeHari = (studentExchangeHari + (studentExchangeBulan * 30) % 30) * -1;



        System.out.println("Lama student exchange adalah " + studentExchangeTahun + " tahun , " + studentExchangeBulan + " bulan, " + studentExchangeHari + " hari");











    }
}