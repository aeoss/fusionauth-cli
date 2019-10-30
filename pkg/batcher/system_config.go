package batcher

import (
	"fmt"
	"io/ioutil"

	"github.com/aeoss/fusionauth-cli/pkg/client"
	"gopkg.in/yaml.v2"
)

// SystemConfigBatcher exports or configures items on the `sytem-configuration`
// endpoint.
type SystemConfigBatcher struct {
	Batcher
}

// SystemConfigAuditLogSpec is the audit-log specific settings.
type SystemConfigAuditLogSpec struct {
	Delete struct {
		Enabled              bool `yaml:"enabled"`
		NumberOfDaysToRetain int  `yaml:"numberOfDaysToRetain"`
	} `yaml:"delete"`
}

// SystemConfigCORSConfigSpec is the
type SystemConfigCORSConfigSpec struct {
	AllowCredentials         bool     `yaml:"allowCredentials"`
	AllowedHeaders           []string `yaml:"allowedHeaders"`
	AllowedMethods           []string `yaml:"allowedMethods"`
	Enabled                  bool     `yaml:"enabled"`
	ExposedHeaders           []string `yaml:"exposedHeaders"`
	PreflightMaxAgeInSeconds int      `yaml:"preflightMaxAgeInSeconds"`
}

// SystemConfigEventLogSpec is the spec from for the Event Log configuration.
type SystemConfigEventLogSpec struct {
	NumberToRetain int `yaml:"numberToRetain"`
}

// SystemConfigLoginRecordConfigSpec is the login record config.
type SystemConfigLoginRecordConfigSpec struct {
	Delete struct {
		Enabled              bool `yaml:"enabled"`
		NumberOfDaysToRetain int  `yaml:"numberOfDaysToRetain"`
	} `yaml:"delete"`
}

// SystemConfigUIConfigSpec is the user interface config spec.
type SystemConfigUIConfigSpec struct {
	HeaderColor   string `yaml:"headerColor"`
	LogoURL       string `yaml:"logoURL"`
	MenuFontColor string `yaml:"menuFontColor"`
}

// SystemConfigSpec is the spec for the `system-configuration` endpoint
type SystemConfigSpec struct {
	AuditLogConfiguration    SystemConfigAuditLogSpec          `yaml:"auditLogConfiguration"`
	CORSConfiguration        SystemConfigCORSConfigSpec        `yaml:"corsConfiguration"`
	EventLogConfiguration    SystemConfigEventLogSpec          `yaml:"eventLogConfiguration"`
	LoginRecordConfiguration SystemConfigLoginRecordConfigSpec `yaml:"loginRecordConfiguration"`
	ReportTimezone           string                            `yaml:"reportTimezone"`
	UIConfiguration          SystemConfigUIConfigSpec          `yaml:"uiConfiguration"`
}

// SystemConfigRequest is the request body sent to the `system-configuration` endpoint.
type SystemConfigRequest struct {
	SystemConfiguration SystemConfigSpec `yaml:"systemConfiguration"`
}

// GetName returns the Kinds that are supported by this Batcher.
func (b *SystemConfigBatcher) GetName() string {
	return "SystemConfig"
}

// GetKinds returns the Kinds that are supported by this Batcher.
func (b *SystemConfigBatcher) GetKinds() []string {
	return []string{"SystemConfig"}
}

// ConvertSpec converts the provided bytes into a statically typed struct
// specific to this Batcher.
func (b *SystemConfigBatcher) ConvertSpec(data []byte) (interface{}, error) {
	typed := &SystemConfigSpec{}

	err := yaml.Unmarshal(data, typed)
	if err != nil {
		return nil, err
	}

	return typed, nil
}

// ExportAll exports all data for this Batcher into a struct.
func (b *SystemConfigBatcher) ExportAll(client *client.Client) ([]interface{}, error) {
	ret := make([]interface{}, 0)

	res, err := client.Get("/system-configuration")
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	typed, err := b.ConvertSpec(data)
	if err != nil {
		return nil, err
	}

	return append(ret, typed), nil
}

// ExportByID exports an item within this batcher by ID.
func (b *SystemConfigBatcher) ExportByID(client *client.Client, id string) (interface{}, error) {
	return nil, fmt.Errorf("not possible")
}

// ExportByName exports an item within this batcher by Name.
func (b *SystemConfigBatcher) ExportByName(client *client.Client, name string) (interface{}, error) {
	return nil, fmt.Errorf("not possible")
}

// ImportAll imports all items within this batcher by ID.
func (b *SystemConfigBatcher) ImportAll(client *client.Client, spec []interface{}) error {
	reqBody := &SystemConfigRequest{
		SystemConfiguration: spec[0].(SystemConfigSpec),
	}

	data, err := yaml.Marshal(reqBody)
	if err != nil {
		return err
	}

	_, err = client.Put("/system-configuration", data)
	if err != nil {
		return err
	}

	// more error handling
	// 400
	// 401
	// 500
	// 503

	return nil
}

// ImportByID imports an item within this batcher by ID.
func (b *SystemConfigBatcher) ImportByID(client *client.Client, id string, spec interface{}) error {
	return fmt.Errorf("not possible")
}

// ImportByName imports an item within this batcher by Name.
func (b *SystemConfigBatcher) ImportByName(client *client.Client, name string, spec interface{}) error {
	return fmt.Errorf("not possible")
}

// Save the data to a file. Used for exports.
func (b *SystemConfigBatcher) Save(spec []interface{}) error {
	return nil
}

func init() {
	registeredBatchers = append(registeredBatchers, &SystemConfigBatcher{})
}
