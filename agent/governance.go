package agent

import (
	"context"
	"time"

	"github.com/MagalixCorp/magalix-agent/v3/proto"
)

type Match struct {
	Namespaces []string
	Kinds      []string
	Labels     []map[string]string
}

type Constraint struct {
	Id         string
	TemplateId string
	AccountId  string
	ClusterId  string

	Name         string
	TemplateName string
	Parameters   map[string]interface{}
	Match        Match
	Code         string
	Description  string
	HowToSolve   string

	UpdatedAt  time.Time
	CategoryId string
	Severity   string
	Controls   []string
	Standards  []string
}

type AuditResultStatus string

const (
	AuditResultStatusViolating = "Violation"
	AuditResultStatusCompliant = "Compliance"
	AuditResultStatusIgnored   = "Ignored"
)

type AuditResult struct {
	TemplateID   *string
	ConstraintID *string
	CategoryID   *string
	Severity     *string
	Controls     []string
	Standards    []string

	Description string
	HowToSolve  string

	Status AuditResultStatus
	Msg    *string

	EntityName    *string
	EntityKind    *string
	NamespaceName *string
	ParentName    *string
	ParentKind    *string
	EntitySpec    map[string]interface{}
	Trigger       string
}

func (r *AuditResult) ToPacket() *proto.PacketAuditResultItem {
	item := proto.PacketAuditResultItem{
		TemplateID:    r.TemplateID,
		ConstraintID:  r.ConstraintID,
		CategoryID:    r.CategoryID,
		Severity:      r.Severity,
		Controls:      r.Controls,
		Standards:     r.Standards,
		Description:   r.Description,
		HowToSolve:    r.HowToSolve,
		Msg:           r.Msg,
		EntityName:    r.EntityName,
		EntityKind:    r.EntityKind,
		NamespaceName: r.NamespaceName,
		ParentName:    r.ParentName,
		ParentKind:    r.ParentKind,
		EntitySpec:    r.EntitySpec,
		Trigger:       r.Trigger,
	}
	switch r.Status {
	case AuditResultStatusViolating:
		item.Status = proto.AuditResultStatusViolating
	case AuditResultStatusCompliant:
		item.Status = proto.AuditResultStatusCompliant
	case AuditResultStatusIgnored:
		item.Status = proto.AuditResultStatusIgnored
	}
	return &item
}

type AuditResultHandler func(auditResult []*AuditResult) error

type Auditor interface {
	Start(ctx context.Context) error
	Stop() error

	HandleConstraints(constraint []*Constraint) map[string]error
	HandleAuditCommand() error
	SetAuditResultHandler(handler AuditResultHandler)
}
