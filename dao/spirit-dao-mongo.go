package dao

import (
	"github.com/Sfeir/handsongo/model"
	logger "github.com/Sirupsen/logrus"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	collection = "spirits"
	index      = "id"
)

// SpiritDAOMongo is the mongo implementation of the SpiritDAO
type SpiritDAOMongo struct {
	session *mgo.Session
}

// NewSpiritDAOMongo creates a new SpiritDAO mongo implementation
func NewSpiritDAOMongo(session *mgo.Session) SpiritDAO {
	// create index
	err := session.DB("").C(collection).EnsureIndex(mgo.Index{
		Key:        []string{index},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	})

	if err != nil {
		logger.WithField("error", err).Warn("mongo db connection")
	}

	return &SpiritDAOMongo{
		session: session,
	}
}

// GetSpiritByID returns a spirit by its ID
func (s *SpiritDAOMongo) GetSpiritByID(ID string) (*model.Spirit, error) {

	// TODO use the bson library to check if the ID is a well formed ObjectId
	// if not return a new error "Invalid input to ObjectIdHex"
	// check ID

	// copying the session and defering the close
	session := s.session.Copy()
	defer session.Close()

	// instantiate a new spirit to be hydrated
	spirit := model.Spirit{}

	// retrieve the collection
	c := session.DB("").C(collection)

	// find one spirit by its ObjectId
	err := c.Find(bson.M{"_id": bson.ObjectIdHex(ID)}).One(&spirit)
	return &spirit, err
}

// getAllSpiritsByQuery returns spirits by query and paging capability
func (s *SpiritDAOMongo) getAllSpiritsByQuery(query interface{}, start, end int) ([]model.Spirit, error) {

	// TODO copy the session and also defer the closing

	// TODO retrieve the collection

	// check param
	hasPaging := start > NoPaging && end > NoPaging && end > start

	// perform request
	var err error
	spirits := []model.Spirit{}
	if hasPaging {
		// TODO if paging is defined, use the Find, Skip, Limit and All method to execute the query
	} else {
		// TODO if no paging is defined simply query All the spirits
	}

	return spirits, err
}

// GetAllSpirits returns all spirits with paging capability
func (s *SpiritDAOMongo) GetAllSpirits(start, end int) ([]model.Spirit, error) {
	return s.getAllSpiritsByQuery(nil, start, end)
}

// GetSpiritsByName returns all spirits by name
func (s *SpiritDAOMongo) GetSpiritsByName(name string) ([]model.Spirit, error) {
	// TODO with previous method, query all spirits by their name "bson.M{"name": name}" without Paging
	return nil, nil
}

// GetSpiritsByType returns all spirits by type
func (s *SpiritDAOMongo) GetSpiritsByType(spiritType string) ([]model.Spirit, error) {
	// TODO query all spirits by their type without Paging
	return nil, nil
}

// GetSpiritsByTypeAndScore returns all spirits by type and score greater than parameter
func (s *SpiritDAOMongo) GetSpiritsByTypeAndScore(spiritType string, score uint8) ([]model.Spirit, error) {
	// TODO query all spirits by their type and score greater than score parameter ""score": bson.M{"$gte": score}"
	return nil, nil
}

// SaveSpirit saves the spirit
func (s *SpiritDAOMongo) SaveSpirit(spirit *model.Spirit) error {
	// TODO copy the session and also defer the closing
	// TODO retrieve the collection
	// TODO use Insert on the collection to save a spirit
	return nil
}

// UpsertSpirit updates or creates a spirit
func (s *SpiritDAOMongo) UpsertSpirit(ID string, spirit *model.Spirit) (bool, error) {
	session := s.session.Copy()
	defer session.Close()
	c := session.DB("").C(collection)
	chg, err := c.Upsert(bson.M{"_id": bson.ObjectIdHex(ID)}, spirit)
	if err != nil {
		return false, err
	}
	return chg.Updated > 0, err
}

// DeleteSpirit deletes a spirits by its ID
func (s *SpiritDAOMongo) DeleteSpirit(ID string) error {
	// TODO copy the session and also defer the closing
	// TODO retrieve the collection
	// TODO use Remove on collection to delete a spirit by its ObjectId
	return nil
}