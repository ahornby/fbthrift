// Autogenerated by Thrift for thrift/compiler/test/fixtures/optionals/src/module.thrift
//
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
//  @generated

package module

import (
    "fmt"
    "strings"

    thrift "github.com/facebook/fbthrift/thrift/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction)
var _ = fmt.Printf
var _ = strings.Split
var _ = thrift.ZERO


type PersonID = int64

func NewPersonID() PersonID {
    return 0
}

func WritePersonID(item PersonID, p thrift.Encoder) error {
    if err := p.WriteI64(item); err != nil {
    return err
}
    return nil
}

func ReadPersonID(p thrift.Decoder) (PersonID, error) {
    var decodeResult PersonID
    decodeErr := func() error {
        result, err := p.ReadI64()
if err != nil {
    return err
}
        decodeResult = result
        return nil
    }()
    return decodeResult, decodeErr
}

type Animal int32

const (
    Animal_DOG Animal = 1
    Animal_CAT Animal = 2
    Animal_TARANTULA Animal = 3
)

// Enum value maps for Animal
var (
    AnimalToName = map[Animal]string {
        Animal_DOG: "DOG",
        Animal_CAT: "CAT",
        Animal_TARANTULA: "TARANTULA",
    }

    AnimalToValue = map[string]Animal {
        "DOG": Animal_DOG,
        "CAT": Animal_CAT,
        "TARANTULA": Animal_TARANTULA,
    }
)

func (x Animal) String() string {
    if v, ok := AnimalToName[x]; ok {
        return v
    }
    return "<UNSET>"
}

func (x Animal) Ptr() *Animal {
    return &x
}

// Deprecated: Use AnimalToValue instead (e.g. `x, ok := AnimalToValue["name"]`).
func AnimalFromString(s string) (Animal, error) {
    if v, ok := AnimalToValue[s]; ok {
        return v, nil
    }
    return Animal(0), fmt.Errorf("not a valid Animal string")
}


type Color struct {
    Red float64 `thrift:"red,1" json:"red" db:"red"`
    Green float64 `thrift:"green,2" json:"green" db:"green"`
    Blue float64 `thrift:"blue,3" json:"blue" db:"blue"`
    Alpha float64 `thrift:"alpha,4" json:"alpha" db:"alpha"`
}
// Compile time interface enforcer
var _ thrift.Struct = (*Color)(nil)

func NewColor() *Color {
    return (&Color{}).
        SetRedNonCompat(0.0).
        SetGreenNonCompat(0.0).
        SetBlueNonCompat(0.0).
        SetAlphaNonCompat(0.0)
}

func (x *Color) GetRed() float64 {
    return x.Red
}

func (x *Color) GetGreen() float64 {
    return x.Green
}

func (x *Color) GetBlue() float64 {
    return x.Blue
}

func (x *Color) GetAlpha() float64 {
    return x.Alpha
}

func (x *Color) SetRedNonCompat(value float64) *Color {
    x.Red = value
    return x
}

func (x *Color) SetRed(value float64) *Color {
    x.Red = value
    return x
}

func (x *Color) SetGreenNonCompat(value float64) *Color {
    x.Green = value
    return x
}

func (x *Color) SetGreen(value float64) *Color {
    x.Green = value
    return x
}

func (x *Color) SetBlueNonCompat(value float64) *Color {
    x.Blue = value
    return x
}

func (x *Color) SetBlue(value float64) *Color {
    x.Blue = value
    return x
}

func (x *Color) SetAlphaNonCompat(value float64) *Color {
    x.Alpha = value
    return x
}

func (x *Color) SetAlpha(value float64) *Color {
    x.Alpha = value
    return x
}

func (x *Color) writeField1(p thrift.Encoder) error {  // Red
    if err := p.WriteFieldBegin("red", thrift.DOUBLE, 1); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.Red
    if err := p.WriteDouble(item); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *Color) writeField2(p thrift.Encoder) error {  // Green
    if err := p.WriteFieldBegin("green", thrift.DOUBLE, 2); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.Green
    if err := p.WriteDouble(item); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *Color) writeField3(p thrift.Encoder) error {  // Blue
    if err := p.WriteFieldBegin("blue", thrift.DOUBLE, 3); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.Blue
    if err := p.WriteDouble(item); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *Color) writeField4(p thrift.Encoder) error {  // Alpha
    if err := p.WriteFieldBegin("alpha", thrift.DOUBLE, 4); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.Alpha
    if err := p.WriteDouble(item); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *Color) readField1(p thrift.Decoder) error {  // Red
    result, err := p.ReadDouble()
if err != nil {
    return err
}

    x.Red = result
    return nil
}

func (x *Color) readField2(p thrift.Decoder) error {  // Green
    result, err := p.ReadDouble()
if err != nil {
    return err
}

    x.Green = result
    return nil
}

func (x *Color) readField3(p thrift.Decoder) error {  // Blue
    result, err := p.ReadDouble()
if err != nil {
    return err
}

    x.Blue = result
    return nil
}

func (x *Color) readField4(p thrift.Decoder) error {  // Alpha
    result, err := p.ReadDouble()
if err != nil {
    return err
}

    x.Alpha = result
    return nil
}

func (x *Color) toString1() string {  // Red
    return fmt.Sprintf("%v", x.Red)
}

func (x *Color) toString2() string {  // Green
    return fmt.Sprintf("%v", x.Green)
}

func (x *Color) toString3() string {  // Blue
    return fmt.Sprintf("%v", x.Blue)
}

func (x *Color) toString4() string {  // Alpha
    return fmt.Sprintf("%v", x.Alpha)
}



func (x *Color) Write(p thrift.Encoder) error {
    if err := p.WriteStructBegin("Color"); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", x), err)
    }

    if err := x.writeField1(p); err != nil {
        return err
    }

    if err := x.writeField2(p); err != nil {
        return err
    }

    if err := x.writeField3(p); err != nil {
        return err
    }

    if err := x.writeField4(p); err != nil {
        return err
    }

    if err := p.WriteFieldStop(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", x), err)
    }

    if err := p.WriteStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", x), err)
    }
    return nil
}

func (x *Color) Read(p thrift.Decoder) error {
    if _, err := p.ReadStructBegin(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read error: ", x), err)
    }

    for {
        _, wireType, id, err := p.ReadFieldBegin()
        if err != nil {
            return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", x, id), err)
        }

        if wireType == thrift.STOP {
            break;
        }

        switch {
        case (id == 1 && wireType == thrift.Type(thrift.DOUBLE)):  // red
            if err := x.readField1(p); err != nil {
                return err
            }
        case (id == 2 && wireType == thrift.Type(thrift.DOUBLE)):  // green
            if err := x.readField2(p); err != nil {
                return err
            }
        case (id == 3 && wireType == thrift.Type(thrift.DOUBLE)):  // blue
            if err := x.readField3(p); err != nil {
                return err
            }
        case (id == 4 && wireType == thrift.Type(thrift.DOUBLE)):  // alpha
            if err := x.readField4(p); err != nil {
                return err
            }
        default:
            if err := p.Skip(wireType); err != nil {
                return err
            }
        }

        if err := p.ReadFieldEnd(); err != nil {
            return err
        }
    }

    if err := p.ReadStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", x), err)
    }

    return nil
}

func (x *Color) String() string {
    if x == nil {
        return "<nil>"
    }

    var sb strings.Builder

    sb.WriteString("Color({")
    sb.WriteString(fmt.Sprintf("Red:%s ", x.toString1()))
    sb.WriteString(fmt.Sprintf("Green:%s ", x.toString2()))
    sb.WriteString(fmt.Sprintf("Blue:%s ", x.toString3()))
    sb.WriteString(fmt.Sprintf("Alpha:%s", x.toString4()))
    sb.WriteString("})")

    return sb.String()
}

type Vehicle struct {
    Color *Color `thrift:"color,1" json:"color" db:"color"`
    LicensePlate *string `thrift:"licensePlate,2,optional" json:"licensePlate,omitempty" db:"licensePlate"`
    Description *string `thrift:"description,3,optional" json:"description,omitempty" db:"description"`
    Name *string `thrift:"name,4,optional" json:"name,omitempty" db:"name"`
    HasAC *bool `thrift:"hasAC,5,optional" json:"hasAC,omitempty" db:"hasAC"`
}
// Compile time interface enforcer
var _ thrift.Struct = (*Vehicle)(nil)

func NewVehicle() *Vehicle {
    return (&Vehicle{}).
        SetColorNonCompat(*NewColor()).
        SetHasACNonCompat(false)
}

func (x *Vehicle) GetColor() *Color {
    if !x.IsSetColor() {
        return nil
    }

    return x.Color
}

func (x *Vehicle) GetLicensePlate() string {
    if !x.IsSetLicensePlate() {
        return ""
    }

    return *x.LicensePlate
}

func (x *Vehicle) GetDescription() string {
    if !x.IsSetDescription() {
        return ""
    }

    return *x.Description
}

func (x *Vehicle) GetName() string {
    if !x.IsSetName() {
        return ""
    }

    return *x.Name
}

func (x *Vehicle) GetHasAC() bool {
    if !x.IsSetHasAC() {
        return false
    }

    return *x.HasAC
}

func (x *Vehicle) SetColorNonCompat(value Color) *Vehicle {
    x.Color = &value
    return x
}

func (x *Vehicle) SetColor(value *Color) *Vehicle {
    x.Color = value
    return x
}

func (x *Vehicle) SetLicensePlateNonCompat(value string) *Vehicle {
    x.LicensePlate = &value
    return x
}

func (x *Vehicle) SetLicensePlate(value *string) *Vehicle {
    x.LicensePlate = value
    return x
}

func (x *Vehicle) SetDescriptionNonCompat(value string) *Vehicle {
    x.Description = &value
    return x
}

func (x *Vehicle) SetDescription(value *string) *Vehicle {
    x.Description = value
    return x
}

func (x *Vehicle) SetNameNonCompat(value string) *Vehicle {
    x.Name = &value
    return x
}

func (x *Vehicle) SetName(value *string) *Vehicle {
    x.Name = value
    return x
}

func (x *Vehicle) SetHasACNonCompat(value bool) *Vehicle {
    x.HasAC = &value
    return x
}

func (x *Vehicle) SetHasAC(value bool) *Vehicle {
    x.HasAC = &value
    return x
}

func (x *Vehicle) IsSetColor() bool {
    return x != nil && x.Color != nil
}

func (x *Vehicle) IsSetLicensePlate() bool {
    return x != nil && x.LicensePlate != nil
}

func (x *Vehicle) IsSetDescription() bool {
    return x != nil && x.Description != nil
}

func (x *Vehicle) IsSetName() bool {
    return x != nil && x.Name != nil
}

func (x *Vehicle) IsSetHasAC() bool {
    return x != nil && x.HasAC != nil
}

func (x *Vehicle) writeField1(p thrift.Encoder) error {  // Color
    if !x.IsSetColor() {
        return nil
    }

    if err := p.WriteFieldBegin("color", thrift.STRUCT, 1); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.Color
    if err := item.Write(p); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *Vehicle) writeField2(p thrift.Encoder) error {  // LicensePlate
    if !x.IsSetLicensePlate() {
        return nil
    }

    if err := p.WriteFieldBegin("licensePlate", thrift.STRING, 2); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := *x.LicensePlate
    if err := p.WriteString(item); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *Vehicle) writeField3(p thrift.Encoder) error {  // Description
    if !x.IsSetDescription() {
        return nil
    }

    if err := p.WriteFieldBegin("description", thrift.STRING, 3); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := *x.Description
    if err := p.WriteString(item); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *Vehicle) writeField4(p thrift.Encoder) error {  // Name
    if !x.IsSetName() {
        return nil
    }

    if err := p.WriteFieldBegin("name", thrift.STRING, 4); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := *x.Name
    if err := p.WriteString(item); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *Vehicle) writeField5(p thrift.Encoder) error {  // HasAC
    if !x.IsSetHasAC() {
        return nil
    }

    if err := p.WriteFieldBegin("hasAC", thrift.BOOL, 5); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := *x.HasAC
    if err := p.WriteBool(item); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *Vehicle) readField1(p thrift.Decoder) error {  // Color
    result := *NewColor()
err := result.Read(p)
if err != nil {
    return err
}

    x.Color = &result
    return nil
}

func (x *Vehicle) readField2(p thrift.Decoder) error {  // LicensePlate
    result, err := p.ReadString()
if err != nil {
    return err
}

    x.LicensePlate = &result
    return nil
}

func (x *Vehicle) readField3(p thrift.Decoder) error {  // Description
    result, err := p.ReadString()
if err != nil {
    return err
}

    x.Description = &result
    return nil
}

func (x *Vehicle) readField4(p thrift.Decoder) error {  // Name
    result, err := p.ReadString()
if err != nil {
    return err
}

    x.Name = &result
    return nil
}

func (x *Vehicle) readField5(p thrift.Decoder) error {  // HasAC
    result, err := p.ReadBool()
if err != nil {
    return err
}

    x.HasAC = &result
    return nil
}

func (x *Vehicle) toString1() string {  // Color
    return fmt.Sprintf("%v", x.Color)
}

func (x *Vehicle) toString2() string {  // LicensePlate
    if x.IsSetLicensePlate() {
        return fmt.Sprintf("%v", *x.LicensePlate)
    }
    return fmt.Sprintf("%v", x.LicensePlate)
}

func (x *Vehicle) toString3() string {  // Description
    if x.IsSetDescription() {
        return fmt.Sprintf("%v", *x.Description)
    }
    return fmt.Sprintf("%v", x.Description)
}

func (x *Vehicle) toString4() string {  // Name
    if x.IsSetName() {
        return fmt.Sprintf("%v", *x.Name)
    }
    return fmt.Sprintf("%v", x.Name)
}

func (x *Vehicle) toString5() string {  // HasAC
    if x.IsSetHasAC() {
        return fmt.Sprintf("%v", *x.HasAC)
    }
    return fmt.Sprintf("%v", x.HasAC)
}

// Deprecated: Use NewVehicle().GetColor() instead.
func (x *Vehicle) DefaultGetColor() *Color {
    if !x.IsSetColor() {
        return NewColor()
    }
    return x.Color
}







func (x *Vehicle) Write(p thrift.Encoder) error {
    if err := p.WriteStructBegin("Vehicle"); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", x), err)
    }

    if err := x.writeField1(p); err != nil {
        return err
    }

    if err := x.writeField2(p); err != nil {
        return err
    }

    if err := x.writeField3(p); err != nil {
        return err
    }

    if err := x.writeField4(p); err != nil {
        return err
    }

    if err := x.writeField5(p); err != nil {
        return err
    }

    if err := p.WriteFieldStop(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", x), err)
    }

    if err := p.WriteStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", x), err)
    }
    return nil
}

func (x *Vehicle) Read(p thrift.Decoder) error {
    if _, err := p.ReadStructBegin(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read error: ", x), err)
    }

    for {
        _, wireType, id, err := p.ReadFieldBegin()
        if err != nil {
            return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", x, id), err)
        }

        if wireType == thrift.STOP {
            break;
        }

        switch {
        case (id == 1 && wireType == thrift.Type(thrift.STRUCT)):  // color
            if err := x.readField1(p); err != nil {
                return err
            }
        case (id == 2 && wireType == thrift.Type(thrift.STRING)):  // licensePlate
            if err := x.readField2(p); err != nil {
                return err
            }
        case (id == 3 && wireType == thrift.Type(thrift.STRING)):  // description
            if err := x.readField3(p); err != nil {
                return err
            }
        case (id == 4 && wireType == thrift.Type(thrift.STRING)):  // name
            if err := x.readField4(p); err != nil {
                return err
            }
        case (id == 5 && wireType == thrift.Type(thrift.BOOL)):  // hasAC
            if err := x.readField5(p); err != nil {
                return err
            }
        default:
            if err := p.Skip(wireType); err != nil {
                return err
            }
        }

        if err := p.ReadFieldEnd(); err != nil {
            return err
        }
    }

    if err := p.ReadStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", x), err)
    }

    return nil
}

func (x *Vehicle) String() string {
    if x == nil {
        return "<nil>"
    }

    var sb strings.Builder

    sb.WriteString("Vehicle({")
    sb.WriteString(fmt.Sprintf("Color:%s ", x.toString1()))
    sb.WriteString(fmt.Sprintf("LicensePlate:%s ", x.toString2()))
    sb.WriteString(fmt.Sprintf("Description:%s ", x.toString3()))
    sb.WriteString(fmt.Sprintf("Name:%s ", x.toString4()))
    sb.WriteString(fmt.Sprintf("HasAC:%s", x.toString5()))
    sb.WriteString("})")

    return sb.String()
}

type Person struct {
    Id PersonID `thrift:"id,1" json:"id" db:"id"`
    Name string `thrift:"name,2" json:"name" db:"name"`
    Age *int16 `thrift:"age,3,optional" json:"age,omitempty" db:"age"`
    Address *string `thrift:"address,4,optional" json:"address,omitempty" db:"address"`
    FavoriteColor *Color `thrift:"favoriteColor,5,optional" json:"favoriteColor,omitempty" db:"favoriteColor"`
    Friends []PersonID `thrift:"friends,6,optional" json:"friends,omitempty" db:"friends"`
    BestFriend *PersonID `thrift:"bestFriend,7,optional" json:"bestFriend,omitempty" db:"bestFriend"`
    PetNames map[Animal]string `thrift:"petNames,8,optional" json:"petNames,omitempty" db:"petNames"`
    AfraidOfAnimal *Animal `thrift:"afraidOfAnimal,9,optional" json:"afraidOfAnimal,omitempty" db:"afraidOfAnimal"`
    Vehicles []*Vehicle `thrift:"vehicles,10,optional" json:"vehicles,omitempty" db:"vehicles"`
}
// Compile time interface enforcer
var _ thrift.Struct = (*Person)(nil)

func NewPerson() *Person {
    return (&Person{}).
        SetIdNonCompat(NewPersonID()).
        SetNameNonCompat("")
}

func (x *Person) GetId() PersonID {
    return x.Id
}

func (x *Person) GetName() string {
    return x.Name
}

func (x *Person) GetAge() int16 {
    if !x.IsSetAge() {
        return 0
    }

    return *x.Age
}

func (x *Person) GetAddress() string {
    if !x.IsSetAddress() {
        return ""
    }

    return *x.Address
}

func (x *Person) GetFavoriteColor() *Color {
    if !x.IsSetFavoriteColor() {
        return nil
    }

    return x.FavoriteColor
}

func (x *Person) GetFriends() []PersonID {
    if !x.IsSetFriends() {
        return make([]PersonID, 0)
    }

    return x.Friends
}

func (x *Person) GetBestFriend() PersonID {
    if !x.IsSetBestFriend() {
        return NewPersonID()
    }

    return *x.BestFriend
}

func (x *Person) GetPetNames() map[Animal]string {
    if !x.IsSetPetNames() {
        return make(map[Animal]string)
    }

    return x.PetNames
}

func (x *Person) GetAfraidOfAnimal() Animal {
    if !x.IsSetAfraidOfAnimal() {
        return 0
    }

    return *x.AfraidOfAnimal
}

func (x *Person) GetVehicles() []*Vehicle {
    if !x.IsSetVehicles() {
        return make([]*Vehicle, 0)
    }

    return x.Vehicles
}

func (x *Person) SetIdNonCompat(value PersonID) *Person {
    x.Id = value
    return x
}

func (x *Person) SetId(value PersonID) *Person {
    x.Id = value
    return x
}

func (x *Person) SetNameNonCompat(value string) *Person {
    x.Name = value
    return x
}

func (x *Person) SetName(value string) *Person {
    x.Name = value
    return x
}

func (x *Person) SetAgeNonCompat(value int16) *Person {
    x.Age = &value
    return x
}

func (x *Person) SetAge(value *int16) *Person {
    x.Age = value
    return x
}

func (x *Person) SetAddressNonCompat(value string) *Person {
    x.Address = &value
    return x
}

func (x *Person) SetAddress(value *string) *Person {
    x.Address = value
    return x
}

func (x *Person) SetFavoriteColorNonCompat(value Color) *Person {
    x.FavoriteColor = &value
    return x
}

func (x *Person) SetFavoriteColor(value *Color) *Person {
    x.FavoriteColor = value
    return x
}

func (x *Person) SetFriendsNonCompat(value []PersonID) *Person {
    x.Friends = value
    return x
}

func (x *Person) SetFriends(value []PersonID) *Person {
    x.Friends = value
    return x
}

func (x *Person) SetBestFriendNonCompat(value PersonID) *Person {
    x.BestFriend = &value
    return x
}

func (x *Person) SetBestFriend(value *PersonID) *Person {
    x.BestFriend = value
    return x
}

func (x *Person) SetPetNamesNonCompat(value map[Animal]string) *Person {
    x.PetNames = value
    return x
}

func (x *Person) SetPetNames(value map[Animal]string) *Person {
    x.PetNames = value
    return x
}

func (x *Person) SetAfraidOfAnimalNonCompat(value Animal) *Person {
    x.AfraidOfAnimal = &value
    return x
}

func (x *Person) SetAfraidOfAnimal(value *Animal) *Person {
    x.AfraidOfAnimal = value
    return x
}

func (x *Person) SetVehiclesNonCompat(value []*Vehicle) *Person {
    x.Vehicles = value
    return x
}

func (x *Person) SetVehicles(value []*Vehicle) *Person {
    x.Vehicles = value
    return x
}

func (x *Person) IsSetAge() bool {
    return x != nil && x.Age != nil
}

func (x *Person) IsSetAddress() bool {
    return x != nil && x.Address != nil
}

func (x *Person) IsSetFavoriteColor() bool {
    return x != nil && x.FavoriteColor != nil
}

func (x *Person) IsSetFriends() bool {
    return x != nil && x.Friends != nil
}

func (x *Person) IsSetBestFriend() bool {
    return x != nil && x.BestFriend != nil
}

func (x *Person) IsSetPetNames() bool {
    return x != nil && x.PetNames != nil
}

func (x *Person) IsSetAfraidOfAnimal() bool {
    return x != nil && x.AfraidOfAnimal != nil
}

func (x *Person) IsSetVehicles() bool {
    return x != nil && x.Vehicles != nil
}

func (x *Person) writeField1(p thrift.Encoder) error {  // Id
    if err := p.WriteFieldBegin("id", thrift.I64, 1); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.Id
    err := WritePersonID(item, p)
if err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *Person) writeField2(p thrift.Encoder) error {  // Name
    if err := p.WriteFieldBegin("name", thrift.STRING, 2); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.Name
    if err := p.WriteString(item); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *Person) writeField3(p thrift.Encoder) error {  // Age
    if !x.IsSetAge() {
        return nil
    }

    if err := p.WriteFieldBegin("age", thrift.I16, 3); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := *x.Age
    if err := p.WriteI16(item); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *Person) writeField4(p thrift.Encoder) error {  // Address
    if !x.IsSetAddress() {
        return nil
    }

    if err := p.WriteFieldBegin("address", thrift.STRING, 4); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := *x.Address
    if err := p.WriteString(item); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *Person) writeField5(p thrift.Encoder) error {  // FavoriteColor
    if !x.IsSetFavoriteColor() {
        return nil
    }

    if err := p.WriteFieldBegin("favoriteColor", thrift.STRUCT, 5); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.FavoriteColor
    if err := item.Write(p); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *Person) writeField6(p thrift.Encoder) error {  // Friends
    if !x.IsSetFriends() {
        return nil
    }

    if err := p.WriteFieldBegin("friends", thrift.SET, 6); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.Friends
    if err := p.WriteSetBegin(thrift.I64, len(item)); err != nil {
    return thrift.PrependError("error writing set begin: ", err)
}
for _, v := range item {
    {
        item := v
        err := WritePersonID(item, p)
if err != nil {
    return err
}
    }
}
if err := p.WriteSetEnd(); err != nil {
    return thrift.PrependError("error writing set end: ", err)
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *Person) writeField7(p thrift.Encoder) error {  // BestFriend
    if !x.IsSetBestFriend() {
        return nil
    }

    if err := p.WriteFieldBegin("bestFriend", thrift.I64, 7); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := *x.BestFriend
    err := WritePersonID(item, p)
if err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *Person) writeField8(p thrift.Encoder) error {  // PetNames
    if !x.IsSetPetNames() {
        return nil
    }

    if err := p.WriteFieldBegin("petNames", thrift.MAP, 8); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.PetNames
    if err := p.WriteMapBegin(thrift.I32, thrift.STRING, len(item)); err != nil {
    return thrift.PrependError("error writing map begin: ", err)
}
for k, v := range item {
    {
        item := k
        if err := p.WriteI32(int32(item)); err != nil {
    return err
}
    }

    {
        item := v
        if err := p.WriteString(item); err != nil {
    return err
}
    }
}
if err := p.WriteMapEnd(); err != nil {
    return thrift.PrependError("error writing map end: ", err)
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *Person) writeField9(p thrift.Encoder) error {  // AfraidOfAnimal
    if !x.IsSetAfraidOfAnimal() {
        return nil
    }

    if err := p.WriteFieldBegin("afraidOfAnimal", thrift.I32, 9); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := *x.AfraidOfAnimal
    if err := p.WriteI32(int32(item)); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *Person) writeField10(p thrift.Encoder) error {  // Vehicles
    if !x.IsSetVehicles() {
        return nil
    }

    if err := p.WriteFieldBegin("vehicles", thrift.LIST, 10); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.Vehicles
    if err := p.WriteListBegin(thrift.STRUCT, len(item)); err != nil {
    return thrift.PrependError("error writing list begin: ", err)
}
for _, v := range item {
    {
        item := v
        if err := item.Write(p); err != nil {
    return err
}
    }
}
if err := p.WriteListEnd(); err != nil {
    return thrift.PrependError("error writing list end: ", err)
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *Person) readField1(p thrift.Decoder) error {  // Id
    result, err := ReadPersonID(p)
if err != nil {
    return err
}

    x.Id = result
    return nil
}

func (x *Person) readField2(p thrift.Decoder) error {  // Name
    result, err := p.ReadString()
if err != nil {
    return err
}

    x.Name = result
    return nil
}

func (x *Person) readField3(p thrift.Decoder) error {  // Age
    result, err := p.ReadI16()
if err != nil {
    return err
}

    x.Age = &result
    return nil
}

func (x *Person) readField4(p thrift.Decoder) error {  // Address
    result, err := p.ReadString()
if err != nil {
    return err
}

    x.Address = &result
    return nil
}

func (x *Person) readField5(p thrift.Decoder) error {  // FavoriteColor
    result := *NewColor()
err := result.Read(p)
if err != nil {
    return err
}

    x.FavoriteColor = &result
    return nil
}

func (x *Person) readField6(p thrift.Decoder) error {  // Friends
    _ /* elemType */, size, err := p.ReadSetBegin()
if err != nil {
    return thrift.PrependError("error reading set begin: ", err)
}

setResult := make([]PersonID, 0, size)
for i := 0; i < size; i++ {
    var elem PersonID
    {
        result, err := ReadPersonID(p)
if err != nil {
    return err
}
        elem = result
    }
    setResult = append(setResult, elem)
}

if err := p.ReadSetEnd(); err != nil {
    return thrift.PrependError("error reading set end: ", err)
}
result := setResult

    x.Friends = result
    return nil
}

func (x *Person) readField7(p thrift.Decoder) error {  // BestFriend
    result, err := ReadPersonID(p)
if err != nil {
    return err
}

    x.BestFriend = &result
    return nil
}

func (x *Person) readField8(p thrift.Decoder) error {  // PetNames
    _ /* keyType */, _ /* valueType */, size, err := p.ReadMapBegin()
if err != nil {
    return thrift.PrependError("error reading map begin: ", err)
}

mapResult := make(map[Animal]string, size)
for i := 0; i < size; i++ {
    var key Animal
    {
        enumResult, err := p.ReadI32()
if err != nil {
    return err
}
result := Animal(enumResult)
        key = result
    }

    var value string
    {
        result, err := p.ReadString()
if err != nil {
    return err
}
        value = result
    }

    mapResult[key] = value
}

if err := p.ReadMapEnd(); err != nil {
    return thrift.PrependError("error reading map end: ", err)
}
result := mapResult

    x.PetNames = result
    return nil
}

func (x *Person) readField9(p thrift.Decoder) error {  // AfraidOfAnimal
    enumResult, err := p.ReadI32()
if err != nil {
    return err
}
result := Animal(enumResult)

    x.AfraidOfAnimal = &result
    return nil
}

func (x *Person) readField10(p thrift.Decoder) error {  // Vehicles
    _ /* elemType */, size, err := p.ReadListBegin()
if err != nil {
    return thrift.PrependError("error reading list begin: ", err)
}

listResult := make([]*Vehicle, 0, size)
for i := 0; i < size; i++ {
    var elem Vehicle
    {
        result := *NewVehicle()
err := result.Read(p)
if err != nil {
    return err
}
        elem = result
    }
    listResult = append(listResult, &elem)
}

if err := p.ReadListEnd(); err != nil {
    return thrift.PrependError("error reading list end: ", err)
}
result := listResult

    x.Vehicles = result
    return nil
}

func (x *Person) toString1() string {  // Id
    return fmt.Sprintf("%v", x.Id)
}

func (x *Person) toString2() string {  // Name
    return fmt.Sprintf("%v", x.Name)
}

func (x *Person) toString3() string {  // Age
    if x.IsSetAge() {
        return fmt.Sprintf("%v", *x.Age)
    }
    return fmt.Sprintf("%v", x.Age)
}

func (x *Person) toString4() string {  // Address
    if x.IsSetAddress() {
        return fmt.Sprintf("%v", *x.Address)
    }
    return fmt.Sprintf("%v", x.Address)
}

func (x *Person) toString5() string {  // FavoriteColor
    return fmt.Sprintf("%v", x.FavoriteColor)
}

func (x *Person) toString6() string {  // Friends
    return fmt.Sprintf("%v", x.Friends)
}

func (x *Person) toString7() string {  // BestFriend
    if x.IsSetBestFriend() {
        return fmt.Sprintf("%v", *x.BestFriend)
    }
    return fmt.Sprintf("%v", x.BestFriend)
}

func (x *Person) toString8() string {  // PetNames
    return fmt.Sprintf("%v", x.PetNames)
}

func (x *Person) toString9() string {  // AfraidOfAnimal
    if x.IsSetAfraidOfAnimal() {
        return fmt.Sprintf("%v", *x.AfraidOfAnimal)
    }
    return fmt.Sprintf("%v", x.AfraidOfAnimal)
}

func (x *Person) toString10() string {  // Vehicles
    return fmt.Sprintf("%v", x.Vehicles)
}



// Deprecated: Use NewPerson().GetFavoriteColor() instead.
func (x *Person) DefaultGetFavoriteColor() *Color {
    if !x.IsSetFavoriteColor() {
        return NewColor()
    }
    return x.FavoriteColor
}





func (x *Person) Write(p thrift.Encoder) error {
    if err := p.WriteStructBegin("Person"); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", x), err)
    }

    if err := x.writeField1(p); err != nil {
        return err
    }

    if err := x.writeField2(p); err != nil {
        return err
    }

    if err := x.writeField3(p); err != nil {
        return err
    }

    if err := x.writeField4(p); err != nil {
        return err
    }

    if err := x.writeField5(p); err != nil {
        return err
    }

    if err := x.writeField6(p); err != nil {
        return err
    }

    if err := x.writeField7(p); err != nil {
        return err
    }

    if err := x.writeField8(p); err != nil {
        return err
    }

    if err := x.writeField9(p); err != nil {
        return err
    }

    if err := x.writeField10(p); err != nil {
        return err
    }

    if err := p.WriteFieldStop(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", x), err)
    }

    if err := p.WriteStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", x), err)
    }
    return nil
}

func (x *Person) Read(p thrift.Decoder) error {
    if _, err := p.ReadStructBegin(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read error: ", x), err)
    }

    for {
        _, wireType, id, err := p.ReadFieldBegin()
        if err != nil {
            return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", x, id), err)
        }

        if wireType == thrift.STOP {
            break;
        }

        switch {
        case (id == 1 && wireType == thrift.Type(thrift.I64)):  // id
            if err := x.readField1(p); err != nil {
                return err
            }
        case (id == 2 && wireType == thrift.Type(thrift.STRING)):  // name
            if err := x.readField2(p); err != nil {
                return err
            }
        case (id == 3 && wireType == thrift.Type(thrift.I16)):  // age
            if err := x.readField3(p); err != nil {
                return err
            }
        case (id == 4 && wireType == thrift.Type(thrift.STRING)):  // address
            if err := x.readField4(p); err != nil {
                return err
            }
        case (id == 5 && wireType == thrift.Type(thrift.STRUCT)):  // favoriteColor
            if err := x.readField5(p); err != nil {
                return err
            }
        case (id == 6 && wireType == thrift.Type(thrift.SET)):  // friends
            if err := x.readField6(p); err != nil {
                return err
            }
        case (id == 7 && wireType == thrift.Type(thrift.I64)):  // bestFriend
            if err := x.readField7(p); err != nil {
                return err
            }
        case (id == 8 && wireType == thrift.Type(thrift.MAP)):  // petNames
            if err := x.readField8(p); err != nil {
                return err
            }
        case (id == 9 && wireType == thrift.Type(thrift.I32)):  // afraidOfAnimal
            if err := x.readField9(p); err != nil {
                return err
            }
        case (id == 10 && wireType == thrift.Type(thrift.LIST)):  // vehicles
            if err := x.readField10(p); err != nil {
                return err
            }
        default:
            if err := p.Skip(wireType); err != nil {
                return err
            }
        }

        if err := p.ReadFieldEnd(); err != nil {
            return err
        }
    }

    if err := p.ReadStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", x), err)
    }

    return nil
}

func (x *Person) String() string {
    if x == nil {
        return "<nil>"
    }

    var sb strings.Builder

    sb.WriteString("Person({")
    sb.WriteString(fmt.Sprintf("Id:%s ", x.toString1()))
    sb.WriteString(fmt.Sprintf("Name:%s ", x.toString2()))
    sb.WriteString(fmt.Sprintf("Age:%s ", x.toString3()))
    sb.WriteString(fmt.Sprintf("Address:%s ", x.toString4()))
    sb.WriteString(fmt.Sprintf("FavoriteColor:%s ", x.toString5()))
    sb.WriteString(fmt.Sprintf("Friends:%s ", x.toString6()))
    sb.WriteString(fmt.Sprintf("BestFriend:%s ", x.toString7()))
    sb.WriteString(fmt.Sprintf("PetNames:%s ", x.toString8()))
    sb.WriteString(fmt.Sprintf("AfraidOfAnimal:%s ", x.toString9()))
    sb.WriteString(fmt.Sprintf("Vehicles:%s", x.toString10()))
    sb.WriteString("})")

    return sb.String()
}

// RegisterTypes registers types found in this file that have a thrift_uri with the passed in registry.
func RegisterTypes(registry interface {
  RegisterType(name string, initializer func() any)
}) {

}
