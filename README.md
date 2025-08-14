# Cara Menjalankan Project

1. **Buat Database di PostgreSQL**  
   Buat database baru untuk project ini.  
   Contoh:  
   ```sql
   CREATE DATABASE ayo_indonesia;
   ```

2. **Sesuaikan File `.env` dengan Credential Database**  
   Ubah isi file `.env` sesuai konfigurasi PostgreSQL di komputer kamu:  
   ```
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=postgres
   DB_PASSWORD=your_password
   DB_NAME=ayo_indonesia
   ```

3. **Jalankan Project Go**  
   Pastikan sudah berada di root project, lalu jalankan:  
   ```bash
   go run main.go
   ```

4. **Masukkan Data Contoh dari `data.sql`**  
   Import file `data.sql` ke database:  
   ```bash
   psql -U postgres -d ayo_indonesia -f data.sql
   ```

5. **Uji API di Postman**  
   Gunakan Postman untuk mencoba endpoint API yang sudah tersedia.
