# conGo

Go package seeking to replace solidity as a way of developing safe and secure smart contracts for EVM based blockchains. 

    
    syntax : 
    
~keywords:
    called, 
    emit,
    event, 
    name, string
    metadata, = json,
    by [address or list of addresses] 

    address[type] -> address types = deployer, owner, user[allowed, restricted, blocked] , spender
        
        address_list_name := []address[allowed]{

        }


    IF( condition ) {} -> example : if (fn() called by address_list_name) {do this stuff}



    LOCK([lock logic]) //deposit token 
    UNLOCK([unlock logic])  //withdrawal token

    CALL([function name])




~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

type contractName contract {

        GO SYNTAXed CONTRACT

}

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

example : 
package main 

type myToken contract {
    event("eventName as first argument", ...)
    emit("emitName as first argument", ...)

    payable.Receive() 
}

