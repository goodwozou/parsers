package parsers

// ZombieMap map zombie service to gogo finger
var ZombieMap = map[string]string{
	"mariadb":   "MYSQL",
	"mysql":     "MYSQL",
	"rdp":       "RDP",
	"oracle":    "ORACLE",
	"sqlserver": "MSSQL",
	"mssql":     "MSSQL",
	"smb":       "SMB",
	"redis":     "REDIS",
	"vnc":       "VNC",
	//"elasticsearch": "ELASTICSEARCH",
	"postgresql": "POSTGRESQL",
	"mongo":      "MONGO",
	"ssh":        "SSH",
	"ftp":        "FTP",
	"socks5":     "SOCKS5",
	"rsync":      "RSYNC",
	"telnet":     "TELNET",
}

type ZombieInput struct {
	IP      string            `json:"ip"`
	Port    string            `json:"port"`
	Service string            `json:"service"`
	Param   map[string]string `json:"param"`
}
