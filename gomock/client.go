package gomock

type sshClient interface {
	Close() error
}
