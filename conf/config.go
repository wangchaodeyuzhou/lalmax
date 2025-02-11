package config

import (
	"encoding/json"
	"io/ioutil"
)

var defaultConfig Config

type Config struct {
	SrtConfig        SrtConfig      `json:"srt_config"`      // srt配置
	RtcConfig        RtcConfig      `json:"rtc_config"`      // rtc配置
	HttpConfig       HttpConfig     `json:"http_config"`     // http/https配置
	HttpFmp4Config   HttpFmp4Config `json:"httpfmp4_config"` // http-fmp4配置
	HlsConfig        HlsConfig      `json:"hls_config"`      // hls-fmp4/llhls配置
	GB28181Config    GB28181Config  `json:"gb28181_config"`  // gb28181配置
	OnvifConfig      OnvifConfig    `json:"onvif_config"`    //
	LalSvrConfigPath string         `json:"lal_config_path"` // lal配置目录
}

type SrtConfig struct {
	Enable bool   `json:"enable"` // srt服务使能配置
	Addr   string `json:"addr"`   // srt服务监听地址
}

type RtcConfig struct {
	Enable          bool     `json:"enable"`          // rtc服务使能配置
	ICEHostNATToIPs []string `json:"iceHostNatToIps"` // rtc服务公网IP，未设置使用内网
	ICEUDPMuxPort   int      `json:"iceUdpMuxPort"`   // rtc udp mux port
	ICETCPMuxPort   int      `json:"iceTcpMuxPort"`   // rtc tcp mux port
}

type HttpConfig struct {
	ListenAddr      string `json:"http_listen_addr"`  // http服务监听地址
	EnableHttps     bool   `json:"enable_https"`      // https使能标志
	HttpsListenAddr string `json:"https_listen_addr"` // https监听地址
	HttpsCertFile   string `json:"https_cert_file"`   // https cert 文件
	HttpsKeyFile    string `json:"https_key_file"`    // https key 文件
}

type HttpFmp4Config struct {
	Enable bool `json:"enable"` // http-fmp4使能标志
}

type HlsConfig struct {
	Enable          bool `json:"enable"`          // hls使能标志
	SegmentCount    int  `json:"segmentCount"`    // 分片个数,llhls默认7个
	SegmentDuration int  `json:"segmentDuration"` // hls分片时长,默认1s
	PartDuration    int  `json:"partDuration"`    // llhls part时长,默认200ms
	LowLatency      bool `json:"lowLatency"`      // 是否开启llhls
}

type GB28181Config struct {
	Enable     bool   `json:"enable"`     // gb28181使能标志
	ListenAddr string `json:"listenAddr"` // gb28181监听地址
	SipNetwork string `json:"sipNetwork"` // 传输协议，默认UDP，可选TCP
	SipIP      string `json:"sipIp"`      // sip 服务器公网IP
	SipPort    uint16 `json:"sipPort"`    // sip 服务器端口，默认 5060
	Serial     string `json:"serial"`     // sip 服务器 id, 默认 34020000002000000001
	Realm      string `json:"realm"`      // sip 服务器域，默认 3402000000
	Username   string `json:"username"`   // sip 服务器账号
	Password   string `json:"password"`   // sip 服务器密码
}

type OnvifConfig struct {
	Enable bool `json:"enable"`
}

func Open(filepath string) error {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &defaultConfig)
	if err != nil {
		return err
	}
	return nil
}

func GetConfig() *Config {
	return &defaultConfig
}
