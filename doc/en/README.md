# WIP

The english translation is work in progress and someone is working on ([Jorropo](https://github.com/Jorropo/))

# Rhizome the meta cryptocurrency

Rhizome as for goal to allow the creation of full autonomous cryptocurrency.

Today, in year 2018, there is 2 way to create cryptocurrency :
- Copy the code of an cryptocurrency like [Bitcoin](https://github.com/bitcoin/bitcoin) or [Ethereum](https://github.com/ethereum/go-ethereum) (fork) and edit some value like genesis block, network ID (for [Cryptonote](https://cryptonote.org/)), ... Changing these value will create an new network. You can also change the main code like PoW algorithm (this is what initially [litecoin](https://github.com/litecoin-project/litecoin) does with [Bitcoin](https://github.com/bitcoin/bitcoin))
- Create a token on an existing blockchain. But this way doesn't create an cryptocurrency, this token is dependent of main cryptocurrency. You can easly do that with [ethereum](https://ethereum.org/token).

So we will reject this second option cause we can't use it for create inovating or interesting cryptocurrency. (In fact this way is used in the most of case for a presale of an future cryptocurrency.)

The first solution is very easy to do, you only need basic programation notion. But if you do that you create a shitcoin (a fork of cryptocurrency that doesn't add somethings except existence), and your coin will be used for only 2 things cause why not use directly the original cryptocurrency with more security, support, user, ... :
- Trading
- Mining (If your coin is over evaluated by trader)
So for create interesting cryptocurrency you will need to change things in core of your fork, but this can be really difficult (it depend of what you would do with, change block time is simple, but add smart contract to a coin without is very difficult).
And you need to be carefull cause you can maybe create security problem.

The second problem is the network, when I write this document [bitcoin](https://github.com/bitcoin/bitcoin) have ~ 10k nodes and ethereum ~ 17k nodes.
That allow a very big bandwith and redundancy for the blockchain.
The other problem is 51% attack, the hashrate avaible in cryptocurrency is very important, more a cryptocurrency have hashrate more the cryptocurrency is safe.
If someone (pool, ...) have 51% or more of the total hashrate he can :
- Roll back in time
- Ban some address (reject evry block containing tx of this address)
- Make an exclude mining (reject all the block not mined by you, that make a slower network but you totaly kill other pool and force miner to stop mine or came mining on your pool)
- Make an extremly slow network
You think getting 51% of a coin is impossible ?
With [nicehash](https://www.nicehash.com/) for ~ 500$ you can get 80% for hashrate of little altcoin and ~ 4000$ for big altcoin (some altcoin like [Monero](https://github.com/monero-project/monero) isn't touched cause [Monero](https://github.com/monero-project/monero) have own safe hashing algo and not a lot people is mining cryptonight V7 on nicehash, so creating your algo or use very exotic algo can be a solution (but require very advanced mathematical knowledge and expose you to use an unsafe hashing algo)).

So creating a new cryptocurrency will require one things and you can't pass over :
**Make people use your cryptocurrency !**
You will need miner mining your cryptocurrency and pool, exchange and people hosting full node (a full node is node serving the blockchain or a node serving the whole blockchain).

The goal of **Rhizome** is to produce an new modular cryptocurrency, that can be used for create new cryptocurrency without affecting the decentralized network.
**Rhizome** is an spiritual descendant of [Bitcoin](https://github.com/bitcoin/bitcoin) (0.15.1)

# Glossary

This lexicon is not general and is valid only in this context.

**Cryptocurrency** : Is à protocol and some data, like genesis block, block time, ...
Exemple : [Monero](https://github.com/monero-project/monero) is the cryptocurrency and [Cryptonote](https://cryptonote.org/) is the protocol.
Sometimes people use cryptocurrency for say protocol.

**Secondary Cryptocurrency** (when talking of a cryptocurrency or cryptocurrency protocol) : Is a cryptocurrency or cryptocurrency protocol based on an other.
Exemple : [Litecoin](https://github.com/litecoin-project/litecoin) is a fork of [Bitcoin](https://github.com/bitcoin/bitcoin)

**Node** : A computer running the program that allows to access and interact with the blockchain.

**Full Node** : A node which is accessible from the WAN (a node which is directly on the internet or which redirects the traffic of the ports of connection to itself is with the upnp or a static redirection), that allows the other nodes to contact it to obtain blockchain or send tx, it allows to have a better responsiveness in the link to the network and to contribute to it.

**Boot Node** : This is a special node that is often maintained by the main developers, this node does not necessarily serve the blockchain but is used as a meeting point by newcomers.

**Protocol** : Set of rules that states how a network should communicate. This includes the block structure, the transaction structure, and the consensus rule that synchronizes the decentralized network.
Exemple : [Cryptonote](https://cryptonote.org/) is a protocol.

**Secondary Protocol** : A Protocol created by a user of **Rhizome**.

**Master Protocol** : The main protocol of **Rhizome**.

**Master Network** : Decentralized network using the master protocol.

**Network Topology** : All the nodes.

# Modularity

The innovation of **Rhizome** is to bring a modularity to cryptocurrency.
A creator of cryptocurrency must be able to define his cryptocurrency, with more or less complexities and without having to modify the original code.
This will create a split in the master network

The separation is only ephemeral.
The master network lists the new cryptocurrencies and allows the actors of this network to also be able to act on each new cryptocurrency, it is up to them to choose, the master protocol to take care of them to learn about the existence of cryptocurrencies. This makes it possible to benefit from a part of the topology of the master network.

The modularity is delimited as follows :
- Settings
- consensus algorithm
- Transaction structure
- Structure of the blocks
- Scripts

## Settings

The parameters serve as global variables that can be used in the consensus algorithm for example.
They _can_ be defined by the creator of the secondary protocol.

Settings list :
- Base of tokens issued by reward
- Period of division of the issue
- PoW limit
- PoW duration (time to mine a block)
- PoW period before adjustment
- Default mainnet port
- Default testnet port
- Token name
- Ticker of the token (ex: BTC)
- Address prefix

## Consensus Algorithm

The consensus algorithm _can_ be written by the creator of the secondary protocol.
The default consensus algorithm is the same implemented in Bitcoin: the proof
work.

## Transaction structure

The transaction structure _can_ be modified by the creator of the secondary protocol.

The creator is able to produce scripts specific to his cryptocurrency.

All value flaged _absolute_ can be changed.

By default transactions have the following attributes:
- Input counter (_absolute_): number of inputs
- Inputs
  - Prev Tx (_absolute_): hash of the transaction that precedes the input
  - Signature: signature script
- Output counter: number of outputs
- Outputs
  - Value (_absolute_): integer positive number (value sent to the output)
  - Script: conditions to be spent
- Additional fields ...

## Block structure and genesis block

The block structure _can_ be modified by the creator of the secondary protocol.

The default blocks have the following attributes:
- Header
   - Hash of the previous block (_absolute_)
   - Hash of the Merkle tree root (_absolute_)
   - Timestamp (_absolute_)
   - Bits : define the difficulty
   - Nonce : solution of the mined block
- Transaction counter (_absolute_): number of transactions
- Transactions (_absolute_)
  - List of transactions (_absolute_)
- Additional fields ...

The genesis block is the first block to launch a new cryptocurrency.

For each new secondary cryptocurrency a genesis block must be generated. Attributes _absolute_ are generated by the master protocol without the creator opinion. But editable attributes and additional fields _can_ be defined by the creator at the generation of the genesis block.

## Scripts

[Bitcoin's](https://en.bitcoin.it/wiki/Script) define a script as the following way :

> Bitcoin uses a scripting system for transactions. The script is simple, stack-based, and procedurally from left to right.
It is deliberately not complete turing, with no way to perform loops.

In other words, the Bitcoin script allows to execute code with transactions.

By default all the functions integrated by Bitcoin are also in this cryptocurrency, these functions are called "OP CODE". So a script is a set of op codes.

The protocol creator appendix _can_ write its op codes as it wishes, it could even create complete turing scripts and produce smart contracts on a cryptocurrency similar to Bitcoin, since the master protocol is close to Bitcoin.
