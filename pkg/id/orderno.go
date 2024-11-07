package id

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type order struct{}

var (
	bits   = []int{16, 8, 4, 2, 1}
	base32 = []byte("0123456789bcdefghjkmnpqrstuvwxyz")
)

// 订单号相关
func Order() order {
	return order{}
}

// GenNO 生成订单号
func (o order) GenNO(longitude, latitude float64, orderID int64) string {
	return o.genNO(longitude, latitude, orderID, time.Now())
}

func (order) genNO(longitude, latitude float64, orderID int64, t time.Time) string {
	geo := geoHashEncode(latitude, longitude, 5)
	year := t.Year() - 2000
	month := strconv.FormatInt(int64(t.Month()), 16)
	orderIDx := strconv.FormatInt(orderID, 36)
	orderNO := strings.ToUpper(fmt.Sprint(year, month, orderIDx, geo))
	return orderNO
}

// NOExtractID 通过订单号获订单ID
func (order) NOExtractID(orderNO string) (int64, error) {
	buf := []rune(orderNO)
	if len(buf) < 9 {
		return 0, fmt.Errorf("order no extract order id error; OrderNO: %s", orderNO)
	}
	orderID := string(buf[3 : len(buf)-5])
	i, err := strconv.ParseInt(orderID, 36, 64)
	if err != nil {
		return 0, fmt.Errorf("order no extract order id error; OrderNO: %s", orderNO)
	}
	return i, nil
}

// NOExtractYear 通过订单号获取订单年份
func (order) NOExtractYear(orderNO string) (int, error) {
	buf := []rune(orderNO)
	if len(buf) < 3 {
		return 0, fmt.Errorf("order no extract order year error; OrderNO: %s", orderNO)
	}
	i, err := strconv.ParseInt(string(buf[0:2]), 10, 32)
	if err != nil {
		return 0, fmt.Errorf("order no extract order year error; OrderNO: %s", orderNO)
	}
	return int(i + 2000), nil
}

// NOExtractMonth 通过订单号获取月份
func (order) NOExtractMonth(orderNO string) (int, error) {
	buf := []rune(orderNO)
	if len(buf) < 3 {
		return 0, fmt.Errorf("order no extract order month error; OrderNO: %s", orderNO)
	}
	i, err := strconv.ParseInt(string(buf[2]), 16, 32)
	if err != nil {
		return 0, fmt.Errorf("order no extract order month error; OrderNO: %s", orderNO)
	}
	return int(i), nil
}

// NOExtractGeo  通过订单号获经纬度
func (order) NOExtractGeo(orderNO string) string {
	buf := []rune(orderNO)
	if len(buf) < 5 {
		return ""
	}
	orderGeo := string(buf[len(buf)-5:])
	return strings.ToLower(orderGeo)
}

// geoHashEncode Create a geohash with 12 positions based on LatLng coordinates
func geoHashEncode(latitude, longitude float64, len int) string {
	return encodeWithPrecision(latitude, longitude, len)
}

// encodeWithPrecision Create a geohash with given precision (number of characters of the resulting
// hash) based on LatLng coordinates
func encodeWithPrecision(latitude, longitude float64, precision int) string {
	isEven := true
	lat := []float64{-90, 90}
	lng := []float64{-180, 180}
	bit := 0
	ch := 0
	var geohash bytes.Buffer
	var mid float64
	for geohash.Len() < precision {
		if isEven {
			mid = (lng[0] + lng[1]) / 2
			if longitude > mid {
				ch |= bits[bit]
				lng[0] = mid
			} else {
				lng[1] = mid
			}
		} else {
			mid = (lat[0] + lat[1]) / 2
			if latitude > mid {
				ch |= bits[bit]
				lat[0] = mid
			} else {
				lat[1] = mid
			}
		}
		isEven = !isEven
		if bit < 4 {
			bit++
		} else {
			geohash.WriteByte(base32[ch])
			bit = 0
			ch = 0
		}
	}
	return geohash.String()
}
