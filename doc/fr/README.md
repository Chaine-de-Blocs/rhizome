# Rhizome la cryptomonaie meta

Rhizome a pour objectif de permettre la création d'une cryptomonnaie
à part entière.

Aujourd'hui, en 2018, il existe deux manières de créer une cryptomonnaie :
- Copier le code de [bitcoin](https://github.com/bitcoin/bitcoin) (fork) en modifiant certaines
parties dont le genesis block pour se séparer du réseau de Bitcoin.
- Créer un jeton depuis un contrat intelligent. Cette méthode, bien qu'en apparence
laisse penser que nous créons une nouvelle cryptomonnaie, dépend intégralement
du protocole qui régie le contrat intelligent. Elle est adaptée au usages simples.

On va d'ores et déjà rejeter la seconde méthode puisqu'elle ne consiste pas à créer une cryptomonnaie,
il ne s'agit que de valeurs gérées dans un contrat dont l'existence est limitée par
le protocole qui le régie (par exemple un contrat Ethereum dépend entièrement de Ethereum).

Il nous reste donc que la première solution. Bien qu'elle puisse parraître pratiquable de prime abord,
il s'avère qu'il y a des considérations à prendre en compte.

Le première est qu'en général, lorsqu'une personne décide de créer sa cryptomonnaie c'est
pour apporter quelque chose de nouveau au protocole, sinon autant utiliser Bitcoin.
Pour ce faire, il est alors nécessaire d'avoir une compréhension profonde du code
source de Bitcoin au risque de créer des failles rendant le protocole obsolète.

La seconde considération est le réseau. Bitcoin fonctionne avec un réseau décentralisé et
dépend de l'activité des noeuds : pas de noeuds, pas de Bitcoin. Alors, créer sa cryptomonnaie
doit nécessairement impliquer de motiver des acteurs pour maintenir ce nouveau
réseau décentralisé. Une nouvelle monnaie n'a en général aucune valeur, donc
les acteurs n'ont aucun intérêt à participer au réseau. Bitcoin a à ses débuts
rencontré aucun franc succès, l'utilisation assez faible du réseau a permis de
connaître une existence avec peu d'acteurs, bien que la récompense était
faible (en monnaie fiduciaire : euro, dollar, yuan). De plus Bitcoin a bénéficié de la primauté :
un nouveau moyen de gérer l'échange dans un réseau trustless.

L'objectif de ce projet est de produire une nouvelle cryptomonnaie modulaire,
sans affecter la topologie de son réseau décentralisé et basée
sur la structure de [Bitcoin](https://bitcoin.org/bitcoin.pdf).

L'implantation des différentes couches est basée sur la dernière version de
[Bitcoin](https://github.com/bitcoin/bitcoin) (0.15.1)

# Faciliter la compréhension

**Cryptomonnaie** : parfois utilisé comme synonyme à "protocole". C'est un terme plus
usuel pour parler d'un protocole qui gère un consensus avec des procédés
cryptographiques, et dont l'utilité est d'être une monnaie.

**Cryptomonnaie annexe** : cryptomonnaie créée par un utilisateur du projet, synonyme de "protocole annexe".

**Protocole** : ensemble de règles qui stipule comment un réseau doit communiquer. Cela
comprend la structure des blocs, la structure des transactions et la règle de
consensus qui permet de synchroniser le réseau décentralisé.

**Protocole annexe** : protocole créé par un utilisateur du projet.

**Protocole maître** : désigne le configuration de base du protocole du projet. Ce
sera le protocole utilisé dans le réseau maître.

**Réseau** : réseau décentralisé utilisant le protocole maître pour communiquer.

**Topologie de réseau** : architecture du réseau. La topologie évoque la localisation
des acteurs du réseau et leur capacité à diffuser les informations selon un protocole
commun. Voir la topologie comme une mapemonde avec des points qui sont les acteurs
et des liens qui sont leurs connexions.

**Tx** : désigne une transaction.

# Modularité

L'apport innovant de ce projet est d'apporter une modularité à la cryptomonnaie.
Un créateur de cryptomonnaie doit pouvoir définir sa cryptomonnaie,
avec plus ou moins de complexités et sans devoir modifier le code d'origine.
Foncièrement cela revient à se séparer du réseau maître.

La séparation n'est qu'éphémère. Le réseau maître recense les cryptomonnaies
nouvelles et permet aux acteurs de ce réseau de pouvoir être également
des acteurs sur chaque nouvelle cryptomonnaie, c'est à eux de choisir, le protocole maître se
charge juste de leur renseigner l'existence de cryptomonnaies annexes.
Cela permet de bénéficier d'une partie de la topologie du réseau maître.

La modularité se délimite ainsi :
- Paramètres,
- Algorithme de consensus,
- Structure des transactions,
- Structure des blocs,
- Scripts.

## Paramètres

Les paramètres servent de variables globales qui peuvent être utilisées dans l'algorithme
de consensus par exemple. Ils _peuvent_ être définis par le créateur du protocole annexe.

Liste des paramètres :
- Base de jetons émis par récompense
- Période de division de l'émission
- PoW limit
- PoW durée (temps pour miner un bloc)
- PoW période avant ajustement
- Port mainnet
- Port testnet
- Nom du token
- Ticker du token (ex: BTC)
- Préfixe des adresses

## Algorithme de consensus

L'algorithme de consensus _peut_ être écrit par le créateur du protocole annexe.

L'algorithme de consensus par défaut est le même implanté dans Bitcoin : la preuve
de travail.

## Structure des transactions

La structure des transactions _peut_ être modifiée par le créateur du protocole annexe.

Le créateur est apte à produire des scripts spécifiques à sa cryptomonnaie.

Les transactions par défaut comportent les attributs suivants :
- Input counter (_inchangeable_) : nombre d'inputs
- Inputs
  - Prev Tx (_inchangeable_) : hash de la transaction qui précède l'input
  - Signature : script de signature
- Output counter : nombre d'outputs
- Outputs
  - Value (_inchangeable_) : nombre positif entier (valeur envoyée à l'output)
  - Script : conditions pour être dépensée
- Champs additionnels...

## Structure des blocs et bloc génésis

La structure des blocs _peut_ être modifiée par le créateur du protocole annexe.

Les blocs par défaut comportent les attributs suivants :
- Header
  - Hash du bloc précédent (_inchangeable_)
  - Hash de la racine de l'arbre de Merkle des transactions (_inchangeable_)
  - Timestamp (_inchangeable_)
  - Bits : permet de définir la difficulté
  - Nonce : solution du bloc miné
- Transaction counter (_inchangeable_) : nombre de transactions
- Transactions (_inchangeable_)
 - Liste des transactions (_inchangeable_)
- Champs additionnels...

Le bloc génésis est le premier bloc au lancement d'une nouvelle cryptomonaie.

Pour chaque nouvelle cryptomonaie annexe un bloc génésis doit être généré. Les attributs
_inchangeables_ sont générés par le protocole maître sans que le créateur en ait
le contrôle. En revanche, les attributs modifiables et les champs additionnels
_peuvent_ être définis par le créateur à la génération du bloc génésis.

## Scripts

[Le wiki de Bitcoin](https://en.bitcoin.it/wiki/Script) définie le script de la manière suivante :

> Bitcoin utilise un système de script pout les transactions. Le script est simple, basé sur la pile, et procéduré de gauche à droite.
C'est volontairement pas turing complet, avec aucun moyen d'effectuer des boucles.

En d'autres termes, le script Bitcoin permet de programmer avec des transactions.

Par défaut toutes les fonctions intégrées par Bitcoin sont également dans cette
cryptomonnaie, ces fonctions sont appelées des "OP CODE". Donc un script
est un ensemble d'op codes.

Le créateur de protocole annexe _peut_ écrire ses op codes comme il le souhaite,
il pourrait même tenter de créer des scripts turing complet et produire des
contrats intelligents sur une base de cryptomonnaie similaire à Bitcoin, puisque
le protocole maître est proche de Bitcoin.

# Algorithme de consensus

Intégration de l'algorithme AlgoRand.

AlgoRand est un algorithme de consenus proposé par Silvo Micali dans le cadre du
MIT. La proposition tente de palier aux problèmes de la preuve de travail de Bitcoin
qui sont la consommation en électricité élevée et le temps de validation trop long.

Algorand n'intègre pas de récompense pour le travail des acteurs, dans sa version actuelle.

AlgoRand permet de produire un bloc toutes les 10 secondes sans que la sécurité soit
négligée. Rhizome a une production de bloc à raison de 1 bloc par 1 minute.

Nous verrons que Rhizome intègre un intérêt pour les acteurs à être actifs sur le réseau.
Cet aspect est crucial puisque la qualité du réseau réside dans le nombre d'acteurs
qui assurent la sécurité.

# Dictionnaire

L'une des partcularités du protocole maître est de recenser un dictionnaire des
protocoles annexes. Ce recensemment se fait par le biais de transactions correspondant
au dictionnaire, exactement de la même manière qu'opère Ethereum pour les transactions
qui concernent un contrat (création ou appel d'une méthode). Une entrée de dictionnaire
est alors une adresse publique.

## Entrée de dictionnaire (protocole annexe publié)

Une entrée inclut les informations suivantes :
- Metadonnées
  - Nom de jeton
  - Nom du créateur
  - Adresse publique du créateur
  - Essence
  - Maximum d'acteurs souhaité
  - % de redistribution par tour de participation
  - Description _optionnel_
- Bloc code
- Transaction code
- Commandes code
- OPCodes code
- Consensus algorithme code

Une entrée est créée par une transaction.

Lorsqu'un créateur publie son protocole annexe il doit inclure un montant RZM qu'on
appelle le carburant. Ce montant sert de récompense aux acteurs du protocole
annexe, expliqué en détail plus bas. Concrètement les jetons inclus ont pour destin
d'être détruits, la redistribution se fait sous forme de _coinbase_ comme le système
de récompense dans Bitcoin.

## Intérêt pour les acteurs

Le protocole annexe, selon sa conception produit également des jetons qui lui sont propres,
cela peut justifier d'un intérêt à être acteur sur ce protocole en dehors de Rhizome.
En d'autres termes, un protocole annexe peut perdurer sans dépendre du réseau
Rhizome, avec sa propre économie et ses acteurs.

Les acteurs sur le protocole annexe doivent prouver leur activité sur ce dernier, afin
d'être élligibles à la consommation des fonds de l'entrée du dictionnaire, recevoir
la coinbase. Etant donné que l'algorithme de consensus est défini par le créateur du protocole
annexe, il n'est alors pas possible de générer une élligibilité sans confiance (trustless)
à partir du protocole annexe. Les parties suivantes détaillent comment Rhizome
pallie à cette problématique.

### Tour de participation

Un tour de participation est une période où tous les acteurs déclarent participer
à x protocole. ils renseignent leur adresse IP avec leur adresse publique Rhizome,
ces informations sont recensées lors de la validation du tour de participation.
Un tour de participation par entrée de dictionnaire. Un tour de participation débute
tous les 59 blocs, après le tour précédent, ce qui correspond à 59 minutes.
Cela signifie qu'un protocole annexe, lorsqu'il est publié, n'est considéré
par le réseau qu'après 59 blocs qui suit sa publication. Cela laisse le temps
aux acteurs de le notifier et de s'inscrire comme participants au prochain tour.

### Déterminer l'activité des acteurs en fin de tour de participation

Avant le début d'un nouveau tour de participation un quorum d'acteurs est choisi
aléatoirement. Son rôle est de contrôler l'activité des participants.
L'activité est contrôlée en réalisant un `PING` vers les participants.
Un `PING` est lancé toutes les minutes et la réponse est recensée par chaque
acteur du quorum indépendemment. Une réponse compte 1 point pour le participant,
une absence de réponse apporte aucun point. En fin de tour de participation,
tous les acteurs du quorum envoient leur rapport aux validateurs.

Les validateurs pondèrent les points de chaque participant du tour afin de déterminer
quels participants ont été suffisamment actifs durant ce tour de participation.
On détermine l'activité suffisante selon le pseudocode suivant :

```
algorithme determiner-activite est
    entrée: Des rapports R
    sortie: Liste des participants déterminés actifs L

    On suppose P une liste vide telle que <adresse, points> où l'adresse
    est l'adresse publique du participant

    tant qu'il existe un rapport r dans R faire
        pour chaque entrée (adresse, points) dans r faire
            P[adresse] += points

    On suppose L une liste vide

    On suppose B := 59 qui correspond au nombre de point total maximal sur un rapport
    On support nb_R le nombre de rapport

    pour chaque entrée (adresse, points) dans P faire
        si (points / nb_R) > (B x 0.8)
            Ajouter adresse à L

    retourner L
```

On considère qu'un acteur a été suffisamment actif à +80% de réponse durant le
tour de participation.

Tous les validateurs doivent arriver à la même liste de participants, ainsi le consensus
peut-être atteint et la _coinbase_ peut être ditribuée à tous les participants
déterminés comme actifs durant le tour de participation. Le montant de la _coinbase_
est définie ainsi :

coinbase = (essence * (redistribution / 100)) / nb_participants

Où
- essence est la quantité de jeton injecté dans le dictionnaire au début du tour de
participation
- redistribution est le % de redistribution indiqué par le dictionnaire au début
du tour de participation
- nb_participants est le nombre de participant déterminés comme actifs

# Conclusion

Ce projet propose un échaffaudage à cryptomonnaies. De nombreuses personnes ressentent le besoin
de faire leur cryptomonnaie pour des besoins particuliers ou même pour s'amuser à
avoir sa cryptomonnaie.

En plus de l'utilité économique au même titre que Bitcoin, avec l'ancienneté en moins,
ce projet s'adresse également à l'académie et à la recherche qui pourrait avoir
besoin de mettre à l'épreuve des algorithmes de consensus. Cette cryptomonnaie
est un gain de temps considérable dans la mesure où elle propose de construire
sa cryptomonnaie et de bénéficier d'une topologie de réseau déjà acquise.
