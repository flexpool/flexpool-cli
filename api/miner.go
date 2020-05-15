package api

import (
	"encoding/json"
	"io/ioutil"
	"math/big"
	"net/http"

	"github.com/flexpool/flexpool-cli/types"
)

type jsonResponseBool struct {
	Result bool  `json:"result"`
	Error  error `json:"error"`
}

// FullStats is a full API stats object struct
type FullStats struct {
	ReportedHashrate  float64 `json:"reported_hashrate"`
	EffectiveHashrate float64 `json:"effective_hashrate"`
	ValidShares       uint64  `json:"valid_shares"`
	StaleShares       uint64  `json:"stale_shares"`
	InvalidShares     uint64  `json:"invalid_shares"`
}

// FullStatsWithTimestamp is a full API stats object struct (with timestamp)
type FullStatsWithTimestamp struct {
	ReportedHashrate       float64 `json:"reported_hashrate"`
	EffectiveHashrate      float64 `json:"effective_hashrate"`
	ValidShares            uint64  `json:"valid_shares"`
	StaleShares            uint64  `json:"stale_shares"`
	InvalidShares          uint64  `json:"invalid_shares"`
	FullStatsWithTimestamp int64   `json:"timestamp"`
}

// HashrateStats is a API hashrate stats object struct
type HashrateStats struct {
	ReportedHashrate  float64 `json:"reported_hashrate"`
	EffectiveHashrate float64 `json:"effective_hashrate"`
}

// DailyCurrentStats is a <miner>/stats result object
type DailyCurrentStats struct {
	Current HashrateStats `json:"current"`
	Daily   FullStats     `json:"daily"`
}

type jsonResponseStats struct {
	Result DailyCurrentStats `json:"result"`
	Error  error             `json:"error"`
}

type jsonResponseStringSlice struct {
	Result []string `json:"result"`
	Error  error    `json:"error"`
}

type jsonResponseChart struct {
	Result []FullStatsWithTimestamp `json:"result"`
	Error  error                    `json:"error"`
}

// WorkersCount is a worker count struct (online/offline)
type WorkersCount struct {
	Online  int `json:"online"`
	Offline int `json:"offline"`
}

type jsonResponseWorkerCount struct {
	Result WorkersCount `json:"result"`
	Error  error        `json:"error"`
}

// Worker is a worker object struct
type Worker struct {
	Name              string  `json:"name"`
	Online            bool    `json:"online"`
	ReportedHashrate  float64 `json:"reported_hashrate"`
	EffectiveHashrate float64 `json:"effective_hashrate"`
	ValidShares       uint64  `json:"valid_shares"`
	StaleShares       uint64  `json:"stale_shares"`
	InvalidShares     uint64  `json:"invalid_shares"`
}

type jsonResponseWorkers struct {
	Result []Worker `json:"result"`
	Error  error    `json:"error"`
}

type jsonResponseBalance struct {
	Result types.JSONBigInt `json:"result"`
	Error  error            `json:"error"`
}

// MinerDetail is a miner details API method response struct
type MinerDetail struct {
	MinPayoutThreshold   types.JSONBigInt `json:"min_payout_threshold"`
	PoolDonation         float64          `json:"pool_donation"`
	CensoredEmail        string           `json:"censored_email"`
	CensoredIP           string           `json:"censored_ip"`
	FirstJoinedTimestamp int64            `json:"first_joined"`
}

type jsonResponseDetail struct {
	Result MinerDetail `json:"result"`
	Error  error       `json:"error"`
}

// APIEndpoint is a v1 flexpool API endpoint
var APIEndpoint = "https://flexpool.io/api/v1"

// MinerBalance returns miner's balance (in weis)
func MinerBalance(address string) (*big.Int, error) {
	resp, err := http.Get(APIEndpoint + "/miner/" + address + "/balance")
	if err != nil {
		return big.NewInt(0), err
	}
	respBytes, _ := ioutil.ReadAll(resp.Body)
	var parsed jsonResponseBalance
	json.Unmarshal(respBytes, &parsed)

	return &parsed.Result.Int, parsed.Error
}

// MinerExists returns true if miner has ever mined
func MinerExists(address string) (bool, error) {
	resp, err := http.Get(APIEndpoint + "/miner/" + address + "/exists")
	if err != nil {
		return false, err
	}
	respBytes, _ := ioutil.ReadAll(resp.Body)
	var parsed jsonResponseBool
	json.Unmarshal(respBytes, &parsed)

	return parsed.Result, parsed.Error
}

// MinerStats returns miner's stats (Reported, Effective, Valid, Stale, Invalid)
func MinerStats(address string) (DailyCurrentStats, error) {
	resp, err := http.Get(APIEndpoint + "/miner/" + address + "/stats")
	if err != nil {
		return DailyCurrentStats{}, err
	}
	respBytes, _ := ioutil.ReadAll(resp.Body)
	var parsed jsonResponseStats
	err = json.Unmarshal(respBytes, &parsed)

	return parsed.Result, parsed.Error
}

// MinerWorkerCount returns online/offline worker count
func MinerWorkerCount(address string) (WorkersCount, error) {
	resp, err := http.Get(APIEndpoint + "/miner/" + address + "/workersCount")
	if err != nil {
		return WorkersCount{}, err
	}
	respBytes, _ := ioutil.ReadAll(resp.Body)
	var parsed jsonResponseWorkerCount
	err = json.Unmarshal(respBytes, &parsed)

	return parsed.Result, parsed.Error
}

// MinerWorkers returns' workers by address
func MinerWorkers(address string) ([]Worker, error) {
	resp, err := http.Get(APIEndpoint + "/miner/" + address + "/workers")
	if err != nil {
		return nil, err
	}
	respBytes, _ := ioutil.ReadAll(resp.Body)
	var parsed jsonResponseWorkers
	err = json.Unmarshal(respBytes, &parsed)

	return parsed.Result, err
}

// MinerChart returns miner's chart data
func MinerChart(address string) ([]FullStatsWithTimestamp, error) {
	resp, err := http.Get(APIEndpoint + "/miner/" + address + "/chart")
	if err != nil {
		return nil, err
	}
	respBytes, _ := ioutil.ReadAll(resp.Body)
	var parsed jsonResponseChart
	err = json.Unmarshal(respBytes, &parsed)

	return parsed.Result, err
}

// MinerDetails returns the miner's detalis
func MinerDetails(address string) (MinerDetail, error) {
	resp, err := http.Get(APIEndpoint + "/miner/" + address + "/details")
	if err != nil {
		return MinerDetail{}, err
	}
	respBytes, _ := ioutil.ReadAll(resp.Body)
	var parsed jsonResponseDetail
	err = json.Unmarshal(respBytes, &parsed)

	return parsed.Result, err
}
