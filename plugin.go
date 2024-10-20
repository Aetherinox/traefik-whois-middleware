/*
	Package
*/

package traefik_whois_middleware

/*
	Imports
*/

import (
	"regexp"
	"io"
	"context"
	"fmt"
	"net"
	"net/http"
	"time"
	"strings"
	"log"
	"strconv"
	"encoding/json"
	"os"
)

/*
	Define > Color Codes
*/

var Reset = "\033[0m"
var Red = "\033[31m"
var RedL = "\033[91m"
var Green = "\033[32m"
var GreenL = "\033[92m"
var Orange = "\033[33m"
var Yellow = "\033[93m"
var Blue = "\033[34m"
var BlueL = "\033[94m"
var PurpleL = "\033[95m"
var Magenta = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var GrayD = "\033[90m"
var White = "\033[97m"

var BBlack = "\033[1;30m"
var BRed = "\033[1;31m"
var BGreen = "\033[1;32m"
var BYellow = "\033[1;33m"
var BBlue = "\033[1;34m"
var BPurple = "\033[1;35m"
var BCyan = "\033[1;36m"
var BWhite = "\033[1;37m"

type logWriter struct {
}

/*
	Logs > Writer
*/

func (writer logWriter) Write(bytes []byte) (int, error) {
	str := GrayD + time.Now().Format("2006-01-02T15:04:05") + Reset + " " + string(bytes)
	return io.WriteString(os.Stderr, str)
}

/*
	Logging
*/

var ( 
	logInfo = log.New(io.Discard, BPurple + "[ WHOIS ] " + BlueL + "[INFO] " + Reset + ": ", log.Ldate|log.Ltime)
	logErr = log.New(io.Discard, BPurple + "[ WHOIS ] " + RedL + "[ERROR] " + Reset + ": ", log.Ldate|log.Ltime)
	logWarn = log.New(io.Discard, BPurple + "[ WHOIS ] " + Orange + "[WARN] " + Reset + ": ", log.Ldate|log.Ltime)
	logDebug = log.New(io.Discard, BPurple + "[ WHOIS ] " + GrayD + "[Debug] " + Reset + ": ", log.Ldate|log.Ltime)
)

/*
	Define > Header Values
*/

const (
	xForwardedFor                      = "X-Forwarded-For"
	xRealIP                            = "X-Real-IP"
	countryHeader                      = "X-IPCountry"
)

/*
	Construct Configurations

	OTP Secret can be generated at		https://it-tools.tech/otp-generator
*/

type Config struct {
	DebugLogs					bool		`json:"debugLogs,omitempty"`
}

/*
	Create Config
*/

func CreateConfig() *Config {
	return &Config{
		DebugLogs:					false,
	}
}

type KeyAuth struct {
	next                     	http.Handler
	debugLogs           		bool
}

/*
	Strings > Slice
*/

func sliceString(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}

	return false
}

/*
	iP > Slice

	returns true of ips match
*/

func sliceIp(a net.IP, list []net.IP) bool {
	for _, b := range list {
		if b.Equal(a) {
			return true
		}
	}

	return false
}

/*
	IP > Parse
*/

func parseIP(addr string) (net.IP, error) {
	ipAddress := net.ParseIP(addr)

	if ipAddress == nil {
		return nil, fmt.Errorf("cant parse IP address from address [%s]", addr)
	}

	return ipAddress, nil
}

func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	fmt.Printf( Red + "[Aetherx-whois]: " + Reset + "Starting Plugin " + Magenta + "%s" + Reset + "\n instance: " + Yellow + "%+v" + Reset + "\n ctx: " + Yellow + "%+v \n\n", name, *config, ctx)

	/*
		@TODO		merge logs
	*/

	logInfo.SetFlags(0)
	logInfo.SetOutput(new(logWriter))
	// logInfo.SetOutput(os.Stdout)

	logErr.SetFlags(0)
	logErr.SetOutput(new(logWriter))

	logWarn.SetFlags(0)
	logWarn.SetOutput(new(logWriter))

	logDebug.SetFlags(0)
	logDebug.SetOutput(new(logWriter))

	/*
		return structure
	*/

	return &KeyAuth {
		next:                     	next,
		debugLogs:					config.DebugLogs,
	}, nil
}

/*
	taken api tokens and compare to list of valid tokens.
	return if specified token is valid
*/

func contains(token string, validTokens []string) bool {
	for _, a := range validTokens {
		if a == token {
			return true
		}
	}

	return false
}

/*
	Get or Default Value
*/

func get(val string, deflt string) string  {
    if len(val) < 1 {
        return deflt
    }
    return val
}

/*
	Serve
*/

func (ka *KeyAuth) ServeHTTP(rw http.ResponseWriter, req *http.Request) {

	reqIPAddr, err := ka.collectRemoteIP(req)

	if err != nil {
		logErr.Printf(Reset + "Error: " + Yellow + "%s" + Reset, err)
	}

	now := time.Now().Format(time.UnixDate) // UnixDate
	urlFull := fmt.Sprintf("%s%s", req.Host, req.URL)
	userAgent := req.UserAgent()
	userIp := "Unknown"

	/*
		logs > output user
		assign user ip to string
	*/

	for _, ipAddress := range reqIPAddr {
		userIp = ipAddress.String()
		logInfo.Printf(Reset + "New connection from " + Yellow + "%s" + Reset + " for url " + Yellow + "%s" + Reset, userIp, urlFull)
	}

	/*
		Debug Logs
		All users should pass this step before being directed to their proper destinations
	*/

	logInfo.Printf(Yellow + "%s :" + Reset + " %s ", "Host", req.Host)
	logInfo.Printf(Yellow + "%s :" + Reset + " %s ", "Remote Address", req.RemoteAddr)
	logInfo.Printf(Yellow + "%s :" + Reset + " %s ", "User Agent", userAgent)
	logInfo.Printf(Yellow + "%s :" + Reset + " %s ", "User IP", userIp)
	logInfo.Printf(Yellow + "%s :" + Reset + " %s ", "User (X-Forwarded-For)",  get(req.Header.Get("X-Forwarded-For"), "none"))
	logInfo.Printf(Yellow + "%s :" + Reset + " %s ", "User (X-Real-IP)", get(req.Header.Get("X-Real-IP"), "none"))
	logInfo.Printf(Yellow + "%s :" + Reset + " %s ", "User (X-IPCountry)", get(req.Header.Get("X-IPCountry"), "none"))
	logInfo.Printf(Yellow + "%s :" + Reset + " %s ", "User (Cf-Connecting-Ip)", get(req.Header.Get("Cf-Connecting-Ip"), "none"))
	logInfo.Printf(Yellow + "%s :" + Reset + " %s ", "User Country (Cf-Ipcountry)", get(req.Header.Get("Cf-Ipcountry"), "none"))
	logInfo.Printf(Yellow + "%s :" + Reset + " %s ", "Target URL", urlFull)
	logInfo.Printf(Yellow + "%s :" + Reset + " %s ", "Request URI", req.RequestURI)
	logInfo.Printf(Yellow + "%s :" + Reset + " %s ", "Request URL", req.URL)
	logInfo.Printf(Yellow + "%s :" + Reset + " %s ", "Method", req.Method)
	logInfo.Printf(Yellow + "%s :" + Reset + " %s ", "Prototype", req.Proto)
	logInfo.Printf(Yellow + "%s :" + Reset + " %s ", "Timestamp", now)

	ka.next.ServeHTTP(rw, req)
}

/*
	Collect Remote IP
*/

func (a *KeyAuth) collectRemoteIP(req *http.Request) ([]*net.IP, error) {
	var ipList []*net.IP

	splitFn := func(c rune) bool {
		return c == ','
	}

	xForwardedForValue := req.Header.Get(xForwardedFor)
	xForwardedForIPs := strings.FieldsFunc(xForwardedForValue, splitFn)

	xRealIPValue := req.Header.Get(xRealIP)
	xRealIPList := strings.FieldsFunc(xRealIPValue, splitFn)

	for _, value := range xForwardedForIPs {
		ipAddress, err := parseIP(value)
		if err != nil {
			return ipList, fmt.Errorf("parsing failed: %s", err)
		}

		ipList = append(ipList, &ipAddress)
	}

	for _, value := range xRealIPList {
		ipAddress, err := parseIP(value)
		if err != nil {
			return ipList, fmt.Errorf("parsing failed: %s", err)
		}

		ipList = append(ipList, &ipAddress)
	}

	return ipList, nil
}
