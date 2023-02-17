package sutils

import (
	"bytes"
	"compress/gzip"
	"io"
	"net"
	"net/http"

	"github.com/sirupsen/logrus"
)

const (
	ContentEncodingHeader = "Content-Encoding"
	ContentEncodingGZIP   = "gzip"
)

func CloneRequestBody(r *http.Request) (body string) {
	if r == nil {
		return ""
	}

	var bodyRC io.ReadCloser
	switch {
	case r.GetBody != nil:
		if rBody, err := r.GetBody(); err == nil {
			bodyRC = rBody
			break
		}
		fallthrough
	case r.Body != nil:
		bodyRC = r.Body
	default:
		return ""
	}

	// read body
	bodyContent, err := io.ReadAll(bodyRC)
	if err != nil {
		logrus.Warnf("error reading request body: %v", err)
		return ""
	}
	err = bodyRC.Close()
	if err != nil {
		logrus.Warnf("error closing request body: %v", err)
	}

	// set body for future reading
	r.Body = io.NopCloser(bytes.NewReader(bodyContent))
	r.GetBody = func() (io.ReadCloser, error) {
		return io.NopCloser(bytes.NewReader(bodyContent)), nil
	}

	return string(bodyContent)
}

func CloneResponseBody(r *http.Response) (body string) {
	if r == nil || r.Body == nil {
		return ""
	}
	bodyContent, err := io.ReadAll(r.Body)
	if err != nil {
		logrus.Warnf("error reading response body: %v", err)
	}
	if err := r.Body.Close(); err != nil {
		logrus.Warnf("error closing response body: %v", err)
	}
	if r.Header.Get(ContentEncodingHeader) == ContentEncodingGZIP {
		if bodyContent, err = UnGzipBody(bodyContent); err != nil {
			logrus.Warnf("ungzip response body error: %v", err)
		}
	}
	r.Body = io.NopCloser(bytes.NewReader(bodyContent))
	return string(bodyContent)
}

func UnGzipBody(bodyContent []byte) ([]byte, error) {
	gzipReader, err := gzip.NewReader(bytes.NewReader(bodyContent))
	if err != nil {
		return bodyContent, err
	}
	gzipContent, err := io.ReadAll(gzipReader)
	if err != nil {
		return bodyContent, err
	}
	return gzipContent, nil
}

const xForwardedForHeader = "X-Forwarded-For"

func GetRequestIP(r *http.Request) (requestIP string) {
	if ip := r.Header.Get(xForwardedForHeader); ip != "" {
		requestIP = ip
	} else {
		requestIP, _, _ = net.SplitHostPort(r.RemoteAddr)
	}
	return
}

func SplitIPVersions(ip string) (ipv4, ipv6 string) {
	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		ip = ""
		parsedIP = net.ParseIP(ip)
	}
	if parsedIP.To4() != nil {
		return ip, ""
	}
	return "", ip
}

func JoinIPVersions(ipv4, ipv6 string) (ip string) {
	switch {
	case net.ParseIP(ipv4) != nil:
		ip = ipv4
	case net.ParseIP(ipv6) != nil:
		ip = ipv6
	}
	return
}

const (
	DefaultIPv4 = "0.0.0.0"
	DefaultIPv6 = "0000:0000:0000:0000:0000:0000:0000:0000"
)

func FillEmptyIP(ipv4, ipv6 string) (string, string) {
	if net.ParseIP(ipv4) == nil {
		ipv4 = DefaultIPv4
	}
	if net.ParseIP(ipv6) == nil {
		ipv6 = DefaultIPv6
	}
	return ipv4, ipv6
}

const (
	accessControlHeader = "Access-Control-Allow-Origin"
	contentTypeHeader   = "Content-Type"
)

type ContentType string

const (
	ContentTypeApplicationJSON ContentType = "application/json"
	ContentTypeTextHTML        ContentType = "text/html"
)

func SetCorsHeader(w http.ResponseWriter) {
	w.Header().Set(accessControlHeader, "*")
}

func SetContentTypeHeader(w http.ResponseWriter, contentType ContentType) {
	w.Header().Set(contentTypeHeader, string(contentType))
}
