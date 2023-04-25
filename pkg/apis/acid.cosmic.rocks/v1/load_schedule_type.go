package v1

// Postgres CRD definition, please use CamelCase for field names.

import (
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Postgresql defines PostgreSQL Custom Resource Definition Object.
type WorkLoadSchedule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WorkLoadScheduleSpec `json:"spec"`
	Status LoadScheduleStatus   `json:"status"`
	Error  string               `json:"-"`
}

// PostgresSpec defines the specification for the PostgreSQL TPR.
type WorkLoadScheduleSpec struct {
	ScheduleParam `json:"loadSchedule"`
	WorkLoad      `json:"workLoad,omitempty"`
	*Resources    `json:"resources,omitempty"`

	TeamID string `json:"teamId"`

	Users map[string]UserFlags `json:"users,omitempty"`

	DowntimeWindows       []DowntimeWindow `json:"downtimeWindows,omitempty"`
	LogicalBackupSchedule string           `json:"logicalBackupSchedule,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PostgresqlList defines a list of PostgreSQL clusters.
type WorkLoadScheduleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []WorkLoadSchedule `json:"items"`
}

// WorkLoad describes the workload that the operator will force the workload schedule on.
type WorkLoad struct {
	Selector    *metav1.LabelSelector `json:"selector,omitempty"`
	LoadType    string                `json:"loadType,omitempty"`
	LoadOptions string                `json:"loadOptions,omitempty"`
}

// OfftimeWindow describes the time window when the load should be turned off.
type DowntimeWindow struct {
	Everyday  bool         `json:"everyday,omitempty"`
	Weekday   time.Weekday `json:"weekday,omitempty"`
	StartTime metav1.Time  `json:"startTime,omitempty"`
	EndTime   metav1.Time  `json:"endTime,omitempty"`
}

// ScheduleParam describes the parameters for the workload schedule.
type ScheduleParam struct {
	Parameters map[string]string `json:"parameters,omitempty"`
}

// LoadScheduleStatus describes the status of the workload schedule.
type LoadScheduleStatus struct {
	LoadScheduleStatus string `json:"PostgresClusterStatus"`
}
