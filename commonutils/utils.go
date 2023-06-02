package commonutils

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/csv"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"hash"
	"io"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/shopspring/decimal"
)

// Const declarations for common.go operations
const (
	HashSHA1 = iota
	HashSHA256
	HashSHA512
	HashSHA512_384
	MD5New
	SatoshisPerBTC = 100000000
	SatoshisPerLTC = 100000000
	WeiPerEther    = 1000000000000000000
)

// NewHTTPClientWithTimeout initialises a new HTTP client with the specified
// timeout duration
func NewHTTPClientWithTimeout(t time.Duration) *http.Client {
	h := &http.Client{Timeout: t}
	return h
}

// GetRandomSalt returns a random salt
func GetRandomSalt(input []byte, saltLen int) ([]byte, error) {
	if saltLen <= 0 {
		return nil, errors.New("salt length is too small")
	}
	salt := make([]byte, saltLen)
	if _, err := io.ReadFull(rand.Reader, salt); err != nil {
		return nil, err
	}

	var result []byte
	if input != nil {
		result = input
	}
	result = append(result, salt...)
	return result, nil
}

// GetMD5 returns a MD5 hash of a byte array
func GetMD5(input []byte) []byte {
	hash := md5.New()
	hash.Write(input)
	return hash.Sum(nil)
}

// GetSHA512 returns a SHA512 hash of a byte array
func GetSHA512(input []byte) []byte {
	sha := sha512.New()
	sha.Write(input)
	return sha.Sum(nil)
}

// GetSHA256 returns a SHA256 hash of a byte array
func GetSHA256(input []byte) []byte {
	sha := sha256.New()
	sha.Write(input)
	return sha.Sum(nil)
}

// GetHMAC returns a keyed-hash message authentication code using the desired
// hashtype
func GetHMAC(hashType int, input, key []byte) []byte {
	var hash func() hash.Hash

	switch hashType {
	case HashSHA1:
		{
			hash = sha1.New
		}
	case HashSHA256:
		{
			hash = sha256.New
		}
	case HashSHA512:
		{
			hash = sha512.New
		}
	case HashSHA512_384:
		{
			hash = sha512.New384
		}
	case MD5New:
		{
			hash = md5.New
		}
	}

	// 使用给定的hash.Hash类型和密钥返回新的HMAC哈希
	hmac := hmac.New(hash, []byte(key))
	hmac.Write(input)
	return hmac.Sum(nil)
}

// Sha1ToHex Sign signs provided payload and returns encoded string sum.
func Sha1ToHex(data string) string {
	h := sha1.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

// HexEncodeToString takes in a hexadecimal byte array and returns a string
func HexEncodeToString(input []byte) string {
	return hex.EncodeToString(input)
}

// HexDecodeToBytes takes in a hexadecimal string and returns a byte array
func HexDecodeToBytes(input string) ([]byte, error) {
	return hex.DecodeString(input)
}

// ByteArrayToString returns a string
func ByteArrayToString(input []byte) string {
	return fmt.Sprintf("%x", input)
}

// Base64Decode takes in a Base64 string and returns a byte array and an error
func Base64Decode(input string) ([]byte, error) {
	result, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Base64Encode takes in a byte array then returns an encoded base64 string
func Base64Encode(input []byte) string {
	return base64.StdEncoding.EncodeToString(input)
}

// StringSliceDifference concatenates slices together based on its index and
// returns an individual string array
func StringSliceDifference(slice1 []string, slice2 []string) []string {
	var diff []string
	for i := 0; i < 2; i++ {
		for _, s1 := range slice1 {
			found := false
			for _, s2 := range slice2 {
				if s1 == s2 {
					found = true
					break
				}
			}
			if !found {
				diff = append(diff, s1)
			}
		}
		if i == 0 {
			slice1, slice2 = slice2, slice1
		}
	}
	return diff
}

// StringContains checks a substring if it contains your input then returns a
// bool
func StringContains(input, substring string) bool {
	return strings.Contains(input, substring)
}

// StringDataContains checks the substring array with an input and returns a bool
func StringDataContains(haystack []string, needle string) bool {
	data := strings.Join(haystack, ",")
	return strings.Contains(data, needle)
}

// StringDataCompare data checks the substring array with an input and returns a bool
func StringDataCompare(haystack []string, needle string) bool {
	for x := range haystack {
		if haystack[x] == needle {
			return true
		}
	}
	return false
}

// TimeDataCompare data checks the substring array with an input and returns a bool
func TimeDataCompare(haystack []time.Time, needle time.Time) bool {
	for x := range haystack {
		if haystack[x] == needle {
			return true
		}
	}
	return false
}

// IntDataCompare data checks the substring array with an input and returns a bool
func IntDataCompare(haystack []int, needle int) bool {
	for x := range haystack {
		if haystack[x] == needle {
			return true
		}
	}
	return false
}

// Int64DataCompare data checks the substring array with an input and returns a bool
func Int64DataCompare(haystack []int64, needle int64) bool {
	for x := range haystack {
		if haystack[x] == needle {
			return true
		}
	}
	return false
}

// StringDataCompareInsensitive data checks the substring array with an input and returns
// a bool irrespective of lower or upper case strings
func StringDataCompareInsensitive(haystack []string, needle string) bool {
	for x := range haystack {
		if strings.EqualFold(haystack[x], needle) {
			return true
		}
	}
	return false
}

// StringDataContainsInsensitive checks the substring array with an input and returns
// a bool irrespective of lower or upper case strings
func StringDataContainsInsensitive(haystack []string, needle string) bool {
	for _, data := range haystack {
		if strings.Contains(StringToUpper(data), StringToUpper(needle)) {
			return true
		}
	}
	return false
}

// StringDataCompareUpper data checks the substring array with an input and returns
// a bool irrespective of lower or upper case strings
func StringDataCompareUpper(haystack []string, needle string) bool {
	for x := range haystack {
		if StringToUpper(haystack[x]) == StringToUpper(needle) {
			return true
		}
	}
	return false
}

// StringDataContainsUpper checks the substring array with an input and returns
// a bool irrespective of lower or upper case strings
func StringDataContainsUpper(haystack []string, needle string) bool {
	for _, data := range haystack {
		if strings.Contains(StringToUpper(data), StringToUpper(needle)) {
			return true
		}
	}
	return false
}

// JoinStrings joins an array together with the required separator and returns
// it as a string
func JoinStrings(input []string, separator string) string {
	return strings.Join(input, separator)
}

// SplitStrings splits blocks of strings from string into a string array using
// a separator ie "," or "_"
func SplitStrings(input, separator string) []string {
	return strings.Split(input, separator)
}

// TrimString trims unwanted prefixes or postfixes
func TrimString(input, cutset string) string {
	return strings.Trim(input, cutset)
}

// ReplaceString replaces a string with another
// 返回将s中前n个不重叠old子串都替换为new的新字符串，如果n<0会替换所有old子串。
func ReplaceString(input, old, new string, n int) string {
	return strings.Replace(input, old, new, n)
}

// StringToUpper changes strings to uppercase
func StringToUpper(input string) string {
	return strings.ToUpper(input)
}

// StringToLower changes strings to lowercase
func StringToLower(input string) string {
	return strings.ToLower(input)
}

// RoundFloat rounds your floating point number to the desired decimal place
func RoundFloat(x float64, prec int) float64 {
	var rounder float64
	pow := math.Pow(10, float64(prec))
	intermed := x * pow
	_, frac := math.Modf(intermed)
	intermed += .5
	x = .5
	if frac < 0.0 {
		x = -.5
		intermed--
	}
	if frac >= x {
		rounder = math.Ceil(intermed)
	} else {
		rounder = math.Floor(intermed)
	}

	return rounder / pow
}

// IsEnabled takes in a boolean param  and returns a string if it is enabled
// or disabled
func IsEnabled(isEnabled bool) string {
	if isEnabled {
		return "Enabled"
	}
	return "Disabled"
}

// IsValidCryptoAddress validates your cryptocurrency address string using the
// regexp package // Validation issues occurring because "3" is contained in
// litecoin and Bitcoin addresses - non-fatal
func IsValidCryptoAddress(address, crypto string) (bool, error) {
	switch StringToLower(crypto) {
	case "btc":
		return regexp.MatchString("^[13][a-km-zA-HJ-NP-Z1-9]{25,34}$", address)
	case "ltc":
		return regexp.MatchString("^[L3M][a-km-zA-HJ-NP-Z1-9]{25,34}$", address)
	case "eth":
		return regexp.MatchString("^0x[a-km-z0-9]{40}$", address)
	default:
		return false, errors.New("Invalid crypto currency")
	}
}

// YesOrNo returns a boolean variable to check if input is "y" or "yes"
func YesOrNo(input string) bool {
	if StringToLower(input) == "y" || StringToLower(input) == "yes" {
		return true
	}
	return false
}

// CalculateAmountWithFee returns a calculated fee included amount on fee
func CalculateAmountWithFee(amount, fee float64) float64 {
	return amount + CalculateFee(amount, fee)
}

// CalculateFee returns a simple fee on amount
func CalculateFee(amount, fee float64) float64 {
	return amount * (fee / 100)
}

// CalculatePercentageGainOrLoss returns the percentage rise over a certain
// period
func CalculatePercentageGainOrLoss(priceNow, priceThen float64) float64 {
	return (priceNow - priceThen) / priceThen * 100
}

// CalculatePercentageDifference returns the percentage of difference between
// multiple time periods
func CalculatePercentageDifference(amount, secondAmount float64) float64 {
	return (amount - secondAmount) / ((amount + secondAmount) / 2) * 100
}

// CalculateNetProfit returns net profit
func CalculateNetProfit(amount, priceThen, priceNow, costs float64) float64 {
	return (priceNow * amount) - (priceThen * amount) - costs
}

// JSONEncode encodes structure data into JSON
func JSONEncode(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

// JSONDecode decodes JSON data into a structure
func JSONDecode(data []byte, to interface{}) error {
	if !StringContains(reflect.ValueOf(to).Type().String(), "*") {
		return errors.New("json decode error - memory address not supplied")
	}
	return json.Unmarshal(data, to)
}

// EncodeURLValues concatenates url values onto a url string and returns a
// string
func EncodeURLValues(url string, values url.Values) string {
	path := url
	if len(values) > 0 {
		path += "?" + values.Encode()
	}
	return path
}

// ExtractHost returns the hostname out of a string
func ExtractHost(address string) string {
	host := SplitStrings(address, ":")[0]
	if host == "" {
		return "localhost"
	}
	return host
}

// ExtractPort returns the port name out of a string
func ExtractPort(host string) int {
	portStr := SplitStrings(host, ":")[1]
	port, _ := strconv.Atoi(portStr)
	return port
}

// OutputCSV dumps data into a file as comma-separated values
func OutputCSV(path string, data [][]string) error {
	_, err := ReadFile(path)
	if err != nil {
		errTwo := WriteFile(path, nil)
		if errTwo != nil {
			return errTwo
		}
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}

	writer := csv.NewWriter(file)

	err = writer.WriteAll(data)
	if err != nil {
		return err
	}

	writer.Flush()
	file.Close()
	return nil
}

func GetCaller() (filePath, fileName, funcName string, lineNum int) {
	var filePathName string
	var pc uintptr
	// Skip this function, and fetch the PC and file for its parent
	pc, filePathName, lineNum, _ = runtime.Caller(1)
	// Retrieve a Function object this functions parent
	functionObject := runtime.FuncForPC(pc)
	// Regex to extract just the function name (and not the module path)
	extractFnName := regexp.MustCompile(`^.*\.(.*)$`)
	funcName = extractFnName.ReplaceAllString(functionObject.Name(), "$1")
	filePath, fileName = filepath.Split(filePathName)
	return
}

// GetFuncName 获取调用方法的名称
func GetFuncName() string {
	// Skip this function, and fetch the PC and file for its parent
	pc, _, _, _ := runtime.Caller(1)
	// Retrieve a Function object this functions parent
	functionObject := runtime.FuncForPC(pc)
	// Regex to extract just the function name (and not the module path)
	extractFnName := regexp.MustCompile(`^.*\.(.*)$`)
	return extractFnName.ReplaceAllString(functionObject.Name(), "$1")
}

// GetFuncLine 获取调用方法的行号
func GetFuncLine() int {
	// Skip this function, and fetch the PC and file for its parent
	_, _, line, _ := runtime.Caller(1)
	return line
}

// UnixTimestampToTime returns time.time
func UnixTimestampToTime(timeint64 int64) time.Time {
	return time.Unix(timeint64, 0)
}

// UnixTimestampStrToTime returns a time.time and an error
func UnixTimestampStrToTime(timeStr string) (time.Time, error) {
	i, err := strconv.ParseInt(timeStr, 10, 64)
	if err != nil {
		return time.Time{}, err
	}

	return time.Unix(i, 0), nil
}

// CreateDir creates a directory based on the supplied parameter
func CreateDir(dir string) error {
	_, err := os.Stat(dir)
	if !os.IsNotExist(err) {
		return nil
	}

	return os.MkdirAll(dir, 0770)
}

// ReadFile reads a file and returns read data as byte array.
func ReadFile(path string) ([]byte, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return file, nil
}

// WriteFile writes selected data to a file and returns an error
func WriteFile(file string, data []byte) error {
	err := ioutil.WriteFile(file, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

// PathExists 判断文件或文件夹是否存在的方法
func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

// RemoveFile removes a file
func RemoveFile(file string) error {
	return os.Remove(file)
}

// GetURIPath returns the path of a URL given a URI
func GetURIPath(uri string) string {
	urip, err := url.Parse(uri)
	if err != nil {
		return ""
	}
	if urip.RawQuery != "" {
		return fmt.Sprintf("%s?%s", urip.Path, urip.RawQuery)
	}
	return urip.Path
}

// GetExecutablePath returns the executables launch path
func GetExecutablePath() (string, error) {
	ex, err := os.Executable()
	if err != nil {
		return "", err
	}
	return filepath.Dir(ex), nil
}

// GetOSPathSlash returns the slash used by the operating systems
// file system
func GetOSPathSlash() string {
	if runtime.GOOS == "windows" {
		return "\\"
	}
	return "/"
}

// FormatMapStringToString 将 map[string]interface{} 格式化成字符串
// commonutils.FormatMapStringToString("买进补仓", logFileds)
func FormatMapStringToString(msg string, fileds map[string]interface{}) string {

	//排序
	var keys []string
	for key := range fileds {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	//字符串拼接
	var buf bytes.Buffer
	buf.WriteByte('[')
	buf.WriteString(msg)
	buf.WriteByte(']')
	for _, key := range keys {
		if buf.Len() > 0 {
			buf.WriteByte(' ')
		}
		buf.WriteString(key)
		buf.WriteByte('=')
		stringVal, ok := fileds[key].(string)
		if !ok {
			stringVal = fmt.Sprint(fileds[key])
		}
		buf.WriteString(fmt.Sprintf("%q", stringVal))
	}
	return buf.String()
}

// FormatDecimalString 格式化后的字符串
// @param  {[type]} this [description]
// @return {[type]}      [description]
// formatDecimal_String("0.000163786", -6) //"0.000164"
func FormatDecimalString(value float64, exp int32) string {
	return decimal.NewFromFloatWithExponent(value, exp).String()
}

// FormatDecimalFloat64 获取精度计算后的数量
//
//	 example
//		  FormatDecimalFloat64(123.456, -2)   // output: 123.46
//		  FormatDecimalFloat64(-500,-2)   // output: -500
//		  FormatDecimalFloat64(1.1001, -2) // output: 1.1
//		  FormatDecimalFloat64(1.454, -1) // output: 1.5
func FormatDecimalFloat64(value float64, exp int32) float64 {
	var val, _ = decimal.NewFromFloatWithExponent(value, exp).Float64()
	return val
}

// FloatFromString format
func FloatFromString(raw interface{}) (float64, error) {
	str, ok := raw.(string)
	if !ok {
		return 0, fmt.Errorf("unable to parse, value not string: %T", raw)
	}
	flt, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0, fmt.Errorf("unable to parse, value not string: %T", raw)
	}
	return flt, nil
}

// FloatFromStringDontRound 不需要小数点四舍五入，直接取位数
//
//	 example
//		  FloatFromStringDontRound(545, -2)   // output: 500
//		  FloatFromStringDontRound(-500,-2)   // output: -500
//		  FloatFromStringDontRound(1.1001, 2) // output: 1.1
//		  FloatFromStringDontRound(1.454, 1) // output: 1.4
func FloatFromStringDontRound(num float64, exp int32) float64 {
	var val, _ = decimal.NewFromFloat(num).RoundFloor(exp).Float64()
	return val
}

// Int32ToString format
//
//	 example
//		  Int32ToString(123)   // output: "123"
//		  Int32ToString(-10)   // output: "-10"
func Int32ToString(n int32) string {
	return decimal.NewFromInt32(n).String()
}

// IntFromString format
func IntFromString(raw interface{}) (int, error) {
	str, ok := raw.(string)
	if !ok {
		return 0, fmt.Errorf("unable to parse, value not string: %T", raw)
	}
	n, err := strconv.Atoi(str)
	if err != nil {
		return 0, fmt.Errorf("unable to parse as int: %T", raw)
	}
	return n, nil
}

// Int32FromString format
func Int32FromString(raw interface{}) (int32, error) {
	str, ok := raw.(string)
	if !ok {
		return 0, fmt.Errorf("unable to parse, value not string: %T", raw)
	}
	n, err := strconv.Atoi(str)
	if err != nil {
		return 0, fmt.Errorf("unable to parse as int: %T", raw)
	}
	return int32(n), nil
}

// Int64FromString format
func Int64FromString(raw interface{}) (int64, error) {
	str, ok := raw.(string)
	if !ok {
		return 0, fmt.Errorf("unable to parse, value not string: %T", raw)
	}
	n, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("unable to parse as int: %T", raw)
	}
	return n, nil
}

// RecvWindow 格式化时间
func RecvWindow(d time.Duration) int64 {
	return int64(d) / int64(time.Millisecond)
}

// TimeFromUnixTimestampFloat format
func TimeFromUnixTimestampFloat(raw interface{}) (time.Time, error) {
	ts, ok := raw.(float64)
	if !ok {
		return time.Time{}, fmt.Errorf("unable to parse, value not int64: %T", raw)
	}
	return time.Unix(0, int64(ts)*int64(time.Millisecond)), nil
}

// UnixNesc 格式化时间
// i := int64(1532246192000)
// t := commonutils.TimeFromUnixNEscInt64(i)
// fmt.Println(t.Format("2006-01-02 15:04:05")) //2018-07-22 15:56:32
// fmt.Println(commonutils.UnixNesc(t))	// 1532246192000
func UnixNesc(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}

// TimeFromUnixNEscInt64 format
// i := int64(1532246192000)
// t := commonutils.TimeFromUnixEscInt64(i)
// fmt.Println(t.Format("2006-01-02 15:04:05")) //2018-07-22 15:56:32
func TimeFromUnixNEscInt64(i int64) time.Time {
	return time.Unix(0, int64(i)*int64(time.Millisecond))
}

// TimeFromUnixEscInt64 format
// i := int64(1530854162)
// fmt.Println(TimeFromUnixNescInt64(i)) //2018-07-06 13:16:02
func TimeFromUnixEscInt64(i int64) time.Time {
	//time.Unix(0, int64(i)*int64(time.Millisecond))
	return time.Unix(i, 0)
}

// ReverseModelLineData 对数组进行反转到一个新数组
// func ReverseModelLineData(s []*commonmodels.Kline) []*commonmodels.Kline {
// 	_length := len(s)
// 	var _tempArr []*commonmodels.Kline

// 	for i := _length - 1; i >= 0; i-- {
// 		_tempArr = append(_tempArr, s[i])
// 	}
// 	return _tempArr
// }
