package secure_manager

const (
	SecureManagerPrefix = "[Secure Manager]"
)

func InitSecureManager() {
	StartSSHGuardWatcher()
}
