import random
import datetime
from customer import Customer

atm = Customer(id)
while True:

    id = int(input("Masukkan pin anda:  "))
    trial = 0
    
    while (id != int(atm.checkPin()) and trial < 3):
        id = int(input("Pin anda salah. Silahkan masukkan lagi:  "))
        
        trial += 1
        
        if trial == 3:
            print("Error. Silahkan ambil kartu dan coba lagi..")
            temp = int(input(""))
            exit()
            
    while True:
        print("Selamat datang di ATM Progate..")
        print("\n1 - Cek Saldo \t 2 - Debet \t 3 - Simpan \t 4 - Ganti Pin \t 5 - Keluar ")
        selectMenu = int(input("\n Silahkan pilih menu:  "))
        if selectMenu == 1:
            print("\nSaldo anda sekarang: Rp. " + str(atm.checkBalance()) + "\n" )
        elif selectMenu == 2:
            nominal = float(input("Masukkan nominal saldo: "))
            verify_withdraw = input("Konfirmasi: Anda akan melakukan debet dengan nominal berikut ? y/n " + str(nominal) + " ")
            if verify_withdraw == "y":
                print("Saldo awal anda adalah: Rp. " + str(atm.checkBalance()) + "")
            else:
                break
            if nominal < atm.checkBalance():
                atm.withdrawBalance(nominal)
                print("Transaksi debet berhasil!")
                print("Saldo sisa sekarang: Rp. " + str(atm.checkBalance()) + "")
            else:
                print("Maaf. Saldo anda tidak cukup untuk melakukan debet!")
                print("Silakan lakukan penambahan nominal saldo")
        elif selectMenu == 3:
            nominal = float(input("Masukkan nominal saldo: "))
            verify_deposit = input("Konfirmasi: Anda akan melakukan penyimpanan dengan nominal berikut ? y/n " + str(nominal) + " ")
            if verify_deposit == "y":
                atm.depositBalance(nominal)
                print("Saldo anda sekarang adalah: Rp." + str(atm.checkBalance()) + "\n" )
            else:
                break
        elif selectMenu == 4:
            verify_pin = int(input("masukkan pin anda: "))
            while verify_pin != int(atm.checkPin()):
                print("pin anda salah, silakan masukkan pin:")
                verify_pin = int(input("masukkan pin anda: "))
            updated_pin = int(input("silakan masukkan pin baru: "))
            print("pin anda berhasil diganti! ")
            atm.changePin(updated_pin)
            verify_newpin = int(input("coba masukkan pin baru: "))
            if verify_newpin == updated_pin:
                print("pin baru anda benar!")
            else:
                print("maaf, pin baru anda salah!")
        elif selectMenu == 5:
            print("Resi tercetak otomatis saat anda keluar. \nHarap simpan tanda terima ini \nsebagai bukti transaksi anda. ")
            print("No. Record: ", random.randint(100000, 1000000))
            print("Tanggal: ", datetime.datetime.now())
            print("Saldo akhir: ", atm.checkBalance())
            print("Terima kasih telah menggunakan ATM Progate! ")
            temp = int(input(""))
            exit()
        else:
            print("Error. Maaf, menu tidak tersedia")
