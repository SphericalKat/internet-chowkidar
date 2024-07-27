// Code generated by ent, DO NOT EDIT.

package blocks

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/gnulinuxindia/internet-chowkidar/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Blocks {
	return predicate.Blocks(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Blocks {
	return predicate.Blocks(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Blocks {
	return predicate.Blocks(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Blocks {
	return predicate.Blocks(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Blocks {
	return predicate.Blocks(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Blocks {
	return predicate.Blocks(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Blocks {
	return predicate.Blocks(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Blocks {
	return predicate.Blocks(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Blocks {
	return predicate.Blocks(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Blocks {
	return predicate.Blocks(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Blocks {
	return predicate.Blocks(sql.FieldEQ(FieldUpdatedAt, v))
}

// SiteID applies equality check predicate on the "site_id" field. It's identical to SiteIDEQ.
func SiteID(v int) predicate.Blocks {
	return predicate.Blocks(sql.FieldEQ(FieldSiteID, v))
}

// IspID applies equality check predicate on the "isp_id" field. It's identical to IspIDEQ.
func IspID(v int) predicate.Blocks {
	return predicate.Blocks(sql.FieldEQ(FieldIspID, v))
}

// BlockReports applies equality check predicate on the "block_reports" field. It's identical to BlockReportsEQ.
func BlockReports(v int) predicate.Blocks {
	return predicate.Blocks(sql.FieldEQ(FieldBlockReports, v))
}

// UnblockReports applies equality check predicate on the "unblock_reports" field. It's identical to UnblockReportsEQ.
func UnblockReports(v int) predicate.Blocks {
	return predicate.Blocks(sql.FieldEQ(FieldUnblockReports, v))
}

// LastReportedAt applies equality check predicate on the "last_reported_at" field. It's identical to LastReportedAtEQ.
func LastReportedAt(v time.Time) predicate.Blocks {
	return predicate.Blocks(sql.FieldEQ(FieldLastReportedAt, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Blocks {
	return predicate.Blocks(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Blocks {
	return predicate.Blocks(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Blocks {
	return predicate.Blocks(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Blocks {
	return predicate.Blocks(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Blocks {
	return predicate.Blocks(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Blocks {
	return predicate.Blocks(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Blocks {
	return predicate.Blocks(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Blocks {
	return predicate.Blocks(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Blocks {
	return predicate.Blocks(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Blocks {
	return predicate.Blocks(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Blocks {
	return predicate.Blocks(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Blocks {
	return predicate.Blocks(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Blocks {
	return predicate.Blocks(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Blocks {
	return predicate.Blocks(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Blocks {
	return predicate.Blocks(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Blocks {
	return predicate.Blocks(sql.FieldLTE(FieldUpdatedAt, v))
}

// SiteIDEQ applies the EQ predicate on the "site_id" field.
func SiteIDEQ(v int) predicate.Blocks {
	return predicate.Blocks(sql.FieldEQ(FieldSiteID, v))
}

// SiteIDNEQ applies the NEQ predicate on the "site_id" field.
func SiteIDNEQ(v int) predicate.Blocks {
	return predicate.Blocks(sql.FieldNEQ(FieldSiteID, v))
}

// SiteIDIn applies the In predicate on the "site_id" field.
func SiteIDIn(vs ...int) predicate.Blocks {
	return predicate.Blocks(sql.FieldIn(FieldSiteID, vs...))
}

// SiteIDNotIn applies the NotIn predicate on the "site_id" field.
func SiteIDNotIn(vs ...int) predicate.Blocks {
	return predicate.Blocks(sql.FieldNotIn(FieldSiteID, vs...))
}

// IspIDEQ applies the EQ predicate on the "isp_id" field.
func IspIDEQ(v int) predicate.Blocks {
	return predicate.Blocks(sql.FieldEQ(FieldIspID, v))
}

// IspIDNEQ applies the NEQ predicate on the "isp_id" field.
func IspIDNEQ(v int) predicate.Blocks {
	return predicate.Blocks(sql.FieldNEQ(FieldIspID, v))
}

// IspIDIn applies the In predicate on the "isp_id" field.
func IspIDIn(vs ...int) predicate.Blocks {
	return predicate.Blocks(sql.FieldIn(FieldIspID, vs...))
}

// IspIDNotIn applies the NotIn predicate on the "isp_id" field.
func IspIDNotIn(vs ...int) predicate.Blocks {
	return predicate.Blocks(sql.FieldNotIn(FieldIspID, vs...))
}

// BlockReportsEQ applies the EQ predicate on the "block_reports" field.
func BlockReportsEQ(v int) predicate.Blocks {
	return predicate.Blocks(sql.FieldEQ(FieldBlockReports, v))
}

// BlockReportsNEQ applies the NEQ predicate on the "block_reports" field.
func BlockReportsNEQ(v int) predicate.Blocks {
	return predicate.Blocks(sql.FieldNEQ(FieldBlockReports, v))
}

// BlockReportsIn applies the In predicate on the "block_reports" field.
func BlockReportsIn(vs ...int) predicate.Blocks {
	return predicate.Blocks(sql.FieldIn(FieldBlockReports, vs...))
}

// BlockReportsNotIn applies the NotIn predicate on the "block_reports" field.
func BlockReportsNotIn(vs ...int) predicate.Blocks {
	return predicate.Blocks(sql.FieldNotIn(FieldBlockReports, vs...))
}

// BlockReportsGT applies the GT predicate on the "block_reports" field.
func BlockReportsGT(v int) predicate.Blocks {
	return predicate.Blocks(sql.FieldGT(FieldBlockReports, v))
}

// BlockReportsGTE applies the GTE predicate on the "block_reports" field.
func BlockReportsGTE(v int) predicate.Blocks {
	return predicate.Blocks(sql.FieldGTE(FieldBlockReports, v))
}

// BlockReportsLT applies the LT predicate on the "block_reports" field.
func BlockReportsLT(v int) predicate.Blocks {
	return predicate.Blocks(sql.FieldLT(FieldBlockReports, v))
}

// BlockReportsLTE applies the LTE predicate on the "block_reports" field.
func BlockReportsLTE(v int) predicate.Blocks {
	return predicate.Blocks(sql.FieldLTE(FieldBlockReports, v))
}

// UnblockReportsEQ applies the EQ predicate on the "unblock_reports" field.
func UnblockReportsEQ(v int) predicate.Blocks {
	return predicate.Blocks(sql.FieldEQ(FieldUnblockReports, v))
}

// UnblockReportsNEQ applies the NEQ predicate on the "unblock_reports" field.
func UnblockReportsNEQ(v int) predicate.Blocks {
	return predicate.Blocks(sql.FieldNEQ(FieldUnblockReports, v))
}

// UnblockReportsIn applies the In predicate on the "unblock_reports" field.
func UnblockReportsIn(vs ...int) predicate.Blocks {
	return predicate.Blocks(sql.FieldIn(FieldUnblockReports, vs...))
}

// UnblockReportsNotIn applies the NotIn predicate on the "unblock_reports" field.
func UnblockReportsNotIn(vs ...int) predicate.Blocks {
	return predicate.Blocks(sql.FieldNotIn(FieldUnblockReports, vs...))
}

// UnblockReportsGT applies the GT predicate on the "unblock_reports" field.
func UnblockReportsGT(v int) predicate.Blocks {
	return predicate.Blocks(sql.FieldGT(FieldUnblockReports, v))
}

// UnblockReportsGTE applies the GTE predicate on the "unblock_reports" field.
func UnblockReportsGTE(v int) predicate.Blocks {
	return predicate.Blocks(sql.FieldGTE(FieldUnblockReports, v))
}

// UnblockReportsLT applies the LT predicate on the "unblock_reports" field.
func UnblockReportsLT(v int) predicate.Blocks {
	return predicate.Blocks(sql.FieldLT(FieldUnblockReports, v))
}

// UnblockReportsLTE applies the LTE predicate on the "unblock_reports" field.
func UnblockReportsLTE(v int) predicate.Blocks {
	return predicate.Blocks(sql.FieldLTE(FieldUnblockReports, v))
}

// LastReportedAtEQ applies the EQ predicate on the "last_reported_at" field.
func LastReportedAtEQ(v time.Time) predicate.Blocks {
	return predicate.Blocks(sql.FieldEQ(FieldLastReportedAt, v))
}

// LastReportedAtNEQ applies the NEQ predicate on the "last_reported_at" field.
func LastReportedAtNEQ(v time.Time) predicate.Blocks {
	return predicate.Blocks(sql.FieldNEQ(FieldLastReportedAt, v))
}

// LastReportedAtIn applies the In predicate on the "last_reported_at" field.
func LastReportedAtIn(vs ...time.Time) predicate.Blocks {
	return predicate.Blocks(sql.FieldIn(FieldLastReportedAt, vs...))
}

// LastReportedAtNotIn applies the NotIn predicate on the "last_reported_at" field.
func LastReportedAtNotIn(vs ...time.Time) predicate.Blocks {
	return predicate.Blocks(sql.FieldNotIn(FieldLastReportedAt, vs...))
}

// LastReportedAtGT applies the GT predicate on the "last_reported_at" field.
func LastReportedAtGT(v time.Time) predicate.Blocks {
	return predicate.Blocks(sql.FieldGT(FieldLastReportedAt, v))
}

// LastReportedAtGTE applies the GTE predicate on the "last_reported_at" field.
func LastReportedAtGTE(v time.Time) predicate.Blocks {
	return predicate.Blocks(sql.FieldGTE(FieldLastReportedAt, v))
}

// LastReportedAtLT applies the LT predicate on the "last_reported_at" field.
func LastReportedAtLT(v time.Time) predicate.Blocks {
	return predicate.Blocks(sql.FieldLT(FieldLastReportedAt, v))
}

// LastReportedAtLTE applies the LTE predicate on the "last_reported_at" field.
func LastReportedAtLTE(v time.Time) predicate.Blocks {
	return predicate.Blocks(sql.FieldLTE(FieldLastReportedAt, v))
}

// HasSite applies the HasEdge predicate on the "site" edge.
func HasSite() predicate.Blocks {
	return predicate.Blocks(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, SiteTable, SiteColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSiteWith applies the HasEdge predicate on the "site" edge with a given conditions (other predicates).
func HasSiteWith(preds ...predicate.Sites) predicate.Blocks {
	return predicate.Blocks(func(s *sql.Selector) {
		step := newSiteStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasIsp applies the HasEdge predicate on the "isp" edge.
func HasIsp() predicate.Blocks {
	return predicate.Blocks(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, IspTable, IspColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasIspWith applies the HasEdge predicate on the "isp" edge with a given conditions (other predicates).
func HasIspWith(preds ...predicate.Isps) predicate.Blocks {
	return predicate.Blocks(func(s *sql.Selector) {
		step := newIspStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Blocks) predicate.Blocks {
	return predicate.Blocks(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Blocks) predicate.Blocks {
	return predicate.Blocks(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Blocks) predicate.Blocks {
	return predicate.Blocks(sql.NotPredicates(p))
}
