package engine

func (t *Tree) Update(key uint64, val string) error {
	t.rwMu.Lock()
	var (
		node *Node
		err  error
	)

	if t.rootOff == INVALID_OFFSET {
		return NotFoundKey
	}

	if node, err = t.newMappingNodeFromPool(INVALID_OFFSET); err != nil {
		return err
	}

	if err = t.findLeaf(node, key); err != nil {
		return err
	}

	t.rwMu.Unlock()
	defer t.NodeUnlock(node.Self)
	for i, nkey := range node.Keys {
		if nkey == key {
			node.Records[i] = val
			return t.flushNodesAndPutNodesPool(node)
		}
	}

	return NotFoundKey
}
