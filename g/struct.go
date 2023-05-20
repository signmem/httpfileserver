package g

type GlobalConfig struct {
	Debug			bool		`json:"debug"`
	LogFile			string		`json:"logfile"`
	LogMaxAge		int			`json:"logmaxage"`
	LogRotateAge	int			`json:"logrotateage"`
	DownloadDir		string		`json:"downloaddir"`
	MaxUploadSize	int64		`json:"maxuploadsize"`
	Http 			*HTTP		`json:"http"`
}

type HTTP struct {
	Address			string		`json:"address"`
	Port			string		`json:"port"`
}


type BackupPeriod struct {
	FSID  			int64 		`json:"id"`
	FSPath 			string 		`json:"fspath"`
	FSName			string		`json:"fsname"`
	FSType			string 		`json:"fstype"`    // temp , perm
	FSStatus 		int 		`json:"status"`    // 0 create 9 done
	FSSize 			int64 		`json:"size"`
	FSClient 		string		`json:"client"`
	FSAgent 		string 		`json:"agent"`
	FSCreateTime 	string 		`json:"create_time"`
	FSUpdateTime 	string 		`json:"update_time"`
	FSToken			string		`json:"token"`
	FSStruct 		string 		`json:"struct"`		// file , directory
	FSRetry 		bool 		`json:"retry"`
	GFSCluster		string		`json:"gfscluster"`
}