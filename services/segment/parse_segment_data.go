package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"runtime"

	"github.com/gocql/gocql"
	"github.com/qedus/osmpbf"
)

type Node struct {
	Lat float64
	Lon float64
}

func main() {
	// Open PBF file
	file, err := os.Open("northern-zone-latest.osm.pbf")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// First pass: decode all nodes
	d := osmpbf.NewDecoder(file)
	if err := d.Start(runtime.GOMAXPROCS(-1)); err != nil {
		log.Fatal(err)
	}
	nodeMap := make(map[int64]Node)
	fmt.Println("Reading nodes...")
	for {
		v, err := d.Decode()
		if err != nil {
			break
		}
		if n, ok := v.(*osmpbf.Node); ok {
			nodeMap[n.ID] = Node{Lat: n.Lat, Lon: n.Lon}
		}
	}

	// Rewind file and restart decoder
	file.Seek(0, 0)
	d = osmpbf.NewDecoder(file)
	if err := d.Start(runtime.GOMAXPROCS(-1)); err != nil {
		log.Fatal(err)
	}

	// Connect to Cassandra
	cluster := gocql.NewCluster("127.0.0.1") // use "cassandra" if connecting from another container
	cluster.Keyspace = "segment_keyspace"
	cluster.Consistency = gocql.Quorum
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatal("Failed to connect to Cassandra:", err)
	}
	defer session.Close()

	fmt.Println("Processing and inserting road segments...")
	for {
		v, err := d.Decode()
		if err != nil {
			break
		}
		way, ok := v.(*osmpbf.Way)
		if !ok || !isHighway(way.Tags) || len(way.NodeIDs) < 2 {
			continue
		}

		start, ok1 := nodeMap[way.NodeIDs[0]]
		end, ok2 := nodeMap[way.NodeIDs[len(way.NodeIDs)-1]]
		if !ok1 || !ok2 {
			continue
		}

		length := haversine(start.Lat, start.Lon, end.Lat, end.Lon)
		roadType := way.Tags["highway"]
		segmentID := fmt.Sprintf("%d", way.ID) // convert int64 to string
		err = session.Query(`
			INSERT INTO road_segments (segment_id, start_lat, start_lon, end_lat, end_lon, road_type, length_km)
			VALUES (?, ?, ?, ?, ?, ?, ?)`,
			segmentID, start.Lat, start.Lon, end.Lat, end.Lon, roadType, length,
		).Exec()
		if err != nil {
			log.Println("Insert failed:", err)
		}
	}

	fmt.Println("Done.")
}

func isHighway(tags map[string]string) bool {
	_, ok := tags["highway"]
	return ok
}

func haversine(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371 // Earth radius in km
	lat1Rad := lat1 * math.Pi / 180
	lat2Rad := lat2 * math.Pi / 180
	dlat := (lat2 - lat1) * math.Pi / 180
	dlon := (lon2 - lon1) * math.Pi / 180
	a := math.Sin(dlat/2)*math.Sin(dlat/2) +
		math.Cos(lat1Rad)*math.Cos(lat2Rad)*
			math.Sin(dlon/2)*math.Sin(dlon/2)
	return R * 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
}