package board

import "github.com/Dmitrev/gangsta-monopoly/player"

const (
	SpaceTypeStart    = iota
	SpaceTypeProperty
	SpaceTypeRailroad
	SpaceTypeJail
	SpaceTypeGotoJail
	SpaceTypeTax
	SpaceTypeChance
	SpaceTypeChest
	SpaceTypeFreeParking
)

type Space struct {
	Info  *SpaceInfo
	Owner *player.Player
}

type SpaceInfo struct {
	Name  string
	Price int
	Type  int
	Rent  int
}

type Board struct {
	Spaces []*Space
}

func NewSpace(name string, spaceType int, price int, rent int) *Space {
	return &Space{
		Info: &SpaceInfo{
			Name:  name,
			Price: price,
			Type:  spaceType,
			Rent:  rent,
		},
		Owner: nil,
	}
}
