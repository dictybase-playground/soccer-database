// IMPORTANT! This is auto generated code by https://github.com/src-d/go-kallax
// Please, do not touch the code below, and if you do, do it under your own
// risk. Take into account that all the code you write here will be completely
// erased from earth the next time you generate the kallax models.
package main

import (
	"database/sql"
	"fmt"

	"gopkg.in/src-d/go-kallax.v1"
	"gopkg.in/src-d/go-kallax.v1/types"
)

var _ types.SQLType
var _ fmt.Formatter

// NewClubs returns a new instance of Clubs.
func NewClubs(clubname string) (record *Clubs, err error) {
	return newClubs(clubname)
}

// GetID returns the primary key of the model.
func (r *Clubs) GetID() kallax.Identifier {
	return (*kallax.NumericID)(&r.CLUBID)
}

// ColumnAddress returns the pointer to the value of the given column.
func (r *Clubs) ColumnAddress(col string) (interface{}, error) {
	switch col {
	case "clubid":
		return (*kallax.NumericID)(&r.CLUBID), nil
	case "clubname":
		return &r.Clubname, nil

	default:
		return nil, fmt.Errorf("kallax: invalid column in Clubs: %s", col)
	}
}

// Value returns the value of the given column.
func (r *Clubs) Value(col string) (interface{}, error) {
	switch col {
	case "clubid":
		return r.CLUBID, nil
	case "clubname":
		return r.Clubname, nil

	default:
		return nil, fmt.Errorf("kallax: invalid column in Clubs: %s", col)
	}
}

// NewRelationshipRecord returns a new record for the relatiobship in the given
// field.
func (r *Clubs) NewRelationshipRecord(field string) (kallax.Record, error) {
	return nil, fmt.Errorf("kallax: model Clubs has no relationships")
}

// SetRelationship sets the given relationship in the given field.
func (r *Clubs) SetRelationship(field string, rel interface{}) error {
	return fmt.Errorf("kallax: model Clubs has no relationships")
}

// ClubsStore is the entity to access the records of the type Clubs
// in the database.
type ClubsStore struct {
	*kallax.Store
}

// NewClubsStore creates a new instance of ClubsStore
// using a SQL database.
func NewClubsStore(db *sql.DB) *ClubsStore {
	return &ClubsStore{kallax.NewStore(db)}
}

// GenericStore returns the generic store of this store.
func (s *ClubsStore) GenericStore() *kallax.Store {
	return s.Store
}

// SetGenericStore changes the generic store of this store.
func (s *ClubsStore) SetGenericStore(store *kallax.Store) {
	s.Store = store
}

// Insert inserts a Clubs in the database. A non-persisted object is
// required for this operation.
func (s *ClubsStore) Insert(record *Clubs) error {

	return s.Store.Insert(Schema.Clubs.BaseSchema, record)

}

// Update updates the given record on the database. If the columns are given,
// only these columns will be updated. Otherwise all of them will be.
// Be very careful with this, as you will have a potentially different object
// in memory but not on the database.
// Only writable records can be updated. Writable objects are those that have
// been just inserted or retrieved using a query with no custom select fields.
func (s *ClubsStore) Update(record *Clubs, cols ...kallax.SchemaField) (updated int64, err error) {

	return s.Store.Update(Schema.Clubs.BaseSchema, record, cols...)

}

// Save inserts the object if the record is not persisted, otherwise it updates
// it. Same rules of Update and Insert apply depending on the case.
func (s *ClubsStore) Save(record *Clubs) (updated bool, err error) {
	if !record.IsPersisted() {
		return false, s.Insert(record)
	}

	rowsUpdated, err := s.Update(record)
	if err != nil {
		return false, err
	}

	return rowsUpdated > 0, nil
}

// Delete removes the given record from the database.
func (s *ClubsStore) Delete(record *Clubs) error {

	return s.Store.Delete(Schema.Clubs.BaseSchema, record)

}

// Find returns the set of results for the given query.
func (s *ClubsStore) Find(q *ClubsQuery) (*ClubsResultSet, error) {
	rs, err := s.Store.Find(q)
	if err != nil {
		return nil, err
	}

	return NewClubsResultSet(rs), nil
}

// MustFind returns the set of results for the given query, but panics if there
// is any error.
func (s *ClubsStore) MustFind(q *ClubsQuery) *ClubsResultSet {
	return NewClubsResultSet(s.Store.MustFind(q))
}

// Count returns the number of rows that would be retrieved with the given
// query.
func (s *ClubsStore) Count(q *ClubsQuery) (int64, error) {
	return s.Store.Count(q)
}

// MustCount returns the number of rows that would be retrieved with the given
// query, but panics if there is an error.
func (s *ClubsStore) MustCount(q *ClubsQuery) int64 {
	return s.Store.MustCount(q)
}

// FindOne returns the first row returned by the given query.
// `ErrNotFound` is returned if there are no results.
func (s *ClubsStore) FindOne(q *ClubsQuery) (*Clubs, error) {
	q.Limit(1)
	q.Offset(0)
	rs, err := s.Find(q)
	if err != nil {
		return nil, err
	}

	if !rs.Next() {
		return nil, kallax.ErrNotFound
	}

	record, err := rs.Get()
	if err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}

	return record, nil
}

// MustFindOne returns the first row retrieved by the given query. It panics
// if there is an error or if there are no rows.
func (s *ClubsStore) MustFindOne(q *ClubsQuery) *Clubs {
	record, err := s.FindOne(q)
	if err != nil {
		panic(err)
	}
	return record
}

// Reload refreshes the Clubs with the data in the database and
// makes it writable.
func (s *ClubsStore) Reload(record *Clubs) error {
	return s.Store.Reload(Schema.Clubs.BaseSchema, record)
}

// Transaction executes the given callback in a transaction and rollbacks if
// an error is returned.
// The transaction is only open in the store passed as a parameter to the
// callback.
func (s *ClubsStore) Transaction(callback func(*ClubsStore) error) error {
	if callback == nil {
		return kallax.ErrInvalidTxCallback
	}

	return s.Store.Transaction(func(store *kallax.Store) error {
		return callback(&ClubsStore{store})
	})
}

// ClubsQuery is the object used to create queries for the Clubs
// entity.
type ClubsQuery struct {
	*kallax.BaseQuery
}

// NewClubsQuery returns a new instance of ClubsQuery.
func NewClubsQuery() *ClubsQuery {
	return &ClubsQuery{
		BaseQuery: kallax.NewBaseQuery(Schema.Clubs.BaseSchema),
	}
}

// Select adds columns to select in the query.
func (q *ClubsQuery) Select(columns ...kallax.SchemaField) *ClubsQuery {
	if len(columns) == 0 {
		return q
	}
	q.BaseQuery.Select(columns...)
	return q
}

// SelectNot excludes columns from being selected in the query.
func (q *ClubsQuery) SelectNot(columns ...kallax.SchemaField) *ClubsQuery {
	q.BaseQuery.SelectNot(columns...)
	return q
}

// Copy returns a new identical copy of the query. Remember queries are mutable
// so make a copy any time you need to reuse them.
func (q *ClubsQuery) Copy() *ClubsQuery {
	return &ClubsQuery{
		BaseQuery: q.BaseQuery.Copy(),
	}
}

// Order adds order clauses to the query for the given columns.
func (q *ClubsQuery) Order(cols ...kallax.ColumnOrder) *ClubsQuery {
	q.BaseQuery.Order(cols...)
	return q
}

// BatchSize sets the number of items to fetch per batch when there are 1:N
// relationships selected in the query.
func (q *ClubsQuery) BatchSize(size uint64) *ClubsQuery {
	q.BaseQuery.BatchSize(size)
	return q
}

// Limit sets the max number of items to retrieve.
func (q *ClubsQuery) Limit(n uint64) *ClubsQuery {
	q.BaseQuery.Limit(n)
	return q
}

// Offset sets the number of items to skip from the result set of items.
func (q *ClubsQuery) Offset(n uint64) *ClubsQuery {
	q.BaseQuery.Offset(n)
	return q
}

// Where adds a condition to the query. All conditions added are concatenated
// using a logical AND.
func (q *ClubsQuery) Where(cond kallax.Condition) *ClubsQuery {
	q.BaseQuery.Where(cond)
	return q
}

// FindByCLUBID adds a new filter to the query that will require that
// the CLUBID property is equal to one of the passed values; if no passed values,
// it will do nothing.
func (q *ClubsQuery) FindByCLUBID(v ...int64) *ClubsQuery {
	if len(v) == 0 {
		return q
	}
	values := make([]interface{}, len(v))
	for i, val := range v {
		values[i] = val
	}
	return q.Where(kallax.In(Schema.Clubs.CLUBID, values...))
}

// FindByClubname adds a new filter to the query that will require that
// the Clubname property is equal to the passed value.
func (q *ClubsQuery) FindByClubname(v string) *ClubsQuery {
	return q.Where(kallax.Eq(Schema.Clubs.Clubname, v))
}

// ClubsResultSet is the set of results returned by a query to the
// database.
type ClubsResultSet struct {
	ResultSet kallax.ResultSet
	last      *Clubs
	lastErr   error
}

// NewClubsResultSet creates a new result set for rows of the type
// Clubs.
func NewClubsResultSet(rs kallax.ResultSet) *ClubsResultSet {
	return &ClubsResultSet{ResultSet: rs}
}

// Next fetches the next item in the result set and returns true if there is
// a next item.
// The result set is closed automatically when there are no more items.
func (rs *ClubsResultSet) Next() bool {
	if !rs.ResultSet.Next() {
		rs.lastErr = rs.ResultSet.Close()
		rs.last = nil
		return false
	}

	var record kallax.Record
	record, rs.lastErr = rs.ResultSet.Get(Schema.Clubs.BaseSchema)
	if rs.lastErr != nil {
		rs.last = nil
	} else {
		var ok bool
		rs.last, ok = record.(*Clubs)
		if !ok {
			rs.lastErr = fmt.Errorf("kallax: unable to convert record to *Clubs")
			rs.last = nil
		}
	}

	return true
}

// Get retrieves the last fetched item from the result set and the last error.
func (rs *ClubsResultSet) Get() (*Clubs, error) {
	return rs.last, rs.lastErr
}

// ForEach iterates over the complete result set passing every record found to
// the given callback. It is possible to stop the iteration by returning
// `kallax.ErrStop` in the callback.
// Result set is always closed at the end.
func (rs *ClubsResultSet) ForEach(fn func(*Clubs) error) error {
	for rs.Next() {
		record, err := rs.Get()
		if err != nil {
			return err
		}

		if err := fn(record); err != nil {
			if err == kallax.ErrStop {
				return rs.Close()
			}

			return err
		}
	}
	return nil
}

// All returns all records on the result set and closes the result set.
func (rs *ClubsResultSet) All() ([]*Clubs, error) {
	var result []*Clubs
	for rs.Next() {
		record, err := rs.Get()
		if err != nil {
			return nil, err
		}
		result = append(result, record)
	}
	return result, nil
}

// One returns the first record on the result set and closes the result set.
func (rs *ClubsResultSet) One() (*Clubs, error) {
	if !rs.Next() {
		return nil, kallax.ErrNotFound
	}

	record, err := rs.Get()
	if err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}

	return record, nil
}

// Err returns the last error occurred.
func (rs *ClubsResultSet) Err() error {
	return rs.lastErr
}

// Close closes the result set.
func (rs *ClubsResultSet) Close() error {
	return rs.ResultSet.Close()
}

// NewCountries returns a new instance of Countries.
func NewCountries(countryname string) (record *Countries, err error) {
	return newCountries(countryname)
}

// GetID returns the primary key of the model.
func (r *Countries) GetID() kallax.Identifier {
	return (*kallax.NumericID)(&r.COUNTRYID)
}

// ColumnAddress returns the pointer to the value of the given column.
func (r *Countries) ColumnAddress(col string) (interface{}, error) {
	switch col {
	case "countryid":
		return (*kallax.NumericID)(&r.COUNTRYID), nil
	case "countryname":
		return &r.Countryname, nil

	default:
		return nil, fmt.Errorf("kallax: invalid column in Countries: %s", col)
	}
}

// Value returns the value of the given column.
func (r *Countries) Value(col string) (interface{}, error) {
	switch col {
	case "countryid":
		return r.COUNTRYID, nil
	case "countryname":
		return r.Countryname, nil

	default:
		return nil, fmt.Errorf("kallax: invalid column in Countries: %s", col)
	}
}

// NewRelationshipRecord returns a new record for the relatiobship in the given
// field.
func (r *Countries) NewRelationshipRecord(field string) (kallax.Record, error) {
	return nil, fmt.Errorf("kallax: model Countries has no relationships")
}

// SetRelationship sets the given relationship in the given field.
func (r *Countries) SetRelationship(field string, rel interface{}) error {
	return fmt.Errorf("kallax: model Countries has no relationships")
}

// CountriesStore is the entity to access the records of the type Countries
// in the database.
type CountriesStore struct {
	*kallax.Store
}

// NewCountriesStore creates a new instance of CountriesStore
// using a SQL database.
func NewCountriesStore(db *sql.DB) *CountriesStore {
	return &CountriesStore{kallax.NewStore(db)}
}

// GenericStore returns the generic store of this store.
func (s *CountriesStore) GenericStore() *kallax.Store {
	return s.Store
}

// SetGenericStore changes the generic store of this store.
func (s *CountriesStore) SetGenericStore(store *kallax.Store) {
	s.Store = store
}

// Insert inserts a Countries in the database. A non-persisted object is
// required for this operation.
func (s *CountriesStore) Insert(record *Countries) error {

	return s.Store.Insert(Schema.Countries.BaseSchema, record)

}

// Update updates the given record on the database. If the columns are given,
// only these columns will be updated. Otherwise all of them will be.
// Be very careful with this, as you will have a potentially different object
// in memory but not on the database.
// Only writable records can be updated. Writable objects are those that have
// been just inserted or retrieved using a query with no custom select fields.
func (s *CountriesStore) Update(record *Countries, cols ...kallax.SchemaField) (updated int64, err error) {

	return s.Store.Update(Schema.Countries.BaseSchema, record, cols...)

}

// Save inserts the object if the record is not persisted, otherwise it updates
// it. Same rules of Update and Insert apply depending on the case.
func (s *CountriesStore) Save(record *Countries) (updated bool, err error) {
	if !record.IsPersisted() {
		return false, s.Insert(record)
	}

	rowsUpdated, err := s.Update(record)
	if err != nil {
		return false, err
	}

	return rowsUpdated > 0, nil
}

// Delete removes the given record from the database.
func (s *CountriesStore) Delete(record *Countries) error {

	return s.Store.Delete(Schema.Countries.BaseSchema, record)

}

// Find returns the set of results for the given query.
func (s *CountriesStore) Find(q *CountriesQuery) (*CountriesResultSet, error) {
	rs, err := s.Store.Find(q)
	if err != nil {
		return nil, err
	}

	return NewCountriesResultSet(rs), nil
}

// MustFind returns the set of results for the given query, but panics if there
// is any error.
func (s *CountriesStore) MustFind(q *CountriesQuery) *CountriesResultSet {
	return NewCountriesResultSet(s.Store.MustFind(q))
}

// Count returns the number of rows that would be retrieved with the given
// query.
func (s *CountriesStore) Count(q *CountriesQuery) (int64, error) {
	return s.Store.Count(q)
}

// MustCount returns the number of rows that would be retrieved with the given
// query, but panics if there is an error.
func (s *CountriesStore) MustCount(q *CountriesQuery) int64 {
	return s.Store.MustCount(q)
}

// FindOne returns the first row returned by the given query.
// `ErrNotFound` is returned if there are no results.
func (s *CountriesStore) FindOne(q *CountriesQuery) (*Countries, error) {
	q.Limit(1)
	q.Offset(0)
	rs, err := s.Find(q)
	if err != nil {
		return nil, err
	}

	if !rs.Next() {
		return nil, kallax.ErrNotFound
	}

	record, err := rs.Get()
	if err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}

	return record, nil
}

// MustFindOne returns the first row retrieved by the given query. It panics
// if there is an error or if there are no rows.
func (s *CountriesStore) MustFindOne(q *CountriesQuery) *Countries {
	record, err := s.FindOne(q)
	if err != nil {
		panic(err)
	}
	return record
}

// Reload refreshes the Countries with the data in the database and
// makes it writable.
func (s *CountriesStore) Reload(record *Countries) error {
	return s.Store.Reload(Schema.Countries.BaseSchema, record)
}

// Transaction executes the given callback in a transaction and rollbacks if
// an error is returned.
// The transaction is only open in the store passed as a parameter to the
// callback.
func (s *CountriesStore) Transaction(callback func(*CountriesStore) error) error {
	if callback == nil {
		return kallax.ErrInvalidTxCallback
	}

	return s.Store.Transaction(func(store *kallax.Store) error {
		return callback(&CountriesStore{store})
	})
}

// CountriesQuery is the object used to create queries for the Countries
// entity.
type CountriesQuery struct {
	*kallax.BaseQuery
}

// NewCountriesQuery returns a new instance of CountriesQuery.
func NewCountriesQuery() *CountriesQuery {
	return &CountriesQuery{
		BaseQuery: kallax.NewBaseQuery(Schema.Countries.BaseSchema),
	}
}

// Select adds columns to select in the query.
func (q *CountriesQuery) Select(columns ...kallax.SchemaField) *CountriesQuery {
	if len(columns) == 0 {
		return q
	}
	q.BaseQuery.Select(columns...)
	return q
}

// SelectNot excludes columns from being selected in the query.
func (q *CountriesQuery) SelectNot(columns ...kallax.SchemaField) *CountriesQuery {
	q.BaseQuery.SelectNot(columns...)
	return q
}

// Copy returns a new identical copy of the query. Remember queries are mutable
// so make a copy any time you need to reuse them.
func (q *CountriesQuery) Copy() *CountriesQuery {
	return &CountriesQuery{
		BaseQuery: q.BaseQuery.Copy(),
	}
}

// Order adds order clauses to the query for the given columns.
func (q *CountriesQuery) Order(cols ...kallax.ColumnOrder) *CountriesQuery {
	q.BaseQuery.Order(cols...)
	return q
}

// BatchSize sets the number of items to fetch per batch when there are 1:N
// relationships selected in the query.
func (q *CountriesQuery) BatchSize(size uint64) *CountriesQuery {
	q.BaseQuery.BatchSize(size)
	return q
}

// Limit sets the max number of items to retrieve.
func (q *CountriesQuery) Limit(n uint64) *CountriesQuery {
	q.BaseQuery.Limit(n)
	return q
}

// Offset sets the number of items to skip from the result set of items.
func (q *CountriesQuery) Offset(n uint64) *CountriesQuery {
	q.BaseQuery.Offset(n)
	return q
}

// Where adds a condition to the query. All conditions added are concatenated
// using a logical AND.
func (q *CountriesQuery) Where(cond kallax.Condition) *CountriesQuery {
	q.BaseQuery.Where(cond)
	return q
}

// FindByCOUNTRYID adds a new filter to the query that will require that
// the COUNTRYID property is equal to one of the passed values; if no passed values,
// it will do nothing.
func (q *CountriesQuery) FindByCOUNTRYID(v ...int64) *CountriesQuery {
	if len(v) == 0 {
		return q
	}
	values := make([]interface{}, len(v))
	for i, val := range v {
		values[i] = val
	}
	return q.Where(kallax.In(Schema.Countries.COUNTRYID, values...))
}

// FindByCountryname adds a new filter to the query that will require that
// the Countryname property is equal to the passed value.
func (q *CountriesQuery) FindByCountryname(v string) *CountriesQuery {
	return q.Where(kallax.Eq(Schema.Countries.Countryname, v))
}

// CountriesResultSet is the set of results returned by a query to the
// database.
type CountriesResultSet struct {
	ResultSet kallax.ResultSet
	last      *Countries
	lastErr   error
}

// NewCountriesResultSet creates a new result set for rows of the type
// Countries.
func NewCountriesResultSet(rs kallax.ResultSet) *CountriesResultSet {
	return &CountriesResultSet{ResultSet: rs}
}

// Next fetches the next item in the result set and returns true if there is
// a next item.
// The result set is closed automatically when there are no more items.
func (rs *CountriesResultSet) Next() bool {
	if !rs.ResultSet.Next() {
		rs.lastErr = rs.ResultSet.Close()
		rs.last = nil
		return false
	}

	var record kallax.Record
	record, rs.lastErr = rs.ResultSet.Get(Schema.Countries.BaseSchema)
	if rs.lastErr != nil {
		rs.last = nil
	} else {
		var ok bool
		rs.last, ok = record.(*Countries)
		if !ok {
			rs.lastErr = fmt.Errorf("kallax: unable to convert record to *Countries")
			rs.last = nil
		}
	}

	return true
}

// Get retrieves the last fetched item from the result set and the last error.
func (rs *CountriesResultSet) Get() (*Countries, error) {
	return rs.last, rs.lastErr
}

// ForEach iterates over the complete result set passing every record found to
// the given callback. It is possible to stop the iteration by returning
// `kallax.ErrStop` in the callback.
// Result set is always closed at the end.
func (rs *CountriesResultSet) ForEach(fn func(*Countries) error) error {
	for rs.Next() {
		record, err := rs.Get()
		if err != nil {
			return err
		}

		if err := fn(record); err != nil {
			if err == kallax.ErrStop {
				return rs.Close()
			}

			return err
		}
	}
	return nil
}

// All returns all records on the result set and closes the result set.
func (rs *CountriesResultSet) All() ([]*Countries, error) {
	var result []*Countries
	for rs.Next() {
		record, err := rs.Get()
		if err != nil {
			return nil, err
		}
		result = append(result, record)
	}
	return result, nil
}

// One returns the first record on the result set and closes the result set.
func (rs *CountriesResultSet) One() (*Countries, error) {
	if !rs.Next() {
		return nil, kallax.ErrNotFound
	}

	record, err := rs.Get()
	if err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}

	return record, nil
}

// Err returns the last error occurred.
func (rs *CountriesResultSet) Err() error {
	return rs.lastErr
}

// Close closes the result set.
func (rs *CountriesResultSet) Close() error {
	return rs.ResultSet.Close()
}

// NewPlayers returns a new instance of Players.
func NewPlayers(playername string, clubname string, countryname string) (record *Players, err error) {
	return newPlayers(playername, clubname, countryname)
}

// GetID returns the primary key of the model.
func (r *Players) GetID() kallax.Identifier {
	return (*kallax.NumericID)(&r.PID)
}

// ColumnAddress returns the pointer to the value of the given column.
func (r *Players) ColumnAddress(col string) (interface{}, error) {
	switch col {
	case "pid":
		return (*kallax.NumericID)(&r.PID), nil
	case "pname":
		return &r.Pname, nil
	case "clubid":
		return &r.Clubid, nil
	case "countryid":
		return &r.Countryid, nil

	default:
		return nil, fmt.Errorf("kallax: invalid column in Players: %s", col)
	}
}

// Value returns the value of the given column.
func (r *Players) Value(col string) (interface{}, error) {
	switch col {
	case "pid":
		return r.PID, nil
	case "pname":
		return r.Pname, nil
	case "clubid":
		return r.Clubid, nil
	case "countryid":
		return r.Countryid, nil

	default:
		return nil, fmt.Errorf("kallax: invalid column in Players: %s", col)
	}
}

// NewRelationshipRecord returns a new record for the relatiobship in the given
// field.
func (r *Players) NewRelationshipRecord(field string) (kallax.Record, error) {
	return nil, fmt.Errorf("kallax: model Players has no relationships")
}

// SetRelationship sets the given relationship in the given field.
func (r *Players) SetRelationship(field string, rel interface{}) error {
	return fmt.Errorf("kallax: model Players has no relationships")
}

// PlayersStore is the entity to access the records of the type Players
// in the database.
type PlayersStore struct {
	*kallax.Store
}

// NewPlayersStore creates a new instance of PlayersStore
// using a SQL database.
func NewPlayersStore(db *sql.DB) *PlayersStore {
	return &PlayersStore{kallax.NewStore(db)}
}

// GenericStore returns the generic store of this store.
func (s *PlayersStore) GenericStore() *kallax.Store {
	return s.Store
}

// SetGenericStore changes the generic store of this store.
func (s *PlayersStore) SetGenericStore(store *kallax.Store) {
	s.Store = store
}

// Insert inserts a Players in the database. A non-persisted object is
// required for this operation.
func (s *PlayersStore) Insert(record *Players) error {

	return s.Store.Insert(Schema.Players.BaseSchema, record)

}

// Update updates the given record on the database. If the columns are given,
// only these columns will be updated. Otherwise all of them will be.
// Be very careful with this, as you will have a potentially different object
// in memory but not on the database.
// Only writable records can be updated. Writable objects are those that have
// been just inserted or retrieved using a query with no custom select fields.
func (s *PlayersStore) Update(record *Players, cols ...kallax.SchemaField) (updated int64, err error) {

	return s.Store.Update(Schema.Players.BaseSchema, record, cols...)

}

// Save inserts the object if the record is not persisted, otherwise it updates
// it. Same rules of Update and Insert apply depending on the case.
func (s *PlayersStore) Save(record *Players) (updated bool, err error) {
	if !record.IsPersisted() {
		return false, s.Insert(record)
	}

	rowsUpdated, err := s.Update(record)
	if err != nil {
		return false, err
	}

	return rowsUpdated > 0, nil
}

// Delete removes the given record from the database.
func (s *PlayersStore) Delete(record *Players) error {

	return s.Store.Delete(Schema.Players.BaseSchema, record)

}

// Find returns the set of results for the given query.
func (s *PlayersStore) Find(q *PlayersQuery) (*PlayersResultSet, error) {
	rs, err := s.Store.Find(q)
	if err != nil {
		return nil, err
	}

	return NewPlayersResultSet(rs), nil
}

// MustFind returns the set of results for the given query, but panics if there
// is any error.
func (s *PlayersStore) MustFind(q *PlayersQuery) *PlayersResultSet {
	return NewPlayersResultSet(s.Store.MustFind(q))
}

// Count returns the number of rows that would be retrieved with the given
// query.
func (s *PlayersStore) Count(q *PlayersQuery) (int64, error) {
	return s.Store.Count(q)
}

// MustCount returns the number of rows that would be retrieved with the given
// query, but panics if there is an error.
func (s *PlayersStore) MustCount(q *PlayersQuery) int64 {
	return s.Store.MustCount(q)
}

// FindOne returns the first row returned by the given query.
// `ErrNotFound` is returned if there are no results.
func (s *PlayersStore) FindOne(q *PlayersQuery) (*Players, error) {
	q.Limit(1)
	q.Offset(0)
	rs, err := s.Find(q)
	if err != nil {
		return nil, err
	}

	if !rs.Next() {
		return nil, kallax.ErrNotFound
	}

	record, err := rs.Get()
	if err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}

	return record, nil
}

// MustFindOne returns the first row retrieved by the given query. It panics
// if there is an error or if there are no rows.
func (s *PlayersStore) MustFindOne(q *PlayersQuery) *Players {
	record, err := s.FindOne(q)
	if err != nil {
		panic(err)
	}
	return record
}

// Reload refreshes the Players with the data in the database and
// makes it writable.
func (s *PlayersStore) Reload(record *Players) error {
	return s.Store.Reload(Schema.Players.BaseSchema, record)
}

// Transaction executes the given callback in a transaction and rollbacks if
// an error is returned.
// The transaction is only open in the store passed as a parameter to the
// callback.
func (s *PlayersStore) Transaction(callback func(*PlayersStore) error) error {
	if callback == nil {
		return kallax.ErrInvalidTxCallback
	}

	return s.Store.Transaction(func(store *kallax.Store) error {
		return callback(&PlayersStore{store})
	})
}

// PlayersQuery is the object used to create queries for the Players
// entity.
type PlayersQuery struct {
	*kallax.BaseQuery
}

// NewPlayersQuery returns a new instance of PlayersQuery.
func NewPlayersQuery() *PlayersQuery {
	return &PlayersQuery{
		BaseQuery: kallax.NewBaseQuery(Schema.Players.BaseSchema),
	}
}

// Select adds columns to select in the query.
func (q *PlayersQuery) Select(columns ...kallax.SchemaField) *PlayersQuery {
	if len(columns) == 0 {
		return q
	}
	q.BaseQuery.Select(columns...)
	return q
}

// SelectNot excludes columns from being selected in the query.
func (q *PlayersQuery) SelectNot(columns ...kallax.SchemaField) *PlayersQuery {
	q.BaseQuery.SelectNot(columns...)
	return q
}

// Copy returns a new identical copy of the query. Remember queries are mutable
// so make a copy any time you need to reuse them.
func (q *PlayersQuery) Copy() *PlayersQuery {
	return &PlayersQuery{
		BaseQuery: q.BaseQuery.Copy(),
	}
}

// Order adds order clauses to the query for the given columns.
func (q *PlayersQuery) Order(cols ...kallax.ColumnOrder) *PlayersQuery {
	q.BaseQuery.Order(cols...)
	return q
}

// BatchSize sets the number of items to fetch per batch when there are 1:N
// relationships selected in the query.
func (q *PlayersQuery) BatchSize(size uint64) *PlayersQuery {
	q.BaseQuery.BatchSize(size)
	return q
}

// Limit sets the max number of items to retrieve.
func (q *PlayersQuery) Limit(n uint64) *PlayersQuery {
	q.BaseQuery.Limit(n)
	return q
}

// Offset sets the number of items to skip from the result set of items.
func (q *PlayersQuery) Offset(n uint64) *PlayersQuery {
	q.BaseQuery.Offset(n)
	return q
}

// Where adds a condition to the query. All conditions added are concatenated
// using a logical AND.
func (q *PlayersQuery) Where(cond kallax.Condition) *PlayersQuery {
	q.BaseQuery.Where(cond)
	return q
}

// FindByPID adds a new filter to the query that will require that
// the PID property is equal to one of the passed values; if no passed values,
// it will do nothing.
func (q *PlayersQuery) FindByPID(v ...int64) *PlayersQuery {
	if len(v) == 0 {
		return q
	}
	values := make([]interface{}, len(v))
	for i, val := range v {
		values[i] = val
	}
	return q.Where(kallax.In(Schema.Players.PID, values...))
}

// FindByPname adds a new filter to the query that will require that
// the Pname property is equal to the passed value.
func (q *PlayersQuery) FindByPname(v string) *PlayersQuery {
	return q.Where(kallax.Eq(Schema.Players.Pname, v))
}

// FindByClubid adds a new filter to the query that will require that
// the Clubid property is equal to the passed value.
func (q *PlayersQuery) FindByClubid(cond kallax.ScalarCond, v int64) *PlayersQuery {
	return q.Where(cond(Schema.Players.Clubid, v))
}

// FindByCountryid adds a new filter to the query that will require that
// the Countryid property is equal to the passed value.
func (q *PlayersQuery) FindByCountryid(cond kallax.ScalarCond, v int64) *PlayersQuery {
	return q.Where(cond(Schema.Players.Countryid, v))
}

// PlayersResultSet is the set of results returned by a query to the
// database.
type PlayersResultSet struct {
	ResultSet kallax.ResultSet
	last      *Players
	lastErr   error
}

// NewPlayersResultSet creates a new result set for rows of the type
// Players.
func NewPlayersResultSet(rs kallax.ResultSet) *PlayersResultSet {
	return &PlayersResultSet{ResultSet: rs}
}

// Next fetches the next item in the result set and returns true if there is
// a next item.
// The result set is closed automatically when there are no more items.
func (rs *PlayersResultSet) Next() bool {
	if !rs.ResultSet.Next() {
		rs.lastErr = rs.ResultSet.Close()
		rs.last = nil
		return false
	}

	var record kallax.Record
	record, rs.lastErr = rs.ResultSet.Get(Schema.Players.BaseSchema)
	if rs.lastErr != nil {
		rs.last = nil
	} else {
		var ok bool
		rs.last, ok = record.(*Players)
		if !ok {
			rs.lastErr = fmt.Errorf("kallax: unable to convert record to *Players")
			rs.last = nil
		}
	}

	return true
}

// Get retrieves the last fetched item from the result set and the last error.
func (rs *PlayersResultSet) Get() (*Players, error) {
	return rs.last, rs.lastErr
}

// ForEach iterates over the complete result set passing every record found to
// the given callback. It is possible to stop the iteration by returning
// `kallax.ErrStop` in the callback.
// Result set is always closed at the end.
func (rs *PlayersResultSet) ForEach(fn func(*Players) error) error {
	for rs.Next() {
		record, err := rs.Get()
		if err != nil {
			return err
		}

		if err := fn(record); err != nil {
			if err == kallax.ErrStop {
				return rs.Close()
			}

			return err
		}
	}
	return nil
}

// All returns all records on the result set and closes the result set.
func (rs *PlayersResultSet) All() ([]*Players, error) {
	var result []*Players
	for rs.Next() {
		record, err := rs.Get()
		if err != nil {
			return nil, err
		}
		result = append(result, record)
	}
	return result, nil
}

// One returns the first record on the result set and closes the result set.
func (rs *PlayersResultSet) One() (*Players, error) {
	if !rs.Next() {
		return nil, kallax.ErrNotFound
	}

	record, err := rs.Get()
	if err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}

	return record, nil
}

// Err returns the last error occurred.
func (rs *PlayersResultSet) Err() error {
	return rs.lastErr
}

// Close closes the result set.
func (rs *PlayersResultSet) Close() error {
	return rs.ResultSet.Close()
}

type schema struct {
	Clubs     *schemaClubs
	Countries *schemaCountries
	Players   *schemaPlayers
}

type schemaClubs struct {
	*kallax.BaseSchema
	CLUBID   kallax.SchemaField
	Clubname kallax.SchemaField
}

type schemaCountries struct {
	*kallax.BaseSchema
	COUNTRYID   kallax.SchemaField
	Countryname kallax.SchemaField
}

type schemaPlayers struct {
	*kallax.BaseSchema
	PID       kallax.SchemaField
	Pname     kallax.SchemaField
	Clubid    kallax.SchemaField
	Countryid kallax.SchemaField
}

var Schema = &schema{
	Clubs: &schemaClubs{
		BaseSchema: kallax.NewBaseSchema(
			"clubs",
			"__clubs",
			kallax.NewSchemaField("clubid"),
			kallax.ForeignKeys{},
			func() kallax.Record {
				return new(Clubs)
			},
			true,
			kallax.NewSchemaField("clubid"),
			kallax.NewSchemaField("clubname"),
		),
		CLUBID:   kallax.NewSchemaField("clubid"),
		Clubname: kallax.NewSchemaField("clubname"),
	},
	Countries: &schemaCountries{
		BaseSchema: kallax.NewBaseSchema(
			"countries",
			"__countries",
			kallax.NewSchemaField("countryid"),
			kallax.ForeignKeys{},
			func() kallax.Record {
				return new(Countries)
			},
			true,
			kallax.NewSchemaField("countryid"),
			kallax.NewSchemaField("countryname"),
		),
		COUNTRYID:   kallax.NewSchemaField("countryid"),
		Countryname: kallax.NewSchemaField("countryname"),
	},
	Players: &schemaPlayers{
		BaseSchema: kallax.NewBaseSchema(
			"players",
			"__players",
			kallax.NewSchemaField("pid"),
			kallax.ForeignKeys{},
			func() kallax.Record {
				return new(Players)
			},
			true,
			kallax.NewSchemaField("pid"),
			kallax.NewSchemaField("pname"),
			kallax.NewSchemaField("clubid"),
			kallax.NewSchemaField("countryid"),
		),
		PID:       kallax.NewSchemaField("pid"),
		Pname:     kallax.NewSchemaField("pname"),
		Clubid:    kallax.NewSchemaField("clubid"),
		Countryid: kallax.NewSchemaField("countryid"),
	},
}
