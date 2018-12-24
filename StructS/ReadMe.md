
Pointer Notes :- 

   1. AddOfMemory := &VarName :- If "&" Mentioned before variable then it will return actual address of variable where it store in RAM of machine.
    
      	* pointerAnil := &anil
    	* pointerAnil.updateData("Neel")

   2. (*ValueOfPointer).variableName :- If "*" Mentioned then it will return actual struct value from AddOfMemory to performe further action.
        * 	(*pointerToPerson).firstName = firstName

    
If  "*Types" mentioned in receiver then it will only accessible by PointerOftype

for e.g.

    package main

    import "fmt"

    type contactinfo struct {
        emailid   string
        postalZip int
    }

    type person struct {
        firstName string
        lastName  string
        contactinfo
    }

    func main() {

        anil := person{
            firstName: "Anil",
            lastName:  "Anand",
            contactinfo: contactinfo{
                emailid:   "Anil@xinfin.org",
                postalZip: 400067,
            },
        }
        pointerAnil := &anil
        fmt.Println(pointerAnil)
        pointerAnil.updateData("Neel")
        anil.print()
    }

    func (p person) print() {

        fmt.Printf("%+v", p)
    }

    func (pointerToPerson *person) updateData(firstName string) {
        (*pointerToPerson).firstName = firstName
        fmt.Println((*pointerToPerson))
    }
