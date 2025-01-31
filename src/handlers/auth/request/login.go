package request

type LoginRequest struct {
	NoTelepon string `json:"no_telp" validate:"required"`
	KataSandi string `json:"kata_sandi" validate:"required"`
}
