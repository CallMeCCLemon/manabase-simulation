// Code generated by proroc-gen-graphql, DO NOT EDIT.
package api

import (
	"context"

	"github.com/graphql-go/graphql"
	"github.com/pkg/errors"
	"github.com/ysugimoto/grpc-graphql-gateway/runtime"
	"google.golang.org/grpc"
)

var (
	gql__enum_ManaColor                 *graphql.Enum        // enum ManaColor in api/manabase-simulation.proto
	gql__enum_LandType                  *graphql.Enum        // enum LandType in api/manabase-simulation.proto
	gql__enum_ConditionType             *graphql.Enum        // enum ConditionType in api/manabase-simulation.proto
	gql__type_ValidateDeckListResponse  *graphql.Object      // message ValidateDeckListResponse in api/manabase-simulation.proto
	gql__type_ValidateDeckListRequest   *graphql.Object      // message ValidateDeckListRequest in api/manabase-simulation.proto
	gql__type_UntappedCondition         *graphql.Object      // message UntappedCondition in api/manabase-simulation.proto
	gql__type_SimulateDeckResponse      *graphql.Object      // message SimulateDeckResponse in api/manabase-simulation.proto
	gql__type_SimulateDeckRequest       *graphql.Object      // message SimulateDeckRequest in api/manabase-simulation.proto
	gql__type_ResultCheckpoint          *graphql.Object      // message ResultCheckpoint in api/manabase-simulation.proto
	gql__type_Objective                 *graphql.Object      // message Objective in api/manabase-simulation.proto
	gql__type_NonLand                   *graphql.Object      // message NonLand in api/manabase-simulation.proto
	gql__type_ManaCost                  *graphql.Object      // message ManaCost in api/manabase-simulation.proto
	gql__type_Land                      *graphql.Object      // message Land in api/manabase-simulation.proto
	gql__type_InvalidCard               *graphql.Object      // message InvalidCard in api/manabase-simulation.proto
	gql__type_GameConfiguration         *graphql.Object      // message GameConfiguration in api/manabase-simulation.proto
	gql__type_EchoResponse              *graphql.Object      // message EchoResponse in api/manabase-simulation.proto
	gql__type_EchoRequest               *graphql.Object      // message EchoRequest in api/manabase-simulation.proto
	gql__type_DeckList                  *graphql.Object      // message DeckList in api/manabase-simulation.proto
	gql__type_ActivationCost            *graphql.Object      // message ActivationCost in api/manabase-simulation.proto
	gql__input_ValidateDeckListResponse *graphql.InputObject // message ValidateDeckListResponse in api/manabase-simulation.proto
	gql__input_ValidateDeckListRequest  *graphql.InputObject // message ValidateDeckListRequest in api/manabase-simulation.proto
	gql__input_UntappedCondition        *graphql.InputObject // message UntappedCondition in api/manabase-simulation.proto
	gql__input_SimulateDeckResponse     *graphql.InputObject // message SimulateDeckResponse in api/manabase-simulation.proto
	gql__input_SimulateDeckRequest      *graphql.InputObject // message SimulateDeckRequest in api/manabase-simulation.proto
	gql__input_ResultCheckpoint         *graphql.InputObject // message ResultCheckpoint in api/manabase-simulation.proto
	gql__input_Objective                *graphql.InputObject // message Objective in api/manabase-simulation.proto
	gql__input_NonLand                  *graphql.InputObject // message NonLand in api/manabase-simulation.proto
	gql__input_ManaCost                 *graphql.InputObject // message ManaCost in api/manabase-simulation.proto
	gql__input_Land                     *graphql.InputObject // message Land in api/manabase-simulation.proto
	gql__input_InvalidCard              *graphql.InputObject // message InvalidCard in api/manabase-simulation.proto
	gql__input_GameConfiguration        *graphql.InputObject // message GameConfiguration in api/manabase-simulation.proto
	gql__input_EchoResponse             *graphql.InputObject // message EchoResponse in api/manabase-simulation.proto
	gql__input_EchoRequest              *graphql.InputObject // message EchoRequest in api/manabase-simulation.proto
	gql__input_DeckList                 *graphql.InputObject // message DeckList in api/manabase-simulation.proto
	gql__input_ActivationCost           *graphql.InputObject // message ActivationCost in api/manabase-simulation.proto
)

func Gql__enum_ManaColor() *graphql.Enum {
	if gql__enum_ManaColor == nil {
		gql__enum_ManaColor = graphql.NewEnum(graphql.EnumConfig{
			Name: "Api_Enum_ManaColor",
			Values: graphql.EnumValueConfigMap{
				"WHITE": &graphql.EnumValueConfig{
					Value: ManaColor(0),
				},
				"BLUE": &graphql.EnumValueConfig{
					Value: ManaColor(1),
				},
				"BLACK": &graphql.EnumValueConfig{
					Value: ManaColor(2),
				},
				"RED": &graphql.EnumValueConfig{
					Value: ManaColor(3),
				},
				"GREEN": &graphql.EnumValueConfig{
					Value: ManaColor(4),
				},
				"COLORLESS": &graphql.EnumValueConfig{
					Value: ManaColor(5),
				},
			},
		})
	}
	return gql__enum_ManaColor
}

func Gql__enum_LandType() *graphql.Enum {
	if gql__enum_LandType == nil {
		gql__enum_LandType = graphql.NewEnum(graphql.EnumConfig{
			Name: "Api_Enum_LandType",
			Values: graphql.EnumValueConfigMap{
				"PLAINS": &graphql.EnumValueConfig{
					Value: LandType(0),
				},
				"MOUNTAIN": &graphql.EnumValueConfig{
					Value: LandType(1),
				},
				"FOREST": &graphql.EnumValueConfig{
					Value: LandType(2),
				},
				"ISLAND": &graphql.EnumValueConfig{
					Value: LandType(3),
				},
				"SWAMP": &graphql.EnumValueConfig{
					Value: LandType(4),
				},
			},
		})
	}
	return gql__enum_LandType
}

func Gql__enum_ConditionType() *graphql.Enum {
	if gql__enum_ConditionType == nil {
		gql__enum_ConditionType = graphql.NewEnum(graphql.EnumConfig{
			Name: "Api_Enum_ConditionType",
			Values: graphql.EnumValueConfigMap{
				"SHOCK_LAND": &graphql.EnumValueConfig{
					Value: ConditionType(0),
				},
				"FAST_LAND": &graphql.EnumValueConfig{
					Value: ConditionType(1),
				},
				"CHECK_LAND": &graphql.EnumValueConfig{
					Value: ConditionType(2),
				},
			},
		})
	}
	return gql__enum_ConditionType
}

func Gql__type_ValidateDeckListResponse() *graphql.Object {
	if gql__type_ValidateDeckListResponse == nil {
		gql__type_ValidateDeckListResponse = graphql.NewObject(graphql.ObjectConfig{
			Name: "Api_Type_ValidateDeckListResponse",
			Fields: graphql.Fields{
				"isValid": &graphql.Field{
					Type: graphql.Boolean,
				},
				"invalidCards": &graphql.Field{
					Type: graphql.NewList(Gql__type_InvalidCard()),
				},
			},
		})
	}
	return gql__type_ValidateDeckListResponse
}

func Gql__type_ValidateDeckListRequest() *graphql.Object {
	if gql__type_ValidateDeckListRequest == nil {
		gql__type_ValidateDeckListRequest = graphql.NewObject(graphql.ObjectConfig{
			Name: "Api_Type_ValidateDeckListRequest",
			Fields: graphql.Fields{
				"deckList": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__type_ValidateDeckListRequest
}

func Gql__type_UntappedCondition() *graphql.Object {
	if gql__type_UntappedCondition == nil {
		gql__type_UntappedCondition = graphql.NewObject(graphql.ObjectConfig{
			Name: "Api_Type_UntappedCondition",
			Fields: graphql.Fields{
				"type": &graphql.Field{
					Type: Gql__enum_ConditionType(),
				},
				"data": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__type_UntappedCondition
}

func Gql__type_SimulateDeckResponse() *graphql.Object {
	if gql__type_SimulateDeckResponse == nil {
		gql__type_SimulateDeckResponse = graphql.NewObject(graphql.ObjectConfig{
			Name: "Api_Type_SimulateDeckResponse",
			Fields: graphql.Fields{
				"message": &graphql.Field{
					Type: graphql.String,
				},
				"successRate": &graphql.Field{
					Type: graphql.Float,
				},
				"checkpoints": &graphql.Field{
					Type: graphql.NewList(Gql__type_ResultCheckpoint()),
				},
			},
		})
	}
	return gql__type_SimulateDeckResponse
}

func Gql__type_SimulateDeckRequest() *graphql.Object {
	if gql__type_SimulateDeckRequest == nil {
		gql__type_SimulateDeckRequest = graphql.NewObject(graphql.ObjectConfig{
			Name:        "Api_Type_SimulateDeckRequest",
			Description: `SimulateDeckRequest Represents the request to simulate a deck.`,
			Fields: graphql.Fields{
				"deckList": &graphql.Field{
					Type: Gql__type_DeckList(),
				},
				"gameConfiguration": &graphql.Field{
					Type: Gql__type_GameConfiguration(),
				},
				"objective": &graphql.Field{
					Type: Gql__type_Objective(),
				},
			},
		})
	}
	return gql__type_SimulateDeckRequest
}

func Gql__type_ResultCheckpoint() *graphql.Object {
	if gql__type_ResultCheckpoint == nil {
		gql__type_ResultCheckpoint = graphql.NewObject(graphql.ObjectConfig{
			Name: "Api_Type_ResultCheckpoint",
			Fields: graphql.Fields{
				"iterations": &graphql.Field{
					Type: graphql.Int,
				},
				"successes": &graphql.Field{
					Type: graphql.Int,
				},
			},
		})
	}
	return gql__type_ResultCheckpoint
}

func Gql__type_Objective() *graphql.Object {
	if gql__type_Objective == nil {
		gql__type_Objective = graphql.NewObject(graphql.ObjectConfig{
			Name: "Api_Type_Objective",
			Fields: graphql.Fields{
				"targetTurn": &graphql.Field{
					Type: graphql.Int,
				},
				"manaCosts": &graphql.Field{
					Type: graphql.NewList(Gql__type_ManaCost()),
				},
			},
		})
	}
	return gql__type_Objective
}

func Gql__type_NonLand() *graphql.Object {
	if gql__type_NonLand == nil {
		gql__type_NonLand = graphql.NewObject(graphql.ObjectConfig{
			Name: "Api_Type_NonLand",
			Fields: graphql.Fields{
				"name": &graphql.Field{
					Type: graphql.String,
				},
				"castingCost": &graphql.Field{
					Type: Gql__type_ManaCost(),
				},
				"quantity": &graphql.Field{
					Type: graphql.Int,
				},
			},
		})
	}
	return gql__type_NonLand
}

func Gql__type_ManaCost() *graphql.Object {
	if gql__type_ManaCost == nil {
		gql__type_ManaCost = graphql.NewObject(graphql.ObjectConfig{
			Name: "Api_Type_ManaCost",
			Fields: graphql.Fields{
				"colorRequirements": &graphql.Field{
					Type: graphql.NewList(Gql__enum_ManaColor()),
				},
				"genericCost": &graphql.Field{
					Type: graphql.Int,
				},
			},
		})
	}
	return gql__type_ManaCost
}

func Gql__type_Land() *graphql.Object {
	if gql__type_Land == nil {
		gql__type_Land = graphql.NewObject(graphql.ObjectConfig{
			Name: "Api_Type_Land",
			Fields: graphql.Fields{
				"name": &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				"colors": &graphql.Field{
					Type: graphql.NewList(Gql__enum_ManaColor()),
				},
				"entersTapped": &graphql.Field{
					Type: graphql.Boolean,
				},
				"activationCost": &graphql.Field{
					Type: Gql__type_ActivationCost(),
				},
				"types": &graphql.Field{
					Type: graphql.NewList(Gql__enum_LandType()),
				},
				"untappedCondition": &graphql.Field{
					Type: Gql__type_UntappedCondition(),
				},
				"quantity": &graphql.Field{
					Type: graphql.Int,
				},
			},
		})
	}
	return gql__type_Land
}

func Gql__type_InvalidCard() *graphql.Object {
	if gql__type_InvalidCard == nil {
		gql__type_InvalidCard = graphql.NewObject(graphql.ObjectConfig{
			Name: "Api_Type_InvalidCard",
			Fields: graphql.Fields{
				"name": &graphql.Field{
					Type: graphql.String,
				},
				"reason": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__type_InvalidCard
}

func Gql__type_GameConfiguration() *graphql.Object {
	if gql__type_GameConfiguration == nil {
		gql__type_GameConfiguration = graphql.NewObject(graphql.ObjectConfig{
			Name: "Api_Type_GameConfiguration",
			Fields: graphql.Fields{
				"initialHandSize": &graphql.Field{
					Type: graphql.Int,
				},
				"cardsDrawnPerTurn": &graphql.Field{
					Type: graphql.Int,
				},
				"onThePlay": &graphql.Field{
					Type: graphql.Boolean,
				},
			},
		})
	}
	return gql__type_GameConfiguration
}

func Gql__type_EchoResponse() *graphql.Object {
	if gql__type_EchoResponse == nil {
		gql__type_EchoResponse = graphql.NewObject(graphql.ObjectConfig{
			Name: "Api_Type_EchoResponse",
			Fields: graphql.Fields{
				"message": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__type_EchoResponse
}

func Gql__type_EchoRequest() *graphql.Object {
	if gql__type_EchoRequest == nil {
		gql__type_EchoRequest = graphql.NewObject(graphql.ObjectConfig{
			Name: "Api_Type_EchoRequest",
			Fields: graphql.Fields{
				"message": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__type_EchoRequest
}

func Gql__type_DeckList() *graphql.Object {
	if gql__type_DeckList == nil {
		gql__type_DeckList = graphql.NewObject(graphql.ObjectConfig{
			Name: "Api_Type_DeckList",
			Fields: graphql.Fields{
				"lands": &graphql.Field{
					Type: graphql.NewList(Gql__type_Land()),
				},
				"nonLands": &graphql.Field{
					Type: graphql.NewList(Gql__type_NonLand()),
				},
			},
		})
	}
	return gql__type_DeckList
}

func Gql__type_ActivationCost() *graphql.Object {
	if gql__type_ActivationCost == nil {
		gql__type_ActivationCost = graphql.NewObject(graphql.ObjectConfig{
			Name: "Api_Type_ActivationCost",
			Fields: graphql.Fields{
				"life": &graphql.Field{
					Type: graphql.Int,
				},
				"manaCost": &graphql.Field{
					Type: Gql__type_ManaCost(),
				},
			},
		})
	}
	return gql__type_ActivationCost
}

func Gql__input_ValidateDeckListResponse() *graphql.InputObject {
	if gql__input_ValidateDeckListResponse == nil {
		gql__input_ValidateDeckListResponse = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Api_Input_ValidateDeckListResponse",
			Fields: graphql.InputObjectConfigFieldMap{
				"isValid": &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
				"invalidCards": &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(Gql__input_InvalidCard()),
				},
			},
		})
	}
	return gql__input_ValidateDeckListResponse
}

func Gql__input_ValidateDeckListRequest() *graphql.InputObject {
	if gql__input_ValidateDeckListRequest == nil {
		gql__input_ValidateDeckListRequest = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Api_Input_ValidateDeckListRequest",
			Fields: graphql.InputObjectConfigFieldMap{
				"deckList": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__input_ValidateDeckListRequest
}

func Gql__input_UntappedCondition() *graphql.InputObject {
	if gql__input_UntappedCondition == nil {
		gql__input_UntappedCondition = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Api_Input_UntappedCondition",
			Fields: graphql.InputObjectConfigFieldMap{
				"type": &graphql.InputObjectFieldConfig{
					Type: Gql__enum_ConditionType(),
				},
				"data": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__input_UntappedCondition
}

func Gql__input_SimulateDeckResponse() *graphql.InputObject {
	if gql__input_SimulateDeckResponse == nil {
		gql__input_SimulateDeckResponse = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Api_Input_SimulateDeckResponse",
			Fields: graphql.InputObjectConfigFieldMap{
				"message": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"successRate": &graphql.InputObjectFieldConfig{
					Type: graphql.Float,
				},
				"checkpoints": &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(Gql__input_ResultCheckpoint()),
				},
			},
		})
	}
	return gql__input_SimulateDeckResponse
}

func Gql__input_SimulateDeckRequest() *graphql.InputObject {
	if gql__input_SimulateDeckRequest == nil {
		gql__input_SimulateDeckRequest = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Api_Input_SimulateDeckRequest",
			Fields: graphql.InputObjectConfigFieldMap{
				"deckList": &graphql.InputObjectFieldConfig{
					Type: Gql__input_DeckList(),
				},
				"gameConfiguration": &graphql.InputObjectFieldConfig{
					Type: Gql__input_GameConfiguration(),
				},
				"objective": &graphql.InputObjectFieldConfig{
					Type: Gql__input_Objective(),
				},
			},
		})
	}
	return gql__input_SimulateDeckRequest
}

func Gql__input_ResultCheckpoint() *graphql.InputObject {
	if gql__input_ResultCheckpoint == nil {
		gql__input_ResultCheckpoint = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Api_Input_ResultCheckpoint",
			Fields: graphql.InputObjectConfigFieldMap{
				"iterations": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
				"successes": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
			},
		})
	}
	return gql__input_ResultCheckpoint
}

func Gql__input_Objective() *graphql.InputObject {
	if gql__input_Objective == nil {
		gql__input_Objective = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Api_Input_Objective",
			Fields: graphql.InputObjectConfigFieldMap{
				"targetTurn": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
				"manaCosts": &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(Gql__input_ManaCost()),
				},
			},
		})
	}
	return gql__input_Objective
}

func Gql__input_NonLand() *graphql.InputObject {
	if gql__input_NonLand == nil {
		gql__input_NonLand = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Api_Input_NonLand",
			Fields: graphql.InputObjectConfigFieldMap{
				"name": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"castingCost": &graphql.InputObjectFieldConfig{
					Type: Gql__input_ManaCost(),
				},
				"quantity": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
			},
		})
	}
	return gql__input_NonLand
}

func Gql__input_ManaCost() *graphql.InputObject {
	if gql__input_ManaCost == nil {
		gql__input_ManaCost = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Api_Input_ManaCost",
			Fields: graphql.InputObjectConfigFieldMap{
				"colorRequirements": &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(Gql__enum_ManaColor()),
				},
				"genericCost": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
			},
		})
	}
	return gql__input_ManaCost
}

func Gql__input_Land() *graphql.InputObject {
	if gql__input_Land == nil {
		gql__input_Land = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Api_Input_Land",
			Fields: graphql.InputObjectConfigFieldMap{
				"name": &graphql.InputObjectFieldConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"colors": &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(Gql__enum_ManaColor()),
				},
				"entersTapped": &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
				"activationCost": &graphql.InputObjectFieldConfig{
					Type: Gql__input_ActivationCost(),
				},
				"types": &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(Gql__enum_LandType()),
				},
				"untappedCondition": &graphql.InputObjectFieldConfig{
					Type: Gql__input_UntappedCondition(),
				},
				"quantity": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
			},
		})
	}
	return gql__input_Land
}

func Gql__input_InvalidCard() *graphql.InputObject {
	if gql__input_InvalidCard == nil {
		gql__input_InvalidCard = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Api_Input_InvalidCard",
			Fields: graphql.InputObjectConfigFieldMap{
				"name": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"reason": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__input_InvalidCard
}

func Gql__input_GameConfiguration() *graphql.InputObject {
	if gql__input_GameConfiguration == nil {
		gql__input_GameConfiguration = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Api_Input_GameConfiguration",
			Fields: graphql.InputObjectConfigFieldMap{
				"initialHandSize": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
				"cardsDrawnPerTurn": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
				"onThePlay": &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
			},
		})
	}
	return gql__input_GameConfiguration
}

func Gql__input_EchoResponse() *graphql.InputObject {
	if gql__input_EchoResponse == nil {
		gql__input_EchoResponse = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Api_Input_EchoResponse",
			Fields: graphql.InputObjectConfigFieldMap{
				"message": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__input_EchoResponse
}

func Gql__input_EchoRequest() *graphql.InputObject {
	if gql__input_EchoRequest == nil {
		gql__input_EchoRequest = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Api_Input_EchoRequest",
			Fields: graphql.InputObjectConfigFieldMap{
				"message": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__input_EchoRequest
}

func Gql__input_DeckList() *graphql.InputObject {
	if gql__input_DeckList == nil {
		gql__input_DeckList = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Api_Input_DeckList",
			Fields: graphql.InputObjectConfigFieldMap{
				"lands": &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(Gql__input_Land()),
				},
				"nonLands": &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(Gql__input_NonLand()),
				},
			},
		})
	}
	return gql__input_DeckList
}

func Gql__input_ActivationCost() *graphql.InputObject {
	if gql__input_ActivationCost == nil {
		gql__input_ActivationCost = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Api_Input_ActivationCost",
			Fields: graphql.InputObjectConfigFieldMap{
				"life": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
				"manaCost": &graphql.InputObjectFieldConfig{
					Type: Gql__input_ManaCost(),
				},
			},
		})
	}
	return gql__input_ActivationCost
}

// graphql__resolver_ManabaseSimulator is a struct for making query, mutation and resolve fields.
// This struct must be implemented runtime.SchemaBuilder interface.
type graphql__resolver_ManabaseSimulator struct {

	// Automatic connection host
	host string

	// grpc dial options
	dialOptions []grpc.DialOption

	// grpc client connection.
	// this connection may be provided by user
	conn *grpc.ClientConn
}

// new_graphql_resolver_ManabaseSimulator creates pointer of service struct
func new_graphql_resolver_ManabaseSimulator(conn *grpc.ClientConn) *graphql__resolver_ManabaseSimulator {
	return &graphql__resolver_ManabaseSimulator{
		conn: conn,
		host: "mtg-mana-sim-app-server-service:9000",
		dialOptions: []grpc.DialOption{
			grpc.WithInsecure(),
		},
	}
}

// CreateConnection() returns grpc connection which user specified or newly connected and closing function
func (x *graphql__resolver_ManabaseSimulator) CreateConnection(ctx context.Context) (*grpc.ClientConn, func(), error) {
	// If x.conn is not nil, user injected their own connection
	if x.conn != nil {
		return x.conn, func() {}, nil
	}

	// Otherwise, this handler opens connection with specified host
	conn, err := grpc.DialContext(ctx, x.host, x.dialOptions...)
	if err != nil {
		return nil, nil, err
	}
	return conn, func() { conn.Close() }, nil
}

// GetQueries returns acceptable graphql.Fields for Query.
func (x *graphql__resolver_ManabaseSimulator) GetQueries(conn *grpc.ClientConn) graphql.Fields {
	return graphql.Fields{
		"echo": &graphql.Field{
			Type: Gql__type_EchoResponse(),
			Args: graphql.FieldConfigArgument{
				"message": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var req EchoRequest
				if err := runtime.MarshalRequest(p.Args, &req, false); err != nil {
					return nil, errors.Wrap(err, "Failed to marshal request for echo")
				}
				client := NewManabaseSimulatorClient(conn)
				resp, err := client.Echo(p.Context, &req)
				if err != nil {
					return nil, errors.Wrap(err, "Failed to call RPC Echo")
				}
				return resp, nil
			},
		},
		"simulate": &graphql.Field{
			Type: Gql__type_SimulateDeckResponse(),
			Args: graphql.FieldConfigArgument{
				"deckList": &graphql.ArgumentConfig{
					Type: Gql__input_DeckList(),
				},
				"gameConfiguration": &graphql.ArgumentConfig{
					Type: Gql__input_GameConfiguration(),
				},
				"objective": &graphql.ArgumentConfig{
					Type: Gql__input_Objective(),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var req SimulateDeckRequest
				if err := runtime.MarshalRequest(p.Args, &req, false); err != nil {
					return nil, errors.Wrap(err, "Failed to marshal request for simulate")
				}
				client := NewManabaseSimulatorClient(conn)
				resp, err := client.SimulateDeck(p.Context, &req)
				if err != nil {
					return nil, errors.Wrap(err, "Failed to call RPC SimulateDeck")
				}
				return resp, nil
			},
		},
		"validate": &graphql.Field{
			Type: Gql__type_ValidateDeckListResponse(),
			Args: graphql.FieldConfigArgument{
				"deckList": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var req ValidateDeckListRequest
				if err := runtime.MarshalRequest(p.Args, &req, false); err != nil {
					return nil, errors.Wrap(err, "Failed to marshal request for validate")
				}
				client := NewManabaseSimulatorClient(conn)
				resp, err := client.ValidateDeckList(p.Context, &req)
				if err != nil {
					return nil, errors.Wrap(err, "Failed to call RPC ValidateDeckList")
				}
				return resp, nil
			},
		},
	}
}

// GetMutations returns acceptable graphql.Fields for Mutation.
func (x *graphql__resolver_ManabaseSimulator) GetMutations(conn *grpc.ClientConn) graphql.Fields {
	return graphql.Fields{}
}

// Register package divided graphql handler "without" *grpc.ClientConn,
// therefore gRPC connection will be opened and closed automatically.
// Occasionally you may worry about open/close performance for each handling graphql request,
// then you can call RegisterManabaseSimulatorGraphqlHandler with *grpc.ClientConn manually.
func RegisterManabaseSimulatorGraphql(mux *runtime.ServeMux) error {
	return RegisterManabaseSimulatorGraphqlHandler(mux, nil)
}

// Register package divided graphql handler "with" *grpc.ClientConn.
// this function accepts your defined grpc connection, so that we reuse that and never close connection inside.
// You need to close it maunally when application will terminate.
// Otherwise, you can specify automatic opening connection with ServiceOption directive:
//
//	service ManabaseSimulator {
//	   option (graphql.service) = {
//	       host: "host:port"
//	       insecure: true or false
//	   };
//
//	   ...with RPC definitions
//	}
func RegisterManabaseSimulatorGraphqlHandler(mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return mux.AddHandler(new_graphql_resolver_ManabaseSimulator(conn))
}
