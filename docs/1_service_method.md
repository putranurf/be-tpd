## Service method naming guide

Gunakan hanya [Create, Get, Put, Delete, Update, List] prefix.
e.g. :

* `Create` untuk membuat resource baru (baik menulis ke storage / tidak)
    * contoh: `CreateUser(ctx context.Context, ID string) (User, error)`
* `Get` untuk mendapatkan _*single*_ resource yang sudah ada sebelumnya di
    datasource
    * contoh: `GetUser(ctx context.Context, ID string) (User, error)`
* `Update` untuk mengupdate *single* resource yang sudah ada sebelumnya secara
    keseluruhan (replace existing resource)
    * contoh: `UpdateUser(ctx context.Context, u *User)eerror`
* `Put` untuk mengupdate *single* resource secara parsial.
* `List` untuk mendapatkan *many* resources yang sudah ada sebelumnya di
    datasrouce
    * contoh `ListUser(ctx context.Context, options FilterOptions) ([]User, error)`
* `Delete` untuk menghapus *single* resource di datasource
    * `DeleteUser(ctx context.Context, ID string) error`
