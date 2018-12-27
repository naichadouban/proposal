package main

import (
	"../util"
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
// 定义默认的配置项
var (
	defaultHomeDir = util.
)
func loadConfig() (*config,[]string,error) {

}
