
utilities for parsing OpenStreetMap PBF files and extracting geographic data

### installation

> tested on go version go1.6

```bash
$ go get github.com/missinglink/pbf
```

### disclaimer

this repo contains a bunch of experiments including a monkey-patched parsing library, it should be considered as permanently in development.

please feel free to fork and contribute, at some stage it might reach a level of maturity and stability where it could be published as a more general-purpose tool.

### running the cli

```bash
$ pbf --help
```

### cli commands

the following commands are included in the CLI binary

#### stats

output some statistic about the entities contained in the file.

> beta: output info about internal block structures

```bash
$ pbf stats paris_france.osm.pbf

Nodes: 18789281
Ways: 3092821
Relations: 36632
NodesWithName: 79214
WaysWithName: 194404
RelationsWithName: 16078
... etc
```

#### json

convert to overpass json format, optionally using bitmask to filter elements.

```bash
$ pbf json london_england.osm.pbf | head -n10

{"id":1796210057,"type":"node","lat":51.914658,"lon":-0.0151269995}
{"id":1796210058,"type":"node","lat":51.919197,"lon":-0.0117898}
{"id":1796210059,"type":"node","lat":51.91228,"lon":-0.008095}
{"id":1796210060,"type":"node","lat":51.911575,"lon":-0.008256}
{"id":1796210061,"type":"node","lat":51.911762,"lon":-0.0073439996}
{"id":1796210062,"type":"node","lat":51.91978,"lon":-0.0113665}
{"id":1796210064,"type":"node","lat":51.916912,"lon":-0.0092377}
{"id":1796210066,"type":"node","lat":51.9156,"lon":-0.011646}
{"id":1796210067,"type":"node","lat":51.919483,"lon":-0.0058625997}
{"id":1796210069,"type":"node","lat":51.915524,"lon":-0.0169407}
```

#### xml

convert to osm xml format, optionally using bitmask to filter elements.

```bash
$ pbf xml london_england.osm.pbf | head -n10

<?xml version="1.0" encoding="UTF-8"?>
<osm version="0.6" generator="missinglink/pbf">
	<node id="540042355" lat="51.010666" lon="-0.817476">
	</node>
	<node id="540042360" lat="51.008556" lon="-0.815922">
	</node>
	<node id="540042383" lat="51.020397" lon="-0.864640">
	</node>
	<node id="540042386" lat="51.021404" lon="-0.866511">
	</node>
```

#### opl

convert to [opl format](http://osmcode.org/opl-file-format/), optionally using bitmask to filter elements.

```bash
$ pbf opl london_england.osm.pbf | head -n10

n1244022050 T x-1.0038857 y51.4467049
n1244022051 T x-1.0000237 y51.4489517
n1244022053 T x-1.0018616 y51.4508781
n1244022054 T x-0.9954746 y51.4506073
n1244022055 T x-1.0035510 y51.4477654
n1244022056 T x-1.0020980 y51.4490585
n1244022057 T x-0.9958197 y51.4497719
n1244022060 T x-1.0022424 y51.4519882
n1244022062 T x-0.9992446 y51.4492874
n1244022063 T x-1.0026369 y51.4486008
```

#### cypher

convert to neo4j cypher format, optionally using bitmask to filter elements.

> beta: this is incomplete and probably needs some more work to be useful

```bash
$ pbf cypher london_england.osm.pbf | head -n10

CREATE (N540042355:Element:Node {});
CREATE (N540042360:Element:Node {});
CREATE (N540042383:Element:Node {});
CREATE (N540042386:Element:Node {});
CREATE (N540042399:Element:Node {});
```

#### sqlite3

import elements in to a sqlite3 database, optionally using bitmask to filter elements.

```bash
$ pbf sqlite3 london_england.osm.pbf london.db
```

see the source code for more info about the database schema

```bash
sqlite3 london.db "SELECT * FROM RELATION_TAGS LIMIT 5"

20663|name|Southern Walkway
20663|type|route
20663|route|foot
20663|network|lwn
20664|type|route
```

#### leveldb

import elements in to a leveldb database, optionally using bitmask to filter elements.

```bash
$ mkdir leveldir

$ pbf leveldb london_england.osm.pbf leveldir
```

see the source code for the encoding format

```bash
$ ls leveldir/ | head -n10
000023.ldb
000024.ldb
000026.ldb
000027.ldb
000028.ldb
000029.ldb
000030.ldb
000039.ldb
000040.ldb
000041.ldb
```

... more to come
