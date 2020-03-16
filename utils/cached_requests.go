package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
	"time"

	"github.com/imroc/req"
)

var lastUncachedRequest = time.Now()

type CachedRequest struct {
	Data        []byte
	RequestedOn time.Time
	ETag        string
}

func (c *CachedRequest) ToJSON(i interface{}) error {
	return json.Unmarshal(c.Data, i)
}

func hash(s string) string {
	x := sha1.New()
	x.Write([]byte(s))
	return "j2gc-" + hex.EncodeToString(x.Sum(nil))
}

func cacheExists(reqHash string) bool {
	data, _ := os.Stat(path.Join(Config.CacheDir, reqHash))
	return data != nil
}

func writeCache(reqHash string, request *CachedRequest) error {
	b, err := json.Marshal(request)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path.Join(Config.CacheDir, reqHash), b, os.ModePerm)
}

func readCache(reqhash string) (*CachedRequest, error) {
	if !cacheExists(reqhash) {
		return nil, fmt.Errorf("Cache file not found")
	}

	cr := &CachedRequest{}
	b, err := ioutil.ReadFile(path.Join(Config.CacheDir, reqhash))
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, cr)
	return cr, err
}

func CachedReqGet(url string, v ...interface{}) (*CachedRequest, error) {
	ttl := time.Now().Sub(lastUncachedRequest)
	if ttl < time.Millisecond*500 {
		sleepTime := (time.Millisecond * 500) - ttl
		log.Println("Sleeping for", sleepTime)
		time.Sleep(sleepTime)
	}

	requestHash := hash(strings.ReplaceAll(url, Config.API, "") + fmt.Sprintf("%v", v))

	cr := &CachedRequest{}
	if !cacheExists(requestHash) {

		data, err := req.Get(url, v...)
		if err != nil {
			fmt.Println("Cached Request Get Error:", err)
			return nil, err
		}

		switch data.Response().StatusCode {

		case 404:
			return nil, fmt.Errorf("resource not found")
		case 400:
			return nil, fmt.Errorf("bad request")
		case 405:
			return nil, fmt.Errorf("method not allowed")
		case 429:
			return nil, fmt.Errorf("invalid request")
		case 500:
			return nil, fmt.Errorf("internal server error")
		}

		cr.Data = data.Bytes()
		cr.RequestedOn = time.Now()
		cr.ETag = data.Response().Header.Get("ETag")

		if strings.Contains(string(cr.Data), "\"request_cached\": false") {
			log.Println("uncached request, setting timeout for next call")
			lastUncachedRequest = time.Now()
		}

		err = writeCache(requestHash, cr)
		if err != nil {
			return nil, fmt.Errorf("[%v] Failed to write cache: %v\n", requestHash, err)
		}
		return cr, nil
	}

	cache, err := readCache(requestHash)
	if err != nil {
		return nil, err
	}

	if time.Now().Sub(cache.RequestedOn) > Config.CacheLifetime {
		v = append(v, req.Header{
			"If-None-Match": cache.ETag,
		})

		data, err := req.Get(url, v...)
		if err != nil {
			log.Printf("[%v] Failed to renew data: %v\n", requestHash, err)
			return nil, fmt.Errorf("[%v] Failed to renew data: %v\n", requestHash, err)
		}

		//log.Printf("[%v] Status Code: %v\n", requestHash, data.Response().StatusCode)
		if data.Response().StatusCode == 304 {
			log.Printf("[%v] Cached data is still synced with remote resource\n", requestHash)
		}

		//log.Printf("[%v] New ETag: %v\n", requestHash, data.Response().Header.Get("ETag"))

		// If the server returns with 304, just refresh the cache with an updated expiry
		if data.Response().StatusCode != 304 && data.Response().StatusCode == 200 {
			cache.Data = data.Bytes()
			cache.ETag = data.Response().Header.Get("ETag")
		}
		cache.RequestedOn = time.Now()
		if strings.Contains(string(cr.Data), "\"request_cached\": false") {
			log.Println("uncached request, setting timeout for next call")
			lastUncachedRequest = time.Now()
		}

		// Write the cache
		err = writeCache(requestHash, cache)
		if err != nil {
			fmt.Printf("[%v] Failed to write cache: %v\n", requestHash, err)
			return nil, fmt.Errorf("[%v] Failed to write cache: %v\n", requestHash, err)
		}
	}

	return cache, err
}
