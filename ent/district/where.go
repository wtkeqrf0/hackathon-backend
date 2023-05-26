// Code generated by ent, DO NOT EDIT.

package district

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/while-act/hackathon-backend/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.District {
	return predicate.District(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.District {
	return predicate.District(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.District {
	return predicate.District(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.District {
	return predicate.District(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.District {
	return predicate.District(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.District {
	return predicate.District(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.District {
	return predicate.District(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.District {
	return predicate.District(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.District {
	return predicate.District(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.District {
	return predicate.District(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.District {
	return predicate.District(sql.FieldContainsFold(FieldID, id))
}

// AvgCadastralVal applies equality check predicate on the "avg_cadastral_val" field. It's identical to AvgCadastralValEQ.
func AvgCadastralVal(v float64) predicate.District {
	return predicate.District(sql.FieldEQ(FieldAvgCadastralVal, v))
}

// AvgCadastralValEQ applies the EQ predicate on the "avg_cadastral_val" field.
func AvgCadastralValEQ(v float64) predicate.District {
	return predicate.District(sql.FieldEQ(FieldAvgCadastralVal, v))
}

// AvgCadastralValNEQ applies the NEQ predicate on the "avg_cadastral_val" field.
func AvgCadastralValNEQ(v float64) predicate.District {
	return predicate.District(sql.FieldNEQ(FieldAvgCadastralVal, v))
}

// AvgCadastralValIn applies the In predicate on the "avg_cadastral_val" field.
func AvgCadastralValIn(vs ...float64) predicate.District {
	return predicate.District(sql.FieldIn(FieldAvgCadastralVal, vs...))
}

// AvgCadastralValNotIn applies the NotIn predicate on the "avg_cadastral_val" field.
func AvgCadastralValNotIn(vs ...float64) predicate.District {
	return predicate.District(sql.FieldNotIn(FieldAvgCadastralVal, vs...))
}

// AvgCadastralValGT applies the GT predicate on the "avg_cadastral_val" field.
func AvgCadastralValGT(v float64) predicate.District {
	return predicate.District(sql.FieldGT(FieldAvgCadastralVal, v))
}

// AvgCadastralValGTE applies the GTE predicate on the "avg_cadastral_val" field.
func AvgCadastralValGTE(v float64) predicate.District {
	return predicate.District(sql.FieldGTE(FieldAvgCadastralVal, v))
}

// AvgCadastralValLT applies the LT predicate on the "avg_cadastral_val" field.
func AvgCadastralValLT(v float64) predicate.District {
	return predicate.District(sql.FieldLT(FieldAvgCadastralVal, v))
}

// AvgCadastralValLTE applies the LTE predicate on the "avg_cadastral_val" field.
func AvgCadastralValLTE(v float64) predicate.District {
	return predicate.District(sql.FieldLTE(FieldAvgCadastralVal, v))
}

// HasHistories applies the HasEdge predicate on the "histories" edge.
func HasHistories() predicate.District {
	return predicate.District(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, HistoriesTable, HistoriesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasHistoriesWith applies the HasEdge predicate on the "histories" edge with a given conditions (other predicates).
func HasHistoriesWith(preds ...predicate.History) predicate.District {
	return predicate.District(func(s *sql.Selector) {
		step := newHistoriesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.District) predicate.District {
	return predicate.District(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.District) predicate.District {
	return predicate.District(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.District) predicate.District {
	return predicate.District(func(s *sql.Selector) {
		p(s.Not())
	})
}