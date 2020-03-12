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

type CachedRequest struct {
	Data        []byte
	RequestedOn time.Time
	ETag        string
}

func (c *CachedRequest) ToJSON(i interface{}) error {
	return json.Unmarshal(c.Data, i)
}

func hash(s string) string {
	return "j2gc-" + hex.EncodeToString(sha1.New().Sum([]byte(s)))
}

func cacheExists(reqHash string) bool {
	data, _ := os.Stat(path.Join(Constants.CacheDir, reqHash))
	return data != nil
}

func writeCache(reqHash string, request *CachedRequest) error {
	b, err := json.Marshal(request)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path.Join(Constants.CacheDir, reqHash), b, os.ModePerm)
}

func readCache(reqhash string) (*CachedRequest, error) {
	if !cacheExists(reqhash) {
		return nil, fmt.Errorf("Cache file not found")
	}

	cr := &CachedRequest{}
	b, err := ioutil.ReadFile(path.Join(Constants.CacheDir, reqhash))
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, cr)
	return cr, err
}

func CachedReqGet(url string, v ...interface{}) (*CachedRequest, error) {
	requestHash := hash(strings.ReplaceAll(url, Constants.API, ""))

	cr := &CachedRequest{}
	if !cacheExists(requestHash) {

		data, err := req.Get(url, v...)
		if err != nil {
			return nil, err
		}

		cr.Data = data.Bytes()
		cr.RequestedOn = time.Now()
		cr.ETag = data.Response().Header.Get("ETag")

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

	if time.Now().Sub(cache.RequestedOn) > Constants.CacheLifetime {
		v = append(v, req.Header{
			"If-None-Match": cache.ETag,
		})

		data, err := req.Get(url, v...)
		if err != nil {
			log.Printf("[%v] Failed to renew data: %v\n", requestHash, err)
			return nil, fmt.Errorf("[%v] Failed to renew data: %v\n", requestHash, err)
		}

		log.Printf("[%v] Status Code: %v\n", requestHash, data.Response().StatusCode)
		if data.Response().StatusCode == 304 {
			log.Printf("[%v] Cached data is still synced with remote resource\n", requestHash)
		}

		log.Printf("[%v] New ETag: %v\n", requestHash, data.Response().Header.Get("ETag"))

		// If the server returns with 304, just refresh the cache with an updated expiry
		if data.Response().StatusCode != 304 && data.Response().StatusCode == 200 {
			cache.Data = data.Bytes()
			cache.ETag = data.Response().Header.Get("ETag")
		}
		cache.RequestedOn = time.Now()

		// Write the cache
		err = writeCache(requestHash, cache)
		if err != nil {
			fmt.Printf("[%v] Failed to write cache: %v\n", requestHash, err)
			return nil, fmt.Errorf("[%v] Failed to write cache: %v\n", requestHash, err)
		}
	}

	return cache, err
}
