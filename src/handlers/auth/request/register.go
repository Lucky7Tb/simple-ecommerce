package request

type RegisterRequest struct {
	Nama         string `json:"nama" validate:"required"`
	KataSandi    string `json:"kata_sandi" validate:"required"`
	NoTelepon    string `json:"no_telp" validate:"required"`
	TanggalLahir string `json:"tanggal_lahir" validate:"required"`
	Pekerjaan    string `json:"pekerjaan" validate:"required"`
	Email        string `json:"email" validate:"required,email"`
	IdProvinsi   string `json:"id_provinsi" validate:"required,numeric"`
	IdKota       string `json:"id_kota" validate:"required,numeric"`
}
