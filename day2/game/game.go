package main

import "fmt"

/*
	Struct definition
	type {Struct Name} struct

	A struct can have fields and methods
	Remember to have Capitalized fields and methods if you want to
*/
type Item struct {
	X int
	Y int
}

type Player struct {
	Name string
	Item // Embed Item struct into player struct. This is called embedding
	Keys []Key
}

func (p *Player) FoundKey(k Key) error {
	if k < Jade || k > invalidKey {
		return fmt.Errorf("invalid key: %#v", k)
	}

	if !containsKey(p.Keys, k) {
		p.Keys = append(p.Keys, k)
	}
	return nil
}

func containsKey(keys []Key, k Key) bool {
	for _, k2 := range keys {
		if k2 == k {
			return true
		}
	}
	return false
}

/*
	This is go's version of an enum
*/
const (
	// Key is a type here of byte
	// iota is like an incrementer, so jade will get 1, Copper 2 and Crystal 3
	Jade Key = iota + 1
	Copper
	Crystal
	invalidKey // unexporetd
)

type Key byte

/*
	IN go you can write a function for how you want an object to look like as a String.

	When this String method exists for the receiver of K, it will use this method
	whenever a print is called on the Key type like fmt.Println(k)

	Since we have defined a string method on the Key type, This Key type is now implementing the
	Stringer interface
*/
func (k Key) String() string {
	switch k {
	case Jade:
		return "jade"
	case Copper:
		return "copper"
	case Crystal:
		return "crystal"
	}

	// 	return fmt.Sprintf("<Key %s>", k) will cause compiling warning since %s is infinitely checking if k has a string method and calling it again and again
	return fmt.Sprintf("<Key %d>", k)
}

/*
	This is how you define an interface.

	this interface has a single function

	This means that any type that has a Move function on it implements this interface.
	If yuou want a type to implement an interface just make sure the type has the function which the interface enforces.

	In our main function in this file, p1, i1 to i3 have the Move function,
	Either by mebedding or otherwise.const

	Using interfaces:
	1. You can create an interface and have it be implemented by types. Then you can just specify
	the interface in the arguments and any type that implements that interface can be passed into that function.
	This also gives the flexibility for ther function to get multiple types and offer decoupling of interfaces from concrete types.

	2. You can have your types implemnet interfaces which will cause Go to look at your code deifferently, as in
	if your type impplement an interface, it will run the String method when the object is being printed.


	interface tell you what we want from a type to do, not what an interface provides to the type like other languages

	Make sure the interfaces are small, meaning require less methods to be implemented by types.
	Smaller interfaces mean better abstraction and fl;exibility, more types can implement the interfaces
*/
type mover interface {
	Move(x, y int)
}

/*
	moveALl moves everything in the mover to the x and y

	Mover here is a list of mover's which adhere to the interface mover
	A rule of thumb for functions is that they shouyld accept interfaces and return types
*/
func moveAll(ms []mover, x int, y int) {
	for _, m := range ms {
		m.Move(x, y)
	}
}

/*
	This is an example of Item having a method.
	i Item is the receiver of the function Move
	This is how we define methods on structs

	But keep in mind that go has created a copy of the Item struct and then called Move on the copy
	So that means when we did i.X = x , the copy is being modified.

	Go passes everything by value, including the receivers.
	So to have the changes that took place on a receiover to be reflected outside of the function we need
	make the receiver be a pointer to the struct.
	If you want mutate the struct you have to use a pointer receiver.


	Rule of thumb for pointer / value receivers
	1. If the reciever is built in than use value receivers
	2. If you want to mutate the receiver pass a pointer in
	3. Whatever you pick stick to the approach throughout. Dont mix and match value and pointer receivers

	The questiopn ask here to yourself when you want to determine the above is that
	whether do you want to share your copy or not share the copy with the function.
	Think about ownership.
*/
func (i *Item) Move(x, y int) {
	/*
		i here is the Item struct
	*/
	i.X = x
	i.Y = y
}

const (
	maxX = 1000
	maxY = 600
)

func main() {
	var i1 Item
	fmt.Println("Item 1 is: ", i1)
	fmt.Printf("Item 1 is %#v\n: ", i1)

	i2 := Item{1, 2} // This constructor must be passed all fields of the struct
	fmt.Println("Item 2 is: ", i2)

	i3 := Item{
		X: 3,
		Y: 4,
	}
	fmt.Println("Item 3 is: ", i3)
	fmt.Println(NewItem(10, 20))
	fmt.Println(NewItem(10, -20))

	/*
		This is how you call a method on the struct
		But remember how the receiver in the Move method is a pointer to Item struct
		and here we call it on the value. This is fine and this a nice thin offered by the go compiler
	*/
	i3.Move(100, 200)
	fmt.Println("Item 3 after move is: ", i3)

	p1 := Player{
		Name: "Parzival",
		Item: Item{500, 500},
		Keys: []Key{Jade, Crystal, Copper},
	}

	fmt.Printf("p1 is: %#v\n", p1)
	/*
		Below p1.X is possible through and because of embedding
		p1 does not have an X field but since the Item struct is embedded in the player
		struct the player struct has access to it fields

		All fields of the embedded type are available in the top level type which is embedding the embedded. All embedded type fields are lifted up to the t9op level.
		If there is a conflict betrween the top level type having the same field name as the embedded type

		2 embedded types having the same field names can be fine if youi dont use the p1.X notation. This will cause
		the go compiler to complain that p1.X is ambiguous selector since the X is present ion 2 embedded types
	*/
	fmt.Printf("p1.X is: %#v\n", p1.X)
	fmt.Printf("p1.Item.X is: %#v\n", p1.Item.X) // This is also possible but not needed unless you have shadowed fields in the embedded type

	/*
		Keep in mind that this trait where the embedding type can access the embedded types,
		fields and methods is not inheritence, this is embedding.

		p1.Move is still getting a pointer to the Item struct and not getting a player.
		The move function does not know which player is the move being called on.
		Atleast with the current setup.
	*/
	p1.Move(400, 600)                  // This is also possible through embedding
	fmt.Printf("p1.X is: %#v\n", p1.X) // This is also possible but not needed unless you have shadowed fields in the embedded type

	/*
		This is creating a slice of movers (objects of types that have a Move function on them)
		Notice that we are passing pointers to those objects.
		This is because the move method has a pointer receiver
	*/
	ms := []mover{
		&i3,
		&p1,
	}

	moveAll(ms, 0, 0)
	for _, m := range ms {
		fmt.Println("m is: ", m)
	}

	k := Jade
	fmt.Println("k is: ", k)
	p1.FoundKey(Jade)
	fmt.Println("Keys: ", p1.Keys)
	p1.FoundKey(Jade)
	fmt.Println("Keys: ", p1.Keys)
}

/*
	NewItem returns a pointer of type Item and an error. The name NewItem follows the function naming convention of
	New{StructName}

	When a program allocates memory it eitehr allocates it on the stack or it does it on the heap.
	Stack allocations are for local variables and dont outlive the function
	Heap allocations are for objects and structs that out live the function
*/
func NewItem(X, Y int) (*Item, error) {
	if X < 0 || X > maxX || Y < 0 || Y > maxY {
		// nil here is fopr the pointer since we encountered an error
		return nil, fmt.Errorf("%d/%d out of bounds %d/%d", X, Y, maxX, maxY)
	}
	/*
		it is hard for us to determine that if the structs field is 0 because of uninitialized
		or was it set to 0 manually
	*/
	i := Item{
		X: X,
		Y: Y,
	}
	/*
		&i returns the pointer to the i struct

		The go compiler do3es "escape analysis" and will allocate i on teh heap instead of the stack
		This is because i will outlive the function.
	*/
	return &i, nil
}
