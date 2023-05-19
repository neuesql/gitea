package setting

var (
	Huggingface = struct {
		Storage
		Enabled bool
	}{
		Enabled: true,
	}
)
