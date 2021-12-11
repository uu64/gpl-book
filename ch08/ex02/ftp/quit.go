package ftp

func (fc *ftpConn) quit() error {
	// TODO: データ転送中ならコネクションはまだ閉じない
	// TODO: コントロール接続に関してやることはあるか
	return fc.reply(status221)
}
