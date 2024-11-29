// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: api/manabase-simulation.proto

package api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type LandType int32

const (
	LandType_PLAINS   LandType = 0
	LandType_MOUNTAIN LandType = 1
	LandType_FOREST   LandType = 2
	LandType_ISLAND   LandType = 3
	LandType_SWAMP    LandType = 4
)

// Enum value maps for LandType.
var (
	LandType_name = map[int32]string{
		0: "PLAINS",
		1: "MOUNTAIN",
		2: "FOREST",
		3: "ISLAND",
		4: "SWAMP",
	}
	LandType_value = map[string]int32{
		"PLAINS":   0,
		"MOUNTAIN": 1,
		"FOREST":   2,
		"ISLAND":   3,
		"SWAMP":    4,
	}
)

func (x LandType) Enum() *LandType {
	p := new(LandType)
	*p = x
	return p
}

func (x LandType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LandType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_manabase_simulation_proto_enumTypes[0].Descriptor()
}

func (LandType) Type() protoreflect.EnumType {
	return &file_api_manabase_simulation_proto_enumTypes[0]
}

func (x LandType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LandType.Descriptor instead.
func (LandType) EnumDescriptor() ([]byte, []int) {
	return file_api_manabase_simulation_proto_rawDescGZIP(), []int{0}
}

type ManaColor int32

const (
	ManaColor_WHITE     ManaColor = 0
	ManaColor_BLUE      ManaColor = 1
	ManaColor_BLACK     ManaColor = 2
	ManaColor_RED       ManaColor = 3
	ManaColor_GREEN     ManaColor = 4
	ManaColor_COLORLESS ManaColor = 5
)

// Enum value maps for ManaColor.
var (
	ManaColor_name = map[int32]string{
		0: "WHITE",
		1: "BLUE",
		2: "BLACK",
		3: "RED",
		4: "GREEN",
		5: "COLORLESS",
	}
	ManaColor_value = map[string]int32{
		"WHITE":     0,
		"BLUE":      1,
		"BLACK":     2,
		"RED":       3,
		"GREEN":     4,
		"COLORLESS": 5,
	}
)

func (x ManaColor) Enum() *ManaColor {
	p := new(ManaColor)
	*p = x
	return p
}

func (x ManaColor) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ManaColor) Descriptor() protoreflect.EnumDescriptor {
	return file_api_manabase_simulation_proto_enumTypes[1].Descriptor()
}

func (ManaColor) Type() protoreflect.EnumType {
	return &file_api_manabase_simulation_proto_enumTypes[1]
}

func (x ManaColor) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ManaColor.Descriptor instead.
func (ManaColor) EnumDescriptor() ([]byte, []int) {
	return file_api_manabase_simulation_proto_rawDescGZIP(), []int{1}
}

type ConditionType int32

const (
	ConditionType_SHOCK_LAND ConditionType = 0
	ConditionType_FAST_LAND  ConditionType = 1
	ConditionType_CHECK_LAND ConditionType = 2
)

// Enum value maps for ConditionType.
var (
	ConditionType_name = map[int32]string{
		0: "SHOCK_LAND",
		1: "FAST_LAND",
		2: "CHECK_LAND",
	}
	ConditionType_value = map[string]int32{
		"SHOCK_LAND": 0,
		"FAST_LAND":  1,
		"CHECK_LAND": 2,
	}
)

func (x ConditionType) Enum() *ConditionType {
	p := new(ConditionType)
	*p = x
	return p
}

func (x ConditionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConditionType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_manabase_simulation_proto_enumTypes[2].Descriptor()
}

func (ConditionType) Type() protoreflect.EnumType {
	return &file_api_manabase_simulation_proto_enumTypes[2]
}

func (x ConditionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConditionType.Descriptor instead.
func (ConditionType) EnumDescriptor() ([]byte, []int) {
	return file_api_manabase_simulation_proto_rawDescGZIP(), []int{2}
}

type SimulateDeckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeckList          *DeckList          `protobuf:"bytes,1,opt,name=deckList,proto3" json:"deckList,omitempty"`
	GameConfiguration *GameConfiguration `protobuf:"bytes,2,opt,name=gameConfiguration,proto3" json:"gameConfiguration,omitempty"`
	Objective         *Objective         `protobuf:"bytes,3,opt,name=objective,proto3" json:"objective,omitempty"`
}

func (x *SimulateDeckRequest) Reset() {
	*x = SimulateDeckRequest{}
	mi := &file_api_manabase_simulation_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SimulateDeckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimulateDeckRequest) ProtoMessage() {}

func (x *SimulateDeckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_manabase_simulation_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimulateDeckRequest.ProtoReflect.Descriptor instead.
func (*SimulateDeckRequest) Descriptor() ([]byte, []int) {
	return file_api_manabase_simulation_proto_rawDescGZIP(), []int{0}
}

func (x *SimulateDeckRequest) GetDeckList() *DeckList {
	if x != nil {
		return x.DeckList
	}
	return nil
}

func (x *SimulateDeckRequest) GetGameConfiguration() *GameConfiguration {
	if x != nil {
		return x.GameConfiguration
	}
	return nil
}

func (x *SimulateDeckRequest) GetObjective() *Objective {
	if x != nil {
		return x.Objective
	}
	return nil
}

type SimulateDeckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message     string  `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	SuccessRate float32 `protobuf:"fixed32,2,opt,name=successRate,proto3" json:"successRate,omitempty"`
}

func (x *SimulateDeckResponse) Reset() {
	*x = SimulateDeckResponse{}
	mi := &file_api_manabase_simulation_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SimulateDeckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimulateDeckResponse) ProtoMessage() {}

func (x *SimulateDeckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_manabase_simulation_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimulateDeckResponse.ProtoReflect.Descriptor instead.
func (*SimulateDeckResponse) Descriptor() ([]byte, []int) {
	return file_api_manabase_simulation_proto_rawDescGZIP(), []int{1}
}

func (x *SimulateDeckResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *SimulateDeckResponse) GetSuccessRate() float32 {
	if x != nil {
		return x.SuccessRate
	}
	return 0
}

type DeckList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lands    []*Land    `protobuf:"bytes,1,rep,name=lands,proto3" json:"lands,omitempty"`
	NonLands []*NonLand `protobuf:"bytes,2,rep,name=nonLands,proto3" json:"nonLands,omitempty"`
}

func (x *DeckList) Reset() {
	*x = DeckList{}
	mi := &file_api_manabase_simulation_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeckList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeckList) ProtoMessage() {}

func (x *DeckList) ProtoReflect() protoreflect.Message {
	mi := &file_api_manabase_simulation_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeckList.ProtoReflect.Descriptor instead.
func (*DeckList) Descriptor() ([]byte, []int) {
	return file_api_manabase_simulation_proto_rawDescGZIP(), []int{2}
}

func (x *DeckList) GetLands() []*Land {
	if x != nil {
		return x.Lands
	}
	return nil
}

func (x *DeckList) GetNonLands() []*NonLand {
	if x != nil {
		return x.NonLands
	}
	return nil
}

type Land struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name              string             `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Colors            []ManaColor        `protobuf:"varint,2,rep,packed,name=colors,proto3,enum=manabase_simulation.ManaColor" json:"colors,omitempty"`
	EntersTapped      bool               `protobuf:"varint,3,opt,name=entersTapped,proto3" json:"entersTapped,omitempty"`
	ActivationCost    *ActivationCost    `protobuf:"bytes,4,opt,name=activationCost,proto3" json:"activationCost,omitempty"`
	Types             []LandType         `protobuf:"varint,5,rep,packed,name=types,proto3,enum=manabase_simulation.LandType" json:"types,omitempty"`
	UntappedCondition *UntappedCondition `protobuf:"bytes,6,opt,name=untappedCondition,proto3" json:"untappedCondition,omitempty"`
	Quantity          int32              `protobuf:"varint,7,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *Land) Reset() {
	*x = Land{}
	mi := &file_api_manabase_simulation_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Land) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Land) ProtoMessage() {}

func (x *Land) ProtoReflect() protoreflect.Message {
	mi := &file_api_manabase_simulation_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Land.ProtoReflect.Descriptor instead.
func (*Land) Descriptor() ([]byte, []int) {
	return file_api_manabase_simulation_proto_rawDescGZIP(), []int{3}
}

func (x *Land) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Land) GetColors() []ManaColor {
	if x != nil {
		return x.Colors
	}
	return nil
}

func (x *Land) GetEntersTapped() bool {
	if x != nil {
		return x.EntersTapped
	}
	return false
}

func (x *Land) GetActivationCost() *ActivationCost {
	if x != nil {
		return x.ActivationCost
	}
	return nil
}

func (x *Land) GetTypes() []LandType {
	if x != nil {
		return x.Types
	}
	return nil
}

func (x *Land) GetUntappedCondition() *UntappedCondition {
	if x != nil {
		return x.UntappedCondition
	}
	return nil
}

func (x *Land) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type NonLand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	CastingCost *ManaCost `protobuf:"bytes,2,opt,name=castingCost,proto3" json:"castingCost,omitempty"`
	Quantity    int32     `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *NonLand) Reset() {
	*x = NonLand{}
	mi := &file_api_manabase_simulation_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NonLand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NonLand) ProtoMessage() {}

func (x *NonLand) ProtoReflect() protoreflect.Message {
	mi := &file_api_manabase_simulation_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NonLand.ProtoReflect.Descriptor instead.
func (*NonLand) Descriptor() ([]byte, []int) {
	return file_api_manabase_simulation_proto_rawDescGZIP(), []int{4}
}

func (x *NonLand) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NonLand) GetCastingCost() *ManaCost {
	if x != nil {
		return x.CastingCost
	}
	return nil
}

func (x *NonLand) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type ActivationCost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Life     int32     `protobuf:"varint,1,opt,name=life,proto3" json:"life,omitempty"`
	ManaCost *ManaCost `protobuf:"bytes,2,opt,name=manaCost,proto3,oneof" json:"manaCost,omitempty"`
}

func (x *ActivationCost) Reset() {
	*x = ActivationCost{}
	mi := &file_api_manabase_simulation_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ActivationCost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivationCost) ProtoMessage() {}

func (x *ActivationCost) ProtoReflect() protoreflect.Message {
	mi := &file_api_manabase_simulation_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivationCost.ProtoReflect.Descriptor instead.
func (*ActivationCost) Descriptor() ([]byte, []int) {
	return file_api_manabase_simulation_proto_rawDescGZIP(), []int{5}
}

func (x *ActivationCost) GetLife() int32 {
	if x != nil {
		return x.Life
	}
	return 0
}

func (x *ActivationCost) GetManaCost() *ManaCost {
	if x != nil {
		return x.ManaCost
	}
	return nil
}

type UntappedCondition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type ConditionType `protobuf:"varint,1,opt,name=type,proto3,enum=manabase_simulation.ConditionType" json:"type,omitempty"`
	Data string        `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *UntappedCondition) Reset() {
	*x = UntappedCondition{}
	mi := &file_api_manabase_simulation_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UntappedCondition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UntappedCondition) ProtoMessage() {}

func (x *UntappedCondition) ProtoReflect() protoreflect.Message {
	mi := &file_api_manabase_simulation_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UntappedCondition.ProtoReflect.Descriptor instead.
func (*UntappedCondition) Descriptor() ([]byte, []int) {
	return file_api_manabase_simulation_proto_rawDescGZIP(), []int{6}
}

func (x *UntappedCondition) GetType() ConditionType {
	if x != nil {
		return x.Type
	}
	return ConditionType_SHOCK_LAND
}

func (x *UntappedCondition) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type ManaCost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ColorRequirements []ManaColor `protobuf:"varint,1,rep,packed,name=colorRequirements,proto3,enum=manabase_simulation.ManaColor" json:"colorRequirements,omitempty"`
	GenericCost       int32       `protobuf:"varint,2,opt,name=genericCost,proto3" json:"genericCost,omitempty"`
}

func (x *ManaCost) Reset() {
	*x = ManaCost{}
	mi := &file_api_manabase_simulation_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ManaCost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManaCost) ProtoMessage() {}

func (x *ManaCost) ProtoReflect() protoreflect.Message {
	mi := &file_api_manabase_simulation_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManaCost.ProtoReflect.Descriptor instead.
func (*ManaCost) Descriptor() ([]byte, []int) {
	return file_api_manabase_simulation_proto_rawDescGZIP(), []int{7}
}

func (x *ManaCost) GetColorRequirements() []ManaColor {
	if x != nil {
		return x.ColorRequirements
	}
	return nil
}

func (x *ManaCost) GetGenericCost() int32 {
	if x != nil {
		return x.GenericCost
	}
	return 0
}

type Objective struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetTurn int32     `protobuf:"varint,1,opt,name=targetTurn,proto3" json:"targetTurn,omitempty"`
	ManaCosts  *ManaCost `protobuf:"bytes,2,opt,name=manaCosts,proto3" json:"manaCosts,omitempty"`
}

func (x *Objective) Reset() {
	*x = Objective{}
	mi := &file_api_manabase_simulation_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Objective) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Objective) ProtoMessage() {}

func (x *Objective) ProtoReflect() protoreflect.Message {
	mi := &file_api_manabase_simulation_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Objective.ProtoReflect.Descriptor instead.
func (*Objective) Descriptor() ([]byte, []int) {
	return file_api_manabase_simulation_proto_rawDescGZIP(), []int{8}
}

func (x *Objective) GetTargetTurn() int32 {
	if x != nil {
		return x.TargetTurn
	}
	return 0
}

func (x *Objective) GetManaCosts() *ManaCost {
	if x != nil {
		return x.ManaCosts
	}
	return nil
}

type GameConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InitialHandSize   int32 `protobuf:"varint,1,opt,name=initialHandSize,proto3" json:"initialHandSize,omitempty"`
	CardsDrawnPerTurn int32 `protobuf:"varint,2,opt,name=cardsDrawnPerTurn,proto3" json:"cardsDrawnPerTurn,omitempty"`
	OnThePlay         bool  `protobuf:"varint,3,opt,name=onThePlay,proto3" json:"onThePlay,omitempty"`
}

func (x *GameConfiguration) Reset() {
	*x = GameConfiguration{}
	mi := &file_api_manabase_simulation_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GameConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameConfiguration) ProtoMessage() {}

func (x *GameConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_api_manabase_simulation_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameConfiguration.ProtoReflect.Descriptor instead.
func (*GameConfiguration) Descriptor() ([]byte, []int) {
	return file_api_manabase_simulation_proto_rawDescGZIP(), []int{9}
}

func (x *GameConfiguration) GetInitialHandSize() int32 {
	if x != nil {
		return x.InitialHandSize
	}
	return 0
}

func (x *GameConfiguration) GetCardsDrawnPerTurn() int32 {
	if x != nil {
		return x.CardsDrawnPerTurn
	}
	return 0
}

func (x *GameConfiguration) GetOnThePlay() bool {
	if x != nil {
		return x.OnThePlay
	}
	return false
}

var File_api_manabase_simulation_proto protoreflect.FileDescriptor

var file_api_manabase_simulation_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2d, 0x73,
	0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x13, 0x6d, 0x61, 0x6e, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0xe4, 0x01, 0x0a, 0x13, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74,
	0x65, 0x44, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x39, 0x0a, 0x08,
	0x64, 0x65, 0x63, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x65, 0x63, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x08, 0x64,
	0x65, 0x63, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x54, 0x0a, 0x11, 0x67, 0x61, 0x6d, 0x65, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x73, 0x69,
	0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x11, 0x67, 0x61, 0x6d, 0x65,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3c, 0x0a,
	0x09, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x73, 0x69, 0x6d, 0x75,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x52, 0x09, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x22, 0x52, 0x0a, 0x14, 0x53,
	0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x44, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x0b, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x61, 0x74, 0x65, 0x22,
	0x75, 0x0a, 0x08, 0x44, 0x65, 0x63, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x05, 0x6c,
	0x61, 0x6e, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6d, 0x61, 0x6e,
	0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x4c, 0x61, 0x6e, 0x64, 0x52, 0x05, 0x6c, 0x61, 0x6e, 0x64, 0x73, 0x12, 0x38, 0x0a, 0x08,
	0x6e, 0x6f, 0x6e, 0x4c, 0x61, 0x6e, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4e, 0x6f, 0x6e, 0x4c, 0x61, 0x6e, 0x64, 0x52, 0x08, 0x6e, 0x6f,
	0x6e, 0x4c, 0x61, 0x6e, 0x64, 0x73, 0x22, 0xea, 0x02, 0x0a, 0x04, 0x4c, 0x61, 0x6e, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x36, 0x0a, 0x06, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x73,
	0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4d, 0x61, 0x6e, 0x61, 0x43, 0x6f,
	0x6c, 0x6f, 0x72, 0x52, 0x06, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x73, 0x54, 0x61, 0x70, 0x70, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0c, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x54, 0x61, 0x70, 0x70, 0x65, 0x64, 0x12,
	0x4b, 0x0a, 0x0e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x73,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x62, 0x61,
	0x73, 0x65, 0x5f, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x73, 0x74, 0x52, 0x0e, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x05,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x6d, 0x61,
	0x6e, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x4c, 0x61, 0x6e, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x05, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x12, 0x54, 0x0a, 0x11, 0x75, 0x6e, 0x74, 0x61, 0x70, 0x70, 0x65, 0x64, 0x43, 0x6f, 0x6e,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6d,
	0x61, 0x6e, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x55, 0x6e, 0x74, 0x61, 0x70, 0x70, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x64, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x11, 0x75, 0x6e, 0x74, 0x61, 0x70, 0x70, 0x65, 0x64, 0x43, 0x6f,
	0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x22, 0x7a, 0x0a, 0x07, 0x4e, 0x6f, 0x6e, 0x4c, 0x61, 0x6e, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x3f, 0x0a, 0x0b, 0x63, 0x61, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x73,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x62, 0x61,
	0x73, 0x65, 0x5f, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4d, 0x61,
	0x6e, 0x61, 0x43, 0x6f, 0x73, 0x74, 0x52, 0x0b, 0x63, 0x61, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x43,
	0x6f, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22,
	0x71, 0x0a, 0x0e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x66, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x6c, 0x69, 0x66, 0x65, 0x12, 0x3e, 0x0a, 0x08, 0x6d, 0x61, 0x6e, 0x61, 0x43, 0x6f, 0x73,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x62, 0x61,
	0x73, 0x65, 0x5f, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4d, 0x61,
	0x6e, 0x61, 0x43, 0x6f, 0x73, 0x74, 0x48, 0x00, 0x52, 0x08, 0x6d, 0x61, 0x6e, 0x61, 0x43, 0x6f,
	0x73, 0x74, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x43, 0x6f,
	0x73, 0x74, 0x22, 0x5f, 0x0a, 0x11, 0x55, 0x6e, 0x74, 0x61, 0x70, 0x70, 0x65, 0x64, 0x43, 0x6f,
	0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x36, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x5f, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x7a, 0x0a, 0x08, 0x4d, 0x61, 0x6e, 0x61, 0x43, 0x6f, 0x73, 0x74, 0x12,
	0x4c, 0x0a, 0x11, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x6d, 0x61, 0x6e,
	0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x4d, 0x61, 0x6e, 0x61, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x52, 0x11, 0x63, 0x6f, 0x6c, 0x6f,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x20, 0x0a,
	0x0b, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x43, 0x6f, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0b, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x43, 0x6f, 0x73, 0x74, 0x22,
	0x68, 0x0a, 0x09, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x54, 0x75, 0x72, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x54, 0x75, 0x72, 0x6e, 0x12, 0x3b, 0x0a, 0x09,
	0x6d, 0x61, 0x6e, 0x61, 0x43, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x73, 0x69, 0x6d, 0x75, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4d, 0x61, 0x6e, 0x61, 0x43, 0x6f, 0x73, 0x74, 0x52, 0x09,
	0x6d, 0x61, 0x6e, 0x61, 0x43, 0x6f, 0x73, 0x74, 0x73, 0x22, 0x89, 0x01, 0x0a, 0x11, 0x47, 0x61,
	0x6d, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x28, 0x0a, 0x0f, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x48, 0x61, 0x6e, 0x64, 0x53, 0x69,
	0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61,
	0x6c, 0x48, 0x61, 0x6e, 0x64, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x63, 0x61, 0x72,
	0x64, 0x73, 0x44, 0x72, 0x61, 0x77, 0x6e, 0x50, 0x65, 0x72, 0x54, 0x75, 0x72, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x63, 0x61, 0x72, 0x64, 0x73, 0x44, 0x72, 0x61, 0x77, 0x6e,
	0x50, 0x65, 0x72, 0x54, 0x75, 0x72, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x6e, 0x54, 0x68, 0x65,
	0x50, 0x6c, 0x61, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x6f, 0x6e, 0x54, 0x68,
	0x65, 0x50, 0x6c, 0x61, 0x79, 0x2a, 0x47, 0x0a, 0x08, 0x4c, 0x61, 0x6e, 0x64, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x4c, 0x41, 0x49, 0x4e, 0x53, 0x10, 0x00, 0x12, 0x0c, 0x0a,
	0x08, 0x4d, 0x4f, 0x55, 0x4e, 0x54, 0x41, 0x49, 0x4e, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x46,
	0x4f, 0x52, 0x45, 0x53, 0x54, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x49, 0x53, 0x4c, 0x41, 0x4e,
	0x44, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x53, 0x57, 0x41, 0x4d, 0x50, 0x10, 0x04, 0x2a, 0x4e,
	0x0a, 0x09, 0x4d, 0x61, 0x6e, 0x61, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x09, 0x0a, 0x05, 0x57,
	0x48, 0x49, 0x54, 0x45, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x42, 0x4c, 0x55, 0x45, 0x10, 0x01,
	0x12, 0x09, 0x0a, 0x05, 0x42, 0x4c, 0x41, 0x43, 0x4b, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x52,
	0x45, 0x44, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x47, 0x52, 0x45, 0x45, 0x4e, 0x10, 0x04, 0x12,
	0x0d, 0x0a, 0x09, 0x43, 0x4f, 0x4c, 0x4f, 0x52, 0x4c, 0x45, 0x53, 0x53, 0x10, 0x05, 0x2a, 0x3e,
	0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x0e, 0x0a, 0x0a, 0x53, 0x48, 0x4f, 0x43, 0x4b, 0x5f, 0x4c, 0x41, 0x4e, 0x44, 0x10, 0x00, 0x12,
	0x0d, 0x0a, 0x09, 0x46, 0x41, 0x53, 0x54, 0x5f, 0x4c, 0x41, 0x4e, 0x44, 0x10, 0x01, 0x12, 0x0e,
	0x0a, 0x0a, 0x43, 0x48, 0x45, 0x43, 0x4b, 0x5f, 0x4c, 0x41, 0x4e, 0x44, 0x10, 0x02, 0x32, 0x78,
	0x0a, 0x11, 0x4d, 0x61, 0x6e, 0x61, 0x62, 0x61, 0x73, 0x65, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61,
	0x74, 0x6f, 0x72, 0x12, 0x63, 0x0a, 0x0c, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x44,
	0x65, 0x63, 0x6b, 0x12, 0x28, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x73,
	0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61,
	0x74, 0x65, 0x44, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e,
	0x6d, 0x61, 0x6e, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x44, 0x65, 0x63, 0x6b,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2f, 0x5a, 0x2d, 0x6d, 0x61, 0x6e, 0x61,
	0x2d, 0x73, 0x69, 0x6d, 0x2e, 0x6c, 0x61, 0x74, 0x65, 0x6e, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x63,
	0x63, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2d, 0x73, 0x69, 0x6d, 0x75, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_manabase_simulation_proto_rawDescOnce sync.Once
	file_api_manabase_simulation_proto_rawDescData = file_api_manabase_simulation_proto_rawDesc
)

func file_api_manabase_simulation_proto_rawDescGZIP() []byte {
	file_api_manabase_simulation_proto_rawDescOnce.Do(func() {
		file_api_manabase_simulation_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_manabase_simulation_proto_rawDescData)
	})
	return file_api_manabase_simulation_proto_rawDescData
}

var file_api_manabase_simulation_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_api_manabase_simulation_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_api_manabase_simulation_proto_goTypes = []any{
	(LandType)(0),                // 0: manabase_simulation.LandType
	(ManaColor)(0),               // 1: manabase_simulation.ManaColor
	(ConditionType)(0),           // 2: manabase_simulation.ConditionType
	(*SimulateDeckRequest)(nil),  // 3: manabase_simulation.SimulateDeckRequest
	(*SimulateDeckResponse)(nil), // 4: manabase_simulation.SimulateDeckResponse
	(*DeckList)(nil),             // 5: manabase_simulation.DeckList
	(*Land)(nil),                 // 6: manabase_simulation.Land
	(*NonLand)(nil),              // 7: manabase_simulation.NonLand
	(*ActivationCost)(nil),       // 8: manabase_simulation.ActivationCost
	(*UntappedCondition)(nil),    // 9: manabase_simulation.UntappedCondition
	(*ManaCost)(nil),             // 10: manabase_simulation.ManaCost
	(*Objective)(nil),            // 11: manabase_simulation.Objective
	(*GameConfiguration)(nil),    // 12: manabase_simulation.GameConfiguration
}
var file_api_manabase_simulation_proto_depIdxs = []int32{
	5,  // 0: manabase_simulation.SimulateDeckRequest.deckList:type_name -> manabase_simulation.DeckList
	12, // 1: manabase_simulation.SimulateDeckRequest.gameConfiguration:type_name -> manabase_simulation.GameConfiguration
	11, // 2: manabase_simulation.SimulateDeckRequest.objective:type_name -> manabase_simulation.Objective
	6,  // 3: manabase_simulation.DeckList.lands:type_name -> manabase_simulation.Land
	7,  // 4: manabase_simulation.DeckList.nonLands:type_name -> manabase_simulation.NonLand
	1,  // 5: manabase_simulation.Land.colors:type_name -> manabase_simulation.ManaColor
	8,  // 6: manabase_simulation.Land.activationCost:type_name -> manabase_simulation.ActivationCost
	0,  // 7: manabase_simulation.Land.types:type_name -> manabase_simulation.LandType
	9,  // 8: manabase_simulation.Land.untappedCondition:type_name -> manabase_simulation.UntappedCondition
	10, // 9: manabase_simulation.NonLand.castingCost:type_name -> manabase_simulation.ManaCost
	10, // 10: manabase_simulation.ActivationCost.manaCost:type_name -> manabase_simulation.ManaCost
	2,  // 11: manabase_simulation.UntappedCondition.type:type_name -> manabase_simulation.ConditionType
	1,  // 12: manabase_simulation.ManaCost.colorRequirements:type_name -> manabase_simulation.ManaColor
	10, // 13: manabase_simulation.Objective.manaCosts:type_name -> manabase_simulation.ManaCost
	3,  // 14: manabase_simulation.ManabaseSimulator.SimulateDeck:input_type -> manabase_simulation.SimulateDeckRequest
	4,  // 15: manabase_simulation.ManabaseSimulator.SimulateDeck:output_type -> manabase_simulation.SimulateDeckResponse
	15, // [15:16] is the sub-list for method output_type
	14, // [14:15] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_api_manabase_simulation_proto_init() }
func file_api_manabase_simulation_proto_init() {
	if File_api_manabase_simulation_proto != nil {
		return
	}
	file_api_manabase_simulation_proto_msgTypes[5].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_manabase_simulation_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_manabase_simulation_proto_goTypes,
		DependencyIndexes: file_api_manabase_simulation_proto_depIdxs,
		EnumInfos:         file_api_manabase_simulation_proto_enumTypes,
		MessageInfos:      file_api_manabase_simulation_proto_msgTypes,
	}.Build()
	File_api_manabase_simulation_proto = out.File
	file_api_manabase_simulation_proto_rawDesc = nil
	file_api_manabase_simulation_proto_goTypes = nil
	file_api_manabase_simulation_proto_depIdxs = nil
}
