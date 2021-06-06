package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Channel is a top-level type
type Channel struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ChannelSpec `json:"spec,omitempty"`
}

// ChannelSpec defines a channel to receive notifications
type ChannelSpec struct {
	// Type is the type of the channel
	Type string `json:"type" yaml:"type"`

	Callback *ChannelCallbackConfig `json:"callback,omitempty" yaml:"callback,omitempty"`
	Dingtalk *ChannelDingtalkConfig `json:"dingtalk,omitempty" yaml:"dingtalk,omitempty"`
	Flock    *ChannelFlockConfig    `json:"flock,omitempty" yaml:"flock,omitempty"`
	Print    *ChannelPrintConfig    `json:"print,omitempty" yaml:"print,omitempty"`
	Slack    *ChannelSlackConfig    `json:"slack,omitempty" yaml:"slack,omitempty"`
	Telegram *ChannelTelegramConfig `json:"telegram,omitempty" yaml:"telegram,omitempty"`

	// Shim is a special field used in new channel development
	Shim map[string]string `json:"shim,omitempty" yaml:"shim,omitempty"`
}

// ChannelCallbackConfig is config for ChannelCallback, read from config file
type ChannelCallbackConfig struct {
	Method          string `json:"method,omitempty" yaml:"method,omitempty"`
	URL             string `json:"url" yaml:"url"`
	Proxy           string `json:"proxy,omitempty" yaml:"proxy,omitempty"`
	UseTemplate     bool   `json:"useTemplate,omitempty" yaml:"useTemplate,omitempty"`
	AddedTemplate   string `json:"addedTemplate,omitempty" yaml:"addedTemplate,omitempty"`
	DeletedTemplate string `json:"deletedTemplate,omitempty" yaml:"deletedTemplate,omitempty"`
	UpdatedTemplate string `json:"updatedTemplate,omitempty" yaml:"updatedTemplate,omitempty"`
}

// ChannelDingtalkConfig is config for ChannelDingtalk, read from config file
type ChannelDingtalkConfig struct {
	WebhookURL      string   `json:"webhookURL" yaml:"webhookURL"`
	Proxy           string   `json:"proxy,omitempty" yaml:"proxy,omitempty"`
	AtMobiles       []string `json:"atMobiles,omitempty" yaml:"atMobiles,omitempty"`
	AtAll           bool     `json:"atAll,omitempty" yaml:"atAll,omitempty"`
	AddedTemplate   string   `json:"addedTemplate,omitempty" yaml:"addedTemplate,omitempty"`
	DeletedTemplate string   `json:"deletedTemplate,omitempty" yaml:"deletedTemplate,omitempty"`
	UpdatedTemplate string   `json:"updatedTemplate,omitempty" yaml:"updatedTemplate,omitempty"`
}

// ChannelFlockConfig is config for ChannelFlock, read from config file
type ChannelFlockConfig struct {
	URL             string `json:"url" yaml:"url"`
	Proxy           string `json:"proxy,omitempty" yaml:"proxy,omitempty"`
	TitleTemplate   string `json:"titleTemplate,omitempty" yaml:"titleTemplate,omitempty"`
	AddedTemplate   string `json:"addedTemplate,omitempty" yaml:"addedTemplate,omitempty"`
	DeletedTemplate string `json:"deletedTemplate,omitempty" yaml:"deletedTemplate,omitempty"`
	UpdatedTemplate string `json:"updatedTemplate,omitempty" yaml:"updatedTemplate,omitempty"`
}

// ChannelPrintConfig is config for ChannelPrint, read from config file
type ChannelPrintConfig struct {
	Writer          string `json:"writer" yaml:"writer"`
	UseTemplate     bool   `json:"useTemplate,omitempty" yaml:"useTemplate,omitempty"`
	AddedTemplate   string `json:"addedTemplate,omitempty" yaml:"addedTemplate,omitempty"`
	DeletedTemplate string `json:"deletedTemplate,omitempty" yaml:"deletedTemplate,omitempty"`
	UpdatedTemplate string `json:"updatedTemplate,omitempty" yaml:"updatedTemplate,omitempty"`
}

// ChannelSlackConfig is config for ChannelSlack, read from config file
type ChannelSlackConfig struct {
	Token           string `json:"token" yaml:"token"`
	Proxy           string `json:"proxy,omitempty" yaml:"proxy,omitempty"`
	WebhookURL      string `json:"webhookURL" yaml:"webhookURL"`
	TitleTemplate   string `json:"titleTemplate,omitempty" yaml:"titleTemplate,omitempty"`
	AddedTemplate   string `json:"addedTemplate,omitempty" yaml:"addedTemplate,omitempty"`
	DeletedTemplate string `json:"deletedTemplate,omitempty" yaml:"deletedTemplate,omitempty"`
	UpdatedTemplate string `json:"updatedTemplate,omitempty" yaml:"updatedTemplate,omitempty"`
}

// ChannelTelegramConfig is config for ChannelTelegram, read from config file
type ChannelTelegramConfig struct {
	Token           string   `json:"token" yaml:"token"`
	ChatIDs         []string `json:"chatIDs" yaml:"chatIDs"`
	Proxy           string   `json:"proxy,omitempty" yaml:"proxy,omitempty"`
	AddedTemplate   string   `json:"addedTemplate,omitempty" yaml:"addedTemplate,omitempty"`
	DeletedTemplate string   `json:"deletedTemplate,omitempty" yaml:"deletedTemplate,omitempty"`
	UpdatedTemplate string   `json:"updatedTemplate,omitempty" yaml:"updatedTemplate,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ChannelList no client needed for list as it's been created in above
type ChannelList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `son:"metadata,omitempty"`

	Items []Channel `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Watcher is a top-level type
type Watcher struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec WatcherSpec `json:"spec,omitempty"`
}

type WatcherSpec struct {
	// Resource is the resource to watch, e.g. "deployments.v1.apps"
	Resource string `json:"resource" yaml:"resource"`

	// NoticeWhenAdded determine whether to notice when a resource is added
	NoticeWhenAdded bool `json:"noticeWhenAdded" yaml:"noticeWhenAdded"`

	// NoticeWhenDeleted determine whether to notice when a resource is deleted
	NoticeWhenDeleted bool `json:"noticeWhenDeleted" yaml:"noticeWhenDeleted"`

	// NoticeWhenUpdated determine whether to notice when a resource is updated,
	// When UpdateOn is not nil, only notice when fields in UpdateOn is changed
	NoticeWhenUpdated bool `json:"noticeWhenUpdated" yaml:"noticeWhenUpdated"`

	// UpdateOn defines fields to watch, used with NoticeWhenUpdated
	UpdateOn []string `json:"updateOn,omitempty" yaml:"updateOn,omitempty"`

	// ChannelNames defines channels to send notification
	ChannelNames []string `json:"channelNames,omitempty" yaml:"channelNames,omitempty"`

	// ResyncPeriod is the resync period in reflectors for this resource
	ResyncPeriod string `json:"resyncPeriod,omitempty" yaml:"resyncPeriod,omitempty"`

	// Workers is the number of workers
	Workers int `json:"workers,omitempty" yaml:"workers,omitempty"`

	// MaxRetries is the max retry times
	MaxRetries int `json:"maxRetries,omitempty" yaml:"maxRetries,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// WatcherList no client needed for list as it's been created in above
type WatcherList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `son:"metadata,omitempty"`

	Items []Watcher `json:"items"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterWatcher is a top-level type
type ClusterWatcher struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec WatcherSpec `json:"spec,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterWatcherList no client needed for list as it's been created in above
type ClusterWatcherList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `son:"metadata,omitempty"`

	Items []Watcher `json:"items"`
}
