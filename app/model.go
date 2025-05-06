package gwm_app

import (
	"database/sql/driver"
	"fmt"
	"reflect"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

// An Identifier wrapper for database
type ID string

func (id *ID) Scan(value any) error {
	if value == nil {
		*id = "0"
		return nil
	}
	switch v := value.(type) {
	case int64, uint64, int, uint:
		*id = ID(fmt.Sprintf("%d", v))
	default:
		return fmt.Errorf("cannot scan type %T into ID", value)
	}
	return nil
}

func (id ID) Value() (driver.Value, error) {
	i, err := strconv.Atoi(string(id))
	if err != nil {
		return 0, nil
	}
	return uint(i), nil
}

type Model struct {
	ID        ID         `json:"id" bson:"_id" gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time  `json:"createdAt" bson:"created_at"`
	UpdatedAt time.Time  `json:"updatedAt" bson:"updated_at"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" bson:"deleted_at"`
}

func FromObjectID(objectID bson.ObjectID) ID {
	return ID(objectID.Hex())
}

func RegisterIDToBsonRegistry(registry *bson.Registry) {
	registry.RegisterTypeEncoder(IDType, &IDEncoder{})
}

var IDType = reflect.TypeOf(ID(bson.NewObjectID().Hex()))

type IDEncoder struct{}

func (e *IDEncoder) EncodeValue(ectx bson.EncodeContext, vw bson.ValueWriter, val reflect.Value) error {
	if val.Type() != IDType {
		return bson.ValueDecoderError{
			Name: "IDEncodeValue", Types: []reflect.Type{IDType}, Received: val,
		}
	}
	if val.String() == "" {
		return vw.WriteNull()
	}
	objectID, err := bson.ObjectIDFromHex(val.String())
	if err != nil {
		return err
	}
	return vw.WriteObjectID(objectID)
}
