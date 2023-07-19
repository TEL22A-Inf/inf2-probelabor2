package aufgabe4

/* ERREICHBARE PUNKTE: 20 */

// In der Datei graph.go finden Sie eine Implementierung eines Knotens in einem Graphen.
// Vervollständigen Sie die Methode ReachableNodes.
// Hinweis: Sie können die Hilfsfunktion Contains verwenden.

/***HILFSFUNKTIONEN***/
func Contains(list []*Node, node *Node) bool {
	for _, n := range list {
		if n == node {
			return true
		}
	}
	return false
}

/***AUFGABE***/
// ReachableNodes soll eine Liste aller von n aus erreichbaren Knoten liefern.
func (n *Node) ReachableNodes() []*Node {
	return nil
}
