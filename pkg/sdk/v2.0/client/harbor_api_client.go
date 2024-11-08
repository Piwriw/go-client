// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"net/url"

	"github.com/go-openapi/runtime"
	rtclient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/piwriw/go-client/pkg/sdk/v2.0/client/artifact"
	"github.com/piwriw/go-client/pkg/sdk/v2.0/client/auditlog"
	"github.com/piwriw/go-client/pkg/sdk/v2.0/client/configure"
	"github.com/piwriw/go-client/pkg/sdk/v2.0/client/gc"
	"github.com/piwriw/go-client/pkg/sdk/v2.0/client/health"
	"github.com/piwriw/go-client/pkg/sdk/v2.0/client/icon"
	"github.com/piwriw/go-client/pkg/sdk/v2.0/client/immutable"
	"github.com/piwriw/go-client/pkg/sdk/v2.0/client/label"
	"github.com/piwriw/go-client/pkg/sdk/v2.0/client/ldap"
	"github.com/piwriw/go-client/pkg/sdk/v2.0/client/member"
	"github.com/piwriw/go-client/pkg/sdk/v2.0/client/oidc"
	"github.com/piwriw/go-client/pkg/sdk/v2.0/client/ping"
	"github.com/piwriw/go-client/pkg/sdk/v2.0/client/preheat"
	"github.com/piwriw/go-client/pkg/sdk/v2.0/client/project"
	"github.com/piwriw/go-client/pkg/sdk/v2.0/client/project_metadata"
	"github.com/piwriw/go-client/pkg/sdk/v2.0/client/quota"
	"github.com/piwriw/go-client/pkg/sdk/v2.0/client/registry"
	"github.com/piwriw/go-client/pkg/sdk/v2.0/client/replication"
	"github.com/piwriw/go-client/pkg/sdk/v2.0/client/repository"
	"github.com/piwriw/go-client/pkg/sdk/v2.0/client/retention"
	"github.com/piwriw/go-client/pkg/sdk/v2.0/client/robot"
	"github.com/piwriw/go-client/pkg/sdk/v2.0/client/robotv1"
	"github.com/piwriw/go-client/pkg/sdk/v2.0/client/scan"
	"github.com/piwriw/go-client/pkg/sdk/v2.0/client/scan_all"
	"github.com/piwriw/go-client/pkg/sdk/v2.0/client/scanner"
	"github.com/piwriw/go-client/pkg/sdk/v2.0/client/search"
	"github.com/piwriw/go-client/pkg/sdk/v2.0/client/statistic"
	"github.com/piwriw/go-client/pkg/sdk/v2.0/client/system_cve_allowlist"
	"github.com/piwriw/go-client/pkg/sdk/v2.0/client/systeminfo"
	"github.com/piwriw/go-client/pkg/sdk/v2.0/client/user"
	"github.com/piwriw/go-client/pkg/sdk/v2.0/client/usergroup"
	"github.com/piwriw/go-client/pkg/sdk/v2.0/client/webhook"
	"github.com/piwriw/go-client/pkg/sdk/v2.0/client/webhookjob"
)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/api/v2.0"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http", "https"}

type Config struct {
	// URL is the base URL of the upstream server
	URL *url.URL
	// Transport is an inner transport for the client
	Transport http.RoundTripper
	// AuthInfo is for authentication
	AuthInfo runtime.ClientAuthInfoWriter
}

// New creates a new harbor API HTTP client.
func New(c Config) *HarborAPI {
	var (
		host     = DefaultHost
		basePath = DefaultBasePath
		schemes  = DefaultSchemes
	)

	if c.URL != nil {
		host = c.URL.Host
		basePath = c.URL.Path
		schemes = []string{c.URL.Scheme}
	}

	transport := rtclient.New(host, basePath, schemes)
	if c.Transport != nil {
		transport.Transport = c.Transport
	}

	cli := new(HarborAPI)
	cli.Transport = transport
	cli.Artifact = artifact.New(transport, strfmt.Default, c.AuthInfo)
	cli.Auditlog = auditlog.New(transport, strfmt.Default, c.AuthInfo)
	cli.Configure = configure.New(transport, strfmt.Default, c.AuthInfo)
	cli.GC = gc.New(transport, strfmt.Default, c.AuthInfo)
	cli.Health = health.New(transport, strfmt.Default, c.AuthInfo)
	cli.Icon = icon.New(transport, strfmt.Default, c.AuthInfo)
	cli.Immutable = immutable.New(transport, strfmt.Default, c.AuthInfo)
	cli.Label = label.New(transport, strfmt.Default, c.AuthInfo)
	cli.Ldap = ldap.New(transport, strfmt.Default, c.AuthInfo)
	cli.Member = member.New(transport, strfmt.Default, c.AuthInfo)
	cli.OIDC = oidc.New(transport, strfmt.Default, c.AuthInfo)
	cli.Ping = ping.New(transport, strfmt.Default, c.AuthInfo)
	cli.Preheat = preheat.New(transport, strfmt.Default, c.AuthInfo)
	cli.Project = project.New(transport, strfmt.Default, c.AuthInfo)
	cli.ProjectMetadata = project_metadata.New(transport, strfmt.Default, c.AuthInfo)
	cli.Quota = quota.New(transport, strfmt.Default, c.AuthInfo)
	cli.Registry = registry.New(transport, strfmt.Default, c.AuthInfo)
	cli.Replication = replication.New(transport, strfmt.Default, c.AuthInfo)
	cli.Repository = repository.New(transport, strfmt.Default, c.AuthInfo)
	cli.Retention = retention.New(transport, strfmt.Default, c.AuthInfo)
	cli.Robot = robot.New(transport, strfmt.Default, c.AuthInfo)
	cli.Robotv1 = robotv1.New(transport, strfmt.Default, c.AuthInfo)
	cli.Scan = scan.New(transport, strfmt.Default, c.AuthInfo)
	cli.ScanAll = scan_all.New(transport, strfmt.Default, c.AuthInfo)
	cli.Scanner = scanner.New(transport, strfmt.Default, c.AuthInfo)
	cli.Search = search.New(transport, strfmt.Default, c.AuthInfo)
	cli.Statistic = statistic.New(transport, strfmt.Default, c.AuthInfo)
	cli.SystemCVEAllowlist = system_cve_allowlist.New(transport, strfmt.Default, c.AuthInfo)
	cli.Systeminfo = systeminfo.New(transport, strfmt.Default, c.AuthInfo)
	cli.User = user.New(transport, strfmt.Default, c.AuthInfo)
	cli.Usergroup = usergroup.New(transport, strfmt.Default, c.AuthInfo)
	cli.Webhook = webhook.New(transport, strfmt.Default, c.AuthInfo)
	cli.Webhookjob = webhookjob.New(transport, strfmt.Default, c.AuthInfo)
	return cli
}

// HarborAPI is a client for harbor API
type HarborAPI struct {
	Artifact           *artifact.Client
	Auditlog           *auditlog.Client
	Configure          *configure.Client
	GC                 *gc.Client
	Health             *health.Client
	Icon               *icon.Client
	Immutable          *immutable.Client
	Label              *label.Client
	Ldap               *ldap.Client
	Member             *member.Client
	OIDC               *oidc.Client
	Ping               *ping.Client
	Preheat            *preheat.Client
	Project            *project.Client
	ProjectMetadata    *project_metadata.Client
	Quota              *quota.Client
	Registry           *registry.Client
	Replication        *replication.Client
	Repository         *repository.Client
	Retention          *retention.Client
	Robot              *robot.Client
	Robotv1            *robotv1.Client
	Scan               *scan.Client
	ScanAll            *scan_all.Client
	Scanner            *scanner.Client
	Search             *search.Client
	Statistic          *statistic.Client
	SystemCVEAllowlist *system_cve_allowlist.Client
	Systeminfo         *systeminfo.Client
	User               *user.Client
	Usergroup          *usergroup.Client
	Webhook            *webhook.Client
	Webhookjob         *webhookjob.Client
	Transport          runtime.ClientTransport
}
