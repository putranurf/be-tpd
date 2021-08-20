# Data Access / Repository

## Storage SQL (Postgres)

Karena menggunakan Postgres sebagai RDBMS, maka akses
data-nya menggunakan [GORM](https://gorm.io/docs/index.html).
Cukup simpan instance `*sql.DB` di
dalam object `service`, dan create query object di setiap fungsi yg
membutuhkan.

```golang
// internal/service/service.go
type Service struct {
    db *gorm.DB      // digunakan untuk akses data ke postgres
    ...
}

// internal/service/user_service.go
func (s *Service) GetUser(ctx context.Context) (*User, error) {
    // create query obj
    if err := db.GetDBInstance().Find(&arrobj).Error; err != nil {
		return res, err
    }

    ...
}
```