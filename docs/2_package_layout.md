# 2 Go Package Layout

Secara garis besar terbagi menjadi 3 package folder utama
* `cmd` berisi main program

## Package Descriptions

1. `/cmd`

* `/cmd/main.go` -> untuk menjalankan echo HTTP di beda port yang sudah di tentukan.
    pilih port sesuai dengan kebutuhan. 

2. `/internal`

Package `internal` untuk logic program yg dibutuhkan internal oleh service, dan
tidak ditujukan bisa di import oleh go program lainnya secara `gomodule`.

* `/internal/service` core logic dari api JSON webservice
* `/internal/http` implementasi `echo.Echo` http server, router & handler (jika
    diperlukan).
* `/internal/repository` logic data access layer. Implementasi refer ke
* [repository design](./3_data_access_repository.md)