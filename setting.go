package main

type Setting struct {
	// Host Setting
	Host string
	Port string

	// Database setting
	DBHost string
	DBName string
	DBUser string
	DBPwd  string
}

func (s *Setting) GetListenAndServe() string {
	return s.Host + ":" + s.Port
}

func (s *Setting) GetDBSetting() (host string, name string, user string, pwd string) {
	return s.DBHost, s.DBName, s.DBUser, s.DBPwd
}
