export default function ListToTree(list, primaryKey = 'id', secondaryKey = 'pid', root = {id: 0, children: []}) {
  const mapper = {}
  list.forEach(node => {
    if (mapper[node[secondaryKey]]) {
      mapper[node[secondaryKey]].push(node)
    } else {
      mapper[node[secondaryKey]] = [node]
    }
  })
  root.children = []
  mapper[root[primaryKey]].forEach(child => {
    root.children.push(treeHelper(child, mapper, primaryKey))
  })
  return { root: root, mapper: mapper }
}

const treeHelper = (node, mapper, primaryKey = 'id') => {
  const temp = { ...node, children: []}
  if (mapper[temp[primaryKey]]) {
    mapper[temp[primaryKey]].forEach(child => {
      const t = treeHelper(child, mapper, primaryKey)
      temp.children.push(t)
    })
  }
  return temp
}
