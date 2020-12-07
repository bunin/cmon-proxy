package api

import (
	"encoding/json"
	"time"
)

type GetStatByNameRequest struct {
	*WithOperation `json:",inline"`
	*WithClusterID `json:",inline"`

	Name            StatType `json:"name"`
	StartDateTime   StatTS   `json:"start_datetime"`
	EndDateTime     StatTS   `json:"end_datetime"`
	WithHosts       bool     `json:"with_hosts"`
	HostPort        string   `json:"host_port,omitempty"`
	HostID          uint64   `json:"hostid,omitempty"`
	CalculatePerSec bool     `json:"calculate_per_sec"`
}

type GetStatByNameResponse struct {
	*WithResponseData `json:",inline"`
	*WithTotal        `json:",inline"`

	Data json.RawMessage `json:"data"`
}

func IsValidStatType(t StatType) bool {
	for _, s := range StatTypes {
		if t == s {
			return true
		}
	}
	return false
}

type StatType string

var (
	StatTypes = []StatType{
		StatTypeNetStat,
		StatTypeMemoryStat,
		StatTypeDiskStat,
		StatTypeCpuStat,
		StatTypeSqlStat,
		StatTypeDbStat,
		StatTypeTcpStat,
		StatTypeNdbStat,
		StatTypeProxysqlStat,
		StatTypeMongoStat,
		StatTypeHaproxyStat,
	}
)

// String implements Stringer interface
func (st StatType) String() string {
	return string(st)
}

const (
	StatTypeNetStat      StatType = "netstat"      // Network statistics (doc: CmonNetworkStats properties)
	StatTypeMemoryStat   StatType = "memorystat"   // Memory statistics (RAM usage) (doc: CmonMemoryStats properties)
	StatTypeDiskStat     StatType = "diskstat"     // Disk usage statistics (doc: CmonDiskStat properties)
	StatTypeCpuStat      StatType = "cpustat"      // CPU statistics (doc: cmoncpustats )
	StatTypeSqlStat      StatType = "sqlstat"      // SQL server (global) statistics (doc: CmonSqlStats properties)
	StatTypeDbStat       StatType = "dbstat"       // Database statistics (doc: -)
	StatTypeTcpStat      StatType = "tcpStat"      // TCP network statistics (doc: CmonTcpStats properties)
	StatTypeNdbStat      StatType = "ndbstat"      // NDB node statistics (doc: CmonNdbStats properties)
	StatTypeProxysqlStat StatType = "proxysqlstat" // ProxySQL node statistics (doc: CmonProxySqlStats properties)
	StatTypeMongoStat    StatType = "mongoStat"    // MongoDB node statistics (doc: CmonMongoStats properties)
	StatTypeHaproxyStat  StatType = "haproxystat"  // HAProxy statistics (doc: CmonHaProxyStats properties)
)

// StatTS is a wrapper to return custom date format for cmon.
type StatTS time.Time

const (
	CmonTimeRFC3339Nano = "2006-01-02T15:04:05.000Z07:00"
)

// MarshalJSON implements json.Marshaler and returns custom date format.
func (st StatTS) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(st).Format(CmonTimeRFC3339Nano))
}
