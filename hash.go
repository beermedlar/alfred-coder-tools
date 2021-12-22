package main

// Package is called aw
import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"net/url"
	"strconv"
	"time"

	aw "github.com/deanishe/awgo"
)

// Workflow is the main API
var wf *aw.Workflow

func init() {
	// Create a new Workflow using default settings.
	// Critical settings are provided by Alfred via environment variables,
	// so this *will* die in flames if not run in an Alfred-like environment.
	wf = aw.New()
}

// Your workflow starts here
func run() {
	operator := wf.Args()[0]
	query := wf.Args()[1]

	var result string
	switch operator {
	case "md5":
		result = md5V(query)
	case "sha1":
		result = sha1V(query)
	case "sha256":
		result = sha256V(query)
	case "sha512":
		result = sha512V(query)
	case "base64_decode":
		result = base64_decode(query)
	case "base64_encode":
		result = base64_encode(query)
	case "url_encode":
		result = url.QueryEscape(query)
	case "url_decode":
		result, _ = url.QueryUnescape(query)
	case "df":
		if query == "now" {
			result = timestampsToTime(time.Now().UnixNano() / int64(time.Millisecond))
		} else {
			timestamps, _ := strconv.ParseInt(query, 10, 64)
			result = timestampsToTime(timestamps)
		}
	default:
	}
	// Add a "Script Filter" result
	wf.NewItem(result).Subtitle("回车进行复制").Arg(result).Valid(true)
	// Send results to Alfred
	wf.SendFeedback()
}

func main() {
	// Wrap your entry point with Run() to catch and log panics and
	// show an error in Alfred instead of silently dying
	wf.Run(run)
}

func md5V(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func sha1V(str string) string {
	h := sha1.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func sha256V(str string) string {
	h := sha256.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func sha512V(str string) string {
	h := sha512.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func base64_encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func base64_decode(str string) string {
	sDec, _ := base64.StdEncoding.DecodeString(str)
	return string(sDec)
}

func timestampsToTime(long int64) string {
	time := time.UnixMilli(long)
	return time.Format("2006-01-02 15:04:05")
}
