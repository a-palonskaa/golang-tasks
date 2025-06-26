package reflect_spell

import (
	"reflect"
)

type Spell interface {
	Name() string
	Char() string
	Value() int
}

type CastReceiver interface {
	ReceiveSpell(s Spell)
}

func CastToAll(spell Spell, objects []interface{}) {
	for _, obj := range objects {
		CastTo(spell, obj)
	}
}

func CastTo(spell Spell, object interface{}) {
	if receiver, ok := object.(CastReceiver); ok {
		receiver.ReceiveSpell(spell)
		return
	}

	objectValue := reflect.ValueOf(object)
	if objectValue.Kind() != reflect.Ptr || objectValue.IsNil() {
		return
	}

	elem := objectValue.Elem()
	if elem.Kind() != reflect.Struct {
		return
	}

	fieldHealth := elem.FieldByName(spell.Char())
	if !fieldHealth.IsValid() || !fieldHealth.CanSet() {
		return
	}

	if fieldHealth.Kind() == reflect.Int {
		fieldHealth.SetInt(fieldHealth.Int() + int64(spell.Value()))
	}
}

type spell struct {
	name string
	char string
	val  int
}

func newSpell(name string, char string, val int) Spell {
	return &spell{name: name, char: char, val: val}
}

func (s spell) Name() string {
	return s.name
}

func (s spell) Char() string {
	return s.char
}

func (s spell) Value() int {
	return s.val
}

type Player struct {
	name   string
	health int
}

func (p *Player) ReceiveSpell(s Spell) {
	if s.Char() == "Health" {
		p.health += s.Value()
	}
}

type Zombie struct {
	Health int
}

type Daemon struct {
	Health int
}

type Orc struct {
	Health int
}

type Wall struct {
	Durability int
}
