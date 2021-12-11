package ftp

func (fc *ftpConn) user(name string) error {
	fc.username = &name
	// TODO: auth
	return fc.reply(status230)
}
