#Considerations
1.Performance : As the problem states strings are large, it's assumed it's paramount. Having said that, code 
shall be easy to understand, use, and extend. So an equilibrium must be found, attending to this the following decisions
have been taken:

* Not adopting a collection library, as only an Stack needs to be implemented, it makes sense to implement a simple stack
that's re-usable, no generic and extendable. Although if more collections are needed a library can be chosen and compared 
with existing bench test.
*Create BenchMark test

##Test
Test, are enough to ensure the functionality, still implementation the Stack can be pass somehow to the validation 
function allowing to use mocks. In case strict unit test are somehow required, but it's not considered now as it 
will complicate the validator usage.

##Considerations
An example of the function use can be found in main.go, it can be executed as follows :

```console
jcabrera@Jorges-MacBook-Pro-3 portainer % go run main.go example.json
valid definition file:example.json%                                                                                                                                                                                        jcabrera@Jorges-MacBook-Pro-3 portainer % go run main.go fail_example.json
invalid definition file:fail_example.json%  
```