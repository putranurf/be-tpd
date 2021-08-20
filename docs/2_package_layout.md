# 2 Go Package Layout

Secara garis besar terbagi menjadi 3 package folder utama
* `cmd` berisi main program
* `internal` app business logic
* `pkg` untuk package2 yang boleh di-import oleh go project / repo lainnya

## Package Descriptions

1. `/cmd`

* `/cmd/api/main.go` -> untuk menjalankan echo HTTP di beda port yang sudah di tentukan.
    pilih port sesuai dengan kebutuhan. 

2. `/internal`

Package `internal` untuk logic program yg dibutuhkan internal oleh service, dan
tidak ditujukan bisa di import oleh go program lainnya secara `gomodule`.

* `/internal/service` core logic dari api JSON webservice
* `/internal/http` implementasi `echo.Echo` http server, router & handler (jika
    diperlukan).
* `/internal/repo` logic data access layer. Implementasi refer ke
* [repository design](./3_data_access_repository.md)

3. `/pkg`

* `/pkg/helpers` berisi definisi library global

* `/pkg/middleware` berisi middleware untuk menjembatani request client terhadap service yang
    akan diakses


## Reference

* [golang standards](https://github.com/golang-standards/project-layout)