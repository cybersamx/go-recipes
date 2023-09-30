# Inheritance

Go doesn't support inheritance. But we can achieve similar effects using composition.

## Improvement

The example follows a pattern that is typical in many object-oriented programming languages. A more canonical Go way of implementing this is using strictly composition.

Consider do this.
   
1. We determine some key characteristics and behaviors that define what a Pet does. While the details differ from one species to another all pets have these behaviors in common: breathe and eat. And we can declare interfaces `Breather` and `Eater` respectively to presents those behaviors.

   ```go
   type Breather interface {
        Breathe()
   }
   
   type Eater interface {
        Eat()
   }
   ```

   And Go recommends that one-method interfaces be named by the method name + "er" suffix. See [interface name](https://go.dev/doc/effective_go#interface_names) in the official [Effective Go write-up](https://go.dev/doc/effective_go).

1. We define concrete structs that implement those interfaces.

   ```go
   type SkinBreather struct {}
   
   func (sb SkinBreather) Breathe {
        fmt.Println("breathe through skin")
   }
   
   type LungBreather struct {}
   
   func (lb LungBreather) Breathe {
        fmt.Println("breathe through lung")
   }
   
   type Herbivore struct {}
   
   func (h Herbivore) Eat {
        fmt.Println("eat plants")
   }
   
   type Carnivore struct {}
   
   func (c Carnivore) Eat {
        fmt.Println("eat animals")
   }
   ```

1. Finally define the pet structs.

   ```go
   // Instead of thinking Frog IS a skin breather and a herbivore. Think Forg HAS
   // skin breathing and plant eating capabilities.
   type Frog type {
        SkinBreather
        Herbivore
   }
   
   type Dog type {
        LungBreather
        Carnivore
   }
   
   type Turtle type {
        LungBreather
        Herbivore
   }
   ```
