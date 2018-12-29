package main

import (
	"../util"
	"path/filepath"
)

// 定义默认的配置项
const (
	defaultConfigFilename = "hctimed.conf"
	defaultDataDirname    = "data"
	defaultLogLevel       = "info"
	defaultLogDirname     = "logs"
	defaultLogFilename    = "hctimed.log"

	defaultMainnetPort = "49152"
	defaultTestnetPort = "59152"
)

var (
	defaultHomeDir       = util.AppDataDir("timed", false)
	defaultConfigFile    = filepath.Join(defaultHomeDir, defaultConfigFilename)
	defaultDataDir       = filepath.Join(defaultHomeDir, defaultDataDirname)
	defaultHPPTSKeyFile  = filepath.Join(defaultHomeDir, "https.key")
	defaultHTTPSCertFile = filepath.Join(defaultHomeDir, "https.cert")
	defaultLogDir        = filepath.Join(defaultHomeDir, defaultLogDirname)
)
// 配置项
type config struct {
	HomeDir           string   `short:"A" long:"appdata" description:"Path to application home directory"`
	ShowVersion       bool     `short:"V" long:"version" description:"Display version information and exit"`
	ConfigFile        string   `short:"C" long:"configfile" description:"Path to configuration file"`
	DataDir           string   `short:"b" long:"datadir" description:"Directory to store data"`
	LogDir            string   `long:"logdir" description:"Directory to log output."`
	TestNet           bool     `long:"testnet" description:"Use the test network"`
	SimNet            bool     `long:"simnet" description:"Use the simulation test network"`
	Profile           string   `long:"profile" description:"Enable HTTP profiling on given port -- NOTE port must be between 1024 and 65536"`
	CPUProfile        string   `long:"cpuprofile" description:"Write CPU profile to the specified file"`
	MemProfile        string   `long:"memprofile" description:"Write mem profile to the specified file"`
	DebugLevel        string   `short:"d" long:"debuglevel" description:"Logging level for all subsystems {trace, debug, info, warn, error, critical} -- You may also specify <subsystem>=<level>,<subsystem2>=<level>,... to set the log level for individual subsystems -- Use show to list available subsystems"`
	Listeners         []string `long:"listen" description:"Add an interface/port to listen for connections (default all interfaces port: 49152, testnet: 59152)"`
	WalletHost        string   `long:"wallethost" description:"Hostname for wallet server"`
	WalletCert        string   `long:"walletcert" description:"Certificate path for wallet server"`
	WalletPassphrase  string   `long:"walletpassphrase" description:"Passphrase for wallet server"`
	Version           string
	HTTPSCert         string `long:"httpscert" description:"File containing the https certificate file"`
	HTTPSKey          string `long:"httpskey" description:"File containing the https certificate key"`
	StoreHost         string `long:"storehost" description:"Enable proxy mode - send requests to the specified ip:port"`
	StoreCert         string `long:"storecert" description:"File containing the https certificate file for storehost"`
	EnableCollections bool   `long:"enablecollections" description:"Allow clienst to query collection timestamps."`
}

// runServiceCommand:仅设置为window上的实际功能，
// 用来解析和执行通过-s标记的服务命令
var runServiceCommand func(string) error

// 初始化并且解析配置，通过配置文件和命令行
// 配置过程如下
// 1） 以比较合理的设置开始默认配置
// 2） 预解析命令行以检查备用的配置文件
// 3） 加载配置文件，并覆盖任何可能的默认值
// 4） 解析命令行选项，并重写或者改变特定的选项
//
// 这样安排，使应用程序在没有任何配置文件的情况下能正常运行，同事也允许用户用配置文件来更改默认配置
// 命令行的参数往往是最优先的
func loadConfig() (*config, []string, error) {
	// 默认配置
	cfg := config{
		HomeDir:    defaultHomeDir,
		ConfigFile: defaultConfigFile,
		DebugLevel: defaultLogLevel,
		DataDir:    defaultDataDir,
		LogDir:     defaultLogDir,
		HTTPSKey:   defaultHPPTSKeyFile,
		HTTPSCert:  defaultHTTPSCertFile,
		Version:    version(),
	}
}
