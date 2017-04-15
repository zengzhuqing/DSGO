package avlt

//成功返回true，没有返回false。
//AVL树删除过程包括：O(log N)的搜索，O(log N)的旋转，O(log N)的平衡因子调整。
func (tr *Tree) Remove(key int32) bool {
	var target = tr.findRemoveTarget(key)
	if target == nil {
		return false
	}
	var victim, orphan = tr.findRemoveVictim(target)

	if tr.path.isEmpty() { //此时victim==target
		tr.root = orphan
	} else {
		target.key = victim.key //李代桃僵
		tr.rebalanceAfterRemove(orphan)
	}
	return true
}

func (tr *Tree) findRemoveTarget(key int32) *node {
	tr.path.clear()
	var target = tr.root
	for target != nil && key != target.key {
		if key < target.key {
			tr.path.push(target, true)
			target = target.left
		} else {
			tr.path.push(target, false)
			target = target.right
		}
	}
	return target
}

func (tr *Tree) findRemoveVictim(target *node) (victim *node, orphan *node) {
	switch {
	case target.left == nil:
		victim, orphan = target, target.right
	case target.right == nil:
		victim, orphan = target, target.left
	default:
		if target.state == 1 {
			tr.path.push(target, true)
			victim = target.left
			for victim.right != nil {
				tr.path.push(victim, false)
				victim = victim.right
			}
			orphan = victim.left
		} else {
			tr.path.push(target, false)
			victim = target.right
			for victim.left != nil {
				tr.path.push(victim, true)
				victim = victim.left
			}
			orphan = victim.right
		}
	}
	return victim, orphan
}

func (tr *Tree) rebalanceAfterRemove(orphan *node) {
	var root, lf = tr.hookSubTree(orphan)
	var state, stop = root.adjust(lf), false
	for state != 0 { //如果原平衡因子为0则子树高度不变
		if root.state != 0 { //2 || -2
			root, stop = root.rotate()
			if tr.path.isEmpty() {
				tr.root = root
			} else {
				root, lf = tr.hookSubTree(root)
				if !stop {
					state = root.adjust(lf)
					continue
				}
			}
		} else if !tr.path.isEmpty() {
			root, lf = tr.path.pop()
			state = root.adjust(lf)
			continue
		}
		break
	}
}
