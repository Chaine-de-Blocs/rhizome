# WIP

The english translation is work in progress and doesn't someone is working on (@Jorropo)

# Rhizome the meta cryptocurrency

Rhizome as for goal to allow the creation of full autonomous cryptocurrency.

Today, in year 2018, there is 2 way to create cryptocurrency :
- Copy the code of an cryptocurrency like [Bitcoin](https://github.com/bitcoin/bitcoin) or [Ethereum](https://github.com/ethereum/go-ethereum) (fork) and edit some value like genesis block, network ID (for cryptonote), ... Changing these value will create an new network. You can also change the main code like PoW algorithm (this is what initially [litecoin](https://github.com/litecoin-project/litecoin) does with [Bitcoin](https://github.com/bitcoin/bitcoin))
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
With [nicehash](https://www.nicehash.com/) for ~ 500$ you can get 80% for hashrate of little altcoin and ~ 4000$ for big altcoin (some altcoin like [monero](https://github.com/monero-project/monero) isn't touched cause monero have own safe hashing algo and not a lot people is mining cryptonight V7 on nicehash, so creating your algo or use very exotic algo can be a solution (but require very advanced mathematical knowledge and expose you to use an unsafe hashing algo)).

So creating a new cryptocurrency will require one things and you can't pass over :
**Make people use your cryptocurrency !**
You will need miner mining your cryptocurrency and pool, exchange and people hosting full node (a full node is node serving the blockchain or a node serving the whole blockchain).

The goal of **Rhizome** is to produce an new modular cryptocurrency, that can be used for create new cryptocurrency without affecting the decentralized network.
**Rhizome** is an spiritual descendant of [Bitcoin](https://github.com/bitcoin/bitcoin) (0.15.1)
