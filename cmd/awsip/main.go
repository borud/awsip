package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/borud/aws-ip-ranges/pkg/awsip"
	"github.com/borud/aws-ip-ranges/pkg/util"
)

//lint:file-ignore SA5008 Ignore duplicate struct tags
type options struct {
	NoCache     bool          `short:"n" long:"no-cache" description:"do not use cache, download list directly"`
	MaxCacheAge time.Duration `short:"a" long:"max-cache-age" default:"72h" description:"max age of cached file before re-download"`
	RangesURL   string        `short:"u" long:"url" default:"https://ip-ranges.amazonaws.com/ip-ranges.json" description:"URL of AWS ranges JSON file"`
	Verbose     bool          `short:"v" long:"verbose" description:"verbose output"`

	Range    rangeCmd    `command:"range" description:"list ranges"`
	Regions  regionsCmd  `command:"regions" description:"list regions"`
	Services servicesCmd `command:"services" description:"list services"`
}

const (
	cacheFilename = ".aws-ranges-cache.json"
)

var opt options

func main() {
	util.FlagParse(&opt)
}

func getRanges(o options) (*awsip.Ranges, error) {
	// Handle the no-cache case first
	if o.NoCache {
		if o.Verbose {
			log.Printf("skipping cache, reading data from %s", o.RangesURL)
		}
		resp, err := http.Get(o.RangesURL)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		return awsip.Read(resp.Body)
	}

	homedir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	cacheFile := path.Join(homedir, cacheFilename)
	info, err := os.Stat(cacheFile)

	// Check if cache exists or is older than MaxCacheAge
	if errors.Is(err, os.ErrNotExist) || info.ModTime().After(time.Now().Add(o.MaxCacheAge)) {
		if o.Verbose {
			log.Printf("fetching %s", o.RangesURL)
		}
		resp, err := http.Get(o.RangesURL)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		writeFile, err := os.OpenFile(cacheFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
		if err != nil {
			return nil, fmt.Errorf("error opening file %s: %v", cacheFile, err)
		}
		defer writeFile.Close()

		rangeJSON, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("error reading HTTP response: %w", err)
		}

		if o.Verbose {
			log.Printf("size of response: %d", len(rangeJSON))
		}

		n, err := writeFile.Write(rangeJSON)
		if err != nil {
			return nil, fmt.Errorf("error writing file %s: %w", cacheFile, err)
		}

		if n < len(rangeJSON) {
			return nil, fmt.Errorf("short write, got %d bytes but only wrote %d bytes to %s", len(rangeJSON), n, cacheFile)
		}

		return awsip.Read(bytes.NewReader(rangeJSON))
	}

	if o.Verbose {
		log.Printf("reading cached data from %s", cacheFile)
	}
	f, err := os.Open(cacheFile)
	if err != nil {
		return nil, fmt.Errorf("error reading cached data from %s: %w", cacheFile, err)
	}
	defer f.Close()
	return awsip.Read(f)
}
